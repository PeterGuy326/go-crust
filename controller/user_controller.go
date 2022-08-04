package controller

import (
	"crust/logic"
	"crust/model"
	"crust/utils"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) *utils.Response {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return utils.NewResponse(0, err.Error(), nil)
	}

	var input model.LoginInput
	if err = json.Unmarshal(body, &input); err != nil {
		return utils.NewResponse(0, err.Error(), nil)
	}

	res, err := logic.Login(input)
	if err != nil {
		return utils.NewResponse(0, err.Error(), nil)
	}

	return utils.NewResponse(1, "success", res)
}

func Register(ctx *gin.Context) *utils.Response {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return utils.NewResponse(0, err.Error(), nil)
	}

	var input model.RegisterInput
	if err = json.Unmarshal(body, &input); err != nil {
		return utils.NewResponse(0, err.Error(), nil)
	}

	err = logic.Register(input)
	if err != nil {
		return utils.NewResponse(0, err.Error(), nil)
	}

	return utils.NewResponse(1, "success", 1)
}
