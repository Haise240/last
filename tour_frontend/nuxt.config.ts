// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  runtimeConfig: {
    public: {
      apiUrl: process.env.API_URL || 'http://localhost:8080/'
    }
  },

  compatibilityDate: '2024-09-26'
})