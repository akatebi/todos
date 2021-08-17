const { fetch } = require("./fetch");

const User = ({email}) => {
    const query = `query TodoAppQuery($email: String!) {
        user(email: $email) {
          id
          email
          totalCount
          ...TodoListFooter_user
          ...TodoList_user
        }
      }
      fragment TodoListFooter_user on User {
        id
        email
        completedCount
        todos(first: 100) {
          edges {
            node {
              id
              complete
              __typename
            }
            cursor
          }
          pageInfo {
            endCursor
            hasNextPage
          }
        }
        totalCount
      }
      fragment TodoList_user on User {
        todos(first: 100) {
          edges {
            node {
              id
              complete
              ...Todo_todo
              __typename
            }
            cursor
          }
          pageInfo {
            endCursor
            hasNextPage
          }
        }
        id
        email
        totalCount
        completedCount
        ...Todo_user
      }
      fragment Todo_todo on Todo {
        complete
        id
        text
      }
      fragment Todo_user on User {
        id
        totalCount
        completedCount
      }`;
      return fetch({query, variables: { email } });
}

const AddTodo = ({text, userId, clientMutationId}) => {
const query = `mutation addTodo($text: String!, $userId: ID!, $clientMutationId: String) {
        addTodo(
          input: { text: $text, userId: $userId, clientMutationId: $clientMutationId }
        ) {
          clientMutationId
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
        }
      }`;
    return fetch({query, variables: {text, userId, clientMutationId}});
}
exports.User = User;
exports.AddTodo = AddTodo; 
