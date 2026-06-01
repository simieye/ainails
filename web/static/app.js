const state = {
  overview: null,
  selectedCampaignId: "cmp-seed-001",
  selectedDesign: 0,
  mode: "b2c",
  printTimer: null,
  recognition: null,
  offline: window.location.protocol === "file:",
};

const nailSkills = [
  {
    id: "openai-whisper-api",
    name: "OpenAI Whisper API",
    kind: "语音转写",
    status: "installed",
    env: "OPENAI_API_KEY",
    command: "openai-whisper-api/scripts/transcribe.sh audio.m4a --language zh",
    description: "把用户语音创意转成中文提示词，适合 Talk to Create 的第一步。",
  },
  {
    id: "openai-image-gen",
    name: "OpenAI Image Gen",
    kind: "美甲图片生成",
    status: "installed",
    env: "OPENAI_API_KEY",
    command: "python3 openai-image-gen/scripts/gen.py --prompt \"...\" --count 4",
    description: "批量生成美甲图案候选图，输出 PNG 与缩略图画廊。",
  },
  {
    id: "nano-banana-pro",
    name: "Nano Banana Pro",
    kind: "高精度图案重构",
    status: "installed",
    env: "GEMINI_API_KEY",
    command: "uv run nano-banana-pro/scripts/generate_image.py --prompt \"...\" --resolution 1K",
    description: "用于高质量图案重构、甲面适配和 1K/2K/4K 交付图。",
  },
  {
    id: "peekaboo",
    name: "Peekaboo",
    kind: "桌面自动化",
    status: "installed",
    env: "macOS permissions",
    command: "peekaboo see --annotate",
    description: "用于捕捉和自动化本机 UI，可辅助采集参考图与调试桌面流程。",
  },
];

const $ = (selector) => document.querySelector(selector);

function buildDemoAgents() {
  const domains = ["创意", "图案", "甲型", "设备", "社区", "推荐", "合规", "财务"];
  const roles = ["洞察", "规划", "执行", "优化", "质检", "本地算力", "云端算力", "复盘"];
  return domains.flatMap((domain, i) =>
    roles.map((role, j) => ({
      id: `agent-${String(i + 1).padStart(2, "0")}-${String(j + 1).padStart(2, "0")}`,
      name: `${domain}${role}Agent`,
      domain,
      role,
      executor: (i + j) % 2 ? "cloud" : "local",
      status: "ready",
    })),
  );
}

