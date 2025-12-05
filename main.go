package main

import (
	"github.com/gin-gonic/gin"
	"log" // <--- Добавили пакет для логирования
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from Go!",
		})
	})

	// ИСПРАВЛЕНИЕ:
	// Мы пытаемся запустить сервер.
	// Если err != nil (ошибка есть), мы выводим её и аварийно завершаем программу.
	if err := r.Run(":8080"); err != nil {

		log.Fatal("Fatal error, failed to start server: ", err)
	}
}
