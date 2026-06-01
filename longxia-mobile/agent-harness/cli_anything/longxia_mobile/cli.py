from __future__ import annotations

from pathlib import Path
import json

import click

from .model import (
    CAPABILITIES,
    DEFAULT_STATE_PATH,
    build_opc_plan,
    build_export_markdown,
    build_pipeline,
    classify_intent,
    comparison,
    estimate_tiktok_plan,
    generate_followup,
    generate_lead_tasks,
    generate_outreach,
    generate_video_brief,
    load_state,
    recommend_package,
    review_opc_run,
    save_state,
    search_knowledge,
    utc_now,
)


def emit(payload: object, as_json: bool) -> None:
    if as_json:
        click.echo(json.dumps(payload, ensure_ascii=False, indent=2))
        return
    if isinstance(payload, dict):
        for key, value in payload.items():
            click.echo(f"{key}: {value}")
        return
    click.echo(payload)


@click.group(context_settings={"help_option_names": ["-h", "--help"]})
@click.option("--state", type=click.Path(path_type=Path), default=DEFAULT_STATE_PATH, show_default=True)
@click.option("--json-output", is_flag=True, help="Emit machine-readable JSON.")
@click.pass_context
def main(ctx: click.Context, state: Path, json_output: bool) -> None:
    """Longxia AI Mobile automation CLI generated from the solution PDF."""
    ctx.obj = {"state_path": state, "json": json_output}


@main.command()
@click.pass_context
def overview(ctx: click.Context) -> None:
    """Show product positioning and generated command surface."""
    payload = {
        "product": "龙虾AI手机 / Longxia AI Mobile Solution",
        "positioning": "AI agent phone for cross-border e-commerce and B2B automation",
        "engine": "OpenClaw local gateway, atomic action editing and edge execution",
        "capabilities": [
            {"id": cap.id, "name": cap.name, "summary": cap.summary, "commands": list(cap.commands)}
            for cap in CAPABILITIES
        ],
    }
    emit(payload, ctx.obj["json"])


@main.group()
def tiktok() -> None:
    """TikTok full-loop automation commands."""


@tiktok.command("plan")
@click.option("--market", default="US", show_default=True)
@click.option("--sku", default="LX-AI-PHONE", show_default=True)
@click.option("--budget", type=float, default=1000.0, show_default=True)
@click.pass_context
def tiktok_plan(ctx: click.Context, market: str, sku: str, budget: float) -> None:
    """Generate a budgeted TikTok automation loop."""
    if budget <= 0:
        raise click.ClickException("--budget must be greater than zero")
    emit(estimate_tiktok_plan(market, sku, budget), ctx.obj["json"])


@main.group()
def social() -> None:
    """Social messaging and conversion commands."""


@social.command("reply")
@click.option("--platform", default="WhatsApp", show_default=True)
@click.option("--message", required=True)
@click.option("--customer", default="prospect", show_default=True)
@click.pass_context
def social_reply(ctx: click.Context, platform: str, message: str, customer: str) -> None:
    """Draft a 24/7 AI reply for social notifications."""
    intent = classify_intent(message)
    reply = (
        f"{customer}, thanks for reaching out on {platform}. "
        "Longxia AI Mobile can keep a real-device agent online 24/7, "
        "answer in under 10 seconds and record the conversion signal locally."
    )
    emit({"platform": platform, "customer": customer, "intent": intent, "reply": reply}, ctx.obj["json"])


@social.command("log-conversion")
@click.option("--platform", required=True)
@click.option("--customer", required=True)
@click.option("--value", type=float, default=0.0, show_default=True)
@click.pass_context
def log_conversion(ctx: click.Context, platform: str, customer: str, value: float) -> None:
    """Persist a social conversion event for marketing feedback."""
    state = load_state(ctx.obj["state_path"])
    event = {"platform": platform, "customer": customer, "value_usd": value, "created_at": utc_now()}
    state.setdefault("conversions", []).insert(0, event)
    save_state(ctx.obj["state_path"], state)
    emit(event, ctx.obj["json"])


@main.group()
def b2b() -> None:
    """B2B inquiry assistant commands."""


@b2b.command("reply")
@click.option("--customer", required=True)
@click.option("--message", required=True)
@click.option("--attach", multiple=True, help="Candidate attachment names, such as PDF quotation docs.")
@click.pass_context
def b2b_reply(ctx: click.Context, customer: str, message: str, attach: tuple[str, ...]) -> None:
    """Parse a trade inquiry and draft an expert response."""
    state = load_state(ctx.obj["state_path"])
    intent = classify_intent(message)
    hits = search_knowledge(state, message, 2)
    selected_attachment = next((name for name in attach if name.lower().endswith(".pdf")), "")
    reply = (
        f"Hi {customer}, Longxia AI Mobile supports real-SIM automation, cross-app control, "
        "local knowledge retrieval and native outbound calls. "
        "For MOQ and tier pricing, I can match your target market and send a quotation sheet."
    )
    emit(
        {
            "customer": customer,
            "intent": intent,
            "reply": reply,
            "matched_knowledge": [item["title"] for item in hits],
            "attachment": selected_attachment or "no_pdf_attachment_matched",
        },
        ctx.obj["json"],
    )


