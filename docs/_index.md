# magento-cli

This project provides a globally-installed cli tool for managing magento locally. This includes serving it locally (`magento serve`), and helpers to install magento. As well as convienences to other tools, such as `magento sql` to access `mysql-cli`, and to install `composer` and `n98-magerun2`. This tool also passes through commands to the `bin/magento` tool in the working directory, allowing it to serve as a one-stop magento shop.

# Rationale 

Docker is typically the prefered method for running magento locally. It's great because it lets you containzerize all the services required to run magento, and even run them concurrently using a reverse-proxy like traefik with goodies like ssl and all that. 

On linux, docker works exceptionally well with near-native performance. On other platforms, however, it's typical for docker to run within a vm which significantly impacts performance. On these platforms, there's typically a lot of performance lossed through volume mounting (liek the Magento codebase, hundreds of megs of text files). This can be relieved somewhat with `mutagen`, a file syncing service that copies files back and forth between your local and the container - causing delay in updates and frustration when things don't reflect on the frontend. Further still, when the docker service containers aren't compiled for your system arch (M1/ARM7/x64) the performance losses mounts. 

Considering other popular languages and frameworks, like this `hugo` documentation site as an example, it's written in golang and when you want to serve the blog locally, you run `hugo serve` which spawns a webserver and provides a url. `pwa-studio` works the same way with `webpack`, it's a popular convention. Let's provide the same convention, and support using php's built-in web server as a way to serve magento. We can still provide docker-backed services, even a php-fpm service (with mutagen, if you insist) to round things out. But this project differentates from others because it has first-class support for pre-compiled php binaries. These are provided by `brew` and `phpbrew` by the way, easy to install, you can still run multiple php versions concurrently, and they are substantially faster than running it in a mult-arch docker container.

# Roadmap

Here's the expected features to be implemented before this project goes to 1.0

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


# Notice

This project is a prototype and isn't intended for use at this stage
