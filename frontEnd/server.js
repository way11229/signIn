"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    Object.defineProperty(o, k2, { enumerable: true, get: function() { return m[k]; } });
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
var http_1 = require("http");
var fs = __importStar(require("fs"));
var path = __importStar(require("path"));
var server = http_1.createServer(function (request, response) {
    if (request.url === '/') {
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
    else if (request.url === '/test') {
        response.write(JSON.stringify({ test: '123' }));
        response.end();
    }
    else {
        response.writeHead(404);
        response.write('Whoops! File not found!');
        response.end();
    }
});
server.listen(5000);
