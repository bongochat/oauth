#!/bin/bash

set -e  # Exit on any error
set -o pipefail

echo "🚀 Starting deployment at $(date)"

PROJECT_DIR="/opt/services/oauth"
SERVICE_DIR="/opt/services"
cd "$PROJECT_DIR" || { echo "❌ Failed to cd into $PROJECT_DIR"; exit 1; }

echo "🔄 Pulling latest changes from Git..."
git pull origin main


cd "$SERVICE_DIR" || { echo "❌ Failed to cd into $SERVICE_DIR"; exit 1; }
echo "📦 Building Project..."
docker compose up -d --build

echo "✅ Deployment complete at $(date)"
