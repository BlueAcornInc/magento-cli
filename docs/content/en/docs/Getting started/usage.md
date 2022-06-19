---
title: Usage
description: Basic usage conventions and information
---

# .magento-cli directories

A project repostiory can override the default configuration by storing replacement services or tasks yaml directories in `.magento-cli`. For example, let's say you would like to change the database docker composition, you canc create `.magento-cli/services/database.yaml` and this file will be interpreted when you run `magento serve` instead of the version sourced by magento-cli.

