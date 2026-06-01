package store

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"openclaw/internal/domain"
)

type Store struct {
	mu        sync.RWMutex
	dataPath  string
	layers    []domain.Layer
	agents    []domain.Agent
	workflows []domain.WorkflowTemplate
	campaigns []domain.Campaign
	runs      []domain.RunResult
	products  []domain.Product
	orders    []domain.Order
	knowledge []domain.KnowledgeAsset
	mcp       []domain.MCPConnector
	partners  []domain.AIPartner
	skillhub  []domain.SkillHubSkill
	paperclip domain.PaperclipCompany
}

type diskState struct {
	Campaigns []domain.Campaign       `json:"campaigns"`
	Runs      []domain.RunResult      `json:"runs"`
	Products  []domain.Product        `json:"products"`
	Orders    []domain.Order          `json:"orders"`
	Knowledge []domain.KnowledgeAsset `json:"knowledge"`
}

func New() *Store {
	return newSeeded("")
}

func NewPersistent(path string) (*Store, error) {
	st := newSeeded(path)
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if err := st.persist(); err != nil {
				return nil, err
			}
			return st, nil
		}
		return nil, err
	}
	if len(data) == 0 {
		return st, nil
	}
	var state diskState
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, err
	}
	if state.Campaigns != nil {
		st.campaigns = state.Campaigns
	}
	if state.Runs != nil {
		st.runs = state.Runs
	}
	if state.Products != nil {
		st.products = state.Products
	}
	if state.Orders != nil {
		st.orders = state.Orders
	}
	if state.Knowledge != nil {
		st.knowledge = state.Knowledge
	}
	return st, nil
}

func newSeeded(path string) *Store {
	agents := seedAgents()
	workflows := seedWorkflows()
	return &Store{
		dataPath:  path,
		layers:    seedLayers(),
		agents:    agents,
		workflows: workflows,
		products:  seedProducts(),
		orders:    seedOrders(),
		knowledge: seedKnowledgeAssets(),
		mcp:       seedMCPConnectors(),
		partners:  seedAIPartners(),
		skillhub:  seedSkillHub(),
		paperclip: seedPaperclipCompany(),
		campaigns: []domain.Campaign{
			{
				ID:         "cmp-seed-001",
				Name:       "上海静安店 45 天回本计划",
				Market:     "Shanghai Flagship",
				Product:    "Cyber butterfly nail set",
				BudgetUSD:  4200,
				Status:     "ready",
				CreatedAt:  time.Now().Add(-4 * time.Hour),
				WorkflowID: "wf-nail-roi-loop",
			},
		},
	}
}

func (s *Store) Overview() domain.Overview {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return domain.Overview{
		Brand:          "AI NAILS",
		Mission:        "AI NAILS APP 以 OpenClaw Master 承接自然语言创意，以 nanobanana 3.0 重构图案，并通过 LK Box 把 1200 DPI 甲面喷印压缩到 10 秒闭环。",
		Layers:         append([]domain.Layer(nil), s.layers...),
		Agents:         append([]domain.Agent(nil), s.agents...),
		Workflows:      append([]domain.WorkflowTemplate(nil), s.workflows...),
		Campaigns:      append([]domain.Campaign(nil), s.campaigns...),
		Products:       append([]domain.Product(nil), s.products...),
		Orders:         append([]domain.Order(nil), s.orders...),
		Knowledge:      append([]domain.KnowledgeAsset(nil), s.knowledge...),
		MCPConnectors:  append([]domain.MCPConnector(nil), s.mcp...),
		AIPartners:     append([]domain.AIPartner(nil), s.partners...),
		SkillHub:       append([]domain.SkillHubSkill(nil), s.skillhub...),
		SkillHubTarget: "100000+",
		Paperclip:      clonePaperclipCompany(s.paperclip),
		RevenuePlan: domain.RevenuePlan{
			Hardware: []string{"AI NAILS Printer", "LK Box 龙虾云盒", "MagSafe 耗材仓"},
			SaaS:     []string{"B2C 家庭版", "B2B 店中店", "Alliance 多店运营"},
			Usage:    []string{"nanobanana 图案生成", "AR 试戴", "1200 DPI 打印任务"},
			Commerce: []string{"Prompt 资产分成", "耗材续订", "店中店流水抽佣"},
		},
	}
}