const demoOverview = {
  brand: "AI NAILS",
  mission: "AI NAILS APP 以 OpenClaw Master 承接自然语言创意，以 nanobanana 3.0 重构图案，并通过 LK Box 把 1200 DPI 甲面喷印压缩到 10 秒闭环。",
  layers: [
    { position: "Layer 1", name: "OpenClaw Master 大脑层", description: "负责自然语言解析、多模态上下文、MCP 状态回传与用户意图路由。", items: ["Talk to Create", "Prompt Refiner", "MCP Session"] },
    { position: "Layer 2", name: "nanobanana 3.0 创意层", description: "亚秒级 AIGC 图案生成，并对不同甲型执行 Nail-Adaptive Deformation。", items: ["4 图候选生成", "甲型拓扑形变", "AR 试戴预览"] },
    { position: "Layer 3", name: "LK Box 硬执行层", description: "连接 AI NAILS Printer、耗材、视觉定位和本地/云端算力调度。", items: ["1200 DPI 喷印", "V-ALIGN 视觉定位", "本地数据主权"] },
  ],
  agents: buildDemoAgents(),
  products: [
    { sku: "AIN-PRINTER-1200DPI", name: "AI NAILS Printer 01", category: "V-ALIGN 对焦成功", margin_rate: 0.92, status: "online" },
    { sku: "LK-BOX-EDGE-01", name: "LK Box 龙虾云盒", category: "本地算力 68%", margin_rate: 0.68, status: "online" },
    { sku: "CMYK-COATING-MAGSAFE", name: "CMYK 墨盒 / 涂层液", category: "耗材仓", margin_rate: 0.54, status: "reorder-soon" },
  ],
  orders: [
    { customer: "Cyber butterfly #A1", product_sku: "AIN-PRINTER-1200DPI", market: "Shanghai Flagship", channel: "APP WebSocket", status: "queued" },
    { customer: "Aurora ink #B4", product_sku: "CMYK-COATING-MAGSAFE", market: "Seoul Creator Bar", channel: "LK Box Local", status: "printing" },
    { customer: "Lunar french #C2", product_sku: "LK-BOX-EDGE-01", market: "Tokyo Pop-up", channel: "Alliance SaaS", status: "completed" },
  ],
  campaigns: [
    { id: "cmp-seed-001", name: "上海静安店 45 天回本计划", market: "Shanghai Flagship", product: "Cyber butterfly nail set", budget_usd: 4200, status: "ready" },
  ],
  knowledge: [
    { title: "Cyber Butterfly Prompt 资产", domain: "Prompt Economy", locale: "zh-CN", source: "creator-upload", summary: "深蓝微光、蝴蝶翼膜、霓虹边缘与短甲拓扑留白的组合资产。" },
    { title: "本地甲面拓扑隐私策略", domain: "Local-first", locale: "en-US", source: "policy-brief", summary: "指甲边缘 3D 拓扑、手部图像与生成图案优先在 LK Box 加密处理。" },
    { title: "店中店 45 天 ROI 运营 SOP", domain: "Alliance", locale: "zh-CN", source: "implementation-playbook", summary: "开机率、客单价、耗材续订、热点图案排行与无人值守财务看板。" },
  ],
  skillhub: [
    { name: "nanobanana 提示词专家", domain: "创意生成", level: "advanced", integrations: ["nanobanana", "OpenClaw"], install_hint: "install-skill nail-prompt-refiner", description: "把用户口语重构为高表现力图案提示词，并输出甲型适配参数。" },
    { name: "甲型自适应拓扑形变", domain: "AR 试戴", level: "advanced", integrations: ["Metal", "OpenGL", "LK Box"], install_hint: "install-skill nail-topology", description: "识别椭圆、方圆、杏仁等甲型边缘，并对图案进行局部拉伸与留白。" },
    { name: "V-ALIGN 打印定位", domain: "硬件控制", level: "pro", integrations: ["AI NAILS Printer", "MCP"], install_hint: "install-skill v-align-print", description: "校验手指位置、喷头距离、墨量和涂层液状态，生成可打印任务。" },
  ],
  mcp_connectors: [
    { name: "AI NAILS Printer 状态通道", category: "Hardware", platforms: ["WebSocket", "JSON-RPC", "MCP"], status: "ready", description: "毫秒级同步物理阻碍、墨尽、对焦成功、喷头温度和打印完成状态。", use_cases: ["V-ALIGN 对焦", "喷印任务下发", "故障诊断"] },
    { name: "LK Box Local-first 连接器", category: "Edge", platforms: ["LK Box", "Private GPU", "Encrypted Store"], status: "ready", description: "在本地处理手部图像、甲面拓扑和私密 Prompt，保留数据主权。", use_cases: ["本地推理", "隐私加密", "边缘缓存"] },
    { name: "nanobanana 3.0 图案引擎", category: "Creative", platforms: ["AIGC", "Nail Deformation", "AR Preview"], status: "ready", description: "把口语创意重构为高表现力图案，并完成不同甲型的拓扑适配。", use_cases: ["4 图候选", "提示词优化", "甲型形变"] },
    { name: "全球创作者社区连接器", category: "Community", platforms: ["13 Locales", "Prompt Assets", "Creator Split"], status: "ready", description: "支持多语言分享、Prompt 资产采用计费和独立美甲师创作分成。", use_cases: ["一键翻译", "资产分成", "趋势推荐"] },
  ],
  skillhub_target: "100000+",
};

async function api(path, options = {}) {
  if (state.offline) {
    throw new Error("offline file mode");
  }
  const response = await fetch(path, {
    headers: { "Content-Type": "application/json" },
    ...options,
  });
  if (!response.ok) {
    const error = await response.json().catch(() => ({ error: response.statusText }));
    throw new Error(error.error || response.statusText);
  }
  return response.json();
}

function card(title, body, items = []) {
  return `
    <div class="card">
      <h3>${title}</h3>
      <p>${body}</p>
      ${items.length ? `<ul>${items.map((item) => `<li>${item}</li>`).join("")}</ul>` : ""}
    </div>
  `;
}

