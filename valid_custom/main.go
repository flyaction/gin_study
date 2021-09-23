package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"time"
)


type Booking struct {
	CheckIn time.Time `form:"check_in" validate:"required,bookableDate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" validate:"required,gtfield=CheckIn" time_format:"2006-01-02"`

}

// Structure
func bookableDate(fl validator.FieldLevel) bool {

	if date,ok := fl.Field().Interface().(time.Time);ok{
		today := time.Now()
		if date.Unix() > today.Unix(){
			return true;
		}
	}

	return false
}

func main()  {
	r := gin.Default()

	//if v,ok := binding.Validator.Engine().(*validator.Validate);ok{
	//	v.RegisterValidation("bookableDate",bookableDate)
	//}

	validate := validator.New();
	_= validate.RegisterValidation("bookableDate",bookableDate)

	r.GET("/bookable", func(c *gin.Context) {
		var b Booking
		//v10不是这样搞了
		if err := c.ShouldBind(&b);err != nil{
			c.JSON(500,gin.H{"error":err.Error()})
			return
		}

		if err := validate.Struct(b);err != nil{
			c.JSON(500,gin.H{"error2":err.Error()})
			return
		}

		c.JSON(200,gin.H{"message":"ok!","booking":b})
	})

	r.Run()
}