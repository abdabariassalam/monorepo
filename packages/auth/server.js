const express = require('express');
const http = require('http');
const authHandler = require('./routes/api/auth');
const app = express();
const fs = require("fs");
if (!fs.existsSync('./.env')) {
    require('dotenv').config({ path: './../../.env' });
} else {require('dotenv').config();}

const server = http.createServer(app);

// body parser jika yang dikirimkan data json
app.use(express.json());

// route auth
app.use('/auth', authHandler);

// run server
server.listen(process.env.PORT, () => {
    console.log('server running on port', process.env.PORT);
});

module.exports=app;