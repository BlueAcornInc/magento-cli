---
categories: ["macos", "linux"]
tags: ["macos","linux","docs","install"] 
title: "Homebrew"
linkTitle: "Homebrew"
weight: 2
description: >
  Getting started with Magento CLI and running things locally
---

Installing with [brew](https://brew.sh/) is the preferred way to install for most use-cases. Homebrew installs the tool globally, and is updated with every release. 

## Environment Variables

Exporting the following will peg you to a specific version of php for your instance, using homebrew. This is a makeshift approach that requires `direnv` or the like, in the future the tool will abstract this away for you. 

```bash
export MAGECLI_PHP_VERSION=8.1.6
export MAGECLI_PHP_BIN=/opt/homebrew/Cellar/php/$MAGECLI_PHP_VERSION/bin/
export PATH=$MAGECLI_PHP_BIN:$PATH%
```
## Installation 

Installing PHP with brew is easy, just trigger: 

```bash
brew install php@8.1
```

If you run `brew link php@8.1` this will set it as the system-wide default. This seems like it should be fine, but keep in mind tools like `magento-cloud` require PHP 7.4 at the time of this writing so having the system-wide as that version is usually simpler. See the environment variables section to set this per-directory.