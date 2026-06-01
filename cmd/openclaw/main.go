package main

import (
	"log"
	"net/http"
	"os"

	"openclaw/internal/agentclient"
	"openclaw/internal/app"
	"openclaw/internal/store"
)

func main() {
	port := env("OPENCLAW_PORT", "8080")
	agentURL := env("OPENCLAW_AGENT_URL", "http://127.0.0.1:8090")
	dataPath := env("OPENCLAW_DATA_PATH", "data/openclaw-state.json")
	static := http.FileServer(http.Dir("."))
	st, err := store.NewPersistent(dataPath)
	if err != nil {
		log.Fatalf("create OpenClaw store: %v", err)
	}
	srv := app.NewServer(st, agentclient.New(agentURL), static, agentURL)
	log.Printf("OpenClaw Go API listening on http://127.0.0.1:%s with data %s", port, dataPath)
	if err := http.ListenAndServe(":"+port, srv.Routes()); err != nil {
		log.Fatal(err)
	}
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
