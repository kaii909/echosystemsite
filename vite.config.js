import { defineConfig } from 'vite';

export default defineConfig({
  root: 'views',
  publicDir: '../static',
  server: {
    port: 3000,
    proxy: {
      '/download': 'http://localhost:8080',
      '/api': 'http://localhost:8080',
    }
  },
  build: {
    outDir: 'dist',
    emptyOutDir: true,
    rollupOptions: {
      input: {
        main: './static/main.js',
        style: './static/main.css',
      },
      output: {
        entryFileNames: 'assets/[name].js',
        chunkFileNames: 'assets/[name].js',
        assetFileNames: 'assets/[name].[ext]',
      }
    }
  }
});
