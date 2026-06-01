# AI NAILS OpenClaw 原生多智能体美甲打印系统

这是一个根据《AI NAILS APP 系统设计与架构白皮书 V3.0》落地的可运行 MVP。系统采用 Go + Python 双服务：

- Go：API 网关、业务模型、设备/运营状态管理、静态 AI NAILS 控制台。
- Python：OpenClaw 64-Agent 矩阵执行服务，模拟 MCP 工具调用、nanobanana 3.0 图案生成、LK Box 调度与 ROI 预测。
- Web：AI NAILS APP 五大一级看板，包括 Create、Gallery、Device、Alliance、Me。
- Data：默认使用本地 JSON 文件 `data/openclaw-state.json` 持久化运营计划与执行结果。

## 快速启动

```bash
chmod +x scripts/run_dev.sh
./scripts/run_dev.sh
```

打开：

```text
http://127.0.0.1:8080/web/static/
```

## 单独启动

```bash
python3 agents/openclaw_agents.py
```

另开一个终端：

```bash
OPENCLAW_AGENT_URL=http://127.0.0.1:8090 go run ./cmd/openclaw
```

可选环境变量：

```bash
OPENCLAW_PORT=8080
OPENCLAW_AGENT_URL=http://127.0.0.1:8090
OPENCLAW_DATA_PATH=data/openclaw-state.json
```

## API

- `GET /api/health`
- `GET /api/overview`
- `GET /api/agents`
- `GET /api/workflows`
- `GET /api/campaigns`
- `GET /api/products`
- `GET /api/orders`
- `GET /api/knowledge`
- `POST /api/campaigns`
- `POST /api/run`

示例：

```bash
curl -X POST http://127.0.0.1:8080/api/run \
  -H 'Content-Type: application/json' \
  -d '{"campaign_id":"cmp-seed-001","workflow_id":"wf-nail-roi-loop"}'
```

## 测试

```bash
GOCACHE="$(pwd)/.gocache" go test ./...
python3 -m unittest agents/test_openclaw_agents.py
```

## 白皮书映射

- 大脑层：OpenClaw Master 负责自然语言解析、多模态会话与 MCP 状态反馈。
- 创意层：nanobanana 3.0 负责 4 图候选生成、Prompt Refiner 与 Nail-Adaptive Deformation。
- 硬执行层：LK Box 与 AI NAILS Printer 负责本地数据主权、V-ALIGN 视觉定位与 1200 DPI 喷印。
- 五大看板：Create、Gallery、Device、Alliance、Me。
- 商业闭环：B2C 耗材续订、B2B 店中店 ROI 看板、Prompt Economy 创作者分成。

## 当前 APP 体验

- Create：支持 B2C/B2B 模式切换、Prompt 预设、Prompt Refiner 输出和 AR 甲型参数调节。
- Gallery：生成 4 张候选图，选择后同步到甲面预览与热点图案。
- Device：展示 AI NAILS Printer、LK Box、耗材状态、打印任务与 10 秒打印进度。
- Alliance：创建店中店运营计划，查看 45 天 ROI、今日打印量和热点图案。
- Me：展示 Prompt 资产、Local-first 隐私策略、ROI SOP 与可安装技能。

## 多端 APP 封装

当前项目已经按“一套 Web 核心，多端原生壳”组织：

- Web/PWA：`web/static`
- macOS/Windows：`apps/desktop-electron`
- iOS/Android：`apps/mobile-capacitor`
- 封装说明：`docs/multi-platform-packaging.md`

检查封装骨架：

```bash
./scripts/prepare_multiplatform.sh
```

桌面 APP：

```bash
cd apps/desktop-electron
npm install
npm run start
npm run dist:mac
npm run dist:win
```

移动 APP：

```bash
cd apps/mobile-capacitor
npm install
npm run add:ios
npm run add:android
npm run sync
```
