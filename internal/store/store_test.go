package store

import (
	"path/filepath"
	"testing"
	"time"

	"openclaw/internal/domain"
)

func TestSeedCreatesSixtyFourAgents(t *testing.T) {
	st := New()
	agents := st.Agents()
	if len(agents) != 64 {
		t.Fatalf("expected 64 agents, got %d", len(agents))
	}
	if agents[0].ID != "agent-01-01" {
		t.Fatalf("unexpected first agent id: %s", agents[0].ID)
	}
}

func TestCreateCampaignValidatesInput(t *testing.T) {
	st := New()
	_, err := st.CreateCampaign(domain.CampaignRequest{})
	if err == nil {
		t.Fatal("expected validation error")
	}
}

func TestPersistentStoreReloadsCampaignsAndRuns(t *testing.T) {
	path := filepath.Join(t.TempDir(), "openclaw-state.json")
	st, err := NewPersistent(path)
	if err != nil {
		t.Fatalf("create persistent store: %v", err)
	}
	campaign, err := st.CreateCampaign(domain.CampaignRequest{
		Name:      "Japan marketplace launch",
		Market:    "Japan",
		Product:   "Smart translation earbuds",
		BudgetUSD: 2400,
	})
	if err != nil {
		t.Fatalf("create campaign: %v", err)
	}
	st.SaveRun(domain.RunResult{
		ID:          "run-persist-001",
		Goal:        "validate persistence",
		WorkflowID:  campaign.WorkflowID,
		CampaignID:  campaign.ID,
		Summary:     "saved run",
		SpendUSD:    1000,
		GMVUSD:      1800,
		ROI:         1.8,
		CreatedAt:   time.Now(),
		CompletedAt: time.Now(),
	})

	reloaded, err := NewPersistent(path)
	if err != nil {
		t.Fatalf("reload persistent store: %v", err)
	}
	campaigns := reloaded.Campaigns()
	if campaigns[0].ID != campaign.ID {
		t.Fatalf("expected latest campaign %s, got %s", campaign.ID, campaigns[0].ID)
	}
	if campaigns[0].LastRunID != "run-persist-001" {
		t.Fatalf("expected last run id to persist, got %q", campaigns[0].LastRunID)
	}
	if len(reloaded.Runs()) != 1 {
		t.Fatalf("expected one persisted run, got %d", len(reloaded.Runs()))
	}
}

func TestSeedCreatesCommerceAndKnowledgeData(t *testing.T) {
	st := New()
	if len(st.Products()) < 2 {
		t.Fatalf("expected seeded products")
	}
	if len(st.Orders()) < 2 {
		t.Fatalf("expected seeded orders")
	}
	if len(st.KnowledgeAssets()) < 2 {
		t.Fatalf("expected seeded knowledge assets")
	}
}

func TestSeedCreatesEcosystemEntrypoints(t *testing.T) {
	st := New()
	overview := st.Overview()
	if len(overview.MCPConnectors) < 4 {
		t.Fatalf("expected seeded MCP connectors, got %d", len(overview.MCPConnectors))
	}
	if len(overview.AIPartners) < 3 {
		t.Fatalf("expected seeded AI partners, got %d", len(overview.AIPartners))
	}
	if len(overview.SkillHub) < 4 {
		t.Fatalf("expected seeded skill hub skills, got %d", len(overview.SkillHub))
	}
	if overview.SkillHubTarget != "100000+" {
		t.Fatalf("expected 100000+ skill hub target, got %q", overview.SkillHubTarget)
	}
}

func TestSeedCreatesPaperclipVirtualCompany(t *testing.T) {
	st := New()
	company := st.PaperclipCompany()
	if company.ID != "simiai-nails-operator" {
		t.Fatalf("unexpected AI NAILS operator id: %s", company.ID)
	}
	if len(company.Roles) < 5 {
		t.Fatalf("expected AI NAILS operator roles, got %d", len(company.Roles))
	}
	if company.Roles[0].Title != "AI CEO" {
		t.Fatalf("expected AI CEO as first role, got %q", company.Roles[0].Title)
	}
	if len(company.Meetings) < 3 {
		t.Fatalf("expected AI NAILS operator meetings, got %d", len(company.Meetings))
	}
	if len(company.ReleasePipeline) < 4 {
		t.Fatalf("expected release pipeline stages, got %d", len(company.ReleasePipeline))
	}
	if company.GitHubURL != "https://github.com/HKUDS/CLI-Anything" {
		t.Fatalf("unexpected github url: %s", company.GitHubURL)
	}
}

func TestPersistentStoreReloadsCreatedProductOrderAndKnowledge(t *testing.T) {
	path := filepath.Join(t.TempDir(), "openclaw-state.json")
	st, err := NewPersistent(path)
	if err != nil {
		t.Fatalf("create persistent store: %v", err)
	}
	product, err := st.CreateProduct(domain.ProductRequest{
		SKU:           "SIM-TEST-SKU",
		Name:          "Test cross-border SKU",
		Category:      "AI Accessory",
		TargetMarkets: []string{"Canada", "Australia"},
		PriceUSD:      88,
		MarginRate:    0.39,
	})
	if err != nil {
		t.Fatalf("create product: %v", err)
	}
	order, err := st.CreateOrder(domain.OrderRequest{
		Market:     "Canada",
		Channel:    "Shopify",
		ProductSKU: product.SKU,
		Customer:   "Maple Labs",
		AmountUSD:  176,
	})
	if err != nil {
		t.Fatalf("create order: %v", err)
	}
	advanced, err := st.AdvanceOrder(order.ID)
	if err != nil {
		t.Fatalf("advance order: %v", err)
	}
	if advanced.Status != "fulfilling" {
		t.Fatalf("expected fulfilling order, got %s", advanced.Status)
	}
	st.SaveRun(domain.RunResult{
		ID:          "run-knowledge-001",
		Goal:        "turn run into knowledge",
		WorkflowID:  "wf-growth-flywheel",
		CampaignID:  "cmp-seed-001",
		Summary:     "Agent generated a reusable market playbook",
		ROI:         2.1,
		CreatedAt:   time.Now(),
		CompletedAt: time.Now(),
	})

	reloaded, err := NewPersistent(path)
	if err != nil {
		t.Fatalf("reload persistent store: %v", err)
	}
	if reloaded.Products()[0].SKU != product.SKU {
		t.Fatalf("expected latest product %s, got %s", product.SKU, reloaded.Products()[0].SKU)
	}
	if reloaded.Orders()[0].Status != "fulfilling" {
		t.Fatalf("expected advanced order to persist, got %s", reloaded.Orders()[0].Status)
	}
	if reloaded.KnowledgeAssets()[0].Source != "agent-run" {
		t.Fatalf("expected latest knowledge source agent-run, got %s", reloaded.KnowledgeAssets()[0].Source)
	}
}
