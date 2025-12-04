// Explicit day module imports to avoid Vite glob issues
export const dayModules = {
  1: () => import('../../d01/d01.ts'),
  2: () => import('../../d02/d02.ts'),
  3: () => import('../../d03/d03.ts'),
  4: () => import('../../d04/d04.ts'),
  // Add more days as they are implemented
}
