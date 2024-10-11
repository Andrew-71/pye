# PYE Auth

**Mission**: Science compels us to create a microservice!

This is the repository for my **JWT auth microservice assignment**
with(out) blazingly fast cloud-native web3 memory-safe blockchain reactive AI
(insert a dozen more buzzwords of your choosing) technologies.

This should be done by **October 17th 2024**, or at the very least,
in a shape that proves I am somewhat competent.

## Course of action

How I currently see this going

1. Make an HTTP Basic Auth -> JWT -> Open key API
2. Create simple frontend (really stretching the definition) to test it
3. Ask myself and others - "Is this a microservice?"
If the answer is yes, rejoice.
If the answer is no, rejoice for a different reason.
4. Once it's technically solid-ish, polish ever-so-slightly

## "Technology stack"

The technology I *intend* on using

1. **Data storage - SQLite**.
Definitely want to avoid a full-sized DB because they're oversized for most
projects. To be honest, even **JSON** would do for this.  
In fact, this might just be the way to go for the proof-of-concept, hm...
2. **Frontend - template/html module**. Duh, I am anti-bloat.
3. **HTTP routing - Chi**.
I'd use `net/http`, but a deadline of 1 week means speed is everything.