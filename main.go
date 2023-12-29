package main

import (
	"github.com/gin-gonic/gin"
)

// type Product struct {
// 	gorm.Model
// 	Code  string
// 	Price uint
// }

func main() {
	router := gin.Default()

	router.GET("/api/user", getUserHandler)
	// router.POST("/user", postUserHandler)

	router.Run(":3000")

	// database, _ := sql.Open("sqlite3", "./data.db")
	// statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS user (username TEXT, password TEXT)")
	// statement.Exec()
}

// type User struct {
// 	Username string `json:"username" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// }

// func postUserHandler(c *gin.Context) {
// 	var userRequest User

// 	err := c.ShouldBindJSON(&userRequest) // parsing json payload
// 	if err != nil {
// 		var errMessages []string
// 		for _, e := range err.(validator.ValidationErrors) {
// 			errMessage := fmt.Sprintf("Error occured on %s, status %s", e.Field(), e.ActualTag())
// 			errMessages = append(errMessages, errMessage)
// 		}
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": errMessages,
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"username": userRequest.Username,
// 		"password": userRequest.Password,
// 	})
// }

func getUserHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hello Jenkins",
	})
}
