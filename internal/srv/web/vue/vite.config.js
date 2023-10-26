import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    proxy: {
      // Define a proxy for the specific API you want to access.
      '/api': {
        target: 'https://svd.fly.dev/api/v1/state', // The target API you want to access
        changeOrigin: true,
        secure: false, // Use 'false' if your target API uses a self-signed certificate (not recommended for production)
        headers: {
          //'Access-Control-Allow-Origin': '*', // Add the CORS headers
        },
      },
    },
  },
})
