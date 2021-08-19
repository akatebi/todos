const { fetch } = require("./fetch");

const AddUser = ({email, clientMutationId}) => {
  const variables = { email, clientMutationId };
  const query = `mutation AddUser($email: String!, $clientMutationId: String ) {
    addUser(input: { email: $email, clientMutationId: $clientMutationId }) {
      id
      clientMutationId
    }}`;
    return fetch({query, variables});
}

const RemoveUser = ({email, clientMutationId}) => {
  const variables = { email, clientMutationId };
  const query = `mutatation RemoveUser($email: String!, $clientMutationId: String) {
    removeUser(input: { email: $email, clientMutationId: $clientMutationId }) {
      clientMutationId
    }}`;
    return fetch({query, variables});
}


const User = ({email}) => {
    const variables = { email }; 
    const query = `query User($email: String!, $clientMutationId: String) {
        user(email: $email, clientMutationId: $clientMutationId) {
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
      return fetch({query, variables});
}

const AddTodo = ({text, userId, clientMutationId}) => {
const variables = {text, userId, clientMutationId};
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
      return fetch({query, variables});
    }

const MarkAllTodos = ({complete, userId, clientMutationId}) => {
  const variables = {complete, userId, clientMutationId};
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
  return fetch({query, variables});
}

const ClearCompletedTodos = ({userId, clientMutationId}) => {
  const variables = {userId, clientMutationId};
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
  return fetch({query, variables});
}

const ChangeTodoStatus = ({complete, id, userId, clientMutationId}) => {
  const variables = {complete, id, userId, clientMutationId};
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
  return fetch({query, variables});
}




exports.AddUser = AddUser;
exports.RemoveUser = RemoveUser;
exports.User = User;
exports.AddTodo = AddTodo; 
exports.ClearCompletedTodos = ClearCompletedTodos; 
exports.ChangeTodoStatus = ChangeTodoStatus; 
exports.MarkAllTodos = MarkAllTodos; 
// exports.RemoveTodo = RemoveTodo; 
// exports.RenameTodo = RenameTodo; 
