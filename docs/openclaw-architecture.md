# OpenClaw MVP Architecture

## Goal

Build a runnable native AI multi-agent matrix system for cross-border ecommerce, based on the SIMIAICLAW business plan.

## Runtime

Go runs the business API and dashboard. Python runs the AI agent matrix. The services communicate through HTTP so the agent engine can later be replaced by model-backed workers, MCP tools, or private hardware executors.

## Modules

- `cmd/openclaw`: Go entrypoint.
- `internal/app`: HTTP routes and API types.
- `internal/store`: seeded in-memory business data.
- `internal/agentclient`: Go client for the Python agent service.
- `agents/openclaw_agents.py`: standard-library Python agent runtime.
- `web/static`: browser dashboard.

## MVP Boundaries

The current system is an executable product skeleton. It does not call paid LLM APIs, payment providers, ecommerce platforms, ad networks, or logistics vendors yet. Those are represented as MCP-style tool call names in agent outputs so real connectors can be added behind stable interfaces.

