import { fileURLToPath, URL } from 'node:url';
import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import nightwatchPlugin from 'vite-plugin-nightwatch';

export default defineConfig({
  plugins: [
    vue(),
    nightwatchPlugin(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
server: {
  proxy: {
    '/api': {
      target: 'http://localhost:8080/api',
      changeOrigin: true,
      rewrite: (path) => path.replace(/^\/api/, ''), // Это должно работать, если у вас правильный маршрут на бэкенде
    },
  },
},
  
});
