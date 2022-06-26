# go_basic_gin_framework

# Developed a basic Video API with Golang gin framework

Gin-Gonic library: github.com/gin-gonic/gin

# Go Module Init

go mod init example.com/go_basic_gin_framework


# Run

go run server.go

This is in-memory video API. First we need to create videos to see them on site.

Steps for creating a videos:

1. hit /api/videos POST call.
2. basic auth: {"username": "pranav", "password": "randompass"}
3. see results on /views/videos
