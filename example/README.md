# Example Installation

This directory provides a working example of a Magento installation. You can run the uncompiled version of magento-cli against this directory to work on the tool.

Be sure to configure PHP Locally via [Homebrew](https://superterran.net/magento-cli/docs/local-php/homebrew/) before running the example.

```bash
$ cd example
$ composer install # must install magento before you can serve
$ go run .. install # install magento before serving
$ go run .. serve

loading config file magento-cli.yaml... done
magento-cli ≫ starting task serve
magento-cli.serve ≫ [Sat Jun 18 12:16:10 2022] PHP 8.1.6 Development Server (http://0.0.0.0:8080) started
``` 
