const {User, AddTodo, ChangeTodoStatus, MarkAllTodos, ClearCompletedTodos} = require("./schema");

describe('Testing Todo GraphQL', () => {
    let userId;
    let todoIds = [];
    beforeAll(async() => {
        const email = "test@test.com";
        const resp = await User({email})
        // console.log("resp", JSON.stringify(resp, 0, 2));
        expect(resp).toMatchSnapshot();
        userId = resp.data.user.id;
    });
    test("AddTodo", async() => {
      const clientMutationId = "1";
      for(let i = 0; i < 3; i++) {
        const text = `Get a customer ${i+1}`;
        const resp = await AddTodo({text, userId, clientMutationId});
        // console.log("resp", JSON.stringify(resp, 0, 2));
        expect(resp).toMatchSnapshot();
        todoIds.push(resp.data.addTodo.todoEdge.node.id);
      }
      console.log("#### todoIds", todoIds);
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
    test("ChangeTodoStatus", async() => {
      const clientMutationId = "4";
      const complete = true;
      for(let i=0; i<3; i++) { 
        const id = todoIds[i];
        const resp = await ChangeTodoStatus({complete, id, userId, clientMutationId});
        // console.log("resp", JSON.stringify(resp, 0, 2));
        expect(resp).toMatchSnapshot();
      }
    })
  });