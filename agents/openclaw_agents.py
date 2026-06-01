from __future__ import annotations

from dataclasses import asdict, dataclass
from datetime import datetime, timezone
from http.server import BaseHTTPRequestHandler, ThreadingHTTPServer
from json import dumps, loads
from math import fmod
from typing import Any
from urllib.parse import urlparse
from uuid import uuid4


DOMAINS = ["创意", "图案", "甲型", "设备", "社区", "推荐", "合规", "财务"]
ROLES = ["洞察", "规划", "执行", "优化", "质检", "本地算力", "云端算力", "复盘"]


@dataclass(frozen=True)
class MatrixAgent:
    id: str
    name: str
    domain: str
    role: str
    executor: str

    def run(self, goal: str, context: dict[str, Any]) -> dict[str, Any]:
        product = str(context.get("product", "AI 美甲图案"))
        market = str(context.get("market", "AI NAILS 门店"))
        budget = float(context.get("budget_usd", 1000) or 1000)
        tool = tool_for(self.domain)
        action = action_for(self.domain, self.role)
        output = (
            f"{self.name} 面向 {market} 的 {product} 执行「{action}」，"
            f"围绕目标「{goal}」生成可落地动作，建议预算权重 {budget_weight(self.domain, budget)}。"
        )
        confidence = round(0.78 + (stable_score(self.id + goal) % 17) / 100, 2)
        return {
            "agent_id": self.id,
            "agent_name": self.name,
            "domain": self.domain,
            "action": action,
            "output": output,
            "tool_calls": [tool, "rag.nail_playbook.search", "mcp.lkbox.state.write"],
            "confidence": confidence,
        }


def build_matrix() -> list[MatrixAgent]:
    agents: list[MatrixAgent] = []
    for i, domain in enumerate(DOMAINS, start=1):
        for j, role in enumerate(ROLES, start=1):
            agents.append(
                MatrixAgent(
                    id=f"agent-{i:02d}-{j:02d}",
                    name=f"{domain}{role}Agent",
                    domain=domain,
                    role=role,
                    executor="local" if (i + j) % 2 == 0 else "cloud",
                )
            )
    return agents


MATRIX = build_matrix()

WORKFLOW_AGENTS = {
    "wf-nail-create-print": ["agent-01-01", "agent-02-03", "agent-03-04", "agent-04-03", "agent-04-05", "agent-06-08"],
    "wf-nail-roi-loop": ["agent-06-01", "agent-05-03", "agent-04-04", "agent-08-02", "agent-08-08", "agent-07-05"],
    "wf-prompt-economy": ["agent-05-01", "agent-01-03", "agent-02-05", "agent-07-04"],
    "wf-growth-flywheel": ["agent-01-01", "agent-02-03", "agent-03-04", "agent-04-03", "agent-04-05", "agent-06-08"],
}


def execute_workflow(payload: dict[str, Any]) -> dict[str, Any]:
    started = datetime.now(timezone.utc)
    goal = str(payload.get("goal") or "运行 OpenClaw 多智能体 AI 美甲任务")
    workflow_id = str(payload.get("workflow_id") or "wf-nail-create-print")
    campaign_id = str(payload.get("campaign_id") or "")
    context = payload.get("context") if isinstance(payload.get("context"), dict) else {}
    agent_ids = WORKFLOW_AGENTS.get(workflow_id, WORKFLOW_AGENTS["wf-nail-create-print"])
    selected = [agent for agent in MATRIX if agent.id in agent_ids]
    steps = [agent.run(goal, context) for agent in selected]

    budget = float(context.get("budget_usd", 1000) or 1000)
    avg_confidence = sum(step["confidence"] for step in steps) / max(len(steps), 1)
    efficiency = 1.18 + fmod(stable_score(goal), 84) / 100
    spend = round(budget * 0.82, 2)
    gmv = round(spend * efficiency * avg_confidence, 2)
    roi = round(gmv / spend, 2) if spend else 0
    completed = datetime.now(timezone.utc)
    return {
        "id": "run-" + uuid4().hex[:12],
        "goal": goal,
        "workflow_id": workflow_id,
        "campaign_id": campaign_id,
        "summary": build_summary(goal, steps, roi),
        "steps": steps,
        "spend_usd": spend,
        "gmv_usd": gmv,
        "roi": roi,
        "created_at": started.isoformat(),
        "completed_at": completed.isoformat(),
    }


