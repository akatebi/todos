const nodeFetch = require('node-fetch');


 export const fetch = async ({ query, variables }) => {
    const resp = await nodeFetch('http://localhost:8080/query', {
            method: 'post',
            body: JSON.stringify({query, variables}),
            headers: { 'Content-Type': 'application/json' },
    })
    .then(resp => resp.json());
    return resp;
}