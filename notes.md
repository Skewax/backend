
# Implementation of Handler Notes

## Get Files
the api will return a next page if its not all loaded
this means you'll need to either load all and store until
done (susceptible to attacks with too many files) or
have a cap and implement a paging system clientside

## User Data Removal
both backend and frontend, need to be able to 
completely clear user data on request

## Caching Session Tokens and Google Access Tokens
when user is given a session token, it should 
1. upload to database with all data
2. cache key value pair with structure
  - key: session token
  - value:
    - google access token
    - ID associated with token (for user verification)
    - caching expiry (some predetermined period like 10 minutes)
    - access token expiry (subtract some time like a minute to guarantee not sending bad)
    - session token expiry (maybe don't need?)

whenever user makes request it should first go through caching data then if token isn't there
it will check database. only when neither contain the token does it return a permissions error

# Database structure

## NOTES
if a session token is attempted to be used
by a different user, it should be deleted from
its actual user

## User
- key: google user ID
- access_token (from google)
- refresh_token (from google)
- sessions:
  - array of pointers to session objects

## Session
- pointer to its corresponding user object
- session_token
- timeout_date



