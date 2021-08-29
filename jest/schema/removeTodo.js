const { fetch } = require("../fetch");

export const removeTodo = ({ id, userId, clientMutationId }) => {
  const variables = { id, userId, clientMutationId };
  const query = `mutation removeTodo($id: ID!, $userId: ID!, $clientMutationId: String) {
        removeTodo(
          input: { id: $id, userId: $userId, clientMutationId: $clientMutationId }
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

