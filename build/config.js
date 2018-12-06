'use strict'

const path = require('path')
require('dotenv').config()

const {HOST, PORT, API_HOST, API_PORT} = process.env

module.exports = {
	dev: {

		// Paths
		assetsSubDirectory: 'assets',
		assetsPublicPath: '/',
		proxyTable: {
			'/v1': {
				logLevel: 'debug',
				target: `http://${API_HOST}:${API_PORT}`,
				changeOrigin: true,
			}
		},

		// Various Dev Server settings
		host: HOST || 'localhost', // can be overwritten by process.env.HOST
		port: PORT || 8080, // can be overwritten by process.env.PORT, if port is in use, a free one will be determined
		apiHost: API_HOST || 'localhost', // can be overwritten by process.env.API_HOST
		apiPort: API_PORT || 3500, // can be overwritten by process.env.API_PORT
		autoOpenBrowser: false,
		errorOverlay: true,
		notifyOnErrors: true,
		poll: false, // https://webpack.js.org/configuration/dev-server/#devserver-watchoptions-

		// Use Eslint Loader?
		// If true, your code will be linted during bundling and
		// linting errors and warnings will be shown in the console.
		useEslint: true,
		// If true, eslint errors and warnings will also be shown in the error overlay
		// in the browser.
		showEslintErrorsInOverlay: false,

		/**
		 * Source Maps
		 */

		// https://webpack.js.org/configuration/devtool/#development
		devtool: 'cheap-module-eval-source-map',

		// If you have problems debugging vue-files in devtools,
		// set this to false - it *may* help
		// https://vue-loader.vuejs.org/en/options.html#cachebusting
		cacheBusting: true,

		cssSourceMap: true
	},

	build: {
		// Template for index.html
		index: path.resolve(__dirname, '../public/index.html'),

		// Paths
		assetsRoot: path.resolve(__dirname, '../public'),
		assetsSubDirectory: 'assets',
		assetsPublicPath: '/',

		/**
		 * Source Maps
		 */

		productionSourceMap: false,
		// https://webpack.js.org/configuration/devtool/#production
		devtool: '#source-map',

		// Gzip off by default as many popular static hosts such as
		// Surge or Netlify already gzip all static assets for you.
		// Before setting to `true`, make sure to:
		// npm install --save-dev compression-webpack-plugin
		productionGzip: false,
		productionGzipExtensions: ['js', 'css'],

		// Run the build command with an extra argument to
		// View the bundle analyzer report after build finishes:
		// `npm run build --report`
		// Set to `true` or `false` to always turn it on or off
		bundleAnalyzerReport: process.env.npm_config_report
	}
}
