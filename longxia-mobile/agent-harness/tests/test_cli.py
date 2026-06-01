from click.testing import CliRunner

from cli_anything.longxia_mobile.cli import main


def test_overview_mentions_product():
    result = CliRunner().invoke(main, ["--json-output", "overview"])
    assert result.exit_code == 0
    assert "Longxia AI Mobile Solution" in result.output
    assert "tiktok-loop" in result.output


def test_tiktok_plan_allocates_budget():
    result = CliRunner().invoke(
        main,
        ["--json-output", "tiktok", "plan", "--market", "US", "--sku", "LX-AI-PHONE", "--budget", "1200"],
    )
    assert result.exit_code == 0
    assert '"budget_usd": 1200.0' in result.output
    assert "dm_convert" in result.output


def test_b2b_reply_selects_pdf_attachment(tmp_path):
    state = tmp_path / "state.json"
    result = CliRunner().invoke(
        main,
        [
            "--state",
            str(state),
            "--json-output",
            "b2b",
            "reply",
            "--customer",
            "Berlin AI Studio",
            "--message",
            "Need MOQ, price tiers and specs",
            "--attach",
            "Longxia-Quote.pdf",
        ],
    )
    assert result.exit_code == 0
    assert '"intent": "quotation"' in result.output
    assert "Longxia-Quote.pdf" in result.output


def test_macro_record_and_play(tmp_path):
    state = tmp_path / "state.json"
    runner = CliRunner()
    record = runner.invoke(
        main,
        [
            "--state",
            str(state),
            "macro",
            "record",
            "--name",
            "dm-flow",
            "--step",
            "open TikTok inbox",
            "--step",
            "paste AI reply",
        ],
    )
    assert record.exit_code == 0
    play = runner.invoke(main, ["--state", str(state), "--json-output", "macro", "play", "--name", "dm-flow", "--devices", "3"])
    assert play.exit_code == 0
    assert '"devices": 3' in play.output


def test_opc_plan_persists_four_step_sop(tmp_path):
    state = tmp_path / "state.json"
    result = CliRunner().invoke(
        main,
        [
            "--state",
            str(state),
            "--json-output",
            "opc",
            "plan",
            "--product",
            "AI phone",
            "--market",
            "US",
            "--competitor-url",
            "https://example.com/competitor",
        ],
    )
    assert result.exit_code == 0
    assert "AnyGen" in result.output
    assert "HeyGen" in result.output
    assert "OpenClaw" in result.output
    assert state.exists()


def test_opc_video_generates_heygen_command():
    result = CliRunner().invoke(
        main,
        ["--json-output", "opc", "video", "--product", "Longxia AI Mobile", "--persona", "Founder", "--language", "en"],
    )
    assert result.exit_code == 0
    assert "call heygen" in result.output
    assert "AnyGen_Script" in result.output


def test_opc_followup_and_review():
    runner = CliRunner()
    followup = runner.invoke(
        main,
        ["--json-output", "opc", "followup", "--customer", "Berlin AI Studio", "--stage", "quoted", "--quote-pdf", "quote.pdf"],
    )
    assert followup.exit_code == 0
    assert "Handle objections" in followup.output
    assert "quote.pdf" in followup.output

    review = runner.invoke(main, ["--json-output", "opc", "review", "--leads", "100", "--replies", "25", "--deals", "5", "--spend", "500"])
    assert review.exit_code == 0
    assert '"reply_rate": 0.25' in review.output
    assert '"cost_per_deal_usd": 100.0' in review.output


def test_opc_packages_recommends_tier_by_budget():
    result = CliRunner().invoke(main, ["--json-output", "opc", "packages", "--budget-rmb", "30000"])
    assert result.exit_code == 0
    assert "定制获客系统版" in result.output
    assert '"price_rmb": 29800' in result.output


def test_opc_pipeline_creates_persisted_run(tmp_path):
    state = tmp_path / "state.json"
    result = CliRunner().invoke(
        main,
        [
            "--state",
            str(state),
            "--json-output",
            "opc",
            "pipeline",
            "--product",
            "Longxia AI Mobile",
            "--market",
            "US",
            "--industry",
            "AI hardware",
            "--lead-count",
            "30",
        ],
    )
    assert result.exit_code == 0
    assert '"status": "ready"' in result.output
    assert "anygen_strategy" in result.output
    assert "heygen_video" in result.output
    assert state.read_text(encoding="utf-8").count("opc-pipe-") == 1


def test_opc_export_outputs_markdown_checklist(tmp_path):
    output = tmp_path / "opc.md"
    result = CliRunner().invoke(
        main,
        [
            "opc",
            "export",
            "--product",
            "Longxia AI Mobile",
            "--market",
            "US",
            "--output",
            str(output),
        ],
    )
    assert result.exit_code == 0
    text = output.read_text(encoding="utf-8")
    assert "# OPC Deal Flywheel" in text
    assert "- [ ] AnyGen" in text
    assert "- [ ] HeyGen" in text
