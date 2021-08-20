const { fetch } = require("../fetch");

const ClearCompletedTodos = ({ userId, clientMutationId }) => {
    const variables = { userId, clientMutationId };
    const query = `mutation clearCompletedTodos($userId: ID!, $clientMutationId: String) {
      clearCompletedTodos(input: {userId: $userId, clientMutationId: $clientMutationId}) {
        deletedTodoIds
        user {
          id
          completedCount
          totalCount
        }
        clientMutationId
      }
    }
    `;
    return fetch({ query, variables });
  };
  
  const ClearCompletedTodosTest = () => () => {
    it("", async () => {
      const userId = global.userId;
      const clientMutationId = "ClearCompletedTodos";
      const resp = await ClearCompletedTodos({ userId, clientMutationId });
      expect(resp).toMatchSnapshot();
    });
  };
  
  export default ClearCompletedTodosTest;