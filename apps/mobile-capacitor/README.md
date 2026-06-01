# AI NAILS Mobile Shell

Capacitor wrapper for iOS and Android. It packages the shared static app from `web/static`.

## Development

```bash
cd apps/mobile-capacitor
npm install
npm run add:ios
npm run add:android
npm run sync
```

## Open Native Projects

```bash
npm run open:ios
npm run open:android
```

## Notes

- iOS requires Xcode on macOS.
- Android requires Android Studio and an Android SDK.
- The current web core has an offline demo mode, so native shells can run before the Go API is connected.
- To use camera, microphone, push, or local files in production, add Capacitor plugins and native permission strings.
