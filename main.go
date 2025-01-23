package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "os"
    "io"
)

func getPingHandler(context *gin.Context) {
    context.JSON(
        http.StatusOK,
        gin.H {
            "message": "pong",
        },
    )
}

func main() {
    file, _ := os.Create("gin.log")

    gin.DisableConsoleColor()
    gin.DefaultWriter = io.MultiWriter(file)

    engine := gin.Default()
    engine.GET("/ping", getPingHandler)
    engine.Run()
}
