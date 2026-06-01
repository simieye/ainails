# Longxia AI Mobile CLI Harness

This CLI was generated from `龙虾手机Longxia AI Mobile Solution.pdf` for CLI-Anything style evaluation.

It models the Longxia AI phone as a local automation control plane for cross-border commerce:

- TikTok full-loop operation planning.
- WhatsApp, Messenger and Instagram 24/7 reply automation.
- B2B inquiry handling with product specs, quotation tiers and PDF attachment matching.
- Multilingual knowledge base search and ingestion.
- Screen-recorded automation macros.
- Native-line outbound call scripts and result classification.
- Cost, SLA and conversion comparison against traditional staffing.
- OPC deal flywheel commands for AnyGen strategy, HeyGen video, OpenClaw distribution, follow-up and review.

## Quick Start

```bash
python3 -m pytest
python3 -m cli_anything.longxia_mobile.cli overview
python3 -m cli_anything.longxia_mobile.cli tiktok plan --market US --sku LX-AI-PHONE --budget 1200
python3 -m cli_anything.longxia_mobile.cli b2b reply --customer "Berlin AI Studio" --message "Need MOQ, price tiers and AI phone specs"
python3 -m cli_anything.longxia_mobile.cli opc plan --product "Longxia AI Mobile" --market US --competitor-url https://example.com
python3 -m cli_anything.longxia_mobile.cli opc video --product "Longxia AI Mobile" --persona Founder --language en
python3 -m cli_anything.longxia_mobile.cli opc review --leads 100 --replies 25 --deals 5 --spend 500
python3 -m cli_anything.longxia_mobile.cli opc packages --budget-rmb 30000
python3 -m cli_anything.longxia_mobile.cli opc pipeline --product "Longxia AI Mobile" --market US --industry "AI hardware" --lead-count 30
python3 -m cli_anything.longxia_mobile.cli opc export --product "Longxia AI Mobile" --market US --output ./opc-checklist.md
```

Installable entry point:

```bash
python3 -m pip install -e .
longxia-mobile overview
```

No GitHub token is required for this generated public harness. A token is only needed when CLI-Anything is asked to generate a CLI from private software repositories.
