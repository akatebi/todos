import { TestWatcher } from "@jest/core";
import query from "./query" ;

import {user, addTodo} from "./schema"

describe('Testing Todo GraphQL', () => {
    let clientMutationId = 1;
    let userId;
    beforeAll(async() => {
        const resp = await query(user({email: "test@test.com"}))
        console.log("resp", JSON.stringify(resp, 0, 2));
        userId = `"${resp.data.user.id}"`;
    });
    test("AddTodo", async() {
      const resp = await query(addTodo({userId, text: "Get a job", clientMutationId++}));
      console.log("resp", JSON.stringify(resp, 0, 2));
    });
  });