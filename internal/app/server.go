package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"openclaw/internal/agentclient"
	"openclaw/internal/domain"
	"openclaw/internal/store"
)

type Server struct {
	store    *store.Store
	agent    *agentclient.Client
	static   http.Handler
	agentURL string
}

func NewServer(st *store.Store, agent *agentclient.Client, static http.Handler, agentURL string) *Server {
	return &Server{store: st, agent: agent, static: static, agentURL: agentURL}
}

func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/health", s.health)
	mux.HandleFunc("GET /api/overview", s.overview)
	mux.HandleFunc("GET /api/agents", s.agents)
	mux.HandleFunc("GET /api/workflows", s.workflows)
	mux.HandleFunc("GET /api/campaigns", s.campaigns)
	mux.HandleFunc("POST /api/campaigns", s.createCampaign)
	mux.HandleFunc("GET /api/products", s.products)
	mux.HandleFunc("GET /api/orders", s.orders)
	mux.HandleFunc("GET /api/knowledge", s.knowledge)
	mux.HandleFunc("GET /api/mcp-connectors", s.mcpConnectors)
	mux.HandleFunc("GET /api/ai-partners", s.aiPartners)
	mux.HandleFunc("GET /api/skillhub", s.skillHub)
	mux.HandleFunc("GET /api/paperclip", s.paperclip)
	mux.HandleFunc("GET /api/runs", s.runs)
	mux.HandleFunc("POST /api/run", s.runWorkflow)
	mux.Handle("/", s.static)
	return logging(mux)
}

func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"status":           "ok",
		"go_api":           "ready",
		"python_agent_url": s.agentURL,
		"product":          "AI NAILS OpenClaw 原生多智能体美甲打印系统",
	})
}

func (s *Server) overview(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, s.store.Overview())
}

func (s *Server) agents(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, s.store.Agents())
}

func (s *Server) workflows(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, s.store.Workflows())
}

func (s *Server) campaigns(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, s.store.Campaigns())
}

func (s *Server) products(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, s.store.Products())
}

func (s *Server) orders(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, s.store.Orders())
}

func (s *Server) knowledge(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, s.store.KnowledgeAssets())
}

func (s *Server) mcpConnectors(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, s.store.MCPConnectors())
}

func (s *Server) aiPartners(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, s.store.AIPartners())
}

func (s *Server) skillHub(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, s.store.SkillHub())
}

func (s *Server) paperclip(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, s.store.PaperclipCompany())
}

func (s *Server) runs(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, s.store.Runs())
}

func (s *Server) createCampaign(w http.ResponseWriter, r *http.Request) {
	var req domain.CampaignRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	campaign, err := s.store.CreateCampaign(req)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, campaign)
}

func (s *Server) runWorkflow(w http.ResponseWriter, r *http.Request) {
	var req domain.RunRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if req.CampaignID != "" {
		campaign, ok := s.store.FindCampaign(req.CampaignID)
		if !ok {
			writeError(w, http.StatusNotFound, "campaign not found")
			return
		}
		if req.WorkflowID == "" {
			req.WorkflowID = campaign.WorkflowID
		}
		if req.Goal == "" {
			req.Goal = "优化 " + campaign.Market + " 的 " + campaign.Product + " AI 美甲创作、打印与 ROI 闭环"
		}
		if req.Context == nil {
			req.Context = map[string]interface{}{}
		}
		req.Context["market"] = campaign.Market
		req.Context["product"] = campaign.Product
		req.Context["budget_usd"] = campaign.BudgetUSD
	}
	if strings.TrimSpace(req.Goal) == "" {
		writeError(w, http.StatusBadRequest, "goal is required")
		return
	}
	result, err := s.agent.Run(r.Context(), req)
	if err != nil {
		writeError(w, http.StatusBadGateway, "agent service unavailable: "+err.Error())
		return
	}
	s.store.SaveRun(result)
	writeJSON(w, http.StatusOK, result)
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
