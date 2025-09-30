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

   Edit `configs/app.env` to set your admin credentials:

   ```env
   ADMIN_USERNAME=admin
   ADMIN_PASSWORD=your_secure_password
   ADMIN_SECRET=your-secret-key-change-this-in-production
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

### Running Tests

```bash
go test ./...
```

### Testing the Dashboard

Use the provided test script:

```bash
bash tests/test_dashboard.sh
```

## TODO - Future Features

### 🔐 Authentication & User Management

- [ ] **Multi-User Support**
  - Add user registration system
  - Implement role-based access control (Admin, Editor, Author, Viewer)
  - User profile management
  - Password reset functionality via email

- [ ] **Enhanced Security**
  - Two-factor authentication (2FA)
  - Rate limiting on login attempts
  - Password strength requirements
  - Account lockout after failed attempts
  - OAuth integration (Google, GitHub)

### 📝 Content Management

- [ ] **Post Editor**
  - Web-based Markdown editor with live preview
  - Image upload and management
  - Draft auto-save
  - Post scheduling (publish at specific date/time)
  - Post versioning/revision history
  - Bulk post operations (delete, change status)

- [ ] **Media Library**
  - Centralized image/file management
  - Image optimization and resizing
  - CDN integration
  - File organization by folders

- [ ] **Categories & Tags**
  - CRUD operations for categories
  - Tag management interface
  - Hierarchical categories
  - Tag suggestions based on content

### 💬 Comments System

- [ ] **Comment Management**
  - Enable/disable comments per post
  - Comment moderation (approve/reject/spam)
  - Nested/threaded comments
  - Comment notifications
  - Anti-spam measures (CAPTCHA, Akismet)
  - Guest commenting with email verification

### 📊 Analytics & Metrics

- [ ] **Enhanced Dashboard**
  - Interactive charts (views over time, shares trends)
  - Top performing posts
  - Visitor analytics (unique vs returning)
  - Traffic sources
  - Geographic distribution of visitors
  - Export reports (CSV, PDF)

- [ ] **Database Persistence**
  - Migrate metrics from memory to database (PostgreSQL)
  - Historical data retention
  - Aggregate statistics by time periods
  - Real-time analytics

### 🎨 Customization

- [ ] **Theme System**
  - Multiple theme support
  - Theme customization interface
  - Dark mode toggle
  - Custom CSS injection
  - Logo and favicon upload

- [ ] **Settings Panel**
  - Site configuration (title, description, URL)
  - SEO settings (meta tags, Open Graph)
  - Social media links
  - Email notifications configuration
  - Backup and restore functionality

### 🔔 Notifications

- [ ] **Email Notifications**
  - New comment notifications
  - Weekly/monthly analytics reports
  - Post publication confirmations
  - System alerts

- [ ] **In-App Notifications**
  - Real-time notification system
  - Notification preferences
  - Notification history

### 🔍 SEO & Performance

- [ ] **SEO Enhancements**
  - Automatic meta tag generation
  - Schema.org markup
  - Social media preview optimization
  - Canonical URLs
  - Breadcrumb navigation

- [ ] **Performance Optimization**
  - Redis caching integration
  - CDN integration
  - Image lazy loading
  - Asset minification and bundling
  - HTTP/2 and compression

### 📱 Mobile & PWA

- [ ] **Progressive Web App**
  - Service worker for offline access
  - Push notifications
  - Install prompt
  - Mobile app-like experience

### 🔄 Import/Export

- [ ] **Content Migration**
  - Import from WordPress
  - Import from Medium
  - Export to JSON/XML
  - Backup scheduling

### 🌐 Internationalization

- [ ] **Multi-Language Support**
  - i18n framework integration
  - Language switcher
  - Translated content management
  - RTL language support

### 🤖 Automation

- [ ] **Automated Tasks**
  - Scheduled post publishing
  - Automatic sitemap generation
  - Automatic backup creation
  - Email digest generation
  - Social media auto-posting

### 🔗 Integrations

- [ ] **Third-Party Services**
  - Newsletter integration (Mailchimp, SendGrid)
  - Social media auto-sharing
  - Analytics integration (Google Analytics, Plausible)
  - Search integration (Algolia, Elasticsearch)
  - Payment gateway for premium content

### 🧪 Testing & Quality

- [ ] **Testing Infrastructure**
  - Unit tests for all services
  - Integration tests
  - E2E tests with Cypress/Selenium
  - Load testing
  - Code coverage reports

### 📚 Documentation

- [ ] **User Documentation**
  - User guide for content creators
  - Video tutorials
  - FAQ section
  - Troubleshooting guide

- [ ] **Developer Documentation**
  - API documentation (Swagger/OpenAPI)
  - Architecture diagrams
  - Contribution guidelines
  - Plugin/extension system documentation

### 🚀 Deployment & DevOps

- [ ] **Deployment Options**
  - Docker containerization
  - Docker Compose for local development
  - Kubernetes deployment manifests
  - CI/CD pipeline (GitHub Actions)
  - Automated deployment scripts

- [ ] **Monitoring & Logging**
  - Application monitoring (Prometheus)
  - Error tracking (Sentry)
  - Centralized logging (ELK stack)
  - Health check endpoints
  - Performance profiling

### 🎯 Priority Roadmap

**Phase 1 (Core Functionality)** - *Next 2-3 months*
- [ ] Post editor with live preview
- [ ] Database persistence for metrics
- [ ] Enhanced dashboard with charts
- [ ] Comment system

**Phase 2 (User Experience)** - *3-6 months*
- [ ] Multi-user support with roles
- [ ] Theme customization
- [ ] Media library
- [ ] Email notifications

**Phase 3 (Advanced Features)** - *6-12 months*
- [ ] Two-factor authentication
- [ ] Analytics integration
- [ ] SEO enhancements
- [ ] PWA support

**Phase 4 (Scalability)** - *12+ months*
- [ ] Kubernetes deployment
- [ ] CDN integration
- [ ] Plugin system
- [ ] API documentation

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License
