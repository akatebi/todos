const { fetch } = require("../fetch");

const RemoveTodo = ({ id, clientMutationId }) => {
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

const RemoveTodoTest = () => () => {
  it("", async () => {
    global.todoIds = [];
    for (let i = 0; i < 3; i++) {
      const id = global.todoIds[i];
      const clientMutationId = `RemoveToDo-${i}`;
      const resp = await RemoveTodo({ id, clientMutationId });
      expect(resp).toMatchSnapshot();
    }
  });
};

export default RemoveTodoTest;
