# PYE Auth

**Mission**: Science compels us to create a microservice!

This is the repository for my **JWT auth microservice assignment**
with(out) blazingly fast cloud-native web3 memory-safe blockchain reactive AI
(insert a dozen more buzzwords of your choosing) technologies.

This should be done by **October 17th 2024**. Or, at the very least,
in a state that proves I am competent Go developer.

Note: **JSON** is used for storage at proof-of-concept stage for ease of use

## Current functionality

* `POST /register` - register a user with Basic Auth
* `POST /login` - get a JWT token by Basic Auth
* `GET /public-key` - get PEM-encoded public HS256 key