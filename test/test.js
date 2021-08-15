const fetch = require("node-fetch")

// https://docs.google.com/document/d/1GJ54IygQ0q4pX7TzOGSb0-56iX4PYo6MimBZkpy4NLc/edit?usp=sharing

 const query = async (text, variables) => {
    const resp = await fetch('https://localhost:8080/query', {
            method: 'post',
            body:    JSON.stringify({test, variables}),
            headers: { 'Content-Type': 'application/json' },
    })
    .then(resp => resp.json());
    console.log(resp);
    return resp;
}

describe('Testing Todo GraphQL', () => {
    beforeAll(async() => {
        text = `query Todo($email: String!) {
                user(email: $email) {
                    id
                }`
        variables = {email: "test@test.com"}
        resp = await query(text, variables);
        cconsole.log("resp", resp);
    });
    test('test 1', () => {
      expect(true).toEqual(true);
    });
  });