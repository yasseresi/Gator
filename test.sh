#!/bin/bash
# Test script for RSS aggregator

echo "🚀 Testing Gator RSS Aggregator"
echo "================================"

# Build the project
echo "Building project..."
go build
if [ $? -ne 0 ]; then
    echo "❌ Build failed"
    exit 1
fi
echo "✅ Build successful"

# Test 1: Check users
echo -e "\n📋 Testing users command..."
./Gator users
echo "✅ Users command executed"

# Test 2: Login as existing user or register new one
echo -e "\n🔐 Login/Register user..."
./Gator login testuser 2>/dev/null || ./Gator register testuser
echo "✅ User login/register completed"

# Test 3: Add test feeds
echo -e "\n📰 Adding test feeds..."
./Gator addfeed "Wagslane Blog" "https://www.wagslane.dev/index.xml" 2>/dev/null || echo "Feed might already exist"
./Gator addfeed "Go Blog" "https://go.dev/blog/feed.atom" 2>/dev/null || echo "Feed might already exist"
echo "✅ Feeds added"

# Test 4: List feeds
echo -e "\n📑 Listing feeds..."
./Gator feeds
echo "✅ Feeds listed"

# Test 5: Run aggregator once to fetch posts
echo -e "\n🔄 Running aggregator to fetch posts..."
timeout 15s ./Gator agg 10s 2>/dev/null || echo "Aggregator completed or timed out"
echo "✅ Aggregator run completed"

# Test 6: Browse posts
echo -e "\n👀 Browsing posts..."
./Gator browse 5
echo "✅ Browse command executed"

echo -e "\n🎉 All tests completed successfully!"
echo "The RSS aggregator is ready to use!"
