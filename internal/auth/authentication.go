package skewauth

import (
	"time"

	"database/sql"

	"github.com/akyoto/cache"
	_ "github.com/lib/pq"
)

var _ = sql.Named

type GoogleAuthToken struct {
	accessToken  string
	refreshToken string
	accessLife   time.Duration
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
	return t.status[5:]
}

func (t AuthToken) SessionToken() string {
	return t.uToken
}

func (t AuthToken) GoogleToken() string {
	return t.gAuth.accessToken
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

func Begin(quit chan bool, useTokenReqs chan *UseTokenReq, genTokenReqs chan *GenTokenReq, clearTokenReqs chan *ClearTokenReq, conf *AuthConfig) {
	//initialization of all authentication connections and such
	db, err := sql.Open("postgres", conf.ConnStr)
	if err != nil {
		//log.Fatal("open database error") //comment for testing
	}
	defer db.Close()
	c := cache.New(conf.CacheTime)
	defer c.Close()

	go beginTokenGenerator(genTokenReqs, db, c, conf, quit)
	go beginTokenValidator(useTokenReqs, db, c, conf, quit)
	go beginTokenClearer(clearTokenReqs, db, c, conf, quit)

	<-quit
}

func beginTokenGenerator(reqs chan *GenTokenReq, db *sql.DB, c *cache.Cache, conf *AuthConfig, quit chan bool) {
	for req := range reqs {
		GenerateToken(req, conf)
	}
	<-quit
}

func beginTokenValidator(reqs chan *UseTokenReq, db *sql.DB, c *cache.Cache, conf *AuthConfig, quit chan bool) {
	for req := range reqs {
		ValidateTokenReq(req, conf)
	}
	<-quit
}

func beginTokenClearer(reqs chan *ClearTokenReq, db *sql.DB, c *cache.Cache, conf *AuthConfig, quit chan bool) {
	for req := range reqs {
		ClearToken(req, conf)
	}
	<-quit
}

func ValidateTokenReq(req *UseTokenReq, conf *AuthConfig) {
	//checking for a valid request
	//TEMP
	req.Result <- AuthToken{GoogleAuthToken{"", "", time.Duration(1 * time.Minute)}, "the previously valid session token", "OK"}
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
	req.Result <- AuthToken{GoogleAuthToken{"", "", time.Duration(1 * time.Minute)}, "a new session token", "OK"}
}

func ClearToken(req *ClearTokenReq, conf *AuthConfig) {
	//TEMP
	req.Result <- true
}
