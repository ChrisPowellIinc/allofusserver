'use strict'
const path = require('path')
const DotEnv = require('dotenv-webpack')
const utils = require('./utils')
const config = require('./config')

function resolve(dir) {
	return path.join(__dirname, '..', dir)
}

const createLintingRule = () => ({
	test: /\.jsx?$/,
	loader: 'eslint-loader',
	enforce: 'pre',
	include: [resolve('src')],
	options: {
		emitWarning: !config.dev.showEslintErrorsInOverlay,
		failOnWarning: true,
		fix: true
	}
})

module.exports = {
	context: path.resolve(__dirname, '../'),
	entry: {
		app: './src/index.js',
	},
	output: {
		path: config.build.assetsRoot,
		filename: '[name].bundle.js',
		chunkFilename: '[name].[chunkhash].js',
		publicPath: process.env.NODE_ENV === 'production'
			? config.build.assetsPublicPath
			: config.dev.assetsPublicPath
	},
	resolve: {
		modules: ['node_modules', 'src'],
		extensions: ['.js', '.json', '.jsx']
	},
	module: {
		rules: [
			...(config.dev.useEslint ? [createLintingRule()] : []),
			{
				test: /\.js$/,
				loader: 'babel-loader',
				include: [resolve('src'), resolve('node_modules/webpack-dev-server/client')]
			},
			{
				test: /\.(png|jpe?g|gif|svg)(\?.*)?$/,
				loader: 'url-loader',
				options: {
					limit: 10000,
					name: utils.assetsPath('img/[name].[hash:7].[ext]')
				}
			},
			{
				test: /\.(mp4|webm|ogg|mp3|wav|flac|aac)(\?.*)?$/,
				loader: 'url-loader',
				options: {
					limit: 10000,
					name: utils.assetsPath('media/[name].[hash:7].[ext]')
				}
			},
			{
				test: /\.(woff2?|eot|ttf|otf)(\?.*)?$/,
				loader: 'url-loader',
				options: {
					limit: 10000,
					name: utils.assetsPath('fonts/[name].[hash:7].[ext]')
				}
			}
		]
	},
	plugins: [
		new DotEnv()
	],
	node: {
		setImmediate: false,
		dgram: 'empty',
		fs: 'empty',
		net: 'empty',
		tls: 'empty',
		child_process: 'empty'
	}
}
