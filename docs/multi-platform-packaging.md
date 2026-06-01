# AI NAILS 多端 APP 封装方案

AI NAILS 采用一套 Web 核心、多端原生壳的封装方式：

- Web/PWA 核心：`web/static`
- macOS/Windows 本地 APP：`apps/desktop-electron`
- iOS/Android APP：`apps/mobile-capacitor`

## 1. Web/PWA 核心

入口：

```text
web/static/index.html
```

能力：

- 支持 `file://` 直接打开，自动进入离线演示模式。
- 支持 HTTP 部署，自动注册 `sw.js` 离线缓存。
- `manifest.webmanifest` 提供 PWA 安装元数据。

本地验证：

```bash
python3 -m http.server 4173 -d web/static
open http://127.0.0.1:4173
```

## 2. macOS / Windows 本地 APP

技术壳：Electron。

目录：

```text
apps/desktop-electron
```

开发运行：

```bash
cd apps/desktop-electron
npm install
npm run start
```

打包：

```bash
npm run dist:mac
npm run dist:win
```

输出目标：

- macOS：`.dmg`、`.zip`
- Windows：`NSIS installer`、`portable exe`

生产签名：

- macOS：配置 Apple Developer ID、notarization。
- Windows：配置代码签名证书。

## 3. iOS / Android APP

技术壳：Capacitor。

目录：

```text
apps/mobile-capacitor
```

初始化：

```bash
cd apps/mobile-capacitor
npm install
npm run add:ios
npm run add:android
npm run sync
```

打开原生工程：

```bash
npm run open:ios
npm run open:android
```

生产准备：

- iOS：在 Xcode 中配置 Bundle ID、Signing、Camera/Microphone 权限。
- Android：在 Android Studio 中配置 applicationId、签名、权限。
- 如果接入真实语音输入，需要麦克风权限。
- 如果接入 AR 指甲识别，需要相机权限。

## 4. 技能与生成能力

当前 Skills 管理页映射：

- `openai-whisper-api`：语音转 Prompt，需要 `OPENAI_API_KEY`。
- `openai-image-gen`：OpenAI 图片生成，需要 `OPENAI_API_KEY`。
- `nano-banana-pro`：Gemini 3 Pro Image 生成，需要 `GEMINI_API_KEY`。
- `peekaboo`：macOS UI 自动化，需要系统权限。

APP 内当前提供本地演示预览；真实 PNG 生图需要在壳层或后端注入对应 API key。

## 5. 推荐发布路线

1. 先发布 PWA 和桌面 Electron 内测版。
2. 接入 API key/后端代理，避免把密钥放进客户端。
3. 用 Capacitor 生成 iOS/Android 工程。
4. 加入相机、麦克风、相册、本地缓存插件。
5. 完成应用商店签名、隐私政策和审核材料。
