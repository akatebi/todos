const { fetch } = require("./fetch");

const User = ({email}) => {
    const variables = { email }; 
    const query = `query TodoAppQuery($email: String!) {
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
            todos(first: 100) {
              edges {
                cursor
                node {
                  id
                  text
                  complete
                }
              }
            }
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








exports.User = User;
exports.AddTodo = AddTodo; 
exports.MarkAllTodos = MarkAllTodos; 
exports.ClearCompletedTodos = ClearCompletedTodos; 
