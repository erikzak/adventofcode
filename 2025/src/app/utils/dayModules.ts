// Explicit day module imports to avoid Vite glob issues
export const dayModules = {
  1: () => import('../../d01/d01.ts'),
  2: () => import('../../d02/d02.ts'),
  3: () => import('../../d03/d03.ts'),
  4: () => import('../../d04/d04.ts'),
  5: () => import('../../d05/d05.ts'),
  6: () => import('../../d06/d06.ts'),
  7: () => import('../../d07/d07.ts'),
  // Add more days as they are implemented
}
