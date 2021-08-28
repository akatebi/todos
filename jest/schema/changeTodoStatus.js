const { fetch } = require("../fetch");

export const changeTodoStatus = ({ complete, id, userId, clientMutationId }) => {
  const variables = { complete, id, userId, clientMutationId };
  const query = `mutation ChangeTodoStatus(
    $complete: Boolean!
    $id: ID!
    $userId: ID!
    $clientMutationId: String
  ) {
    changeTodoStatus(input: { 
      complete: $complete, id: $id, userId: $userId, 
      clientMutationId: $clientMutationId }) {
      todo {
        id
        text
        complete
      }
      user {
        id
        email
        totalCount
        completedCount
      }
      clientMutationId
    }
  }  
  `;
  return fetch({ query, variables });
};