@main.group()
def knowledge() -> None:
    """Local multilingual knowledge base commands."""


@knowledge.command("add")
@click.option("--title", required=True)
@click.option("--body", required=True)
@click.option("--locale", default="zh-CN", show_default=True)
@click.option("--tag", multiple=True)
@click.pass_context
def knowledge_add(ctx: click.Context, title: str, body: str, locale: str, tag: tuple[str, ...]) -> None:
    """Add a product, FAQ, after-sales or customer preference note."""
    state = load_state(ctx.obj["state_path"])
    item = {"title": title, "body": body, "locale": locale, "tags": list(tag), "created_at": utc_now()}
    state.setdefault("knowledge", []).insert(0, item)
    save_state(ctx.obj["state_path"], state)
    emit(item, ctx.obj["json"])


@knowledge.command("search")
@click.argument("query")
@click.option("--limit", type=int, default=3, show_default=True)
@click.pass_context
def knowledge_search(ctx: click.Context, query: str, limit: int) -> None:
    """Search product materials, FAQ, after-sales rules and preferences."""
    state = load_state(ctx.obj["state_path"])
    emit({"query": query, "results": search_knowledge(state, query, limit)}, ctx.obj["json"])


@main.group()
def macro() -> None:
    """Screen-recorded automation macro commands."""


@macro.command("record")
@click.option("--name", required=True)
@click.option("--step", multiple=True, required=True, help="Natural-language or coordinate action step.")
@click.pass_context
def macro_record(ctx: click.Context, name: str, step: tuple[str, ...]) -> None:
    """Store a lightweight macro from captured touch actions."""
    state = load_state(ctx.obj["state_path"])
    macro_payload = {"name": name, "steps": list(step), "created_at": utc_now()}
    state.setdefault("macros", {})[name] = macro_payload
    save_state(ctx.obj["state_path"], state)
    emit(macro_payload, ctx.obj["json"])


@macro.command("play")
@click.option("--name", required=True)
@click.option("--devices", type=int, default=1, show_default=True)
@click.pass_context
def macro_play(ctx: click.Context, name: str, devices: int) -> None:
    """Replay a stored macro across one or more AI phones."""
    state = load_state(ctx.obj["state_path"])
    stored = state.get("macros", {}).get(name)
    if not stored:
        raise click.ClickException(f"macro not found: {name}")
    emit({"name": name, "devices": devices, "status": "queued", "steps": stored["steps"]}, ctx.obj["json"])


@main.group()
def call() -> None:
    """Native-line AI outbound call commands."""


@call.command("script")
@click.option("--lead", required=True)
@click.option("--goal", default="qualify purchase intent", show_default=True)
@click.option("--language", default="en", show_default=True)
@click.pass_context
def call_script(ctx: click.Context, lead: str, goal: str, language: str) -> None:
    """Generate a native-phone outbound call script."""
    payload = {
        "lead": lead,
        "language": language,
        "line": "native_sim_dialer",
        "goal": goal,
        "script": [
            f"Open with identity confirmation for {lead}.",
            "Explain that Longxia AI Mobile uses real SIM lines and local AI memory guard.",
            "Ask one qualification question, then pause for full-duplex understanding.",
            "Classify hang-up, interested, follow-up or not-qualified outcome.",
        ],
    }
    emit(payload, ctx.obj["json"])


@main.command("compare")
@click.pass_context
def compare(ctx: click.Context) -> None:
    """Compare traditional staffing against Longxia AI phone automation."""
    emit(comparison(), ctx.obj["json"])


@main.group()
def opc() -> None:
    """AnyGen + HeyGen + OpenClaw automated deal SOP commands."""


@opc.command("plan")
@click.option("--product", required=True)
@click.option("--market", default="Global", show_default=True)
@click.option("--competitor-url", default="", show_default=True)
@click.option("--channel", default="TikTok", show_default=True)
@click.pass_context
def opc_plan(ctx: click.Context, product: str, market: str, competitor_url: str, channel: str) -> None:
    """Create the 4-step OPC deal flywheel plan from demand to close."""
    plan = build_opc_plan(product, market, competitor_url, channel)
    state = load_state(ctx.obj["state_path"])
    state.setdefault("opc_runs", []).insert(0, {"type": "plan", "payload": plan, "created_at": utc_now()})
    save_state(ctx.obj["state_path"], state)
    emit(plan, ctx.obj["json"])


