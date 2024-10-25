// nuxt.config.ts
export default defineNuxtConfig({
  runtimeConfig: {
    public: {
      apiUrl: process.env.API_URL || 'http://localhost:8080',
    },
  },
  plugins: ['~/plugins/fontawesome.js'],
  compatibilityDate: '2024-09-26',

  modules: ['nuxt-simple-sitemap'],

  sitemap: {
    hostname: process.env.SITE_URL || 'http://localhost:3000', // Адрес сайта
    trailingSlash: true,
    dynamicUrlsApiEndpoint: '/api/sitemap-urls', // Подключение к API для динамических URL
  },
});
