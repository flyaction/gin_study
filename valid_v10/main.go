package main

import (
	"github.com/gin-gonic/gin"
	en2 "github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"

	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type Person struct {
	Age int `form:"age" validate:"required,gt=18"` // required|gt=10 满足其中一个
	Name string `form:"name" validate:"required"`
	Address string `form:"address" validate:"required"`
}


var (
	Uni *ut.UniversalTranslator
	Validate *validator.Validate
)

func main()  {

	Validate = validator.New()

	zh := zh2.New()
	en := en2.New()
	Uni = ut.New(zh,en)


	r := gin.Default()

	r.GET("/testing", func(c *gin.Context) {

		locale := c.DefaultQuery("locale","zh")
		trans,_ := Uni.GetTranslator(locale)

		switch locale {
		case "zh":
			_= zh_translations.RegisterDefaultTranslations(Validate,trans) //
		case "en":
			_= en_translations.RegisterDefaultTranslations(Validate,trans)
		default:
			_= zh_translations.RegisterDefaultTranslations(Validate,trans)
		}

		var person Person
		if err := c.ShouldBind(&person);err != nil{
			c.String(500,"%v",err)
			c.Abort()
			return
		}

		if err := Validate.Struct(person);err != nil{
			errs := err.(validator.ValidationErrors)
			sliceErrs := []string{}
			for _,e := range errs{
				sliceErrs = append(sliceErrs,e.Translate(trans))
			}
			c.String(500,"%v",sliceErrs)
			c.Abort()
			return
		}

		c.String(200,"%v",person)
	})

	r.Run()

}
