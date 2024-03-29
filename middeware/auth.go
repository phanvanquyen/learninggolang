package middeware
import (
"github.com/gin-gonic/gin"
"log"
"time"
)

func IsLogined() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		log.Println("middle authen")
		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}