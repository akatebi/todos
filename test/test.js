const {User, AddTodo} = require("./schema");

describe('Testing Todo GraphQL', () => {
    let userId;
    beforeAll(async() => {
        const email = "test@test.com";
        const resp = await User({email})
        console.log("resp", JSON.stringify(resp, 0, 2));
        userId = resp.data.user.id;
    });
    test("AddTodo", async() => {
      const clientMutationId = "777";
      const text = "Get a job";
      const resp = await AddTodo({text, userId, clientMutationId});
      console.log("resp", JSON.stringify(resp, 0, 2));
    });
  });