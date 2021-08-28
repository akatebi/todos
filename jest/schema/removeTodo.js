const { fetch } = require("../fetch");

export const removeTodo = ({ id, clientMutationId }) => {
  const variables = { id, clientMutationId };
  const query = `mutation removeTodo($id: ID!, $clientMutationId: String) {
        removeTodo(
          input: { id: $id, clientMutationId: $clientMutationId }
        ) {
          user {
            id
            email
            totalCount
            completedCount
          }
          clientMutationId
        }
      }`;
  return fetch({ query, variables });
};

