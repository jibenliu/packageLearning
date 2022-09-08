package validator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type CreateUserHttpBody struct {
	Birthday string `json:"birthday" binding:"required,datetime=01/02"`
	Timezone string `json:"timezone" binding:"omitempty,timezone"`
}

func CreateUser(c *gin.Context) {
	//var httpBody http.Request
	//
	//if err := c.ShouldBindJSON(&httpBody); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	createUser, _ := c.Get("CreateUserHttpBody")
	fmt.Println(createUser)

}

func ValidateJsonBody[BodyType any](fun func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body BodyType

		err := c.ShouldBindJSON(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fun(c)
		c.Set(reflect.ValueOf(body).String(), body)
		c.Next()
	}
}

func GetJsonBody[BodyType any](c *gin.Context) BodyType {
	return c.MustGet("jsonBody").(BodyType)
}

func main() {
	router := gin.Default()
	router.POST("/user", ValidateJsonBody[CreateUserHttpBody](CreateUser))
}
