const { fetch } = require("../fetch");

export const addUser = ({ email, clientMutationId }) => {
  const variables = { email, clientMutationId };
  const query = `mutation AddUser($email: String!, $clientMutationId: String ) {
    addUser(input: { email: $email, clientMutationId: $clientMutationId }) {
      id
      clientMutationId
    }}`;
  return fetch({ query, variables });
};


