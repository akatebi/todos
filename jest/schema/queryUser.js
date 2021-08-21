const { fetch } = require("../fetch");

const queryUser = ({ email }) => {
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

const QueryUserTest = (user) => () => {
    it("", async () => {
      const email = `${user}@test.com`;
      const resp = await queryUser({ email });
      expect(resp).toMatchSnapshot();
    });
  };

export default QueryUserTest;