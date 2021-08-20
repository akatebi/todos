const { fetch } = require("../fetch");

const AddUser = ({ email, clientMutationId }) => {
  const variables = { email, clientMutationId };
  const query = `mutation AddUser($email: String!, $clientMutationId: String ) {
    addUser(input: { email: $email, clientMutationId: $clientMutationId }) {
      id
      clientMutationId
    }}`;
  return fetch({ query, variables });
};

const AddUserTest = (user) => () => {
  it("", async () => {
    const clientMutationId = user;
    const email = `${user}@test.com`;
    const resp = await AddUser({ email, clientMutationId });
    expect(resp).toMatchSnapshot();
    global.userId = resp.data.addUser.id;
  });
};

export default AddUserTest;