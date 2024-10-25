// types/nuxt-simple-sitemap.d.ts
import 'nuxt-simple-sitemap';

declare module 'nuxt-simple-sitemap' {
  interface ModuleOptions {
    hostname?: string; // Укажите, что hostname может быть строкой
    trailingSlash?: boolean; // Укажите, что trailingSlash может быть булевым значением
    dynamicUrlsApiEndpoint?: string; // Для динамических URL
  }
}
