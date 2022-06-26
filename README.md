# go_basic_gin_framework

# Developed a basic Video API with Golang gin framework

Gin-Gonic library: github.com/gin-gonic/gin

# Go Module Init

go mod init github.com/pranavwadekar98/go_basic_gin_framework


# Run

go run server.go

This is in-memory video API. First we need to create videos to see them on site.

Steps for creating a videos:

1. hit /api/videos POST call.
2. basic auth: {"username": "pranav", "password": "randompass"}
4. body: {
    "title": "Erics Speed Run",
    "description": "Speed run on chess.com",
    "url": "https://www.youtube.com/embed/e6PmV-MYdP8",
    "author": {
        "name": "pranav",
        "email": "pranav90111@gmail.com"
    }
}

5. see results on /views/videos
