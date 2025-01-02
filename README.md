# RSS Scraper

A Go-based RSS feed aggregator that fetches and stores content from multiple RSS feeds. Features a REST API for managing feed subscriptions and viewing aggregated posts.

## Features

- Concurrent RSS feed scraping
- User authentication
- PostgreSQL storage
- RESTful API endpoints
- Configurable scraping intervals

## Setup

1. Clone the repository
2. Set environment variables in `.env`:

```env
DATABASE_URL=postgresql://user:password@localhost:5432/dbname
PORT=8000
```

3. Install the deps

```bash
go mod download
```

4. Run the code

```bash
make run
```

## API Endpoints

```plaintext
POST /api/users          # Create user
POST /api/feeds          # Create feed
GET /api/users           # Get user info (auth required)
GET /api/feeds           # Get feed info (auth required)
GET /api/feed-follow     # Get user followed feeds info (auth required)
GET /api/users/posts     # Get user's posts (auth required)
```

## Database Schema

- users: User accounts
- feeds: RSS feed sources
- posts: Scraped content
- feed_follows: user followed feeds

## Configuration

- Scraping interval: Configurable via StartScraping()
- Concurrency: Adjustable worker pool size
- Database connection pool: Managed automatically