func (s *Store) Agents() []domain.Agent {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]domain.Agent(nil), s.agents...)
}

func (s *Store) Workflows() []domain.WorkflowTemplate {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]domain.WorkflowTemplate(nil), s.workflows...)
}

func (s *Store) Campaigns() []domain.Campaign {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]domain.Campaign(nil), s.campaigns...)
}

func (s *Store) Runs() []domain.RunResult {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]domain.RunResult(nil), s.runs...)
}

func (s *Store) Products() []domain.Product {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]domain.Product(nil), s.products...)
}

func (s *Store) Orders() []domain.Order {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]domain.Order(nil), s.orders...)
}

func (s *Store) KnowledgeAssets() []domain.KnowledgeAsset {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]domain.KnowledgeAsset(nil), s.knowledge...)
}

func (s *Store) MCPConnectors() []domain.MCPConnector {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]domain.MCPConnector(nil), s.mcp...)
}

func (s *Store) AIPartners() []domain.AIPartner {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]domain.AIPartner(nil), s.partners...)
}

func (s *Store) SkillHub() []domain.SkillHubSkill {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]domain.SkillHubSkill(nil), s.skillhub...)
}

func (s *Store) PaperclipCompany() domain.PaperclipCompany {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return clonePaperclipCompany(s.paperclip)
}

