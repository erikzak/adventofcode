import { defineConfig, Plugin } from 'vite'
import react from '@vitejs/plugin-react'

// Plugin to replace getDayInput import with browser-compatible version
const replaceDayInputPlugin = (): Plugin => {
  return {
    name: 'replace-day-input',
    transform(code, id) {
      // Only transform day solution files (d01.ts, d02.ts, etc.)
      if (/src[\\\/]d\d{2}[\\\/]d\d{2}\.ts$/.test(id)) {
        // Replace the import statement and the getDayInput call
        let transformedCode = code.replace(
          /import getDayInput from ["']\.\.\/lib\/input-parser\.ts["']/g,
          '// getDayInput not used in browser\nconst getDayInput = () => []'
        )
        // Also stub out the fs and os imports if they exist
        transformedCode = transformedCode.replace(
          /import .* from ["']fs["']/g,
          '// fs module stubbed for browser'
        )
        transformedCode = transformedCode.replace(
          /import .* from ["']os["']/g,
          '// os module stubbed for browser'
        )
        return {
          code: transformedCode,
          map: null
        }
      }
      return null
    }
  }
}

export default defineConfig({
  plugins: [
    react(),
    replaceDayInputPlugin()
  ],
  resolve: {
    extensions: ['.ts', '.tsx', '.js', '.jsx']
  }
})
