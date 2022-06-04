#!/bin/bash

php -d memory_limit=-1 bin/magento setup:install
        --base-url=http://$MAGE_BASE_URL \
        --db-host=$MAGE_DB_HOST \
        --db-name=$MAGE_DB_NAME \
        --db-user=$MAGE_DB_USER \
        --db-password=$MAGE_DB_PASS \
        --admin-firstname=$MAGE_ADMIN_FIRSTNAME \
        --admin-lastname=$MAGE_ADMIN_LASTNAME \
        --admin-email=$MAGE_ADMIN_EMAIL \
        --admin-user=$MAGE_ADMIN_USER \
        --admin-password=$MAGE_ADMIN_PASSWORD \
        --language=$MAGE_LANGUAGE \
        --currency=$MAGE_CURRENCY \
        --timezone=$MAGE_TIMEZONE \
        --use-rewrites=$MAGE_USE_REWRITES \
        --search-engine=$MAGE_SEARCH_ENGINE \
        --elasticsearch-host=$MAGE_ELASTICSEARCH_HOST \
        --elasticsearch-port=$MAGE_ELASTICSEARCH_PORT