func (s *Store) CreateProduct(req domain.ProductRequest) (domain.Product, error) {
	if strings.TrimSpace(req.SKU) == "" || strings.TrimSpace(req.Name) == "" || strings.TrimSpace(req.Category) == "" {
		return domain.Product{}, errors.New("sku, name and category are required")
	}
	if req.PriceUSD <= 0 {
		return domain.Product{}, errors.New("price_usd must be greater than zero")
	}
	if req.MarginRate < 0 || req.MarginRate > 1 {
		return domain.Product{}, errors.New("margin_rate must be between 0 and 1")
	}
	if len(req.TargetMarkets) == 0 {
		req.TargetMarkets = []string{"Global"}
	}
	if req.Status == "" {
		req.Status = "active"
	}
	product := domain.Product{
		ID:            "prd-" + strconv.FormatInt(time.Now().UnixNano(), 10),
		SKU:           strings.TrimSpace(req.SKU),
		Name:          strings.TrimSpace(req.Name),
		Category:      strings.TrimSpace(req.Category),
		TargetMarkets: append([]string(nil), req.TargetMarkets...),
		PriceUSD:      req.PriceUSD,
		MarginRate:    req.MarginRate,
		Status:        req.Status,
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, existing := range s.products {
		if existing.SKU == product.SKU {
			return domain.Product{}, errors.New("product sku already exists")
		}
	}
	s.products = append([]domain.Product{product}, s.products...)
	return product, s.persistLocked()
}

func (s *Store) CreateOrder(req domain.OrderRequest) (domain.Order, error) {
	if strings.TrimSpace(req.Market) == "" || strings.TrimSpace(req.Channel) == "" || strings.TrimSpace(req.ProductSKU) == "" || strings.TrimSpace(req.Customer) == "" {
		return domain.Order{}, errors.New("market, channel, product_sku and customer are required")
	}
	if req.AmountUSD <= 0 {
		return domain.Order{}, errors.New("amount_usd must be greater than zero")
	}
	order := domain.Order{
		ID:         "ord-" + strconv.FormatInt(time.Now().UnixNano(), 10),
		Market:     strings.TrimSpace(req.Market),
		Channel:    strings.TrimSpace(req.Channel),
		ProductSKU: strings.TrimSpace(req.ProductSKU),
		Customer:   strings.TrimSpace(req.Customer),
		AmountUSD:  req.AmountUSD,
		Status:     "paid",
		Logistics:  "awaiting_fulfillment",
		CreatedAt:  time.Now(),
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.productExistsLocked(order.ProductSKU) {
		return domain.Order{}, errors.New("product_sku not found")
	}
	s.orders = append([]domain.Order{order}, s.orders...)
	return order, s.persistLocked()
}

func (s *Store) AdvanceOrder(id string) (domain.Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.orders {
		if s.orders[i].ID == id {
			switch s.orders[i].Status {
			case "paid":
				s.orders[i].Status = "fulfilling"
				s.orders[i].Logistics = "smart_split: warehouse allocated"
			case "fulfilling":
				s.orders[i].Status = "shipped"
				s.orders[i].Logistics = "tracking generated by OpenClaw"
			case "shipped":
				s.orders[i].Status = "delivered"
				s.orders[i].Logistics = "delivered"
			case "delivered":
				s.orders[i].Logistics = "delivered"
			default:
				s.orders[i].Status = "paid"
				s.orders[i].Logistics = "awaiting_fulfillment"
			}
			return s.orders[i], s.persistLocked()
		}
	}
	return domain.Order{}, errors.New("order not found")
}

func (s *Store) CreateCampaign(req domain.CampaignRequest) (domain.Campaign, error) {
	if req.Name == "" || req.Market == "" || req.Product == "" {
		return domain.Campaign{}, errors.New("name, market and product are required")
	}
	if req.BudgetUSD <= 0 {
		return domain.Campaign{}, errors.New("budget_usd must be greater than zero")
	}
	if req.WorkflowID == "" {
		req.WorkflowID = "wf-nail-roi-loop"
	}
	campaign := domain.Campaign{
		ID:         "cmp-" + time.Now().Format("20060102150405"),
		Name:       req.Name,
		Market:     req.Market,
		Product:    req.Product,
		BudgetUSD:  req.BudgetUSD,
		Status:     "ready",
		CreatedAt:  time.Now(),
		WorkflowID: req.WorkflowID,
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.campaigns = append([]domain.Campaign{campaign}, s.campaigns...)
	return campaign, s.persistLocked()
}

func (s *Store) FindCampaign(id string) (domain.Campaign, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, c := range s.campaigns {
		if c.ID == id {
			return c, true
		}
	}
	return domain.Campaign{}, false
}

func (s *Store) SaveRun(result domain.RunResult) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.runs = append([]domain.RunResult{result}, s.runs...)
	if strings.TrimSpace(result.Summary) != "" {
		s.knowledge = append([]domain.KnowledgeAsset{knowledgeFromRun(result)}, s.knowledge...)
	}
	for i := range s.campaigns {
		if s.campaigns[i].ID == result.CampaignID {
			s.campaigns[i].Status = "optimized"
			s.campaigns[i].LastRunID = result.ID
			s.campaigns[i].LastSummary = result.Summary
			s.campaigns[i].LastROI = result.ROI
			s.campaigns[i].LastGMVUSD = result.GMVUSD
			s.campaigns[i].LastSpendUSD = result.SpendUSD
		}
	}
	if err := s.persistLocked(); err != nil {
		// SaveRun is intentionally fire-and-forget for API ergonomics; callers can
		// detect persistence problems at startup via NewPersistent.
		return
	}
}

func (s *Store) persist() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.persistLocked()
}

func (s *Store) persistLocked() error {
	if s.dataPath == "" {
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(s.dataPath), 0o755); err != nil {
		return err
	}
	state := diskState{
		Campaigns: s.campaigns,
		Runs:      s.runs,
		Products:  s.products,
		Orders:    s.orders,
		Knowledge: s.knowledge,
	}
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.dataPath, data, 0o644)
}

func (s *Store) productExistsLocked(sku string) bool {
	for _, product := range s.products {
		if product.SKU == sku {
			return true
		}
	}
	return false
}