async function loadOverview() {
  try {
    state.overview = await api("/api/overview");
    state.offline = false;
  } catch (error) {
    state.overview = demoOverview;
    state.offline = true;
    $("#runStatus").textContent = "离线演示模式";
    $("#skillRunStatus").textContent = "离线演示模式";
  }
  $("#mission").textContent = state.overview.mission;
  $("#metricAgents").textContent = state.overview.agents.length;
  $("#metricMCP").textContent = state.overview.mcp_connectors.length;
  $("#metricSkillHub").textContent = state.overview.skillhub_target;
  $("#layers").innerHTML = state.overview.layers
    .map((layer) => card(`${layer.position} · ${layer.name}`, layer.description, layer.items))
    .join("");
  $("#agentMatrix").innerHTML = state.overview.agents
    .map((agent) => `<div class="agent-cell"><strong>${agent.domain}</strong>${agent.role}<br /><small>${agent.executor}</small></div>`)
    .join("");
  renderDesigns();
  renderProducts(state.overview.products);
  renderOrders(state.overview.orders);
  renderCampaigns(state.overview.campaigns);
  renderKnowledge(state.overview.knowledge);
  renderSkillHub(state.overview.skillhub);
  renderMCPConnectors(state.overview.mcp_connectors);
  renderNailSkillRegistry();
  updateSkillCommandPreview();
}

function renderDesigns(seedText = $("#promptInput")?.value || "") {
  const motifs = [
    ["Cyber Wing", "深蓝微光 / 蝴蝶翼膜 / 镜面渐变", "cyber"],
    ["Aurora Ink", "霓虹紫 / 流体水墨 / 微闪颗粒", "aurora"],
    ["Lunar French", "法式极简 / 月白边缘 / 银色电路", "lunar"],
    ["Jade Pulse", "国风翡翠 / 呼吸绿 / 拓扑云纹", "jade"],
  ];
  $("#designGrid").innerHTML = motifs
    .map(
      ([name, desc, style], index) => `
        <article class="design-card ${index === state.selectedDesign ? "selected" : ""}" data-design="${index}" data-style="${style}">
          <div class="art"><span></span></div>
          <div class="meta">
            <strong>${index + 1}. ${name}</strong>
            <small>${desc}</small>
            <p>${buildRefinedPrompt(seedText, name)}</p>
            <button type="button" data-select-design="${index}">${index === state.selectedDesign ? "已选中" : "选择试戴"}</button>
          </div>
        </article>
      `,
    )
    .join("");
  document.querySelectorAll("[data-select-design]").forEach((button) => {
    button.addEventListener("click", () => selectDesign(Number(button.dataset.selectDesign)));
  });
  updateRefiner(seedText, motifs[state.selectedDesign]);
}

function buildRefinedPrompt(seedText, style) {
  const text = seedText.trim() || "自然语言创意";
  return `nanobanana 3.0 已将“${text.slice(0, 26)}”重构为 ${style} 甲面图案，适配椭圆/方圆/杏仁甲型。`;
}

function updateRefiner(seedText, motif) {
  if (!motif) return;
  const shape = $("#shapeSelect")?.selectedOptions[0]?.textContent || "椭圆";
  $("#refinerOutput").innerHTML = `
    <strong>Prompt Refiner</strong>
    <span>${motif[0]} · ${shape}甲 · 1200 DPI · 安全边距 0.8mm · 局部形变强度 ${$("#glowRange")?.value || 52}%</span>
  `;
  $("#hotDesign").textContent = motif[0];
}

function selectDesign(index) {
  state.selectedDesign = index;
  const cards = document.querySelectorAll(".design-card");
  const card = cards[index];
  if (!card) return;
  const style = card.dataset.style || "cyber";
  $("#nailShape").className = `nail-shape style-${style} shape-${$("#shapeSelect").value}`;
  renderDesigns($("#promptInput").value);
  $("#runStatus").textContent = `${card.querySelector("strong").textContent} 已同步 AR 试戴`;
}

