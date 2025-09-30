# Blog Configuration Guide

This document explains how to configure your blog using environment variables.

## Quick Setup

1. **Copy the example environment file:**
   ```bash
   cp configs/.env.example configs/app.env
   ```

2. **Edit `configs/app.env` with your settings:**
   ```bash
   nano configs/app.env  # or use your preferred editor
   ```

3. **Generate a secure admin secret:**
   ```bash
   openssl rand -base64 32
   ```

## Required Configuration

### Blog Identity

These settings define your blog's identity and should be customized:

```env
# Your blog title (appears in header, meta tags, RSS feed)
BLOG_TITLE=My Awesome Blog

# Your name or the blog author name
BLOG_AUTHOR=Your Name

# Blog description for SEO and RSS
BLOG_DESCRIPTION=A blog about technology and programming

# Your blog URL (important for RSS and sitemap)
BLOG_URL=http://localhost:3100  # Change for production
```

### Admin Access

**⚠️ IMPORTANT: Change these before deploying to production!**

```env
# Admin username
ADMIN_USERNAME=admin

# Admin password - use a strong password!
ADMIN_PASSWORD=change-this-to-a-strong-password

# Secret key for session tokens - use output from: openssl rand -base64 32
ADMIN_SECRET=your-generated-secret-key-here
```

## How Blog Information is Used

### Blog Title
- Appears in the website header
- Used in meta tags for SEO
- Shown in RSS feed title
- Displayed in browser tab

### Blog Author
- Default author name for posts (can be overridden per post)
- Available in templates via context
- Used in RSS feed and metadata

### Blog Description
- Used in meta description tags
- Appears in RSS feed description
- Helps with SEO

### Blog URL
- Used to generate absolute URLs in RSS feed
- Used in sitemap generation
- Important for proper link generation

## Writing Posts with Author Information

Each post can have its own author, or use the default:

```markdown
---
title: 'My Post Title'
author: 'Blog Author'  # Use default from BLOG_AUTHOR env var
date: '2025-09-30'
tags: ['technology']
excerpt: 'Post excerpt'
slug: 'my-post-title'
status: 'published'
---

Post content here...
```

You can override the author per post:

```markdown
---
title: 'Guest Post'
author: 'Guest Writer Name'  # Override default author
# ... rest of frontmatter
---
```

## Environment File Locations

The application looks for environment files in this order:

1. `.env` (root directory)
2. `configs/.env` or `configs/app.env`
3. `../configs/.env` (when running from cmd/)
4. `../../configs/.env` (when running from cmd/server/)

## Production Deployment Checklist

Before deploying to production:

- [ ] Change `BLOG_TITLE` to your actual blog name
- [ ] Change `BLOG_AUTHOR` to your name
- [ ] Update `BLOG_URL` to your production domain
- [ ] Set a strong `ADMIN_PASSWORD`
- [ ] Generate and set a random `ADMIN_SECRET`
- [ ] Update `ADMIN_USERNAME` to something other than "admin"
- [ ] Review and adjust other settings as needed

## Additional Configuration

See `configs/.env.example` for all available configuration options including:

- Server settings (host, port, timeouts)
- Cache configuration
- Logging options
- Feature flags
- Database settings (for future use)

## Getting Help

If you encounter issues with configuration:

1. Check that your `.env` file is in the correct location
2. Verify the syntax of your environment variables
3. Check the server logs for configuration warnings
4. Refer to the example file for correct format

## Security Notes

- Never commit your `configs/app.env` or `.env` files to version control
- Use strong, unique passwords for admin access
- Generate random secrets using cryptographically secure methods
- Keep your environment files backed up securely
- Rotate secrets periodically
