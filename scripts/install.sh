#!/bin/bash

php -d memory_limit=-1 bin/magento setup:install
        --base-url=http://$MAGECLI__BASE_URL \
        --db-host=$MAGECLI__DB_HOST \
        --db-name=$MAGECLI__DB_NAME \
        --db-user=$MAGECLI__DB_USER \
        --db-password=$MAGECLI__DB_PASS \
        --admin-firstname=$MAGECLI__ADMIN_FIRSTNAME \
        --admin-lastname=$MAGECLI__ADMIN_LASTNAME \
        --admin-email=$MAGECLI__ADMIN_EMAIL \
        --admin-user=$MAGECLI__ADMIN_USER \
        --admin-password=$MAGECLI__ADMIN_PASSWORD \
        --language=$MAGECLI__LANGUAGE \
        --currency=$MAGECLI__CURRENCY \
        --timezone=$MAGECLI__TIMEZONE \
        --use-rewrites=$MAGECLI__USE_REWRITES \
        --search-engine=$MAGECLI__SEARCH_ENGINE \
        --elasticsearch-host=$MAGECLI__ELASTICSEARCH_HOST \
        --elasticsearch-port=$MAGECLI__ELASTICSEARCH_PORT