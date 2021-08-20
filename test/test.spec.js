import {
  AddUserTest,
  RemoveUserTest,
  AddTodoTest,
  QueryUserTest,
  MarkAllTodosTest,
  ClearCompletedTodosTest,
  ChangeTodoStatusTest
} from "./schema";

describe("Remove User", RemoveUserTest("user1"));
describe("Add User", AddUserTest("user1"));
describe("Add Todo", AddTodoTest("Get A Customer"));
describe("Query User", QueryUserTest("user1"));
describe("Mark All Todos", MarkAllTodosTest());
describe("Change Todo Status => false", ChangeTodoStatusTest(false));
describe("Change Todo Status => true", ChangeTodoStatusTest(true));
describe("Clear Completed Todos", ClearCompletedTodosTest());
