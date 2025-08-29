import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    port: 5173,
    host: "0.0.0.0",
  },
  plugins: [react()],
  resolve: {
    alias: {
      '@hooks': path.resolve(__dirname, 'src/hooks'), 
      '@types': path.resolve(__dirname, 'src/types') ,
      '@slices': path.resolve(__dirname, 'src/slices'), 
      '@components': path.resolve(__dirname, 'src/components'), 
    }
  }
})
