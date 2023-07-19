package hyperv

import (
	"net/http"
	"wmi-rest/utilities"

	"github.com/gin-gonic/gin"
)

func Processor(c *gin.Context) {
	input := c.Param("machid")

	if input == "" {
		c.Data(returnResponse("No VM ID specified", http.StatusBadRequest, "failure", "error"))
		return
	}

	if input == "all" {
		output, err := utilities.CommandLine(`Get-WmiObject -Namespace 'root\virtualization\v2' -Class Msvm_SummaryInformation | ConvertTo-Json`)
		if err != nil {
			c.Data(returnResponse(err.Error(), http.StatusInternalServerError, "failure", "error"))
			return
		}
		c.Data(returnResponse(output, http.StatusOK, "success", "Processor info is displayed in data field."))
	}

	output, err := utilities.CommandLine(`Get-VM -Id ` + input + ` | Get-VMProcessor | ConvertTo-Json`)
	if err != nil {
		c.Data(returnResponse(err.Error(), http.StatusInternalServerError, "failure", "error"))
		return
	}

	c.Data(returnResponse(output, http.StatusOK, "success", "Processor info is displayed in data field."))
}
