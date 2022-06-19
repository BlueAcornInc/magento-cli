---
categories: ["Getting Started"]
tags: [] 
title: "Getting Started"
linkTitle: "Getting Started"
weight: 2
description: >
  Getting started with Magento CLI and running things locally
---

## Prerequisites

Magento CLI is designed to work wtih `docker` and `docker-compose`, and requires a locally installed verison of `php` to use the built-in php server. If you are running php locally, you will also need to provide tools like `composer`.


## Installation

Binaries are compiled with every release, you can grab it from the [releases](https://github.com/blueacorninc/magento-cli/releases/) page, and use it as-is. These files are fit to be ran directly, from $PATH, or even committed i.e. `/path/to/iac-repo/bin/magento-cli` and invoked with `cd /path/to/iac-repo/ && bin/magento-cli`.

### Homebrew

Installing with [brew](https://brew.sh/) is the preferred way to install for most use-cases. Homebrew installs the tool globally, and is updated with every release. 

```/bin/bash
    brew tap blueacorninc/magento-cli
    brew install magento-cli
```

### Compiling From Source

If you prefer to compile from source, the Makefile can be used:

```bash
    git clone git@github.com:blueacorninc/magento-cli.git 
    cd magento-cli
    make install # runs `go build .` and copies to /usr/local/bin
```

### Running Source Locally

```bash
    cd example/
    composer install
    go run .. serve
```

## Try it out!

Once installed, you can test out the tool against any project!

```bash
cd example
composer install # lets install the php deps
magento install # runs setup:install and creates config.php and env.php

magento serve # That's it!
```
