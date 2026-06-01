# AI NAILS Desktop Shell

Electron wrapper for macOS and Windows. It loads the shared web core from `web/static`.

## Development

```bash
cd apps/desktop-electron
npm install
npm run start
```

## Build

```bash
npm run dist:mac
npm run dist:win
```

The packaged app includes `../../web/static` as an app resource and runs fully offline with the built-in demo data.
