# graphql-demo

GraphQL demo for standalone and federated scenarios

## Scenario

* We have some standalone services:

  * **Todo** is a service with a simple list of items.

  * **TodoWork** is a service which uses the same id from "Todo", but adds more info regarding to some work that was done related to one todo item.

* Each one has its own GraphQL endpoints for exposing their services

* We will create federated GraphQL that merges all those endpoints in a single one

### Request path

```
Client
  \==> GraphQL Federation endpoint
         \==> Todo Application GraphQL
               \==> Todo Rest API
         \==> TodoWork GraphQL
               \==> TodoWork Rest API
```