function renderProducts(products) {
  $("#productsTable").innerHTML = [
    `<div class="row"><span>设备/耗材</span><span>状态</span><span>余量</span></div>`,
    ...products.map(
      (product) =>
        `<div class="row"><span>${product.name}<br /><small>${product.sku}</small></span><span>${product.category} · ${product.status}</span><span>${Math.round(product.margin_rate * 100)}%</span></div>`,
    ),
  ].join("");
}

function renderOrders(orders) {
  $("#ordersTable").innerHTML = [
    `<div class="row"><span>打印任务</span><span>通道</span><span>状态</span></div>`,
    ...orders.map(
      (order) =>
        `<div class="row"><span>${order.customer}<br /><small>${order.product_sku}</small></span><span>${order.market} · ${order.channel}</span><span>${order.status}</span></div>`,
    ),
  ].join("");
}

function renderCampaigns(campaigns) {
  $("#campaignList").innerHTML = campaigns
    .map(
      (campaign) => `
        <div class="campaign">
          <div>
            <strong>${campaign.name}</strong>
            <small>${campaign.market} · ${campaign.product} · $${campaign.budget_usd}</small>
            ${campaign.last_summary ? `<p>${campaign.last_summary}</p>` : ""}
          </div>
          <button data-run="${campaign.id}">运行计划</button>
        </div>
      `,
    )
    .join("");
  document.querySelectorAll("[data-run]").forEach((button) => {
    button.addEventListener("click", () => runCampaign(button.dataset.run));
  });
  if (campaigns[0]) {
    state.selectedCampaignId = campaigns[0].id;
  }
}

function renderKnowledge(assets) {
  $("#knowledgeList").innerHTML = assets
    .map(
      (asset) => `
        <article class="knowledge-item">
          <strong>${asset.title}</strong>
          <p>${asset.summary}</p>
          <small>${asset.domain} · ${asset.locale} · ${asset.source}</small>
        </article>
      `,
    )
    .join("");
}

function renderSkillHub(skills) {
  $("#skillHubList").innerHTML = skills
    .map(
      (skill) => `
        <article class="skill-row">
          <div>
            <strong>${skill.name}</strong>
            <p>${skill.description}</p>
            <small>${skill.domain} · ${skill.level} · ${skill.integrations.join(" + ")}</small>
          </div>
          <code>${skill.install_hint}</code>
        </article>
      `,
    )
    .join("");
}

function renderMCPConnectors(connectors) {
  $("#mcpConnectors").innerHTML = connectors
    .map(
      (connector) => `
        <article class="connector-card">
          <div class="card-head">
            <strong>${connector.name}</strong>
            <span class="status ${connector.status}">${connector.status}</span>
          </div>
          <small>${connector.category}</small>
          <p>${connector.description}</p>
          <div class="chips">${connector.platforms.map((platform) => `<span>${platform}</span>`).join("")}</div>
          <ul>${connector.use_cases.map((item) => `<li>${item}</li>`).join("")}</ul>
        </article>
      `,
    )
    .join("");
}

function renderNailSkillRegistry() {
  $("#nailSkillRegistry").innerHTML = nailSkills
    .map(
      (skill) => `
        <article class="nail-skill-card">
          <div class="card-head">
            <strong>${skill.name}</strong>
            <span class="status ready">${skill.status}</span>
          </div>
          <small>${skill.kind} · ${skill.env}</small>
          <p>${skill.description}</p>
          <code>${skill.command}</code>
        </article>
      `,
    )
    .join("");
}

async function refreshCampaigns() {
  if (state.offline) {
    renderCampaigns(state.overview.campaigns);
    return;
  }
  const campaigns = await api("/api/campaigns");
  renderCampaigns(campaigns);
}

async function createCampaign(event) {
  event.preventDefault();
  const form = new FormData(event.currentTarget);
  const campaign = {
    id: `cmp-local-${Date.now()}`,
    name: form.get("name"),
    market: form.get("market"),
    product: form.get("product"),
    budget_usd: Number(form.get("budget_usd")),
    workflow_id: "wf-nail-roi-loop",
    status: "ready",
  };
  if (state.offline) {
    state.overview.campaigns = [campaign, ...state.overview.campaigns];
    renderCampaigns(state.overview.campaigns);
    $("#runStatus").textContent = "离线运营计划已创建";
    return;
  }
  await api("/api/campaigns", {
    method: "POST",
    body: JSON.stringify({
      name: campaign.name,
      market: campaign.market,
      product: campaign.product,
      budget_usd: campaign.budget_usd,
      workflow_id: "wf-nail-roi-loop",
    }),
  });
  await refreshCampaigns();
}

