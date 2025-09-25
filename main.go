package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
	"gopkg.in/yaml.v2"
)


type BlogPost struct {
    ID      int      `json:"id"`
    Title   string   `json:"title"`
    Content template.HTML `json:"content"` 
    Excerpt string   `json:"excerpt"`
    Author  string   `json:"author"`
    Date    string   `json:"date"`
    Tags    []string `json:"tags"`
    Slug    string   `json:"slug"`
}

type PostMetadata struct {
    Title   string   `yaml:"title"`
    Author  string   `yaml:"author"`
    Date    string   `yaml:"date"`
    Tags    []string `yaml:"tags"`
    Excerpt string   `yaml:"excerpt"`
    Slug    string   `yaml:"slug"`
}

func loadPostsFromMarkdown() ([]BlogPost, error) {
    var posts []BlogPost
    
    files, err := filepath.Glob("posts/*.md")
    if err != nil {
        return nil, err
    }

    for i, file := range files {
        content, err := os.ReadFile(file)
        if err != nil {
            continue
        }

        parts := strings.Split(string(content), "---")
        if len(parts) < 3 {
            continue
        }

        var metadata PostMetadata
        if err := yaml.Unmarshal([]byte(parts[1]), &metadata); err != nil {
            continue
        }

        markdownContent := strings.Join(parts[2:], "---")
        htmlContent := blackfriday.Run([]byte(markdownContent))

        post := BlogPost{
            ID:      i + 1,
            Title:   metadata.Title,
            Content: template.HTML(htmlContent), 
            Excerpt: metadata.Excerpt,
            Author:  metadata.Author,
            Date:    metadata.Date,
            Tags:    metadata.Tags,
            Slug:    metadata.Slug,
        }

        posts = append(posts, post)
    }

    return posts, nil
}

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*")

    router.Static("/images", "./images")

    blogPosts, err := loadPostsFromMarkdown()
    if err != nil {
        blogPosts = getDefaultPosts()
    }

    router.GET("/", func(c *gin.Context) {
        c.Redirect(http.StatusFound, "/posts")
    })

    router.GET("/posts", func(c *gin.Context) {
        c.HTML(http.StatusOK, "list.tmpl", gin.H{
            "SiteTitle":       "Igor's Blog",
            "SiteDescription": "A personal blog about programming, technology, and life",
            "Posts":           blogPosts,
        })
    })

    router.GET("/post/:id", func(c *gin.Context) {
        idParam := c.Param("id")
        id, err := strconv.Atoi(idParam)
        if err != nil {
            c.HTML(http.StatusNotFound, "detail.tmpl", gin.H{
                "Error": "Invalid post ID",
            })
            return
        }

        var post *BlogPost
        for _, p := range blogPosts {
            if p.ID == id {
                post = &p
                break
            }
        }

        if post == nil {
            c.HTML(http.StatusNotFound, "detail.tmpl", gin.H{
                "Error": "Post not found",
            })
            return
        }

        var relatedPosts []BlogPost
        for _, p := range blogPosts {
            if p.ID != id && len(relatedPosts) < 3 {
                relatedPosts = append(relatedPosts, p)
            }
        }
        fmt.Println(post)
        c.HTML(http.StatusOK, "detail.tmpl", gin.H{
            "SiteTitle":    "Igor's Blog",
            "Post":         post,
            "RelatedPosts": relatedPosts,
        })
    })

    router.Run(":8080")
}

func getDefaultPosts() []BlogPost {
	panic("unimplemented")
}
