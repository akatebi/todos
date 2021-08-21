const { fetch } = require("../fetch");

const RenameTodo = ({ id, text, clientMutationId }) => {
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

const RenameTodoTest = () => () => {
  it("", async () => {
    global.todoIds = [];
    for (let i = 0; i < 3; i++) {
      const id = global.todoIds[i];
      const text = `renamed ${i}`;
      const clientMutationId = `RemoveToDo-${i}`;
      const resp = await RenameTodo({ id, text, clientMutationId });
      expect(resp).toMatchSnapshot();
    }
  });
};

export default RenameTodoTest;
