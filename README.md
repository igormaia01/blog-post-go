# Go Markdown Blog Engine

A modern, feature-rich blog engine built with Go that supports Markdown content, admin panel, search functionality, and more.

## Features

- **Markdown Support**: Write posts in Markdown with frontmatter metadata
- **Admin Dashboard**: View post statistics and metrics
- **Authentication**: Secure admin panel with session-based authentication
- **Metrics Tracking**: Track views and social shares per post
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

2. **Configure Environment**

   Copy the example environment file and customize it:

   ```bash
   cp configs/.env.example configs/app.env
   ```

   Edit `configs/app.env` to set your blog and admin configuration:

   ```env
   # Blog Configuration - Customize these values
   BLOG_TITLE=My Blog
   BLOG_AUTHOR=Your Name
   BLOG_DESCRIPTION=Your blog description
   BLOG_URL=http://localhost:3100

   # Admin Configuration - IMPORTANT: Change these!
   ADMIN_USERNAME=admin
   ADMIN_PASSWORD=your_secure_password
   ADMIN_SECRET=your-secret-key-change-this-in-production
   ```

   Generate a secure admin secret:

   ```bash
   openssl rand -base64 32
   ```

3. **Run the Application**

   ```bash
   go run cmd/server/main.go
   ```

4. **Access the Blog**
   - Blog: http://localhost:3100
   - Admin Panel: http://localhost:3100/admin/login
   - Default admin credentials: admin/admin123 (change these!)

## Writing Posts

Create new posts in the `posts/` directory using Markdown format with frontmatter:

```markdown
---
title: 'My Blog Post'
author: 'Blog Author' # You can use the BLOG_AUTHOR from your .env or specify per post
date: '2025-01-15'
tags: ['go', 'programming', 'blog']
excerpt: 'This is a sample blog post.'
slug: 'my-blog-post'
status: 'published' # or 'draft'
---

# My Blog Post

Write your content here in Markdown...
```

The `author` field in each post can be customized per post, or you can use the default author name set in your environment configuration (`BLOG_AUTHOR`).

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

### Running Tests

```bash
go test ./...
```

### Testing the Dashboard

Use the provided test script:

```bash
bash tests/test_dashboard.sh
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License
