from __future__ import annotations

from dataclasses import dataclass
from datetime import datetime, timezone
from pathlib import Path
import json
from typing import Any


DEFAULT_STATE_PATH = Path("data") / "longxia-mobile-state.json"


@dataclass(frozen=True)
class Capability:
    id: str
    name: str
    summary: str
    commands: tuple[str, ...]


CAPABILITIES = (
    Capability(
        "tiktok-loop",
        "TikTok全自动运营闭环",
        "热点采集、真人化互动、策略发布和私信转化。",
        ("tiktok plan", "social reply"),
    ),
    Capability(
        "social-agent",
        "24/7 AI静默接管",
        "覆盖WhatsApp、Messenger、Instagram等平台的自动通知回复与营销回传。",
        ("social reply", "social log-conversion"),
    ),
    Capability(
        "b2b-service",
        "外贸B2B专业客服代理",
        "识别询盘、检索报价库、匹配规格和附件。",
        ("b2b reply", "knowledge search"),
    ),
    Capability(
        "knowledge-brain",
        "多语言自定义知识库",
        "向量化风格检索产品资料、FAQ、售后规则和客户偏好。",
        ("knowledge add", "knowledge search"),
    ),
    Capability(
        "macro-workshop",
        "录屏即能力自动化工坊",
        "把触控事件拆解为可同步回放的轻量宏文件。",
        ("macro record", "macro play"),
    ),
    Capability(
        "native-call",
        "全拟真AI外呼",
        "通过原生线路拨打，生成全双工话术并识别挂断意图。",
        ("call script",),
    ),
    Capability(
        "opc-deal-flywheel",
        "AnyGen+HeyGen+OpenClaw自动化成交SOP",
        "从客户认知、策略脚本、数字人视频、全平台触达到私域成交复盘。",
        ("opc plan", "opc leads", "opc video", "opc outreach", "opc followup", "opc review"),
    ),
)


DEFAULT_KNOWLEDGE = [
    {
        "title": "Longxia AI Phone product definition",
        "locale": "zh-CN",
        "body": "ClawVS is an absolutely controlled AI agent phone with cross-app task execution, semantic understanding and system-level mapping control.",
        "tags": ["product", "control", "clawvs"],
    },
    {
        "title": "TikTok operation loop",
        "locale": "zh-CN",
        "body": "AI analyzes trends, simulates browsing, likes, comments, follows prospects, schedules posts and replies to DMs in real time.",
        "tags": ["tiktok", "growth", "dm"],
    },
    {
        "title": "B2B inquiry playbook",
        "locale": "zh-CN",
        "body": "Parse price, specification and trade terms, search the enterprise quotation library and match PDF attachments.",
        "tags": ["b2b", "quote", "pdf"],
    },
    {
        "title": "Security and memory guard",
        "locale": "zh-CN",
        "body": "Isolated memory allocation keeps AI processes and system data separated; business data remains locally owned by the enterprise.",
        "tags": ["security", "local", "memory"],
    },
]


MEMBERSHIP_PACKAGES = [
    {
        "id": "community",
        "name": "跨境龙虾社社群会员",
        "price_rmb": 5999,
        "best_for": "跨境新手、零基础创业者",
        "includes": ["资深外贸导师一对一指导", "行业干货社群", "全程避坑指引"],
    },
    {
        "id": "starter",
        "name": "龙虾入门基础版",
        "price_rmb": 11800,
        "best_for": "初创团队、小额试水跨境创业者",
        "includes": ["灵活产品选择", "基础获客", "专属客服", "智能运营系统"],
    },
    {
        "id": "acquisition",
        "name": "定制获客系统版",
        "price_rmb": 29800,
        "best_for": "追求精准获客的外贸从业者或中小团队",
        "includes": ["行业获客SOP", "24小时AI智能获客系统", "主动引流转化"],
    },
    {
        "id": "diamond",
        "name": "私董钻石会员",
        "price_rmb": 59800,
        "best_for": "企业级外贸老板",
        "includes": ["全流程AGENTS智能系统", "获客询盘转化履约售后闭环", "企业级自动化运营"],
    },
]


def default_state() -> dict[str, Any]:
    return {
        "knowledge": DEFAULT_KNOWLEDGE,
        "macros": {},
        "conversions": [],
        "opc_runs": [],
        "created_at": utc_now(),
    }


def recommend_package(budget_rmb: int) -> dict[str, Any]:
    affordable = [package for package in MEMBERSHIP_PACKAGES if package["price_rmb"] <= budget_rmb]
    recommended = affordable[-1] if affordable else MEMBERSHIP_PACKAGES[0]
    return {"budget_rmb": budget_rmb, "recommended": recommended, "packages": MEMBERSHIP_PACKAGES}


def utc_now() -> str:
    return datetime.now(timezone.utc).isoformat()


