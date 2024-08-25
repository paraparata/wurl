# 🚧 wurl 🚧

> `wurl` stands for What URL to cURL?

Would be nice if it can be pipeable like this:

```sh
echo '{ "query":
  "{
    viewer {
      zones(filter: { zoneTag: $zoneTag }) {
        firewallEventsAdaptive(
          filter: $filter
          limit: 10
          orderBy: [datetime_DESC]
        ) {
          action
          clientAsn
          clientCountryName
          clientIP
          clientRequestPath
          clientRequestQuery
          datetime
          source
          userAgent
        }
      }
    }
  }",
  "variables": {
    "zoneTag": "<zone-tag>",
    "filter": {
      "datetime_geq": "2022-07-24T11:00:00Z",
      "datetime_leq": "2022-07-24T12:00:00Z"
    }
  }
}' | tr -d '\n' | curl --silent \
https://api.cloudflare.com/client/v4/graphql \
--header "Authorization: Bearer <API_TOKEN>" \
--header "Content-Type: application/json" \
--data @-
```

- [execute-graphql-query](https://developers.cloudflare.com/analytics/graphql-api/getting-started/execute-graphql-query/)

## Layout

### User Story

app:

1. Run `wurl --openapi ./schema.yml`
2. Quit by `q`
3. Show path/endpoint list

Show path/endpoint list:

1.  Show endpoint by `[mehtod] path_name \n pathItem.operation.description`
2.  Move cursor by `up/k` , `down/j`
3.  Search by `/`
4.  Choose endpoint by `o`, `enter`
    1. Default show schema
    2. Show schema: Request | Response
    3. Show form request
    4. Back to endpoint list with `esc`
    5. Show schema with `c-s`
    6. Show form request with `i`

Show schema: Request | Response

1. Back to endpoint list with `esc`

Show form request:

1. Back to endpoint list with `esc`
2. Layout
   1. Endpoint: `[method]` + recursive: endpoint_text -> query ? textinput : endpoint_text
   2. Header
   3. Body
3. Execute form with `enter`
4. wurl quit
5. Show curl command

## Reference

- [libopenapi](https://pb33f.io/libopenapi/model/)
- [undescore - side-effect import](https://stackoverflow.com/questions/21220077/what-does-an-underscore-in-front-of-an-import-statement-mean)
- [learn go w/ test](https://quii.gitbook.io/learn-go-with-tests)
- [effective go](https://go.dev/doc/effective_go)
- [extend existing struct](https://stackoverflow.com/questions/28800672/how-to-add-new-methods-to-an-existing-type-in-go)
- [where x=go](https://learnxinyminutes.com/docs/go/)
