# Pye Auth

**Mission**: Science compels us to create a microservice!

This is the repository for my **JWT authentication microservice**
with(out) blazingly fast cloud-native web3 memory-safe blockchain reactive AI
(insert a dozen more buzzwords of your choosing) technologies.

## Usage

```
Usage:
  pye [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  find        Find a user
  help        Help about any command
  serve       Start JWT service
  verify      Verify a JWT token

Flags:
  -c, --config string   config file (default "config.json")
      --db string       database to use
  -d, --debug           enable debug mode
  -h, --help            help for pye

Use "pye [command] --help" for more information about a command.
```

## Technologies used

* **Storage** - [SQLite](https://github.com/mattn/go-sqlite3) and a
[PEM](https://en.wikipedia.org/wiki/Privacy-Enhanced_Mail) file
* **HTTP routing** - [Chi](https://go-chi.io), just for logging...
* **CLI management** - [Cobra](https://cobra.dev/)