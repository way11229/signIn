
const path = require('path');
const Dotenv = require('dotenv-webpack');
const HtmlWebpackPlugin =require('html-webpack-plugin');

module.exports = {
    entry: ['./src/ts/index.ts'],
    plugins: [
        new Dotenv(),
        new HtmlWebpackPlugin({
            template: './src/html/index.html'
        })
    ],
    module: {
        rules: [
            {
                test: /\.scss$/,
                use: [ 'style-loader', 'css-loader', 'sass-loader' ]
            },
            {
                test: /\.tsx?$/,
                use: 'ts-loader',
                exclude: /node_modules/,
            }
        ]
    },
    resolve: {
        extensions: ['.tsx', '.ts', '.js'],
    },
    output: {
        path: path.join(__dirname, 'dist'),
        filename: 'index.bundle.js'
    },
}