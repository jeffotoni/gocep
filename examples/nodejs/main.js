const https = require('http');

https.get('http://localhost:8080/v1/cep/08226021', res => {
    let data = [];
    console.log('Status Code:', res.statusCode);

    res.on('data', chunk => {
        data.push(chunk);
    });

    res.on('end', () => {
        console.log('Response: ');
        const endereco = JSON.parse(Buffer.concat(data).toString());
        console.log(endereco);
    });
}).on('error', err => {
    console.log('Error: ', err.message);
});