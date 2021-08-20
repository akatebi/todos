const { fetch } = require("../fetch");

const ChangeTodoStatus = ({ complete, id, userId, clientMutationId }) => {
  const variables = { complete, id, userId, clientMutationId };
  const query = `mutation ChangeTodoStatus(
    $complete: Boolean!
    $id: ID!
    $userId: ID!
    $clientMutationId: String
  ) {
    changeTodoStatus(input: { 
      complete: $complete, id: $id, userId: $userId, 
      clientMutationId: $clientMutationId }) {
      todo {
        id
        text
        complete
      }
      user {
        id
        email
        totalCount
        completedCount
      }
      clientMutationId
    }
  }  
  `;
  return fetch({ query, variables });
};

const ChangeTodoStatusTest = (complete) => () => {
  it("", async () => {
    const userId = global.userId;
    for (let i = 0; i < 3; i++) {
      const id = todoIds[i];
      const clientMutationId = `ChangeTodoStatus-${i}`;
      const resp = await ChangeTodoStatus({
        complete,
        id,
        userId,
        clientMutationId,
      });
      expect(resp).toMatchSnapshot();
    }
  });
};

export default ChangeTodoStatusTest;