func knowledgeFromRun(result domain.RunResult) domain.KnowledgeAsset {
	updatedAt := result.CompletedAt
	if updatedAt.IsZero() {
		updatedAt = time.Now()
	}
	return domain.KnowledgeAsset{
		ID:        "kb-" + result.ID,
		Title:     "Agent 复盘：" + result.Goal,
		Domain:    "数据",
		Locale:    "zh-CN",
		Source:    "agent-run",
		Summary:   result.Summary,
		UpdatedAt: updatedAt,
	}
}

func seedLayers() []domain.Layer {
	return []domain.Layer{
		{ID: "brain", Name: "OpenClaw Master 大脑层", Position: "Layer 1", Description: "负责自然语言解析、多模态上下文、MCP 状态回传与用户意图路由。", Items: []string{"Talk to Create", "Prompt Refiner", "MCP Session"}},
		{ID: "creative", Name: "nanobanana 3.0 创意层", Position: "Layer 2", Description: "亚秒级 AIGC 图案生成，并对不同甲型执行 Nail-Adaptive Deformation。", Items: []string{"4 图候选生成", "甲型拓扑形变", "AR 试戴预览"}},
		{ID: "execution", Name: "LK Box 硬执行层", Position: "Layer 3", Description: "连接 AI NAILS Printer、耗材、视觉定位和本地/云端算力调度。", Items: []string{"1200 DPI 喷印", "V-ALIGN 视觉定位", "本地数据主权"}},
	}
}

func seedAgents() []domain.Agent {
	domains := []string{"创意", "图案", "甲型", "设备", "社区", "推荐", "合规", "财务"}
	roles := []string{"洞察", "规划", "执行", "优化", "质检", "本地算力", "云端算力", "复盘"}
	agents := make([]domain.Agent, 0, 64)
	for i, domainName := range domains {
		for j, role := range roles {
			id := "agent-" + two(i+1) + "-" + two(j+1)
			agents = append(agents, domain.Agent{
				ID:       id,
				Name:     domainName + role + "Agent",
				Domain:   domainName,
				Role:     role,
				Executor: []string{"local", "cloud"}[(i+j)%2],
				Status:   "ready",
			})
		}
	}
	return agents
}

func seedWorkflows() []domain.WorkflowTemplate {
	return []domain.WorkflowTemplate{
		{ID: "wf-nail-create-print", Name: "对答即创作打印闭环", Domain: "AI NAILS", Description: "自然语言 -> Prompt Refiner -> nanobanana 4 图生成 -> 甲型自适应形变 -> V-ALIGN 打印。", AgentIDs: []string{"agent-01-01", "agent-02-03", "agent-03-04", "agent-04-03", "agent-04-05", "agent-06-08"}, Inputs: []string{"prompt", "nail_shape", "printer_id"}},
		{ID: "wf-nail-roi-loop", Name: "店中店 45 天 ROI 闭环", Domain: "Alliance SaaS", Description: "热点图案推荐、打印流水、耗材续订、开机率和多店营收并联复盘。", AgentIDs: []string{"agent-06-01", "agent-05-03", "agent-04-04", "agent-08-02", "agent-08-08", "agent-07-05"}, Inputs: []string{"store", "theme", "budget_usd"}},
		{ID: "wf-prompt-economy", Name: "Prompt 资产分成工作流", Domain: "Creator Economy", Description: "创作者上传 Prompt 资产，社区采用并打印后自动计算分成与多语言传播。", AgentIDs: []string{"agent-05-01", "agent-01-03", "agent-02-05", "agent-07-04"}, Inputs: []string{"creator_id", "prompt_asset", "locale"}},
	}
}

func seedProducts() []domain.Product {
	return []domain.Product{
		{ID: "prd-001", SKU: "AIN-PRINTER-1200DPI", Name: "AI NAILS Printer 01", Category: "V-ALIGN 对焦成功", TargetMarkets: []string{"B2C", "B2B"}, PriceUSD: 1899, MarginRate: 0.92, Status: "online"},
		{ID: "prd-002", SKU: "LK-BOX-EDGE-01", Name: "LK Box 龙虾云盒", Category: "本地算力 68%", TargetMarkets: []string{"Local-first", "Shop-in-Shop"}, PriceUSD: 1299, MarginRate: 0.68, Status: "online"},
		{ID: "prd-003", SKU: "CMYK-COATING-MAGSAFE", Name: "CMYK 墨盒 / 涂层液", Category: "耗材仓", TargetMarkets: []string{"Global"}, PriceUSD: 39, MarginRate: 0.54, Status: "reorder-soon"},
	}
}

