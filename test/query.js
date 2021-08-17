const fetch = require('node-fetch');


 export default async ({ query, variables }) => {
    const resp = await fetch('http://localhost:8080/query', {
            method: 'post',
            body:    JSON.stringify({query, variables}),
            headers: { 'Content-Type': 'application/json' },
    })
    .then(resp => resp.json());
    console.log(resp);
    return resp;
}