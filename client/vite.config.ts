import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'
import path from 'path'
import compression from "vite-plugin-compression"

// https://vitejs.dev/config/
export default defineConfig({
  build: {
    cssCodeSplit: true,
    rollupOptions: {
      output: {
        // Customize the filename pattern for dynamic import chunks
        chunkFileNames: 'js/chunks/[name]-[hash].js',
  
        // Customize the filename pattern for entry points (main bundle)
        entryFileNames: 'js/main/[name]-[hash].js',
  
        // Customize the filename pattern for static assets (images, fonts, etc.)
        assetFileNames: 'assets/[name]-[hash].[ext]',
      }
    },
    outDir: 'dist',
    minify: 'esbuild',
  },
  server: {
    port: 5173,
    host: "0.0.0.0",
  },
  plugins: [
    //Gzip plugin with some defaults.
    compression({
      verbose: true, // Optional: logs the compression results
      disable: false, // Optional: set to true to disable compression
      threshold: 10240, // Optional: only compress files larger than 10kb
      algorithm: 'gzip', // Optional: set to 'brotliCompress' for Brotli
      ext: '.gz', // Optional: adds .gz to the file extension
    }),
    react(),
  ],
  resolve: {
    alias: {
      '@hooks': path.resolve(__dirname, 'src/hooks'),
      '@types': path.resolve(__dirname, 'src/types'),
      '@slices': path.resolve(__dirname, 'src/slices'),
      '@components': path.resolve(__dirname, 'src/components'),
    }
  }
})