func seedOrders() []domain.Order {
	now := time.Now()
	return []domain.Order{
		{ID: "ord-1001", Market: "Shanghai Flagship", Channel: "APP WebSocket", ProductSKU: "AIN-PRINTER-1200DPI", Customer: "Cyber butterfly #A1", AmountUSD: 18, Status: "queued", Logistics: "waiting_for_v_align", CreatedAt: now.Add(-3 * time.Minute)},
		{ID: "ord-1002", Market: "Seoul Creator Bar", Channel: "LK Box Local", ProductSKU: "CMYK-COATING-MAGSAFE", Customer: "Aurora ink #B4", AmountUSD: 22, Status: "printing", Logistics: "1200dpi pass 2/4", CreatedAt: now.Add(-90 * time.Second)},
		{ID: "ord-1003", Market: "Tokyo Pop-up", Channel: "Alliance SaaS", ProductSKU: "LK-BOX-EDGE-01", Customer: "Lunar french #C2", AmountUSD: 16, Status: "completed", Logistics: "printed_in_10s", CreatedAt: now.Add(-26 * time.Minute)},
	}
}

func seedKnowledgeAssets() []domain.KnowledgeAsset {
	now := time.Now()
	return []domain.KnowledgeAsset{
		{ID: "kb-001", Title: "Cyber Butterfly Prompt 资产", Domain: "Prompt Economy", Locale: "zh-CN", Source: "creator-upload", Summary: "深蓝微光、蝴蝶翼膜、霓虹边缘与短甲拓扑留白的组合资产。", UpdatedAt: now.Add(-48 * time.Hour)},
		{ID: "kb-002", Title: "本地甲面拓扑隐私策略", Domain: "Local-first", Locale: "en-US", Source: "policy-brief", Summary: "指甲边缘 3D 拓扑、手部图像与生成图案优先在 LK Box 加密处理。", UpdatedAt: now.Add(-24 * time.Hour)},
		{ID: "kb-003", Title: "店中店 45 天 ROI 运营 SOP", Domain: "Alliance", Locale: "zh-CN", Source: "implementation-playbook", Summary: "开机率、客单价、耗材续订、热点图案排行与无人值守财务看板。", UpdatedAt: now.Add(-12 * time.Hour)},
	}
}

func seedMCPConnectors() []domain.MCPConnector {
	return []domain.MCPConnector{
		{ID: "mcp-device", Name: "AI NAILS Printer 状态通道", Category: "Hardware", Platforms: []string{"WebSocket", "JSON-RPC", "MCP"}, Status: "ready", Description: "毫秒级同步物理阻碍、墨尽、对焦成功、喷头温度和打印完成状态。", UseCases: []string{"V-ALIGN 对焦", "喷印任务下发", "故障诊断"}},
		{ID: "mcp-lkbox", Name: "LK Box Local-first 连接器", Category: "Edge", Platforms: []string{"LK Box", "Private GPU", "Encrypted Store"}, Status: "ready", Description: "在本地处理手部图像、甲面拓扑和私密 Prompt，保留数据主权。", UseCases: []string{"本地推理", "隐私加密", "边缘缓存"}},
		{ID: "mcp-nanobanana", Name: "nanobanana 3.0 图案引擎", Category: "Creative", Platforms: []string{"AIGC", "Nail Deformation", "AR Preview"}, Status: "ready", Description: "把口语创意重构为高表现力图案，并完成不同甲型的拓扑适配。", UseCases: []string{"4 图候选", "提示词优化", "甲型形变"}},
		{ID: "mcp-community", Name: "全球创作者社区连接器", Category: "Community", Platforms: []string{"13 Locales", "Prompt Assets", "Creator Split"}, Status: "ready", Description: "支持多语言分享、Prompt 资产采用计费和独立美甲师创作分成。", UseCases: []string{"一键翻译", "资产分成", "趋势推荐"}},
		{ID: "mcp-alliance", Name: "Alliance SaaS 运营连接器", Category: "Business Ops", Platforms: []string{"ROI Dashboard", "Consumables", "Multi-store"}, Status: "planned", Description: "承接多店流水、开机率、热点排行、耗材续订和 45 天回本分析。", UseCases: []string{"财务看板", "多店管理", "耗材补货"}},
	}
}

