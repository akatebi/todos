const fetch = require('node-fetch');

// https://docs.google.com/document/d/1GJ54IygQ0q4pX7TzOGSb0-56iX4PYo6MimBZkpy4NLc/edit?usp=sharing

 const query = async (query, variables) => {
    const resp = await fetch('http://localhost:8080/query', {
            method: 'post',
            body:    JSON.stringify({query, variables}),
            headers: { 'Content-Type': 'application/json' },
    })
    .then(resp => resp.json());
    console.log(resp);
    return resp;
}

describe('Testing Todo GraphQL', () => {
    beforeAll(async() => {
        const text = `
            query Todo($email: String!) {
                user(email: $email) {
                    id
                    email
                    totalCount
                    todos {
                        edges {
                            node {
                                id
                                text
                            }
                        }
                    }
                }
            }
        `;
        const variables = {"email": "test@test.com"}
        const resp = await query(text, variables);
        console.log("resp", JSON.stringify(resp, 0, 2));
    });
    test('test 1', () => {
      expect(true).toEqual(true);
    });
  });