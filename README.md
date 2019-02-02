# :closed_lock_with_key: Secret

This project is based on [the Gophercises lesson](https://gophercises.com/exercises/secret). I have made my own [changes](#changes-from-gophercises) though.

## DANGER

This project implements cryptography and **has not been audited**. I also do not claim to know what I'm doing when it comes to cryptography, so **under no circumstances should you assume that this is safe to use in production**! If you are looking for a proper secret storage solution, use [Vault](https://www.vaultproject.io/) instead.

## Usage

Secret can be used as a library in Go, or it can be used as a CLI application, using the wrapper.

Adding and retrieving a secret from Go is as simple as the following:

```go
//TODO
```

The same can be done on the command line like this:

```bash
$ echo "TODO"
TODO
```

## Installing

`go get -u github.com/kscarlett/secret`

## Changes from Gophercises

- Use SHA256 instead of MD5 to turn the passphrase into the AES key in order to use AES-256 instead of AES-128.