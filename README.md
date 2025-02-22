# secret-yaml

[![Maintainability](https://api.codeclimate.com/v1/badges/e047b9311147b1e8b419/maintainability)](https://codeclimate.com/github/antonmarin/secret-yaml/maintainability)
[![Coverage Status](https://coveralls.io/repos/github/antonmarin/secret-yaml/badge.svg?branch=master)](https://coveralls.io/github/antonmarin/secret-yaml?branch=master)
[![Build Status](https://travis-ci.org/antonmarin/secret-yaml.svg?branch=master)](https://travis-ci.org/antonmarin/secret-yaml)
[![GoDoc](https://godoc.org/github.com/antonmarin/secret-yaml?status.svg)](https://godoc.org/github.com/antonmarin/secret-yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/antonmarin/secret-yaml)](https://goreportcard.com/report/github.com/antonmarin/secret-yaml)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fantonmarin%2Fsecret-yaml.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fantonmarin%2Fsecret-yaml?ref=badge_shield)

## Goal

Lightweight tool to encrypt yaml values,
which you can quickly install right inside your pipeline.
Usable when you don't need centralized secrets management.

## Install

```
export OS=$(uname | tr '[:upper:]' '[:lower:]')
curl -LsSo /usr/local/bin/syml https://github.com/antonmarin/secret-yaml/releases/latest/download/syml-$OS
chmod +x /usr/local/bin/syml
```

## Usage

- `export SYML_SECRET=$(syml generateSecretKey)`
  generate secret and store inside env variable
- `syml encrypt --secret=${SYML_SECRET}
  ~/decryptedSecrets/secret.yaml > ~/encryptedSecrets/secret.yaml`
  encrypt values inside yaml-file and save to new file
- `syml decrypt --secret=${SYML_SECRET}
  ~/encryptedSecrets/secret.yaml > ~/decryptedSecrets/secret.yaml`
  decrypt values inside yaml-file and save to new file

[![asciicast](https://asciinema.org/a/256378.svg)](https://asciinema.org/a/256378)

`syml help` to get more about usage


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fantonmarin%2Fsecret-yaml.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fantonmarin%2Fsecret-yaml?ref=badge_large)