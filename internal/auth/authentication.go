package skewauth

import (
	"log"
	"time"

	"database/sql"

	"github.com/akyoto/cache"
	_ "github.com/lib/pq"
)

type GoogleAuthToken struct {
	accessToken   string
	refreshToken  string
	accessTimeout time.Time
}

type AuthToken struct {
	gAuth  GoogleAuthToken
	uToken string
	status string
}

func (t AuthToken) OK() bool {
	return t.status == "OK"
}

func (t AuthToken) GetErrorCode() string {
	return t.status[:3]
}

func (t AuthToken) GetErrorDesc() string {
	if len(t.status) > 6 {
		return t.status[5:]
	}
	return ""
}

func (t AuthToken) SessionToken() string {
	return t.uToken
}

func (t AuthToken) GoogleToken() string {
	return t.gAuth.accessToken
}

func (t AuthToken) GoogleTimeout() int64 {
	return t.gAuth.accessTimeout.Unix()
}

type AuthConfig struct {
	CacheTime           time.Duration
	ConnStr             string
	GClientId           string
	GClientSecret       string
	authorizedEndpoints []string
}

type UseTokenReq struct {
	Uid        string
	Token      string
	CanReplace bool // if endpoint returns the to be used token it can replace
	Result     chan AuthToken
}

type GenTokenReq struct {
	AuthCode     string
	GrantType    string
	HostedDomain string
	Result       chan AuthToken
}

type ClearTokenReq struct {
	Uid             string
	Token           string
	Result          chan bool
	ShouldClearUser bool //will remove access and refresh tokens as well
}

type DBUserSession struct {
	SessionToken string `field:"sessionkey"`
	Uid          string `field:"userid"`
	Timeout      int64  `field:"expiry"`
}

type DBUser struct {
	id           string
	accesstoken  string
	refreshtoken string
	expiry       int64
}

func Begin(quit chan bool, useTokenReqs chan *UseTokenReq, genTokenReqs chan *GenTokenReq, clearTokenReqs chan *ClearTokenReq, conf *AuthConfig) {
	//initialization of all authentication connections and such
	db, err := sql.Open("postgres", conf.ConnStr)
	if err != nil {
		log.Fatal("open database error") //comment for testing
	}
	defer db.Close()

	c := cache.New(conf.CacheTime)
	defer c.Close()

	go beginTokenGenerator(genTokenReqs, db, c, conf, quit)
	go beginTokenValidator(useTokenReqs, db, c, conf, quit)
	go beginTokenClearer(clearTokenReqs, db, c, conf, quit)

	<-quit
}

func ErrAuthToken(err string) AuthToken {
	return AuthToken{
		gAuth:  GoogleAuthToken{},
		uToken: "",
		status: err,
	}
}

func beginTokenGenerator(reqs chan *GenTokenReq, db *sql.DB, c *cache.Cache, conf *AuthConfig, quit chan bool) {
	for req := range reqs {
		GenerateToken(req, conf)
	}
	<-quit
}

func beginTokenValidator(reqs chan *UseTokenReq, db *sql.DB, c *cache.Cache, conf *AuthConfig, quit chan bool) {
	for req := range reqs {
		ValidateTokenReq(req, db, conf)
	}
	<-quit
}

func beginTokenClearer(reqs chan *ClearTokenReq, db *sql.DB, c *cache.Cache, conf *AuthConfig, quit chan bool) {
	for req := range reqs {
		ClearToken(req, conf)
	}
	<-quit
}

func ValidateTokenReq(req *UseTokenReq, db *sql.DB, conf *AuthConfig) {
	rows, err := db.Query("SELECT sessionkey, userid, expiry FROM usersessions WHERE sessionkey = '" + req.Token + "'")
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
		req.Result <- ErrAuthToken("500: internal server error. " + err.Error())
		return
	}
	//!rows.next implies no session in db
	if !rows.Next() {
		req.Result <- ErrAuthToken("400: Invalid Session")
		return
	}
	//throwing first row result of query into a usersession
	//this is guaranteed to be the only result because sessionkeys are unique
	var uSession DBUserSession
	rows.Scan(&uSession.SessionToken, &uSession.Uid, &uSession.Timeout)

	//error checking to match posted userID with sessionID
	//NOTE error method should be identical to no session in db
	//above for security reasons. If they're different an attacker
	//knows that they've found a valid session by guessing
	if uSession.Uid != req.Uid {
		req.Result <- ErrAuthToken("400: Invalid Session")
	}

	req.Result <- AuthToken{
		gAuth: GoogleAuthToken{
			accessToken:   "Access",
			refreshToken:  "Refresh Token",
			accessTimeout: time.Now(),
		},
		uToken: "",
		status: "OK",
	}
	return
}

func GenerateToken(req *GenTokenReq, conf *AuthConfig) {
	//checking request
	authorized := false
	for _, endpoint := range conf.authorizedEndpoints {
		if req.HostedDomain == endpoint {
			authorized = true
			break
		}
	}
	if !authorized {
		req.Result <- AuthToken{status: "500: Unauthorized Domain"}
		return
	}
	//TEMP
	req.Result <- AuthToken{
		gAuth: GoogleAuthToken{
			accessToken:   "Access Token",
			refreshToken:  "Refresh Token",
			accessTimeout: time.Now(),
		},
		uToken: "",
		status: "",
	}
}

func ClearToken(req *ClearTokenReq, conf *AuthConfig) {
	//TEMP
	req.Result <- true
}
