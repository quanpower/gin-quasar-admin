/*
 * This file runs in a Node context (it's NOT transpiled by Babel), so use only
 * the ES6 features that are supported by your Node version. https://node.green/
 */

// Configuration for your app
// https://v2.quasar.dev/quasar-cli/quasar-conf-js

const { configure } = require('quasar/wrappers')
const path = require('path')

module.exports = configure(function (ctx) {
    return {
        // https://v2.quasar.dev/quasar-cli/supporting-ts
        supportTS: false,

        // https://v2.quasar.dev/quasar-cli/prefetch-feature
        // preFetch: true,

        // app boot file (/src/boot)
        // --> boot files are part of "main.js"
        // https://v2.quasar.dev/quasar-cli/boot-files
        boot: [
            'set-default',
            'theme',
            'version',
            'i18n',
            'axios',
            'permission',
            'echarts',
            'bus',
        ],

        // https://v2.quasar.dev/quasar-cli/quasar-conf-js#Property%3A-css
        css: [
            'app.scss'
        ],

        // https://github.com/quasarframework/quasar/tree/dev/extras
        extras: [
            // 'ionicons-v4',
            // 'mdi-v5',
            'fontawesome-v5',
            // 'eva-icons',
            // 'themify',
            // 'line-awesome',
            // 'roboto-font-latin-ext', // this or either 'roboto-font', NEVER both!

            'roboto-font', // optional, you are not bound to it
            'material-icons', // optional, you are not bound to it
        ],

        // Full list of options: https://v2.quasar.dev/quasar-cli/quasar-conf-js#Property%3A-build
        build: {
            vueRouterMode: 'hash', // available values: 'hash', 'history'
            env: {
                API: ctx.dev
                    // 测试代理地址
                    ? "http://127.0.0.1:8888/"
                    // 正式代理地址
                    : "http://192.168.35.125:8888/"
            },

            // transpile: false,

            // Add dependencies for transpiling with Babel (Array of string/regex)
            // (from node_modules, which are by default not transpiled).
            // Applies only if "transpile" is set to true.
            // transpileDependencies: [],

            // rtl: true, // https://v2.quasar.dev/options/rtl-support
            // preloadChunks: true,
            // showProgress: false,
            // gzip: true,
            // analyze: true,

            // Options below are automatically set depending on the env, set them if you want to override
            // extractCSS: false,

            // https://v2.quasar.dev/quasar-cli/handling-webpack
            // "chain" is a webpack-chain object https://github.com/neutrinojs/webpack-chain
            chainWebpack: chain => {
                chain.module
                    .rule('i18n-resource')
                    .test(/\.(json5?|ya?ml)$/)
                    .include.add(path.resolve(__dirname, './src/i18n'))
                    .end()
                    .type('javascript/auto')
                    .use('i18n-resource')
                    .loader('@intlify/vue-i18n-loader')
                chain.module
                    .rule('i18n')
                    .resourceQuery(/blockType=i18n/)
                    .type('javascript/auto')
                    .use('i18n')
                    .loader('@intlify/vue-i18n-loader')
            }
        },

        // Full list of options: https://v2.quasar.dev/quasar-cli/quasar-conf-js#Property%3A-devServer
        // 改为后端跨域，devServer移除
        // devServer: {
        //     https: false,
        //     port: 8080,
        //     open: true, // opens browser window automatically
        //     proxy: {
        //         '/gqa-api/': {
        //             // 测试后台地址
        //             target: "http://127.0.0.1:8888/",
        //             changeOrigin: true,
        //             pathRewrite: {
        //                 '^/gqa-api/': ''
        //             }
        //         }
        //     }
        // },

        // https://v2.quasar.dev/quasar-cli/quasar-conf-js#Property%3A-framework
        framework: {
            config: {},

            iconSet: 'material-icons', // Quasar icon set
            lang: 'zh-CN', // Quasar language pack

            // For special cases outside of where the auto-import strategy can have an impact
            // (like functional components as one of the examples),
            // you can manually specify Quasar components/directives to be available everywhere:
            //
            // components: [],
            // directives: [],

            // Quasar plugins
            plugins: [
                'Notify',
                'Cookies',
                'SessionStorage',
                'LocalStorage',
                'Loading',
                'LoadingBar',
                'Dialog',
                'AppFullscreen',
            ]
        },

        // animations: 'all', // --- includes all animations
        // https://v2.quasar.dev/options/animations
        animations: [],

        // https://v2.quasar.dev/quasar-cli/developing-ssr/configuring-ssr
        ssr: {
            pwa: false,

            // manualStoreHydration: true,
            // manualPostHydrationTrigger: true,

            prodPort: 3000, // The default port that the production server should use
            // (gets superseded if process.env.PORT is specified at runtime)

            maxAge: 1000 * 60 * 60 * 24 * 30,
            // Tell browser when a file from the server should expire from cache (in ms)

            chainWebpackWebserver( /* chain */) {
                //
            },

            middlewares: [
                ctx.prod ? 'compression' : '',
                'render' // keep this as last one
            ]
        },

        // https://v2.quasar.dev/quasar-cli/developing-pwa/configuring-pwa
        pwa: {
            workboxPluginMode: 'GenerateSW', // 'GenerateSW' or 'InjectManifest'
            workboxOptions: {}, // only for GenerateSW

            // for the custom service worker ONLY (/src-pwa/custom-service-worker.[js|ts])
            // if using workbox in InjectManifest mode
            chainWebpackCustomSW( /* chain */) {
                //
            },

            manifest: {
                name: `gin-quasar-admin`,
                short_name: `gin-quasar-admin`,
                description: `gin-quasar-admin`,
                display: 'standalone',
                orientation: 'portrait',
                background_color: '#ffffff',
                theme_color: '#027be3',
                icons: [{
                    src: 'icons/icon-128x128.png',
                    sizes: '128x128',
                    type: 'image/png'
                },
                {
                    src: 'icons/icon-192x192.png',
                    sizes: '192x192',
                    type: 'image/png'
                },
                {
                    src: 'icons/icon-256x256.png',
                    sizes: '256x256',
                    type: 'image/png'
                },
                {
                    src: 'icons/icon-384x384.png',
                    sizes: '384x384',
                    type: 'image/png'
                },
                {
                    src: 'icons/icon-512x512.png',
                    sizes: '512x512',
                    type: 'image/png'
                }
                ]
            }
        },

        // Full list of options: https://v2.quasar.dev/quasar-cli/developing-cordova-apps/configuring-cordova
        cordova: {
            // noIosLegacyBuildFlag: true, // uncomment only if you know what you are doing
        },

        // Full list of options: https://v2.quasar.dev/quasar-cli/developing-capacitor-apps/configuring-capacitor
        capacitor: {
            hideSplashscreen: true
        },

        // Full list of options: https://v2.quasar.dev/quasar-cli/developing-electron-apps/configuring-electron
        electron: {
            bundler: 'packager', // 'packager' or 'builder'

            packager: {
                // https://github.com/electron-userland/electron-packager/blob/master/docs/api.md#options

                // OS X / Mac App Store
                // appBundleId: '',
                // appCategoryType: '',
                // osxSign: '',
                // protocol: 'myapp://path',

                // Windows only
                // win32metadata: { ... }
            },

            builder: {
                // https://www.electron.build/configuration/configuration

                appId: 'gin-quasar-admin'
            },

            // "chain" is a webpack-chain object https://github.com/neutrinojs/webpack-chain
            chainWebpackMain( /* chain */) {
                // do something with the Electron main process Webpack cfg
                // extendWebpackMain also available besides this chainWebpackMain
            },

            // "chain" is a webpack-chain object https://github.com/neutrinojs/webpack-chain
            chainWebpackPreload( /* chain */) {
                // do something with the Electron main process Webpack cfg
                // extendWebpackPreload also available besides this chainWebpackPreload
            },
        }
    }
});