def load_state(path: Path) -> dict[str, Any]:
    if not path.exists():
        state = default_state()
        save_state(path, state)
        return state
    with path.open("r", encoding="utf-8") as handle:
        return json.load(handle)


def save_state(path: Path, state: dict[str, Any]) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    with path.open("w", encoding="utf-8") as handle:
        json.dump(state, handle, ensure_ascii=False, indent=2)


def estimate_tiktok_plan(market: str, sku: str, budget: float) -> dict[str, Any]:
    content = round(budget * 0.28, 2)
    interaction = round(budget * 0.24, 2)
    dm = round(budget * 0.18, 2)
    ads = round(budget * 0.22, 2)
    reserve = round(budget - content - interaction - dm - ads, 2)
    return {
        "market": market,
        "sku": sku,
        "budget_usd": budget,
        "loop": ["content_collect", "human_like_interact", "schedule_publish", "dm_convert"],
        "allocation": {
            "content_collect": content,
            "interaction": interaction,
            "dm_conversion": dm,
            "paid_boost": ads,
            "risk_reserve": reserve,
        },
        "sla_seconds": 10,
        "guardrails": ["real_sim_card", "native_network", "system_level_isolation"],
    }


def classify_intent(message: str) -> str:
    text = message.lower()
    if any(key in text for key in ("price", "quote", "报价", "价格", "moq")):
        return "quotation"
    if any(key in text for key in ("spec", "规格", "参数", "datasheet")):
        return "specification"
    if any(key in text for key in ("ship", "lead time", "delivery", "交期", "物流")):
        return "delivery"
    return "general"


def search_knowledge(state: dict[str, Any], query: str, limit: int) -> list[dict[str, Any]]:
    tokens = [token.lower() for token in query.replace(",", " ").split() if token.strip()]
    scored: list[tuple[int, dict[str, Any]]] = []
    for item in state.get("knowledge", []):
        haystack = " ".join(
            [
                str(item.get("title", "")),
                str(item.get("body", "")),
                " ".join(item.get("tags", [])),
            ]
        ).lower()
        score = sum(1 for token in tokens if token in haystack)
        if score or not tokens:
            scored.append((score, item))
    scored.sort(key=lambda pair: pair[0], reverse=True)
    return [item for _, item in scored[:limit]]


def comparison() -> dict[str, Any]:
    return {
        "traditional": {
            "staff_per_group": "3-5",
            "sla_seconds": 1800,
            "conversion": "1.0x",
            "monthly_cost_usd": 3000,
        },
        "longxia_mobile": {
            "staff_per_group": "0.1 (1 operator manages 100 phones)",
            "sla_seconds": "<10",
            "conversion": "3.5x-5.0x",
            "monthly_cost_usd": "<150 including hardware",
        },
    }


def build_opc_plan(product: str, market: str, competitor_url: str, channel: str) -> dict[str, Any]:
    return {
        "product": product,
        "market": market,
        "competitor_url": competitor_url,
        "primary_channel": channel,
        "systems": [
            {
                "name": "客户认知系统",
                "modules": ["品牌USP定位工作流", "品牌矩阵内容+SEO", "LinkedIn/Newsletter私域护城河"],
            },
            {
                "name": "成交推进系统",
                "modules": ["深度客户开发", "龙虾AI业务谈判机器人", "标准化跟进体系"],
            },
            {
                "name": "运营/服务系统",
                "modules": ["自动化SOP", "销售飞轮模型", "获客-转化-交付闭环"],
            },
        ],
        "four_step_sop": [
            "OpenClaw部署与anygen-connector/heygen-video-agent技能挂载",
            "AnyGen生成SWOT竞品分析、12页PPT框架和3套60秒营销脚本",
            "HeyGen将脚本转成多语种数字人视频并humanizer去AI味",
            "OpenClaw同步分发到TikTok/Instagram/YouTube并触发私域成交",
        ],
        "install_commands": ["install-skill anygen", "install-skill heygen-poster"],
    }


def generate_lead_tasks(industry: str, market: str, count: int) -> dict[str, Any]:
    channels = ["LinkedIn", "TikTok", "Instagram", "YouTube", "WhatsApp"]
    return {
        "industry": industry,
        "market": market,
        "target_count": count,
        "channels": channels,
        "tasks": [
            f"Search {market} {industry} buyer and competitor keywords on LinkedIn and TikTok.",
            "Score leads by profile activity, buying signal, comment intent and private-domain fit.",
            "Enrich company pain points, decision role, likely SKU interest and preferred contact channel.",
            "Queue high-intent leads for AnyGen negotiation script generation.",
        ],
        "qualification_fields": ["company", "role", "pain_point", "intent_signal", "channel", "next_touch"],
    }


