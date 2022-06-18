---
categories: ["roadmap"]
tags: ["roadmap"]
title: "Project Roadmap"
linkTitle: "Roadmap"
date: 2022-06-18
description: >
  Here are the expected features to be implemented before this project goes to 1.0
---

* Using system-provided php runtime for executing magento
    * currently uses `php -S` provided by phpbrew or homebrew
    * supports concurrency with different versions through environment variables
* using dockerized php runtime for executing magento
    * should provide a default local setup with all-docker
* dockerized backend services
    * currently `elasticsearch` and `databases`
    * variablized, can use any version or image of services
    * ability to add other services through yaml
    * expected to add `mailhog`, `ampq`, `redis`, and friends
* `magento-cli` runs as a global cli app
    * all yaml configurations packaged with executable
    * `magento serve` kicks off dev instance with zero configuration
* Documentation
    * hugo-based documentation site hosted on [github.io](https://superterran.net/magento-cli/)
* Available on Popular Package Managers
    * brew (macos, linux)
    * chocolatey (windows)?
* Build Targets
    * macos (arm/intel)
    * linux (arm/intel)
    * windows (intel)
    * chromeos/crostini (intel and probably arm)