async function runCampaign(id) {
  $("#runStatus").textContent = "SIMIAIOS 64 Agent 调度中...";
  $("#runResult").innerHTML = "";
  simulatePrintProgress();
  if (state.offline) {
    const result = buildOfflineRunResult(id);
    await wait(850);
    renderRunResult(result);
    finishPrintProgress();
    $("#runStatus").textContent = `离线完成：${result.id}`;
    return;
  }
  const result = await api("/api/run", {
    method: "POST",
    body: JSON.stringify({
      campaign_id: id,
      workflow_id: "wf-nail-roi-loop",
      context: { prompt: $("#promptInput").value },
    }),
  });
  renderRunResult(result);
  $("#runStatus").textContent = `完成：${result.id}`;
  finishPrintProgress();
  await refreshCampaigns();
}

function renderRunResult(result) {
  $("#metricROI").textContent = result.roi;
  $("#metricGMV").textContent = `$${result.gmv_usd}`;
  $("#metricSpend").textContent = `$${result.spend_usd}`;
  $("#runResult").innerHTML = `
    <h3>${result.summary}</h3>
    <p>Budget $${result.spend_usd} · Revenue $${result.gmv_usd} · ROI ${result.roi}</p>
    ${result.steps
      .map(
        (step) => `
          <div class="step">
            <strong>${step.agent_name}</strong>
            <p>${step.output}</p>
            <small>Tools: ${step.tool_calls.join(" / ")} · Confidence ${step.confidence}</small>
          </div>
        `,
      )
      .join("")}
  `;
}

function buildOfflineRunResult(campaignId) {
  const campaign = state.overview.campaigns.find((item) => item.id === campaignId) || state.overview.campaigns[0];
  const prompt = $("#promptInput").value;
  return {
    id: `run-local-${Date.now().toString().slice(-6)}`,
    summary: `OpenClaw 离线调度 6 个 Agent，围绕「${campaign.product}」完成创作、试戴、打印与 ROI 闭环模拟。`,
    spend_usd: 3444,
    gmv_usd: 6387,
    roi: 1.85,
    steps: [
      { agent_name: "创意洞察Agent", output: `解析用户创意「${prompt.slice(0, 34)}」，生成美甲视觉方向。`, tool_calls: ["mcp.openclaw.prompt_refiner"], confidence: 0.91 },
      { agent_name: "图案执行Agent", output: "输出 4 张可打印图案，并为甲面边缘保留 0.8mm 安全边距。", tool_calls: ["mcp.nanobanana.generate"], confidence: 0.88 },
      { agent_name: "甲型优化Agent", output: "完成椭圆/方圆/杏仁甲型拓扑形变参数。", tool_calls: ["mcp.nail.topology_deform"], confidence: 0.86 },
      { agent_name: "设备执行Agent", output: "模拟 LK Box 下发 V-ALIGN 打印任务，喷头温度与耗材状态正常。", tool_calls: ["mcp.lkbox.printer_control"], confidence: 0.9 },
      { agent_name: "推荐复盘Agent", output: "将当前图案写入热点推荐池，提升 Gallery 采用率。", tool_calls: ["mcp.recommendation.taiji64"], confidence: 0.84 },
      { agent_name: "财务复盘Agent", output: "根据今日打印量与耗材成本预测 45 天 ROI 进度。", tool_calls: ["mcp.finance.roi_dashboard"], confidence: 0.87 },
    ],
  };
}

