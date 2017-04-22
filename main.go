package main 

import(
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

func main(){
	port := os.Getenv("PORT")

	if port == "" {
		port = "8081"
	}

	router := gin.Default()

	// This handler will match /user/john but will not match neither /user/ or /user
	router.GET("/codebreaker/setup/:number", func(c *gin.Context) {
		number := c.Param("number")
		setSecret(number)
		c.String(http.StatusOK, "Secret number configured %s", number)
	})

	router.GET("/codebreaker/guess/:number", func(c *gin.Context) {
		number := c.Param("number")
		result := codebreaker(number)
		c.String(http.StatusOK, "Answer: " + result)
	})

	router.Run(":"+port)
}
