const { fetch } = require("../fetch");

export const markAllTodos = ({ complete, userId, clientMutationId }) => {
    const variables = { complete, userId, clientMutationId };
    const query = `mutation markAllTodos($complete: Boolean!, $userId: ID!, $clientMutationId: String) {
      markAllTodos(input: { complete: $complete, userId: $userId, clientMutationId: $clientMutationId}) {
        user {
          id
          completedCount
          totalCount
        }
        changedTodos {
          id
          text
          complete
        }
        clientMutationId
      }
    }
    `;
    return fetch({ query, variables });
  };
  
