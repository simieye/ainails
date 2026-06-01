package domain

import "time"

type Layer struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Position    string   `json:"position"`
	Description string   `json:"description"`
	Items       []string `json:"items"`
}

type Agent struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Domain   string `json:"domain"`
	Role     string `json:"role"`
	Executor string `json:"executor"`
	Status   string `json:"status"`
}

type WorkflowTemplate struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Domain      string   `json:"domain"`
	Description string   `json:"description"`
	AgentIDs    []string `json:"agent_ids"`
	Inputs      []string `json:"inputs"`
}

type Campaign struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Market       string    `json:"market"`
	Product      string    `json:"product"`
	BudgetUSD    float64   `json:"budget_usd"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	WorkflowID   string    `json:"workflow_id"`
	LastRunID    string    `json:"last_run_id,omitempty"`
	LastSummary  string    `json:"last_summary,omitempty"`
	LastROI      float64   `json:"last_roi,omitempty"`
	LastGMVUSD   float64   `json:"last_gmv_usd,omitempty"`
	LastSpendUSD float64   `json:"last_spend_usd,omitempty"`
}

type Product struct {
	ID            string   `json:"id"`
	SKU           string   `json:"sku"`
	Name          string   `json:"name"`
	Category      string   `json:"category"`
	TargetMarkets []string `json:"target_markets"`
	PriceUSD      float64  `json:"price_usd"`
	MarginRate    float64  `json:"margin_rate"`
	Status        string   `json:"status"`
}

type ProductRequest struct {
	SKU           string   `json:"sku"`
	Name          string   `json:"name"`
	Category      string   `json:"category"`
	TargetMarkets []string `json:"target_markets"`
	PriceUSD      float64  `json:"price_usd"`
	MarginRate    float64  `json:"margin_rate"`
	Status        string   `json:"status"`
}

type Order struct {
	ID         string    `json:"id"`
	Market     string    `json:"market"`
	Channel    string    `json:"channel"`
	ProductSKU string    `json:"product_sku"`
	Customer   string    `json:"customer"`
	AmountUSD  float64   `json:"amount_usd"`
	Status     string    `json:"status"`
	Logistics  string    `json:"logistics"`
	CreatedAt  time.Time `json:"created_at"`
}

type OrderRequest struct {
	Market     string  `json:"market"`
	Channel    string  `json:"channel"`
	ProductSKU string  `json:"product_sku"`
	Customer   string  `json:"customer"`
	AmountUSD  float64 `json:"amount_usd"`
}

type KnowledgeAsset struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Domain    string    `json:"domain"`
	Locale    string    `json:"locale"`
	Source    string    `json:"source"`
	Summary   string    `json:"summary"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MCPConnector struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Category    string   `json:"category"`
	Platforms   []string `json:"platforms"`
	Status      string   `json:"status"`
	Description string   `json:"description"`
	UseCases    []string `json:"use_cases"`
}

type AIPartner struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Role        string   `json:"role"`
	Avatar      string   `json:"avatar"`
	Channels    []string `json:"channels"`
	MemoryScope string   `json:"memory_scope"`
	Status      string   `json:"status"`
	Description string   `json:"description"`
}

type SkillHubSkill struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Domain       string   `json:"domain"`
	Level        string   `json:"level"`
	Integrations []string `json:"integrations"`
	InstallHint  string   `json:"install_hint"`
	Description  string   `json:"description"`
}

type PaperclipRole struct {
	ID             string   `json:"id"`
	Title          string   `json:"title"`
	Name           string   `json:"name"`
	Mission        string   `json:"mission"`
	Permissions    []string `json:"permissions"`
	DecisionRights []string `json:"decision_rights"`
}

type PaperclipMeeting struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Cadence     string   `json:"cadence"`
	Chair       string   `json:"chair"`
	Participants []string `json:"participants"`
	Outputs     []string `json:"outputs"`
}

type PaperclipCompany struct {
	ID              string             `json:"id"`
	Name            string             `json:"name"`
	GitHubURL       string             `json:"github_url"`
	Positioning     string             `json:"positioning"`
	IntegrationMode string             `json:"integration_mode"`
	Roles           []PaperclipRole    `json:"roles"`
	Meetings        []PaperclipMeeting `json:"meetings"`
	ReleasePipeline []string           `json:"release_pipeline"`
	Governance      []string           `json:"governance"`
}

type CampaignRequest struct {
	Name       string  `json:"name"`
	Market     string  `json:"market"`
	Product    string  `json:"product"`
	BudgetUSD  float64 `json:"budget_usd"`
	WorkflowID string  `json:"workflow_id"`
}

type RunRequest struct {
	Goal       string                 `json:"goal"`
	CampaignID string                 `json:"campaign_id"`
	WorkflowID string                 `json:"workflow_id"`
	Context    map[string]interface{} `json:"context"`
}

type AgentStep struct {
	AgentID    string   `json:"agent_id"`
	AgentName  string   `json:"agent_name"`
	Domain     string   `json:"domain"`
	Action     string   `json:"action"`
	Output     string   `json:"output"`
	ToolCalls  []string `json:"tool_calls"`
	Confidence float64  `json:"confidence"`
}

type RunResult struct {
	ID          string      `json:"id"`
	Goal        string      `json:"goal"`
	WorkflowID  string      `json:"workflow_id"`
	CampaignID  string      `json:"campaign_id,omitempty"`
	Summary     string      `json:"summary"`
	Steps       []AgentStep `json:"steps"`
	SpendUSD    float64     `json:"spend_usd"`
	GMVUSD      float64     `json:"gmv_usd"`
	ROI         float64     `json:"roi"`
	CreatedAt   time.Time   `json:"created_at"`
	CompletedAt time.Time   `json:"completed_at"`
}

type RevenuePlan struct {
	Hardware []string `json:"hardware"`
	SaaS     []string `json:"saas"`
	Usage    []string `json:"usage"`
	Commerce []string `json:"commerce"`
}

type Overview struct {
	Brand          string             `json:"brand"`
	Mission        string             `json:"mission"`
	Layers         []Layer            `json:"layers"`
	Agents         []Agent            `json:"agents"`
	Workflows      []WorkflowTemplate `json:"workflows"`
	Campaigns      []Campaign         `json:"campaigns"`
	Products       []Product          `json:"products"`
	Orders         []Order            `json:"orders"`
	Knowledge      []KnowledgeAsset   `json:"knowledge"`
	MCPConnectors  []MCPConnector     `json:"mcp_connectors"`
	AIPartners     []AIPartner        `json:"ai_partners"`
	SkillHub       []SkillHubSkill    `json:"skillhub"`
	SkillHubTarget string             `json:"skillhub_target"`
	Paperclip      PaperclipCompany   `json:"paperclip"`
	RevenuePlan    RevenuePlan        `json:"revenue_plan"`
}
