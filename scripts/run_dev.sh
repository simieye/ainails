#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

python3 agents/openclaw_agents.py &
AGENT_PID=$!
trap 'kill "$AGENT_PID" 2>/dev/null || true' EXIT

OPENCLAW_AGENT_URL="${OPENCLAW_AGENT_URL:-http://127.0.0.1:8090}" go run ./cmd/openclaw

