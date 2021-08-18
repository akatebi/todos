const {User, AddTodo, MarkAllTodos, ClearCompletedTodos} = require("./schema");

describe('Testing Todo GraphQL', () => {
    let userId;
    beforeAll(async() => {
        const email = "test@test.com";
        const resp = await User({email})
        // console.log("resp", JSON.stringify(resp, 0, 2));
        expect(resp).toMatchSnapshot();
        userId = resp.data.user.id;
    });
    test("AddTodo", async() => {
      const clientMutationId = "1";
      const text = "Get a customer";
      const resp = await AddTodo({text, userId, clientMutationId});
      // console.log("resp", JSON.stringify(resp, 0, 2));
      expect(resp).toMatchSnapshot();
    });
    test("MarkAllTodos", async() => {
      const clientMutationId = "2";
      const complete = true;
      const resp = await MarkAllTodos({complete, userId, clientMutationId});
      // console.log("resp", JSON.stringify(resp, 0, 2));
      expect(resp).toMatchSnapshot();
    });
    test("ClearCompletedTodos", async() => {
      const clientMutationId = "3";
      const resp = await ClearCompletedTodos({userId, clientMutationId});
      // console.log("resp", JSON.stringify(resp, 0, 2));
      expect(resp).toMatchSnapshot();
    });
  });