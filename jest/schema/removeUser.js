const { fetch } = require("../fetch");

const RemoveUser = ({ email, clientMutationId }) => {
  const variables = { email, clientMutationId };
  const query = `mutation RemoveUser($email: String!, $clientMutationId: String) {
    removeUser(input: { email: $email, clientMutationId: $clientMutationId }) {
      clientMutationId
    }}`;
  return fetch({ query, variables });
};

export const RemoveUserTest = (user) => () => {
  it("", async () => {
    const clientMutationId = user;
    const email = `${user}@test.com`;
    const resp = await RemoveUser({ email, clientMutationId });
    expect(resp).toMatchSnapshot();
  });
};

export default RemoveUserTest;