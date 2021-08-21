const { fetch } = require("../fetch");

const MarkAllTodos = ({ complete, userId, clientMutationId }) => {
    const variables = { complete, userId, clientMutationId };
    const query = `mutation markAllTodos($complete: Boolean!, $userId: ID!, $clientMutationId: String) {
      markAllTodos(input: { complete: $complete, userId: $userId, clientMutationId: $clientMutationId}) {
        user {
          id
          completedCount
          totalCount
        }
        changedTodos {
          id
          text
          complete
        }
        clientMutationId
      }
    }
    `;
    return fetch({ query, variables });
  };
  
  const MarkAllTodosTest = () => () => {
    it("", async () => {
      const userId = global.userId;
      const clientMutationId = "MarkAllTodos";
      const complete = true;
      const resp = await MarkAllTodos({ complete, userId, clientMutationId });
      expect(resp).toMatchSnapshot();
    });
  };
  
export default MarkAllTodosTest;