/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest } from "relay-runtime";

export type ClearCompletedTodosInput = {
    userId: string;
    clientMutationId?: string | null;
};
export type useClearCompleteTodosMutationVariables = {
    input: ClearCompletedTodosInput;
};
export type useClearCompleteTodosMutationResponse = {
    readonly clearCompletedTodos: {
        readonly deletedTodoIds: ReadonlyArray<string> | null;
        readonly user: {
            readonly completedCount: number;
            readonly totalCount: number;
        };
    } | null;
};
export type useClearCompleteTodosMutation = {
    readonly response: useClearCompleteTodosMutationResponse;
    readonly variables: useClearCompleteTodosMutationVariables;
};



/*
mutation useClearCompleteTodosMutation(
  $input: ClearCompletedTodosInput!
) {
  clearCompletedTodos(input: $input) {
    deletedTodoIds
    user {
      completedCount
      totalCount
      id
    }
  }
}
*/

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "defaultValue": null,
    "kind": "LocalArgument",
    "name": "input"
  }
],
v1 = [
  {
    "kind": "Variable",
    "name": "input",
    "variableName": "input"
  }
],
v2 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "deletedTodoIds",
  "storageKey": null
},
v3 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "completedCount",
  "storageKey": null
},
v4 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "totalCount",
  "storageKey": null
};
return {
  "fragment": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Fragment",
    "metadata": null,
    "name": "useClearCompleteTodosMutation",
    "selections": [
      {
        "alias": null,
        "args": (v1/*: any*/),
        "concreteType": "ClearCompletedTodosPayload",
        "kind": "LinkedField",
        "name": "clearCompletedTodos",
        "plural": false,
        "selections": [
          (v2/*: any*/),
          {
            "alias": null,
            "args": null,
            "concreteType": "User",
            "kind": "LinkedField",
            "name": "user",
            "plural": false,
            "selections": [
              (v3/*: any*/),
              (v4/*: any*/)
            ],
            "storageKey": null
          }
        ],
        "storageKey": null
      }
    ],
    "type": "Mutation",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "useClearCompleteTodosMutation",
    "selections": [
      {
        "alias": null,
        "args": (v1/*: any*/),
        "concreteType": "ClearCompletedTodosPayload",
        "kind": "LinkedField",
        "name": "clearCompletedTodos",
        "plural": false,
        "selections": [
          (v2/*: any*/),
          {
            "alias": null,
            "args": null,
            "concreteType": "User",
            "kind": "LinkedField",
            "name": "user",
            "plural": false,
            "selections": [
              (v3/*: any*/),
              (v4/*: any*/),
              {
                "alias": null,
                "args": null,
                "kind": "ScalarField",
                "name": "id",
                "storageKey": null
              }
            ],
            "storageKey": null
          }
        ],
        "storageKey": null
      }
    ]
  },
  "params": {
    "cacheID": "f9d07553b9b1e6139c3c8fa90c095ec3",
    "id": null,
    "metadata": {},
    "name": "useClearCompleteTodosMutation",
    "operationKind": "mutation",
    "text": "mutation useClearCompleteTodosMutation(\n  $input: ClearCompletedTodosInput!\n) {\n  clearCompletedTodos(input: $input) {\n    deletedTodoIds\n    user {\n      completedCount\n      totalCount\n      id\n    }\n  }\n}\n"
  }
};
})();
(node as any).hash = 'c281eb9a4ae7be8c62330b4286a19aa3';
export default node;