func seedAIPartners() []domain.AIPartner {
	return []domain.AIPartner{
		{ID: "partner-prompt-artist", Name: "Prompt 美甲师", Role: "创意数字员工", Avatar: "Nail Artist Avatar", Channels: []string{"APP", "Community", "Gallery"}, MemoryScope: "用户偏好、甲型、风格禁忌", Status: "online", Description: "把自然语言创意转成可打印 Prompt 资产与候选图案。"},
		{ID: "partner-device-keeper", Name: "设备管家", Role: "LK Box 运维数字员工", Avatar: "Ops Assistant", Channels: []string{"LK Box", "Printer", "SaaS"}, MemoryScope: "墨量、喷头温度、对焦记录、故障历史", Status: "online", Description: "监控设备状态、耗材续订、远程 OTA 和异常诊断。"},
		{ID: "partner-community-curator", Name: "社区策展人", Role: "多语言社区数字员工", Avatar: "Creator Host", Channels: []string{"13 Locales", "Feeds", "Creator Split"}, MemoryScope: "热点图案、区域趋势、创作者资产", Status: "online", Description: "组织全球图案推荐、翻译、Prompt 采用和创作者分成。"},
		{ID: "partner-roi-analyst", Name: "ROI 分析师", Role: "店中店运营数字员工", Avatar: "Finance Analyst", Channels: []string{"Alliance", "Finance", "Dashboard"}, MemoryScope: "开机率、流水、客单价、耗材成本", Status: "standby", Description: "跟踪 45 天回本进度，生成无人值守财务看板。"},
	}
}

func seedSkillHub() []domain.SkillHubSkill {
	return []domain.SkillHubSkill{
		{ID: "skill-prompt-refiner", Name: "nanobanana 提示词专家", Domain: "创意生成", Level: "advanced", Integrations: []string{"nanobanana", "OpenClaw"}, InstallHint: "install-skill nail-prompt-refiner", Description: "把用户口语重构为高表现力图案提示词，并输出甲型适配参数。"},
		{ID: "skill-nail-topology", Name: "甲型自适应拓扑形变", Domain: "AR 试戴", Level: "advanced", Integrations: []string{"Metal", "OpenGL", "LK Box"}, InstallHint: "install-skill nail-topology", Description: "识别椭圆、方圆、杏仁等甲型边缘，并对图案进行局部拉伸与留白。"},
		{ID: "skill-v-align", Name: "V-ALIGN 打印定位", Domain: "硬件控制", Level: "pro", Integrations: []string{"AI NAILS Printer", "MCP"}, InstallHint: "install-skill v-align-print", Description: "校验手指位置、喷头距离、墨量和涂层液状态，生成可打印任务。"},
		{ID: "skill-roi-dashboard", Name: "店中店 ROI 看板", Domain: "Alliance", Level: "pro", Integrations: []string{"LK Box", "Finance"}, InstallHint: "install-skill nail-roi", Description: "复盘打印流水、客单价、耗材成本和 45 天回本进度。"},
		{ID: "skill-prompt-economy", Name: "Prompt 资产分成", Domain: "Creator Economy", Level: "enterprise", Integrations: []string{"Community", "Billing"}, InstallHint: "install-skill prompt-economy", Description: "跟踪创作者 Prompt 被采用与打印的次数，并计算分成。"},
	}
}