def build_summary(goal: str, steps: list[dict[str, Any]], roi: float) -> str:
    domains = "、".join(dict.fromkeys(step["domain"] for step in steps))
    return f"OpenClaw 已调度 {len(steps)} 个 Agent 覆盖 {domains}，围绕「{goal}」形成创作、试戴、打印与运营闭环，预测 ROI {roi}。"


def action_for(domain: str, role: str) -> str:
    actions = {
        "创意": "自然语言意图解析与 Prompt Refiner",
        "图案": "nanobanana 3.0 图案重构与 4 图候选生成",
        "甲型": "Nail-Adaptive Deformation 与 AR 试戴",
        "设备": "LK Box 调度、V-ALIGN 对焦与喷印执行",
        "社区": "多语言创作者社区与 Prompt 资产分发",
        "推荐": "太极 64 卦偏好矩阵与热点图案推荐",
        "合规": "市场合规、平台政策与风控校验",
        "财务": "45 天 ROI、耗材续订与创作者分成测算",
    }
    return f"{actions.get(domain, '业务自动化')} / {role}"


def tool_for(domain: str) -> str:
    tools = {
        "创意": "mcp.openclaw.prompt_refiner",
        "图案": "mcp.nanobanana.generate",
        "甲型": "mcp.nail.topology_deform",
        "设备": "mcp.lkbox.printer_control",
        "社区": "mcp.community.prompt_asset",
        "推荐": "mcp.recommendation.taiji64",
        "合规": "mcp.privacy.local_first_guardian",
        "财务": "mcp.finance.roi_dashboard",
    }
    return tools.get(domain, "mcp.agent.tool")


def budget_weight(domain: str, budget: float) -> str:
    weights = {
        "创意": 0.18,
        "图案": 0.22,
        "甲型": 0.14,
        "设备": 0.2,
        "社区": 0.08,
        "推荐": 0.1,
        "合规": 0.02,
        "财务": 0.06,
    }
    return f"${round(budget * weights.get(domain, 0.05), 2)}"


def stable_score(value: str) -> int:
    total = 0
    for ch in value:
        total = (total * 33 + ord(ch)) % 100000
    return total


class AgentHandler(BaseHTTPRequestHandler):
    def do_GET(self) -> None:
        path = urlparse(self.path).path
        if path == "/health":
            self.write_json({"status": "ok", "agents": len(MATRIX), "service": "openclaw-python-agent-matrix"})
            return
        if path == "/agents":
            self.write_json([asdict(agent) for agent in MATRIX])
            return
        self.write_json({"error": "not found"}, status=404)

    def do_POST(self) -> None:
        path = urlparse(self.path).path
        if path != "/run":
            self.write_json({"error": "not found"}, status=404)
            return
        try:
            length = int(self.headers.get("Content-Length", "0"))
            payload = loads(self.rfile.read(length).decode("utf-8")) if length else {}
            self.write_json(execute_workflow(payload))
        except Exception as exc:
            self.write_json({"error": str(exc)}, status=400)

    def write_json(self, payload: Any, status: int = 200) -> None:
        body = dumps(payload, ensure_ascii=False).encode("utf-8")
        self.send_response(status)
        self.send_header("Content-Type", "application/json; charset=utf-8")
        self.send_header("Content-Length", str(len(body)))
        self.end_headers()
        self.wfile.write(body)

    def log_message(self, format: str, *args: Any) -> None:
        return


def main() -> None:
    server = ThreadingHTTPServer(("127.0.0.1", 8090), AgentHandler)
    print("OpenClaw Python Agent Matrix listening on http://127.0.0.1:8090", flush=True)
    server.serve_forever()


if __name__ == "__main__":
    main()
