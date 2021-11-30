# how to bitquery

## fetch bitquery schema

* install graphqurl [guide](https://github.com/hasura/graphqurl#steps-to-install-cli)
    gq https://graphql.bitquery.io -H 'X-API-KEY: BQYvhnv04csZHaprIBZNwtpRiDIwEIW9' -H 'content-type: application/json' -q 'query MyQuery { ethereum(network: ethereum) { blocks { count } } }'
* export schema
    gq https://graphql.bitquery.io -H 'X-API-KEY: BQYvhnv04csZHaprIBZNwtpRiDIwEIW9' -H 'content-type: application/json' --introspect > schema.graphqls