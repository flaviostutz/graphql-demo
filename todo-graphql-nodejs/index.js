const { ApolloServer, UserInputError, gql } = require("apollo-server");
const { buildFederatedSchema } = require("@apollo/federation");
const fetch = require("node-fetch");
const { readFileSync } = require('fs')

const port = 4000;
const apiUrl = `${process.env.TODO_SERVICE_URL}`;

const typeDefs = gql(readFileSync('schema.graphql', {encoding: "utf8"}))

const resolvers = {
  Todo: {
    async __resolveReference(ref) {
      var res = await fetch(`${apiUrl}/todos/${ref.id}`)
      if (res.ok) {
        return await res.json()
    } else {
        throw new Error("Error calling service. status=" + res.status)
    }
}
  },
  Query: {
    todo: async (_, {id}) => {
        console.log("TODO")
        var res = await fetch(`${apiUrl}/todos/${id}`)
        if (res.ok) {
            return await res.json()
        } else {
            throw new Error("Error calling service. status=" + res.status)
        }
    },
    todos: async (_) => {
        console.log("TODOS")
        res = await fetch(`${apiUrl}/todos`)
        if (res.ok) {
            return await res.json()
        } else {
            throw new Error("Error calling service. status=" + res.status)
        }
    }
  },
  Mutation: {
    createTodo: async (_, args) => {
        console.log("CREATE TODO " + JSON.stringify(args))
        var res = await fetch(`${apiUrl}/todos`, 
                            { 
                                method: 'POST', 
                                body: JSON.stringify(args),
                                headers: {'Content-Type': 'application/json'}
                            })
        if (res.ok) {
            return await res.json()
        } else if (res.status == 400) {
            var b = await res.json()
            throw new UserInputError("Missing input fields. message=" + b.message)
        } else {
            throw new Error("Couldn't create Todo. status=" + res.status)
        }
    },
    updateTodo: async (_, args) => {
        console.log("UPDATE TODO")
        var res = await fetch(`${apiUrl}/todos/${args.id}`,
                            { 
                                method: 'POST', 
                                body: JSON.stringify(args),
                                headers: {'Content-Type': 'application/json'}
                            })
        if (res.ok) {
            return await res.json()
        } else {
            throw new Error("Error calling service. status=" + res.status)
        }
    },
    deleteTodo: async (_, {id}) => {
        console.log("DELETE TODO")
        var res = await fetch(`${apiUrl}/todos/${args.id}`, { method: 'DELETE'})
        if (res.ok) {
            return await res.json()
        } else {
            throw new Error("Error calling service. status=" + res.status)
        }
    }
  }
}

const server = new ApolloServer({
    schema: buildFederatedSchema([{ typeDefs, resolvers }]),
    // typeDefs, 
    // resolvers,
    tracing: true,
    debug: true
});

server.listen({ port }).then(({ url }) => {
  console.log(`Todo GraphQL ready at ${url}`);
});

