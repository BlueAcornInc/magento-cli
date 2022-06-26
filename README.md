# magento-cli

[Documentation](https://magento-cli.com/docs/) |
[Discussions](https://github.com/blueacorninc/magento-cli/discussions) |
[Twitter](https://twitter.com/superterran) |
[Contribution Guide](CONTRIBUTING.md)

Running magento and adobe commerce locally with a cli tool

[![GoDoc](https://godoc.org/github.com/gohugoio/hugo?status.svg)](https://pkg.go.dev/github.com/blueacorninc/magento-cli)
[![Go](https://github.com/blueacorninc/magento-cli/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/blueacorninc/magento-cli/actions/workflows/go.yml) 
[![Go Report Card](https://goreportcard.com/badge/github.com/blueacorninc/magento-cli)](https://goreportcard.com/report/github.com/blueacorninc/magneto-cli)
[![codecov](https://codecov.io/gh/blueacorninc/magento-cli/branch/main/graph/badge.svg?token=S48U2MJP9I)](https://codecov.io/gh/blueacorninc/magento-cli)

# Installation 

Binaries are compiled with every release, you can grab it from the [releases](https://github.com/blueacorninc/magento-cli/releases/) page, and use it as-is. These files are fit to be ran directly, from $PATH, or even committed i.e. `/path/to/iac-repo/bin/magento-cli` and invoked with `cd /path/to/iac-repo/ && bin/magento-cli`.

## Homebrew

Installing with [brew](https://brew.sh/) is the preferred way to install for most use-cases. Homebrew installs the tool globally, and is updated with every release. 

```/bin/bash
    brew tap blueacorninc/magento-cli
    brew install magento-cli
```

## Compiling From Source

If you prefer to compile from source, the Makefile can be used:

```bash
    git clone git@github.com:blueacorninc/magento-cli.git 
    cd magento-cli
    make install # runs `go build .` and copies to /usr/local/bin
```

## Running Source Locally

```bash
    cd example/
    composer install
    go run .. serve
```
## PHP Concurrency 

Exporting the following will peg you to a specific version of php for your instance, using homebrew. This is a makeshift approach, in the future the tool will abstract this away for you. 

```bash
export MAGECLI_PHP_VERSION=8.1.6
export MAGECLI_PHP_BIN=/opt/homebrew/Cellar/php/$MAGECLI_PHP_VERSION/bin/
export PATH=$MAGECLI_PHP_BIN:$PATH%
```

# Contributing

For a complete guide to contributing to magento-cli, see the [Contribution Guide](CONTRIBUTING.md)

Bug reports and pull requests are welcome on GitHub at https://github.com/blueacorninc/magento-cli/issues. 

# License
Mach is released under the [MIT License](LICENSE)

# Author Information
This project was started in 2022 by Doug Hatcher <doug.hatcher@blueaornici.com>.
