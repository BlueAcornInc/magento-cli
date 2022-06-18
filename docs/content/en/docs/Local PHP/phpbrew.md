---
categories: ["macos", "linux"]
tags: ["macos","linux","docs","install"] 
title: "phpbrew"
linkTitle: "phpbrew"
weight: 2
description: >
  Getting started with Magento CLI and running things locally
---

## Enviroment Variables

This is stubbed out but hasn't been defined yet, `homebrew` is the best way to use this tool at this point.

```bash
export MAGECLI_PHP_VERSION=8.1.6
export MAGECLI_PHP_BIN=/opt/phpbre/--fixme--/$MAGECLI_PHP_VERSION/bin/
export PATH=$MAGECLI_PHP_BIN:$PATH%
```

## Installation

Installing with [phpbrew](https://github.com/phpbrew/phpbrew) will also supported. This will compile php with any kind of modules or flags you need. Please review the official [Installation Documentation](https://github.com/phpbrew/phpbrew#getting-started) to set this up, but we will summerize below.

```bash
# install
curl -L -O https://github.com/phpbrew/phpbrew/releases/latest/download/phpbrew.phar
chmod +x phpbrew.phar

# Move the file to some directory within your $PATH
sudo mv phpbrew.phar /usr/local/bin/phpbrew

phpbrew init

echo "[[ -e ~/.phpbrew/bashrc ]] && source ~/.phpbrew/bashrc" >> ~/.bashrc

phpbrew install 8.1.6 +pdo +mysql

```