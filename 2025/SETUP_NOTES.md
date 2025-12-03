# Setup Notes

## Known Issues

### Vite Dependency Scanning Warning

You may see a warning during `npm run dev` about failing to resolve the "fs" package:

```
Failed to scan for dependencies from entries...
Failed to resolve entry for package "fs"
The plugin "vite:dep-scan" was triggered by this import
  src/lib/input-parser.ts:2:29
```

**This is expected and safe to ignore.** The warning appears because:

1. The day solution modules (`d01.ts`, `d02.ts`, etc.) import `input-parser.ts` which uses Node.js `fs` module
2. Vite's dependency scanner sees these imports during initial scan
3. However, our Vite plugin (`replaceDayInputPlugin`) strips out the `fs` imports when the modules are actually loaded in the browser
4. The calendar app works correctly despite this warning

The app functionality is not affected - you can still click on calendar doors and run solutions on test input.

### Alternative Solution

If you want to eliminate the warning entirely, you would need to:
- Create separate browser-safe versions of each day module without `getDayInput` import
- Or restructure the code to not import `input-parser.ts` at all in the day modules
- Or use a build step that pre-processes the day modules before Vite sees them

For now, the simplest approach is to acknowledge the warning and proceed - the app works fine!

## Adding New Days

When you add a new day (e.g., day 4):

1. Create the day module as usual: `src/d04/d04.ts`
2. Add test and input files: `src/d04/test.txt`, `src/d04/input.txt`
3. **Important:** Add the new day to `src/app/utils/dayModules.ts`:
   ```typescript
   export const dayModules = {
     1: () => import('../../d01/d01.ts'),
     2: () => import('../../d02/d02.ts'),
     3: () => import('../../d03/d03.ts'),
     4: () => import('../../d04/d04.ts'),  // Add this line
     // ... etc
   } as const
   ```

Without step 3, clicking that day's door will show "Day X has not been implemented yet!"
