const { fetch } = require("../fetch");

export const removeUser = ({ email, clientMutationId }) => {
  const variables = { email, clientMutationId };
  const query = `mutation RemoveUser($email: String!, $clientMutationId: String) {
    removeUser(input: { email: $email, clientMutationId: $clientMutationId }) {
      clientMutationId
    }}`;
  return fetch({ query, variables });
};

