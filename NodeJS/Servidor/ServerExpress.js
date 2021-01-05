const express = require('express')
const app = express()
const MongoClient = require('mongodb').MongoClient;
const redis = require("redis");
module.exports.MongoClient = MongoClient;
module.exports.redis = redis;
const http = require('http');
const path = require('path');
const socketIO = require('socket.io');
const port = 3000;
let server = http.createServer(app);
const publicPath = path.resolve(__dirname, '../public');
app.use(express.static(publicPath));
module.exports.io = socketIO(server);
require('./socket');
server.listen(port, (err) => {
    if (err) throw new Error(err);
    console.log(`Running server on port ${ port }`);
});