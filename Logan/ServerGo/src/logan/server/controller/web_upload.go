package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"logan/server/env"
	"logan/server/model"
	"logan/server/service"
)

func WebLogUpload(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic:%+v", r)
			ok(c, true)
		}
	}()
	// l := model.NewWebTaskFromJson(c.Request.Body)
	l := &model.WebTask{}
	c.Bind(l)

	fmt.Printf("c:%+v\nl:%+v\n", c, l)
	d := service.WebProtocolDecoder{RsaPrivateKey: env.RsaPrivateKey}
	var err error

	fmt.Printf("key:%+v\nl:%+v\nl.LogArray:%+v\n", d, l, l.LogArray)
	l.Content, err = d.Decode(l.LogArray)
	if err != nil {
		fmt.Println("decode,err:", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	_, err = model.WebTasks.SaveTask(l)
	if err != nil {
		fmt.Println("save_task", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	ok(c, true)
}