@opc.command("leads")
@click.option("--industry", required=True)
@click.option("--market", default="US", show_default=True)
@click.option("--count", type=int, default=50, show_default=True)
@click.pass_context
def opc_leads(ctx: click.Context, industry: str, market: str, count: int) -> None:
    """Generate deep-customer-development lead tasks."""
    if count <= 0:
        raise click.ClickException("--count must be greater than zero")
    emit(generate_lead_tasks(industry, market, count), ctx.obj["json"])


@opc.command("video")
@click.option("--product", required=True)
@click.option("--persona", default="专属品牌形象", show_default=True)
@click.option("--language", default="en", show_default=True)
@click.option("--style", default="navy business", show_default=True)
@click.pass_context
def opc_video(ctx: click.Context, product: str, persona: str, language: str, style: str) -> None:
    """Prepare AnyGen script prompt and HeyGen render command."""
    emit(generate_video_brief(product, persona, language, style), ctx.obj["json"])


@opc.command("outreach")
@click.option("--platform", default="WhatsApp", show_default=True)
@click.option("--customer", required=True)
@click.option("--offer", required=True)
@click.option("--asset", default="./media/opc/latest.mp4", show_default=True)
@click.pass_context
def opc_outreach(ctx: click.Context, platform: str, customer: str, offer: str, asset: str) -> None:
    """Draft a private-domain outreach message and trigger list."""
    emit(generate_outreach(platform, customer, offer, asset), ctx.obj["json"])


@opc.command("followup")
@click.option("--customer", required=True)
@click.option("--stage", default="new", show_default=True, type=click.Choice(["new", "interested", "quoted", "closing"]))
@click.option("--quote-pdf", default="private-quote.pdf", show_default=True)
@click.pass_context
def opc_followup(ctx: click.Context, customer: str, stage: str, quote_pdf: str) -> None:
    """Generate the standardized follow-up path for one prospect."""
    emit(generate_followup(customer, stage, quote_pdf), ctx.obj["json"])


@opc.command("review")
@click.option("--leads", type=int, required=True)
@click.option("--replies", type=int, required=True)
@click.option("--deals", type=int, required=True)
@click.option("--spend", type=float, default=0.0, show_default=True)
@click.pass_context
def opc_review(ctx: click.Context, leads: int, replies: int, deals: int, spend: float) -> None:
    """Review one OPC flywheel run and suggest optimization actions."""
    if min(leads, replies, deals, spend) < 0:
        raise click.ClickException("metrics cannot be negative")
    emit(review_opc_run(leads, replies, deals, spend), ctx.obj["json"])


@opc.command("packages")
@click.option("--budget-rmb", type=int, default=0, show_default=True)
@click.pass_context
def opc_packages(ctx: click.Context, budget_rmb: int) -> None:
    """List membership packages and recommend the best fit for a budget."""
    if budget_rmb < 0:
        raise click.ClickException("--budget-rmb cannot be negative")
    emit(recommend_package(budget_rmb), ctx.obj["json"])


@opc.command("pipeline")
@click.option("--product", required=True)
@click.option("--market", default="US", show_default=True)
@click.option("--industry", required=True)
@click.option("--lead-count", type=int, default=50, show_default=True)
@click.pass_context
def opc_pipeline(ctx: click.Context, product: str, market: str, industry: str, lead_count: int) -> None:
    """Create and persist an end-to-end OPC pipeline run."""
    if lead_count <= 0:
        raise click.ClickException("--lead-count must be greater than zero")
    pipeline = build_pipeline(product, market, industry, lead_count)
    state = load_state(ctx.obj["state_path"])
    state.setdefault("opc_runs", []).insert(0, {"type": "pipeline", "payload": pipeline, "created_at": utc_now()})
    save_state(ctx.obj["state_path"], state)
    emit(pipeline, ctx.obj["json"])


@opc.command("export")
@click.option("--product", required=True)
@click.option("--market", default="US", show_default=True)
@click.option("--output", type=click.Path(path_type=Path), required=True)
@click.pass_context
def opc_export(ctx: click.Context, product: str, market: str, output: Path) -> None:
    """Export a Markdown OPC execution checklist."""
    markdown = build_export_markdown(product, market)
    output.parent.mkdir(parents=True, exist_ok=True)
    output.write_text(markdown, encoding="utf-8")
    emit({"output": str(output), "bytes": len(markdown.encode("utf-8"))}, ctx.obj["json"])


if __name__ == "__main__":
    main()
