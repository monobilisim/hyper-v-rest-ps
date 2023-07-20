package hyperv

import (
	"net/http"
	"wmi-rest/utilities"

	"github.com/gin-gonic/gin"
)

func VHD(c *gin.Context) {
	input := c.Param("machid")

	if input == "" {
		c.Data(returnResponse("No VM ID specified", http.StatusBadRequest, "failure", "error"))
		return
	}

	if input == "all" {
		output, err := utilities.CommandLine(`Get-VHD | ConvertTo-Json`)
		if err != nil {
			c.Data(returnResponse(err.Error(), http.StatusInternalServerError, "failure", "error"))
			return
		}
		c.Data(returnResponse(output, http.StatusOK, "success", "VHD info is displayed in data field."))
		return
	}

	output, err := utilities.CommandLine(`Get-VHD -Id ` + input + ` | ConvertTo-Json`)
	if err != nil {
		c.Data(returnResponse(err.Error(), http.StatusInternalServerError, "failure", "error"))
		return
	}

	if len(output) < 1 {
		c.Data(returnResponse("No Disk found.", http.StatusOK, "failure", "error"))
		return
	}

	c.Data(returnResponse(output, http.StatusOK, "success", "VHD info is displayed in data field"))
}
