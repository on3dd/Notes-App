module.exports = {
  /*
  ** Headers of the page
  */
  head: {
    title: 'Конспекты',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: 'Client for application' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
      { rel: 'stylesheet', href: 'https://cdn.jsdelivr.net/npm/@mdi/font@4.x/css/materialdesignicons.min.css' }

    ]

  },
  /*
  ** Customize the progress bar color
  */
  loading: {
    color: '#2980b9',
    height: '3px'
  },
  modules: [
    '@nuxtjs/axios',
    '@nuxtjs/vuetify',
  ],
  /*
  ** Build configuration
  */
  // buildModules: [
  //
  //
  // ],

  axios: {
    // proxyHeaders: false
  },
  vuetify: {
    /* module options */
  },
  build: {
    /*
    ** Run ESLint on save
    */
    extend (config, { isDev, isClient }) {
      if (isDev && isClient) {
        config.module.rules.push({
          enforce: 'pre',
          test: /\.(js|vue)$/,
          loader: 'eslint-loader',
          exclude: /(node_modules)/
        })
      }
    },
  }
}

