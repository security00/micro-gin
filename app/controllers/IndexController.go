package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (i *IndexController) Index(c *gin.Context) {
	fmt.Println("index ...")
}
