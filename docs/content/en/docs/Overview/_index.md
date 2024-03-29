---
title: "Overview"
linkTitle: "Overview"
weight: 1
description: >
  How does Magento CLI work, how is it better?
---

This project provides a globally-installed cli tool for managing magento locally. This includes serving it locally (`magento serve`), and helpers to install magento. As well as convienences to other tools, such as `magento sql` to access `mysql-cli`, and to install `composer` and `n98-magerun2`. This tool also passes through commands to the `bin/magento` tool in the working directory, allowing it to serve as a one-stop magento shop.

# Rationale 

Docker is typically the preferred method for running Magento locally. It's great because it lets you containerize all the services required to run Magento, and even run them concurrently using a reverse proxy like traefik with goodies like tls and all that. 

On Linux, docker works exceptionally well with near-native performance. On other platforms, however, it's typical for docker to run within a VM which significantly impacts performance. On these platforms, there's typically a lot of performance lost through volume mounting (the Magento codebase is hundreds of megs of small text files). This can be relieved somewhat with `mutagen`, a file syncing service that copies files back and forth between your local and the container - causing a delay in updates and frustration when things don't reflect on the frontend. Further still, when the docker service containers aren't compiled for your system arch (M1/ARM7/x64) the performance losses mounts. 

Considering other popular languages and frameworks, like this `hugo` documentation site as an example, it's written in golang and when you want to serve the blog locally, you run `hugo serve` which spawns a web server and provides a url. `pwa-studio` works the same way with `webpack`, it's a popular convention. Let's provide the same convention, and support using php's built-in web server as a way to serve Magento. We can still provide docker-backed services, even a php-fpm service (with mutagen, if you insist) to round things out. But this project differentiates from others because it has first-class support for pre-compiled php binaries. These are provided by `brew` and `phpbrew` by the way, easy to install, you can still run multiple php versions concurrently, and they are substantially faster than running it in a multi-arch docker container.