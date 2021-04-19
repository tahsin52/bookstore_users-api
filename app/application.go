package app

import "github.com/gin-gonic/gin"

var(
	rooter = gin.Default()
)

func StartApplication()  {
	mapUrls()
	err := rooter.Run(":8080")
	if err != nil {
		return 
	}
}