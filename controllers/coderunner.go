package controllers

import (
	"GooooooooG/model"
	"GooooooooG/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type CodeRunnerController struct {
	beego.Controller
}

// @router /code post
func (code *CodeRunnerController) Post() {

	var ob model.CodeInfo
	_ = json.Unmarshal(code.Ctx.Input.RequestBody, &ob)
	codeparam := &model.CodeInfo{FilePath: ob.FilePath, OutPath: ob.OutPath, FileName: ob.FileName, TestSet: ob.TestSet, LimiteTime: ob.LimiteTime}
	logs.Info("file path ", codeparam.LimiteTime)
	logs.Info("data is ", codeparam.FilePath)
	message := services.JudgeCode(codeparam.FilePath, codeparam.OutPath, codeparam.FileName, codeparam.TestSet, codeparam.LimiteTime)
	code.Data["json"] = message
	code.ServeJSON()
}
