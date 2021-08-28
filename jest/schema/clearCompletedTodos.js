const { fetch } = require("../fetch");

export const clearCompletedTodos = ({ userId, clientMutationId }) => {
    const variables = { userId, clientMutationId };
    const query = `mutation clearCompletedTodos($userId: ID!, $clientMutationId: String) {
      clearCompletedTodos(input: {userId: $userId, clientMutationId: $clientMutationId}) {
        deletedTodoIds
        user {
          id
          completedCount
          totalCount
        }
        clientMutationId
      }
    }
    `;
    return fetch({ query, variables });
  };
  
  