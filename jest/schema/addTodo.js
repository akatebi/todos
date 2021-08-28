const { fetch } = require("../fetch");

export const addTodo = ({ text, userId, clientMutationId }) => {
  const variables = { text, userId, clientMutationId };
  const query = `mutation addTodo($text: String!, $userId: ID!, $clientMutationId: String) {
        addTodo(
          input: { text: $text, userId: $userId, clientMutationId: $clientMutationId }
        ) {
          user {
            email
            totalCount
            completedCount
          }
          todoEdge {
            cursor
            node {
              id
              text
              complete
            }
          }
          clientMutationId
        }
      }`;
  return fetch({ query, variables });
};

