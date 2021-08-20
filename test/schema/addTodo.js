const { fetch } = require("../fetch");

const AddTodo = ({ text, userId, clientMutationId }) => {
  const variables = { text, userId, clientMutationId };
  const query = `mutation addTodo($text: String!, $userId: ID!, $clientMutationId: String) {
        addTodo(
          input: { text: $text, userId: $userId, clientMutationId: $clientMutationId }
        ) {
          user {
            email
            totalCount
            completedCount
          }
          todoEdge {
            cursor
            node {
              id
              text
              complete
            }
          }
          clientMutationId
        }
      }`;
  return fetch({ query, variables });
};

const AddTodoTest = (txt) => () => {
  it("", async () => {
    global.todoIds = [];
    for (let i = 0; i < 3; i++) {
      const text = `${txt} ${i + 1}`;
      const clientMutationId = `AddToDo-${i}`;
      const resp = await AddTodo({ text, userId, clientMutationId });
      expect(resp).toMatchSnapshot();
      global.todoIds.push(resp.data.addTodo.todoEdge.node.id);
    }
  });
};

export default AddTodoTest;