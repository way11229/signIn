"use strict";
exports.__esModule = true;
var http_1 = require("http");
var fs = require("fs");
var path = require("path");
var server = http_1.createServer(function (request, response) {
    var _a;
    if ((request.url === '/') || ((_a = request.url) === null || _a === void 0 ? void 0 : _a.includes('/?'))) {
        fs.readFile(path.join(__dirname, '/dist/index.html'), 'utf8', function (error, data) {
            if (error) {
                response.writeHead(404);
                response.write('Whoops! File not found!');
            }
            else {
                response.write(data);
            }
            response.end();
        });
    }
    else if (/(.*)\.(css|js)/.test(request.url)) {
        fs.readFile(path.join(__dirname, '/dist/' + request.url), 'utf8', function (error, data) {
            if (error) {
                response.writeHead(404);
                response.write('Whoops! File not found!');
            }
            else {
                response.write(data);
            }
            response.end();
        });
    }
    else {
        response.writeHead(404);
        response.write('Whoops! File not found!');
        response.end();
    }
});
server.listen(80);
