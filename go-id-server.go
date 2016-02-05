package main

import (
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

type IdGenerator struct {
	timer, counter int64
	mutex          sync.Mutex
}

func (idg *IdGenerator) GenerateNewId() int64 {
	if idg.timer != time.Now().Unix() {
		idg.timer = time.Now().Unix()
		idg.counter = 0
	}

	idg.mutex.Lock()
	idg.counter++
	newId := (idg.timer << 32) + idg.counter
	idg.mutex.Unlock()

	return newId
}

var idG IdGenerator

func main() {
	idG = IdGenerator{}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/id", generateId)

	router.Run("127.0.0.1:8080")
}

func generateId(c *gin.Context) {
	c.String(200, "%d\n", idG.GenerateNewId())
}
