/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ReaderFragment } from "relay-runtime";

import { FragmentRefs } from "relay-runtime";
export type Todo_user = {
    readonly id: string;
    readonly email: string;
    readonly totalCount: number;
    readonly completedCount: number;
    readonly " $refType": "Todo_user";
};
export type Todo_user$data = Todo_user;
export type Todo_user$key = {
    readonly " $data"?: Todo_user$data;
    readonly " $fragmentRefs": FragmentRefs<"Todo_user">;
};



const node: ReaderFragment = {
  "argumentDefinitions": [],
  "kind": "Fragment",
  "metadata": null,
  "name": "Todo_user",
  "selections": [
    {
      "alias": null,
      "args": null,
      "kind": "ScalarField",
      "name": "id",
      "storageKey": null
    },
    {
      "alias": null,
      "args": null,
      "kind": "ScalarField",
      "name": "email",
      "storageKey": null
    },
    {
      "alias": null,
      "args": null,
      "kind": "ScalarField",
      "name": "totalCount",
      "storageKey": null
    },
    {
      "alias": null,
      "args": null,
      "kind": "ScalarField",
      "name": "completedCount",
      "storageKey": null
    }
  ],
  "type": "User",
  "abstractKey": null
};
(node as any).hash = '35184cb71937e375f968db4d0b6a3ce3';
export default node;
