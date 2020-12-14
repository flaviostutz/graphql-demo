const { ApolloServer } = require("apollo-server");
const { ApolloGateway } = require("@apollo/gateway");

const port = 5000;

const gateway = new ApolloGateway({
  serviceList: [
    { name: "test1", url: process.env.FEDERATED_GRAPHQL_URLS },
  ]
});

const server = new ApolloServer({
  gateway,
  subscriptions: false
});

server.listen({ port }).then(({ url }) => {
  console.log(`Server ready at ${url}`);
});
