# Go Markdown Blog Engine

A modern, feature-rich blog engine built with Go that supports Markdown content, admin panel, search functionality, and more.

## Features

- **Markdown Support**: Write posts in Markdown with frontmatter metadata
- **Admin Panel**: Web-based interface for managing posts and content
- **Search Functionality**: Full-text search across all posts
- **Tag System**: Organize posts with tags and categories
- **Responsive Design**: Mobile-friendly templates
- **Caching**: In-memory caching for improved performance
- **RSS Feed**: Automatic RSS feed generation
- **Sitemap**: XML sitemap for SEO

## Project Structure

```
blog/
├── cmd/server/          # Main application entry point
├── internal/            # Private application code
│   ├── config/         # Configuration management
│   ├── handlers/       # HTTP request handlers
│   ├── models/         # Data models
│   ├── services/       # Business logic services
│   └── middleware/     # HTTP middleware
├── web/                # Web assets
│   ├── static/         # CSS, JS, images
│   └── templates/      # HTML templates
├── posts/              # Markdown blog posts
├── configs/            # Configuration files
└── docs/               # Documentation
```

## Quick Start

1. **Install Dependencies**

   ```bash
   go mod tidy
   ```

2. **Run the Application**

   ```bash
   go run cmd/server/main.go
   ```

3. **Access the Blog**
   - Blog: http://localhost:3100
   - Admin Panel: http://localhost:3100/admin
   - Default admin credentials: admin/admin123

## Configuration

Copy `configs/app.env` and modify the settings as needed:

```bash
cp configs/app.env .env
```

## Writing Posts

Create new posts in the `posts/` directory using Markdown format with frontmatter:

```markdown
---
title: 'My Blog Post'
author: 'Igor'
date: '2025-01-15'
tags: ['go', 'programming', 'blog']
excerpt: 'This is a sample blog post.'
slug: 'my-blog-post'
---

# My Blog Post

Write your content here in Markdown...
```

## Admin Panel

Access the admin panel at `/admin` to:

- View post statistics
- Manage posts
- Edit post content
- Manage tags
- Configure settings

## API Endpoints

- `GET /` - Redirects to posts list
- `GET /posts` - List all published posts
- `GET /post/:id` - View specific post
- `GET /search?q=query` - Search posts
- `GET /tag/:tag` - Posts by tag
- `GET /tags` - All tags
- `GET /rss.xml` - RSS feed
- `GET /sitemap.xml` - XML sitemap

## Development

### Adding New Features

1. **Models**: Add new data structures in `internal/models/`
2. **Services**: Implement business logic in `internal/services/`
3. **Handlers**: Add HTTP handlers in `internal/handlers/`
4. **Templates**: Create HTML templates in `web/templates/`

### Testing

```bash
go test ./...
```

## License

MIT License
