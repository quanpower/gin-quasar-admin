package public_api

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"github.com/Junvary/gqa-plugin-xk/model"
	"github.com/Junvary/gqa-plugin-xk/service/public_service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetHonour(c *gin.Context)  {
	var getHonourList model.RequestHonourList
	if err := c.ShouldBindJSON(&getHonourList); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, honour, total := public_service.GetHonour(getHonourList); err!=nil{
		global.GqaLog.Error("获取荣誉认证列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取荣誉认证列表失败！"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:  honour,
			Page:     getHonourList.Page,
			PageSize: getHonourList.PageSize,
			Total:    total,
		}, c)
	}
}
