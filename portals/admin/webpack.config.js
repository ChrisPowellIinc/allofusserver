const path = require("path");
const HtmlWebpack = require("html-webpack-plugin");

module.exports = {
  mode: "development",
  entry: "src/index.js",
  output: {
    path: path.resolve(__dirname, "./public"),
    filename: "assets/js/app.js",
  },
  resolve: {
    modules: [path.resolve(__dirname, './src'), path.resolve(__dirname, './node_modules')],
    extensions: ['.js', '.jsx', '.json'],
  },
  module: {
    rules: [{
      test: /\.js$/,
      exclude: /node_modules/,
      use: [{
        loader: "babel-loader",
      },{
        loader: "eslint-loader",
        options: {
          failOnError: false,
          fix: true
        }
      }],
    }, {
      test: /\.scss$/,
      use: [{
        loader: "style-loader"
      }, {
        loader: "css-loader",
        options: {
          sourceMap: true
        }
      }, {
        loader: "sass-loader",
        options: {
          sourceMap: true
        }
      }]
    }, {
      test: /\.css$/,
      use: [{
        loader: "style-loader"
      }, {
        loader: "css-loader",
        options: {
          sourceMap: true
        }
      }]
    }]
  },
  plugins: [
    new HtmlWebpack({
      template: "./src/index.html",
      filename: "./index.html"
    })
  ],
  serve: {
    content: "./public",
    clipboard: false,
    port: 3000,
    host: "0.0.0.0"
  }
};
