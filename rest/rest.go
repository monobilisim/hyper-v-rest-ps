package rest

import (
	"net/http"
	"strconv"
	"wmi-rest/hyperv"

	"github.com/gin-gonic/gin"
)

func StartServer(port int, version string) {
	//gin.SetMode(gin.ReleaseMode)
	//r := gin.New()
	r := gin.Default()
	r.GET("/vms", hyperv.VMS)
	r.GET("/vms/:machid/memory", hyperv.Memory)
	r.GET("/vms/:machid/network", hyperv.Network)
	r.GET("/vms/:machid/processor", hyperv.Processor)
	r.GET("/vms/:machid/vhd", hyperv.VHD)

	r.GET("/version", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Result":  "success",
			"Message": "Application version",
			"Data":    version,
		})
	})
	r.Run(":" + strconv.Itoa(port))
}
