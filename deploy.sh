#!/bin/bash

set -e  # Exit on any error
set -o pipefail

echo "🚀 Starting deployment at $(date)"

PROJECT_DIR="/opt/oauth"

cd "$PROJECT_DIR" || { echo "❌ Failed to cd into $PROJECT_DIR"; exit 1; }

echo "🔄 Pulling latest changes from Git..."
git pull origin main

echo "📦 Building Project..."
go build cmd/main.go

echo "♻️ Restarting the oauth service..."
sudo systemctl restart oauth.service

echo "✅ Deployment complete at $(date)"
