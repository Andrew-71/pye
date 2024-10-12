# Auth microservice

**Mission**: Science compels us to create a microservice!

This is the repository for my **JWT auth microservice assignment**
with(out) blazingly fast cloud-native web3 memory-safe blockchain reactive AI
(insert a dozen more buzzwords of your choosing) technologies.

This should be done by **October 17th 2024**. Or, at the very least,
in a state that proves I am competent Go developer.

## Commands

## `pye serve [--config] [--port] [--db]`

Serve a simple JWT auth system

* `POST /register` - register a user with Basic Auth
* `POST /login` - get a JWT token by Basic Auth
* `GET /pem` - get PEM-encoded public RS256 key
* Data and RS256 key persistently stored in an SQLite database and a PEM file

## `pye verify <jwt> <pem file>`

Verify a JWT with a public key from a PEM file