def generate_video_brief(product: str, persona: str, language: str, style: str) -> dict[str, Any]:
    return {
        "product": product,
        "persona": persona,
        "language": language,
        "style": style,
        "anygen_prompt": (
            f"Analyze the strongest selling points for {product}; create one 60-second {language} "
            f"marketing script for {persona}, with a sharp hook, proof, offer and CTA."
        ),
        "heygen_command": (
            f'call heygen --avatar "{persona}" --script AnyGen_Script '
            f'--style "{style}" --output_path ./media/opc/{slugify(product)}'
        ),
        "quality_gates": ["humanized tone", "clear offer", "native-language CTA", "brand-safe visual style"],
    }


def generate_outreach(platform: str, customer: str, offer: str, asset: str) -> dict[str, Any]:
    return {
        "platform": platform,
        "customer": customer,
        "asset": asset,
        "message": (
            f"Hi {customer}, I made a short localized video showing how {offer}. "
            "Worth a quick look? I can also send the pricing/spec PDF if useful."
        ),
        "automation": ["publish_asset", "watch_comments", "detect_price_intent", "handoff_to_whatsapp_or_email"],
        "intent_triggers": ["how much", "price", "MOQ", "spec", "shipping", "sample"],
    }


def generate_followup(customer: str, stage: str, quote_pdf: str) -> dict[str, Any]:
    playbooks = {
        "new": "Confirm pain point and send a concise proof-oriented intro.",
        "interested": "Send private quotation, SKU fit and one clear next action.",
        "quoted": "Handle objections around MOQ, delivery, payment and after-sales.",
        "closing": "Summarize agreed value, deadline and payment path.",
    }
    return {
        "customer": customer,
        "stage": stage,
        "quote_pdf": quote_pdf,
        "next_action": playbooks.get(stage, playbooks["new"]),
        "cadence": ["T+0 instant reply", "T+1 proof/video", "T+3 objection handling", "T+7 final offer"],
        "private_skill_memory": ["negotiation tone", "price boundary", "industry objections", "customer preference"],
    }


def review_opc_run(leads: int, replies: int, deals: int, spend: float) -> dict[str, Any]:
    reply_rate = round(replies / leads, 4) if leads else 0
    deal_rate = round(deals / replies, 4) if replies else 0
    cost_per_deal = round(spend / deals, 2) if deals else 0
    return {
        "leads": leads,
        "replies": replies,
        "deals": deals,
        "spend_usd": spend,
        "reply_rate": reply_rate,
        "deal_rate": deal_rate,
        "cost_per_deal_usd": cost_per_deal,
        "optimization": [
            "Move high-reply hooks into the AnyGen script template.",
            "Package best objections as OpenClaw private skills.",
            "Re-render winning scripts with HeyGen avatars for the next market.",
        ],
    }


def build_pipeline(product: str, market: str, industry: str, lead_count: int) -> dict[str, Any]:
    run_id = "opc-pipe-" + datetime.now(timezone.utc).strftime("%Y%m%d%H%M%S%f")
    return {
        "id": run_id,
        "status": "ready",
        "product": product,
        "market": market,
        "industry": industry,
        "lead_count": lead_count,
        "stages": [
            {"id": "environment", "owner": "OpenClaw", "status": "ready", "action": "mount anygen and heygen skills"},
            {"id": "lead_mining", "owner": "OpenClaw", "status": "queued", "action": f"mine {lead_count} {market} leads"},
            {"id": "anygen_strategy", "owner": "AnyGen", "status": "queued", "action": "generate SWOT, deck outline and scripts"},
            {"id": "heygen_video", "owner": "HeyGen", "status": "queued", "action": "render avatar videos"},
            {"id": "distribution", "owner": "OpenClaw", "status": "queued", "action": "post to TikTok, Instagram and YouTube"},
            {"id": "private_close", "owner": "OpenClaw", "status": "queued", "action": "reply with quotation PDF and WhatsApp/email follow-up"},
            {"id": "review", "owner": "OpenClaw", "status": "queued", "action": "write reply-rate and deal-rate optimization loop"},
        ],
    }


def build_export_markdown(product: str, market: str) -> str:
    return "\n".join(
        [
            "# OPC Deal Flywheel",
            "",
            f"Product: {product}",
            f"Market: {market}",
            "",
            "## Checklist",
            "- [ ] OpenClaw: deploy local runtime and mount private skills",
            "- [ ] AnyGen: generate competitor SWOT, PPT outline and 3 video scripts",
            "- [ ] HeyGen: render multilingual avatar videos",
            "- [ ] OpenClaw: distribute assets to TikTok, Instagram and YouTube",
            "- [ ] OpenClaw: detect price/MOQ/spec intent and send private quotation",
            "- [ ] Review: package winning hooks and objections as private skills",
            "",
        ]
    )


def slugify(value: str) -> str:
    clean = "".join(ch.lower() if ch.isalnum() else "-" for ch in value)
    return "-".join(part for part in clean.split("-") if part) or "asset"