func seedPaperclipCompany() domain.PaperclipCompany {
	return domain.PaperclipCompany{
		ID:              "simiai-nails-operator",
		Name:            "SIMIAI AI NAILS Operator",
		GitHubURL:       "https://github.com/HKUDS/CLI-Anything",
		Positioning:     "面向 AI NAILS 的无人美妆运营组织，协调创意、硬件、社区、财务与合规智能体。",
		IntegrationMode: "OpenClaw 作为执行层，SIMIAIOS 64 Agent 作为美甲创作与打印运营编排层。",
		Roles: []domain.PaperclipRole{
			{ID: "nails-ceo", Title: "AI CEO", Name: "Beauty Strategist", Mission: "定义门店目标、预算、优先级和 AI 美甲商业增长方向。", Permissions: []string{"approve_budget", "set_okrs", "call_board_meeting"}, DecisionRights: []string{"产品方向", "预算阈值", "上市节奏"}},
			{ID: "nails-cto", Title: "AI CTO", Name: "Technical Operator", Mission: "拆解 OpenClaw、LK Box、设备协议和发布风险。", Permissions: []string{"create_specs", "assign_engineering_tasks", "review_code"}, DecisionRights: []string{"技术方案", "发布门禁", "集成策略"}},
			{ID: "nails-artist", Title: "AI Artist", Name: "Prompt Designer", Mission: "实现图案 Prompt、AR 试戴和趋势图案资产复盘。", Permissions: []string{"create_prompt_asset", "review_design", "publish_gallery"}, DecisionRights: []string{"图案方向", "资产质量"}},
			{ID: "nails-ops", Title: "AI COO", Name: "Device Coordinator", Mission: "维护设备心跳、耗材、打印任务和 LK Box 交付复盘。", Permissions: []string{"schedule_maintenance", "track_heartbeat", "summarize_runs"}, DecisionRights: []string{"任务节奏", "复盘格式"}},
			{ID: "nails-finance", Title: "AI CFO", Name: "ROI Analyst", Mission: "测算 45 天 ROI、流水、分成和耗材续订策略。", Permissions: []string{"analyze_funnel", "forecast_roi", "approve_reorder"}, DecisionRights: []string{"ROI 模型", "补货策略"}},
		},
		Meetings: []domain.PaperclipMeeting{
			{ID: "meeting-board", Name: "AI NAILS 董事会", Cadence: "weekly", Chair: "AI CEO", Participants: []string{"AI CEO", "AI CTO", "AI Artist", "AI COO", "AI CFO"}, Outputs: []string{"OKR 更新", "预算决策", "产品优先级"}},
			{ID: "meeting-device", Name: "设备与打印站会", Cadence: "daily", Chair: "AI COO", Participants: []string{"AI CTO", "AI Artist", "设备管家"}, Outputs: []string{"打印任务", "耗材状态", "发布风险"}},
			{ID: "meeting-roi", Name: "店中店 ROI 复盘", Cadence: "per campaign", Chair: "AI CFO", Participants: []string{"AI CEO", "AI CFO", "ROI 分析师"}, Outputs: []string{"流水质量", "热点图案", "回本进度"}},
		},
		ReleasePipeline: []string{"需求立项", "AI CTO 技术拆解", "AI Artist 图案验证", "OpenClaw 自动验证", "AI CEO 发布批准", "Alliance 分发"},
		Governance:      []string{"预算上限", "人工批准门槛", "设备测试门禁", "用户数据本地化", "会议纪要归档"},
	}
}

func clonePaperclipCompany(company domain.PaperclipCompany) domain.PaperclipCompany {
	company.Roles = append([]domain.PaperclipRole(nil), company.Roles...)
	company.Meetings = append([]domain.PaperclipMeeting(nil), company.Meetings...)
	company.ReleasePipeline = append([]string(nil), company.ReleasePipeline...)
	company.Governance = append([]string(nil), company.Governance...)
	return company
}

func two(n int) string {
	if n < 10 {
		return "0" + string(rune('0'+n))
	}
	return "10"
}
