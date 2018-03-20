const ExtractTextPlugin = require('extract-text-webpack-plugin');

module.exports = {
    context: `${__dirname}/client`,
    entry: {
        scripts: ['regenerator-runtime/runtime', './app.jsx'],
        css: './app.scss',
    },
    output: {
        path: `${__dirname}/public`,
        filename: 'scripts.js',
    },
    externals: {
    },
    devtool: '#inline-source-map',
    module: {
        loaders: [
            {
                test: /\.jsx?$/,
                exclude: /(node_modules)/,
                loader: 'babel-loader',
            }, {
                test: /\.scss/,
                loader: ExtractTextPlugin.extract(
                    'style-loader',
                    'css-loader!postcss-loader!sass-loader'
                ),
            }, {
                test: /\.elm$/,
                exclude: [/elm-stuff/, /node_modules/],
                loader: 'elm-webpack?verbose=true&warn=true',
            },
            {
                test: /\.jpg|.png/,
                loader: 'file-loader',
            },
        ],
        noParse: /\.elm$/,
    },
    resolve: {
        extensions: ['', '.js', '.jsx'],
    },
    plugins: [
        new ExtractTextPlugin('styles.css', { allChunks: true }),
    ],
};