function wait(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

function simulatePrintProgress() {
  clearInterval(state.printTimer);
  const phases = ["V-ALIGN 对焦", "涂层校验", "1200 DPI 喷印", "固化完成"];
  let value = 0;
  $("#printPhase").textContent = phases[0];
  $("#printTelemetry").textContent = "V-ALIGN 正在定位甲面边缘，LK Box 延迟 12ms。";
  state.printTimer = setInterval(() => {
    value = Math.min(value + 9 + Math.round(Math.random() * 8), 96);
    const phase = phases[Math.min(Math.floor(value / 28), phases.length - 1)];
    $("#printPhase").textContent = phase;
    $("#printPercent").textContent = `${value}%`;
    $("#printBar").style.width = `${value}%`;
    $("#printTelemetry").textContent = `${phase}中，喷头温度 ${36 + Math.floor(value / 18)}°C，CMYK 微滴校准 ${Math.min(value + 3, 99)}%。`;
    if (value >= 96) clearInterval(state.printTimer);
  }, 260);
}

function finishPrintProgress() {
  clearInterval(state.printTimer);
  $("#printPhase").textContent = "打印完成";
  $("#printPercent").textContent = "100%";
  $("#printBar").style.width = "100%";
  $("#printTelemetry").textContent = "10 秒喷印闭环完成，打印记录已回写 LK Box 与创作者分成账本。";
}

function updateShapeControls() {
  const nail = $("#nailShape");
  const currentStyle = [...nail.classList].find((name) => name.startsWith("style-")) || "style-cyber";
  nail.className = `nail-shape ${currentStyle} shape-${$("#shapeSelect").value}`;
  nail.style.width = `${$("#lengthRange").value}px`;
  nail.style.boxShadow = `0 0 ${$("#glowRange").value}px rgba(176, 124, 255, 0.58)`;
  updateRefiner($("#promptInput").value, getSelectedMotif());
}

function getSelectedMotif() {
  const motifs = [
    ["Cyber Wing", "深蓝微光 / 蝴蝶翼膜 / 镜面渐变", "cyber"],
    ["Aurora Ink", "霓虹紫 / 流体水墨 / 微闪颗粒", "aurora"],
    ["Lunar French", "法式极简 / 月白边缘 / 银色电路", "lunar"],
    ["Jade Pulse", "国风翡翠 / 呼吸绿 / 拓扑云纹", "jade"],
  ];
  return motifs[state.selectedDesign] || motifs[0];
}

function setMode(mode) {
  state.mode = mode;
  document.querySelectorAll(".mode-option").forEach((button) => {
    button.classList.toggle("active", button.dataset.mode === mode);
  });
  $("#runStatus").textContent = mode === "b2b" ? "B2B 店中店运营视图已启用" : "B2C 家庭版创作视图已启用";
  $("#roiProgress").textContent = mode === "b2b" ? "63%" : "18%";
}

function refineSkillPrompt() {
  const rawPrompt = $("#skillPromptInput").value.trim() || "自然语言美甲创意";
  const engine = $("#imageSkillSelect").value;
  const refined = [
    `AI NAILS premium nail art pattern, ${rawPrompt}`,
    "single elegant fingernail preview, nail-adaptive deformation safe margins",
    "1200 DPI printable surface, glossy gel texture, clean dark beauty-tech studio",
    "avoid text, logos, extra fingers, distorted nail edges",
  ].join(", ");
  $("#skillRefinerOutput").innerHTML = `
    <strong>OpenClaw Skill Refiner</strong>
    <span>${refined}</span>
  `;
  $("#skillRunStatus").textContent = `${engine} Prompt 已精修`;
  updateSkillCommandPreview(refined);
  return refined;
}

function updateSkillCommandPreview(prompt = $("#skillPromptInput")?.value || "") {
  const engine = $("#imageSkillSelect")?.value || "openai-image-gen";
  const size = $("#imageSizeSelect")?.value || "1K";
  const compactPrompt = (prompt || "AI NAILS nail art").replace(/\s+/g, " ").slice(0, 96);
  if (engine === "nano-banana-pro") {
    $("#skillCommandPreview").textContent = `uv run ~/.openclaw/workspace/skills/nano-banana-pro/scripts/generate_image.py --prompt "${compactPrompt}" --filename outputs/ai-nails/generated.png --resolution ${size}`;
    return;
  }
  $("#skillCommandPreview").textContent = `python3 ~/.openclaw/workspace/skills/openai-image-gen/scripts/gen.py --prompt "${compactPrompt}" --count 4 --out-dir outputs/ai-nails`;
}

function generateSkillImage() {
  const refined = refineSkillPrompt();
  const engine = $("#imageSkillSelect").value;
  const styles = ["style-cyber", "style-aurora", "style-lunar", "style-jade"];
  const style = styles[Math.abs(hashString(refined)) % styles.length];
  $("#skillImagePreview").innerHTML = `
    <div class="generated-nail-art ${style}"><span></span></div>
    <strong id="generatedImageTitle">${engine} 预览图已生成</strong>
    <p id="generatedImageMeta">本地演示预览已同步。真实 PNG 生成需要配置 ${engine === "nano-banana-pro" ? "GEMINI_API_KEY" : "OPENAI_API_KEY"} 后运行命令预览。</p>
    <button type="button" id="applyGeneratedDesign">同步到 AR 试戴</button>
  `;
  $("#skillRunStatus").textContent = "美甲图预览已生成";
  $("#applyGeneratedDesign").addEventListener("click", () => {
    const nail = $("#nailShape");
    const shape = $("#shapeSelect").value;
    nail.className = `nail-shape ${style} shape-${shape}`;
    $("#runStatus").textContent = "Skill 生成图已同步 AR 试戴";
  });
}

function startVoicePrompt() {
  const SpeechRecognition = window.SpeechRecognition || window.webkitSpeechRecognition;
  if (!SpeechRecognition) {
    $("#skillRunStatus").textContent = "当前浏览器不支持 Web Speech API";
    $("#skillRefinerOutput").innerHTML = `
      <strong>语音输入不可用</strong>
      <span>可使用 openai-whisper-api：录制音频文件后运行 transcribe.sh，再把文本粘贴到 Prompt。</span>
    `;
    return;
  }
  if (!state.recognition) {
    state.recognition = new SpeechRecognition();
    state.recognition.lang = "zh-CN";
    state.recognition.interimResults = false;
    state.recognition.maxAlternatives = 1;
    state.recognition.addEventListener("result", (event) => {
      const transcript = event.results[0][0].transcript;
      $("#skillPromptInput").value = transcript;
      $("#skillRunStatus").textContent = "语音已转成 Prompt";
      refineSkillPrompt();
    });
    state.recognition.addEventListener("error", (event) => {
      $("#skillRunStatus").textContent = `语音输入失败：${event.error}`;
    });
  }
  $("#skillRunStatus").textContent = "正在聆听语音创意...";
  state.recognition.start();
}

function hashString(value) {
  let hash = 0;
  for (const char of value) {
    hash = (hash * 31 + char.charCodeAt(0)) | 0;
  }
  return hash;
}

$("#refresh").addEventListener("click", loadOverview);
$("#generateDesign").addEventListener("click", () => {
  renderDesigns($("#promptInput").value);
  $("#runStatus").textContent = "nanobanana 3.0 已输出 4 张候选图";
});
$("#sendPrint").addEventListener("click", async () => {
  await runCampaign(state.selectedCampaignId);
});
$("#campaignForm").addEventListener("submit", createCampaign);
$("#shapeSelect").addEventListener("change", updateShapeControls);
$("#lengthRange").addEventListener("input", updateShapeControls);
$("#glowRange").addEventListener("input", updateShapeControls);
$("#promptInput").addEventListener("input", () => updateRefiner($("#promptInput").value, getSelectedMotif()));
document.querySelectorAll("[data-prompt]").forEach((button) => {
  button.addEventListener("click", () => {
    $("#promptInput").value = button.dataset.prompt;
    renderDesigns(button.dataset.prompt);
  });
});
document.querySelectorAll(".mode-option").forEach((button) => {
  button.addEventListener("click", () => setMode(button.dataset.mode));
});
$("#voicePromptButton").addEventListener("click", startVoicePrompt);
$("#refineSkillPrompt").addEventListener("click", refineSkillPrompt);
$("#generateSkillImage").addEventListener("click", generateSkillImage);
$("#imageSkillSelect").addEventListener("change", () => updateSkillCommandPreview());
$("#imageSizeSelect").addEventListener("change", () => updateSkillCommandPreview());
$("#skillPromptInput").addEventListener("input", () => updateSkillCommandPreview());

loadOverview().catch((error) => {
  $("#mission").textContent = error.message;
});

if ("serviceWorker" in navigator && window.location.protocol.startsWith("http")) {
  window.addEventListener("load", () => {
    navigator.serviceWorker.register("sw.js").catch(() => {
      // Offline app still works without service worker registration.
    });
  });
}
