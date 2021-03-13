package router

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
		"data": gin.H{
			"accessToken": "eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJqMmtiQGoya2IuY29tIiwiYXV0aCI6IiIsImV4cCI6MTYxNDAxNTk2Nn0.Yp31VtAyFvfyuZh72Qj_pSF3vYsVr4ZfRrM5Kbk4KAAMJDxIWb0SbYXfY9X1rwTkdTwt5lWn_cRkjldfqZFTrQ",
		},
	})
}
