# Gator RSS Aggregator

A command-line RSS aggregator built in Go that allows you to manage RSS feeds, follow feeds, and browse posts from your terminal.

## Prerequisites

Before running Gator, you'll need to have the following installed:

### Required Software

- **Go** (version 1.19 or later)

  - Download from [golang.org](https://golang.org/download/)
  - Verify installation: `go version`

- **PostgreSQL**
  - Download from [postgresql.org](https://www.postgresql.org/download/)
  - Make sure PostgreSQL is running and accessible
  - You'll need a database created for Gator

## Installation

Install the Gator CLI using Go:

```bash
go install github.com/yasseresi/Gator@latest
```

Make sure your `$GOPATH/bin` (or `$GOBIN`) is in your `$PATH` to run the `Gator` command from anywhere.

## Configuration

### Database Setup

1. Create a PostgreSQL database for Gator
2. Run the database migrations (you may need to use a migration tool like Goose)

### Config File Setup

Create a configuration file at `~/.gatorconfig.json` with the following structure:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

Replace the database URL with your PostgreSQL connection details:

- `username`: Your PostgreSQL username
- `password`: Your PostgreSQL password
- `localhost:5432`: Your PostgreSQL host and port
- `gator`: Your database name

The `current_user_name` field will be automatically managed by the application.

## Usage

### Available Commands

#### User Management

- **`register <username>`** - Register a new user

  ```bash
  Gator register john_doe
  ```

- **`login <username>`** - Login as an existing user

  ```bash
  Gator login john_doe
  ```

- **`users`** - List all registered users and show the current user

  ```bash
  Gator users
  ```

- **`reset`** - Reset the database (removes all users, feeds, and posts)
  ```bash
  Gator reset
  ```

#### Feed Management

- **`addfeed <name> <url>`** - Add a new RSS feed (requires login)

  ```bash
  Gator addfeed "Tech News" "https://example.com/feed.xml"
  ```

- **`feeds`** - List all available feeds in the system

  ```bash
  Gator feeds
  ```

- **`follow <feed_name>`** - Follow an existing feed (requires login)

  ```bash
  Gator follow "Tech News"
  ```

- **`following`** - List all feeds you're currently following (requires login)

  ```bash
  Gator following
  ```

- **`unfollow <feed_name>`** - Unfollow a feed (requires login)
  ```bash
  Gator unfollow "Tech News"
  ```

#### Content Management

- **`agg <duration>`** - Run the RSS aggregator to fetch new posts from all feeds

  ```bash
  Gator agg 30s    # Fetch every 30 seconds
  Gator agg 5m     # Fetch every 5 minutes
  Gator agg 1h     # Fetch every 1 hour
  ```

- **`browse [limit]`** - Browse posts from feeds you follow (requires login)
  ```bash
  Gator browse      # Show 2 posts (default)
  Gator browse 10   # Show 10 posts
  ```

## Getting Started

1. **Install and setup** (see Installation and Configuration sections above)

2. **Register a user:**

   ```bash
   Gator register your_username
   ```

3. **Add some feeds:**

   ```bash
   Gator addfeed "Go Blog" "https://go.dev/blog/feed.atom"
   Gator addfeed "Hacker News" "https://hnrss.org/newest"
   ```

4. **Follow the feeds:**

   ```bash
   Gator follow "Go Blog"
   Gator follow "Hacker News"
   ```

5. **Run the aggregator to fetch posts:**

   ```bash
   Gator agg 1m
   ```

   (This will run continuously. Stop it with Ctrl+C after some posts are fetched)

6. **Browse your posts:**
   ```bash
   Gator browse 5
   ```

## Features

- **Multi-user support**: Multiple users can use the same Gator instance
- **Feed management**: Add, follow, and unfollow RSS feeds
- **Automatic aggregation**: Continuously fetch new posts from RSS feeds
- **Post browsing**: View posts with titles, descriptions, URLs, and publication dates
- **Duplicate handling**: Automatically handles duplicate posts
- **Multiple date formats**: Supports various RSS date formats

## Development

To build from source:

```bash
git clone https://github.com/yasseresi/Gator.git
cd Gator
go build
./Gator --help
```

## Database Schema

The application uses PostgreSQL with the following main tables:

- `users` - User accounts
- `feeds` - RSS feed information
- `feed_follows` - User-feed relationships
- `posts` - Individual RSS posts
