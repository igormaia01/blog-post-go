# Go Blog Engine 📝

Hey there! So I built this simple blog engine in Go because I wanted something lightweight and easy to customize. It reads markdown files and turns them into a nice-looking blog. Pretty neat, right?

## What's this about?

This is a basic blog engine written in Go using the Gin framework. The cool thing is that you just write your posts in markdown files, and the system automatically loads them up. No database needed - just files on disk!

The blog supports:

- Markdown posts with YAML frontmatter
- Clean, responsive design
- Post listings and individual post pages
- Tags and author info
- Images and code syntax highlighting
- Simple navigation

## Project Structure

```
├── cmd/server/          # Main server application (recommended)
├── internal/            # Internal packages
├── web/                 # Web assets and templates
│   ├── templates/       # HTML templates
│   └── static/         # CSS, JS, images
├── posts/              # Your markdown blog posts
├── images/             # Post images
├── main.go             # Simple server (alternative)
└── templates/          # Basic templates (for simple server)
```

## Running it locally

First, make sure you have Go installed (1.23+ recommended).

Clone this repo:

```bash
git clone <your-repo-url>
cd blog
```

Install dependencies:

```bash
go mod tidy
```

Run the server (recommended way):

```bash
go run cmd/server/main.go
```

Or use the simple version:

```bash
go run main.go
```

The blog should be running at `http://localhost:3100` (or `http://localhost:8080` for the simple version).

## Writing posts

Just create a new `.md` file in the `posts/` directory. Here's the format:

````markdown
---
title: 'Your Post Title'
author: 'Your Name'
date: '2025-01-15'
tags: ['go', 'blog', 'programming']
excerpt: 'A short description of your post'
slug: 'your-post-slug'
---

# Your actual content here

Write whatever you want in **markdown**!

## Subheadings work

- Lists work too
- Pretty cool, right?

```go
// Code blocks work great
func main() {
    fmt.Println("Hello, blog!")
}
```
````

````

The system will automatically pick up new posts when you restart the server.

## Making it your own

Want to fork this and build your own blog? Here's what you should probably change:

1. **Update the site info**: Edit the site title and description in the handlers
2. **Customize the design**: Modify the CSS in the template files or add your own stylesheets
3. **Add your content**: Replace the example posts with your own
4. **Configure settings**: Check out `configs/app.env` for server settings

### Customizing templates

The templates are in `web/templates/`:
- `list.tmpl` - Shows all your posts
- `detail.tmpl` - Individual post view
- `index.tmpl` - Homepage
- Admin templates are in `web/templates/admin/`

Just HTML with Go templates - pretty straightforward to modify.

### Adding images

Put your images in the `images/` folder and reference them in your posts like:
```markdown
![Alt text](/images/your-image.png)
````

## Deploying to the internet

Check out the deployment options I mentioned earlier - Railway, Render, Fly.io are all great free options to get started.

For a quick deploy to Railway:

```bash
npm install -g @railway/cli
railway login
railway init
railway up
```

## Contributing

Found a bug? Want to add a feature? PRs are welcome! This is pretty much a personal project, but I'm happy to review contributions.

## License

MIT - do whatever you want with it!

---

That's pretty much it! Let me know if you run into any issues or have questions. Happy blogging! 🚀
