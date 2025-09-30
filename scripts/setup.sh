#!/bin/bash

echo "🚀 Setting up Igor's Blog..."
echo ""

# Check if configs/.env exists
if [ -f "configs/.env" ]; then
    echo "✅ configs/.env already exists"
else
    echo "📝 Creating configs/.env from template..."
    cp configs/.env.example configs/.env
    echo "⚠️  Please edit configs/.env and update:"
    echo "   - ADMIN_PASSWORD"
    echo "   - ADMIN_SECRET (run: openssl rand -base64 32)"
    echo "   - BLOG_URL (for production)"
fi

# Check if .env exists in root
if [ -f ".env" ]; then
    echo "✅ .env already exists"
else
    echo "📝 Creating .env from template..."
    cp .env.example .env
    echo "⚠️  Please edit .env and update the same values"
fi

echo ""
# Install dependencies
echo "📦 Installing Go dependencies..."
go mod tidy

# Generate a random secret if openssl is available
if command -v openssl &> /dev/null; then
    SECRET=$(openssl rand -base64 32)
    echo ""
    echo "🔐 Generated random secret for ADMIN_SECRET:"
    echo "   $SECRET"
    echo ""
    echo "   Add this to your configs/.env file!"
fi

echo ""
echo "✨ Setup complete! Next steps:"
echo "   1. Edit configs/.env with your settings"
echo "   2. Run: go run cmd/server/main.go"
echo "   3. Visit: http://localhost:3100"
echo "   4. Admin: http://localhost:3100/admin/login"
echo ""
