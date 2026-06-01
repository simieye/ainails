#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

required=(
  "web/static/index.html"
  "web/static/app.js"
  "web/static/styles.css"
  "web/static/manifest.webmanifest"
  "web/static/sw.js"
  "web/static/icons/icon.svg"
  "apps/desktop-electron/package.json"
  "apps/mobile-capacitor/capacitor.config.json"
)

for path in "${required[@]}"; do
  if [[ ! -f "$path" ]]; then
    echo "missing required packaging file: $path" >&2
    exit 1
  fi
done

node --check web/static/app.js

echo "AI NAILS multi-platform package skeleton is ready."
echo "Web core:       web/static/index.html"
echo "Desktop shell:  apps/desktop-electron"
echo "Mobile shell:   apps/mobile-capacitor"
echo "Docs:           docs/multi-platform-packaging.md"
