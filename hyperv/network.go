package hyperv

import (
	"net/http"
	"wmi-rest/utilities"

	"github.com/gin-gonic/gin"
)

func Network(c *gin.Context) {
	input := c.Param("machid")

	if input == "" {
		c.Data(returnResponse("No VM ID specified", http.StatusBadRequest, "failure", "error"))
		return
	}
	if input == "all" {
		output, err := utilities.CommandLine(`Get-WmiObject -namespace 'root\virtualization\v2' -class Msvm_GuestNetworkAdapterConfiguration | ConvertTo-Json`)
		if err != nil {
			c.Data(returnResponse(err.Error(), http.StatusInternalServerError, "failure", "error"))
		}
		c.Data(returnResponse(output, http.StatusOK, "success", "Network info is displayed in data field"))
		return
	}
	output, err := utilities.CommandLine(`Get-WmiObject -namespace 'root\virtualization\v2' -class Msvm_GuestNetworkAdapterConfiguration -filter "InstanceID like '%` + input + `%'" | ConvertTo-Json`)

	if err != nil {
		c.Data(returnResponse(err.Error(), http.StatusInternalServerError, "failure", "error"))
		return
	}
	c.Data(returnResponse(output, http.StatusOK, "success", "Network info is displayed in data field"))
}
