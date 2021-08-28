const { fetch } = require("../fetch");

export const renameTodo = ({ id, text, clientMutationId }) => {
  const variables = { id, text, clientMutationId };
  const query = `mutation renameTodo($id: ID!, $text: String!, $clientMutationId: String) {
        renameTodo(
          input: { id: $id, text: $text, clientMutationId: $clientMutationId }
        ) {
          todo {
            id
            text
            complete
          }
          clientMutationId
        }
      }`;
  return fetch({ query, variables });
};


