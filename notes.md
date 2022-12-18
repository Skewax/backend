
# Implementation of Handler Notes

## Get Files
the api will return a next page if its not all loaded
this means you'll need to either load all and store until
done (susceptible to attacks with too many files) or
have a cap and implement a paging system clientside

## User Data Removal
both backend and frontend, need to be able to 
completely clear user data on request

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



