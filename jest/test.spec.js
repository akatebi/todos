import { removeUser } from "./schema/removeUser";
import { addUser } from "./schema/addUser";
import { addTodo } from "./schema/addTodo";
import { renameTodo } from "./schema/renameTodo";
import { removeTodo } from "./schema/removeTodo";
import { queryUser } from "./schema/queryUser";
import { markAllTodos } from "./schema/markAllTodos";
import { clearCompletedTodos } from "./schema/clearCompletedTodos";
import { changeTodoStatus } from "./schema/changeTodoStatus";

describe("Todos GraphQL", () => {
    test("Query & Mutation", async() => {
        let clientMutationId, email, resp;
        clientMutationId = "1";
        email = "user@test.com";
        resp = await removeUser({ email, clientMutationId });
        expect("removeUser").toMatchSnapshot();
        expect(resp).toMatchSnapshot();
        
        clientMutationId = "2";
        resp = await addUser({ email, clientMutationId });
        expect("addUser").toMatchSnapshot();
        expect(resp).toMatchSnapshot();
        const userId = resp.data.addUser.id;

        const todoIds = [];
        clientMutationId = "3";
        for (let i = 0; i < 3; i++) {
            const text = `Get A Customer ${i + 1}`;
            const cid = `${clientMutationId} ${i}`;
            resp = await addTodo({ text, userId, clientMutationId: cid });
            expect(`addTodo ${i+1}`).toMatchSnapshot();
            expect(resp).toMatchSnapshot();
            todoIds.push(resp.data.addTodo.todoEdge.node.id);
        }
        
        resp = await queryUser({ email });
        expect("queryUser").toMatchSnapshot();
        expect(resp).toMatchSnapshot();

        clientMutationId = "4";
        let complete = true;
        resp = await markAllTodos({ complete, userId, clientMutationId });
        expect("markAllTodos").toMatchSnapshot();
        expect(resp).toMatchSnapshot();
        
        complete = false;
        clientMutationId = "5";
        for (let i = 0; i < 3; i++) {
            const id = todoIds[i];
            const cid = `${clientMutationId} ${i}`;
            const resp = await changeTodoStatus({
                complete, id, userId, clientMutationId: cid
            });
            expect(`changeTodoStatus ${i}`).toMatchSnapshot();
            expect(resp).toMatchSnapshot();
        }
        
        complete = true;
        clientMutationId = "6";
        for (let i = 0; i < 3; i++) {
            const id = todoIds[i];
            const cid = `${clientMutationId} ${i}`;
            const resp = await changeTodoStatus({ complete, id, userId, clientMutationId: cid });
            expect(resp).toMatchSnapshot();
        }

        clientMutationId = "7";
        resp = await clearCompletedTodos({ userId, clientMutationId });
        expect(resp).toMatchSnapshot();

        clientMutationId = "8";
        const text = "Get A Job"
        await addTodo({ text, userId, clientMutationId })
        
        clientMutationId = "9";
        for (let i = 0; i < 3; i++) {
            const id = todoIds[i];
            const text = `renamed ${i}`;
            const cid = `${clientMutationId} ${i}`;
            const resp = await renameTodo({ id, text, clientMutationId: cid });
            expect(resp).toMatchSnapshot();
        }
        

        clientMutationId = "10";
        for (let i = 0; i < 3; i++) {
            const id = todoIds[i];
            const cid = `${clientMutationId}-${i}`;
            const resp = await removeTodo({ id, clientMutationId: cid });
            expect(resp).toMatchSnapshot();
        }
        
    });
});
