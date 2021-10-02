package controller

import (
	"Intern-project/database"
	"Intern-project/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//create user
func CreateUser(c *gin.Context) {
	db := database.Connect()

	newUser := user.User{}
	newUser.ID = uuid.New()

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sql := `INSERT INTO user(id,name,age,address) VALUES("` + newUser.ID.String() + `","` + newUser.Name + `","` + fmt.Sprintf("%v", newUser.Age) + `","` + newUser.Address + `");`
	if _, err := db.Query(sql); err != nil {
		c.JSON(200, gin.H{
			"error": err,
		})
	} else {
		c.JSON(200, gin.H{
			"data": newUser,
		})
	}

	defer db.Close()
}

//Get user
func Getuser(c *gin.Context) {
	db := database.Connect()
	defer db.Close()

	s := c.Param("name")

	sql := `SELECT * FROM user WHERE name ="` + s + `"`
	if rs, err := db.Query(sql); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	} else {
		user := user.User{}
		for rs.Next() {
			// for each row, scan the result into our tag composite object
			err = rs.Scan(&user.ID, &user.Name, &user.Age, &user.Address)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
		}
		c.JSON(200, gin.H{
			"data": user,
		})
	}
}

func UpdateUser(c *gin.Context) {
	db := database.Connect()
	defer db.Close()

	user := user.User{}
	s := c.Param("name")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sql := `UPDATE user SET name = "` + user.Name + `",age=` + fmt.Sprintf("%v", user.Age) + `,address="` + user.Address + `" WHERE name = "` + s + `"`
	if _, err := db.Query(sql); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "update success",
		})
	}

}

func DeleteUser(c *gin.Context) {
	db := database.Connect()
	defer db.Close()

	s := c.Param("name")

	sql := `DELETE FROM user WHERE name = "` + s + `"`
	if _, err := db.Query(sql); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "delete success",
		})
	}
}
