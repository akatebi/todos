const nodeFetch = require('node-fetch');


 const fetch = async ({ query, variables }) => {
    const resp = await nodeFetch('http://localhost:8080/query', {
            method: 'post',
            body: JSON.stringify({query, variables}),
            headers: { 'Content-Type': 'application/json' },
    })
    .then(resp => resp.json());
//     console.log(resp);
    return resp;
}

exports.fetch = fetch;