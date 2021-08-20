const { fetch } = require("./fetch");

const AddUser = ({ email, clientMutationId }) => {
  const variables = { email, clientMutationId };
  const query = `mutation AddUser($email: String!, $clientMutationId: String ) {
    addUser(input: { email: $email, clientMutationId: $clientMutationId }) {
      id
      clientMutationId
    }}`;
  return fetch({ query, variables });
};

export const AddUserTest = (user) => () => {
  it("", async () => {
    const clientMutationId = user;
    const email = `${user}@test.com`;
    const resp = await AddUser({ email, clientMutationId });
    expect(resp).toMatchSnapshot();
    global.userId = resp.data.addUser.id;
  });
};

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

const User = ({ email }) => {
  const variables = { email };
  const query = `query User($email: String!) {
        user(email: $email) {
          id
          email
          completedCount
          totalCount
          todos(first: 100) {
            edges {
              cursor
              node {
                id
                complete
                __typename
              }
            }
            pageInfo {
              endCursor
              hasNextPage
            }
          }
        }
      }`;
  return fetch({ query, variables });
};

export const QueryUserTest = (user) => () => {
  it("", async () => {
    const email = `${user}@test.com`;
    const resp = await User({ email });
    expect(resp).toMatchSnapshot();
  });
};

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

export const AddTodoTest = (txt) => () => {
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

export const MarkAllTodosTest = () => () => {
  it("", async () => {
    const userId = global.userId;
    const clientMutationId = "MarkAllTodos";
    const complete = true;
    const resp = await MarkAllTodos({ complete, userId, clientMutationId });
    expect(resp).toMatchSnapshot();
  });
};

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

export const ClearCompletedTodosTest = () => () => {
  it("", async () => {
    const userId = global.userId;
    const clientMutationId = "ClearCompletedTodos";
    const resp = await ClearCompletedTodos({ userId, clientMutationId });
    expect(resp).toMatchSnapshot();
  });
};

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

export const ChangeTodoStatusTest = (complete) => () => {
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
