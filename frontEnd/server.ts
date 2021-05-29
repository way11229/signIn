import { createServer, IncomingMessage, ServerResponse } from 'http';
import * as fs from 'fs'
import * as path from 'path'

const server = createServer((request: IncomingMessage, response: ServerResponse) => {
    if (request.url === '/') {
        fs.readFile(path.join(__dirname, '/dist/index.html'), 'utf8', (error, data) => {
            if (error) {
                response.writeHead(404);
                response.write('Whoops! File not found!');
            } else {
                response.write(data);
            }
            response.end();
        });
    } else if (/(.*)\.(css|js)/.test(request.url!)) {
        fs.readFile(path.join(__dirname, '/dist/' + request.url!), 'utf8', (error, data) => {
            if (error) {
                response.writeHead(404);
                response.write('Whoops! File not found!');
            } else {
                response.write(data);
            }
            response.end();
        });
    } else {
        response.writeHead(404);
        response.write('Whoops! File not found!');
        response.end();
    }
});

server.listen(80);