package main

import(
	"fmt"
	"github.com/gin-gonic/gin"
)

func demo()  {
	r := gin.Default()
	r.POST("/uploads", func(ctx *gin.Context) {
		form ,_:= ctx.MultipartForm()
		files := form.File["file"]
		for _,file := range files{
			_ = ctx.SaveUploadedFile(file,fmt.Sprintf("./%s",file.Filename))
		}
		ctx.JSON(200,gin.H{
			"uploads":"ok",
		})
	})
	_= r.Run(":9090")
}