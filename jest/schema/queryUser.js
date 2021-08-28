const { fetch } = require("../fetch");

export const queryUser = ({ email, first = 100 }) => {
  const variables = { email };
  const query = `query User($email: String!, $status: String, $first: Int, $last: String) {
        user(email: $email) {
          id
          email
          completedCount
          totalCount
          todos(status: $status, first: $first, last: $last) {
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

