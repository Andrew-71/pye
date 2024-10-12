# Auth microservice

**Mission**: Science compels us to create a microservice!

This is the repository for my **JWT auth microservice assignment**
with(out) blazingly fast cloud-native web3 memory-safe blockchain reactive AI
(insert a dozen more buzzwords of your choosing) technologies.

This should be done by **October 17th 2024**. Or, at the very least,
in a state that proves I am competent Go developer.

## Current functionality

## `serve`

* `POST /register` - register a user with Basic Auth
* `POST /login` - get a JWT token by Basic Auth
* `GET /pem` - get PEM-encoded public RS256 key
* Data persistently stored in an SQLite database
* RS256 key loaded from a file or generated on startup if missing

## `verify`

* Verify JWT via public key in a PEM file