package hyperv

import (
	"net/http"
	"hyper-v-rest-ps/utilities"

	"github.com/gin-gonic/gin"
)

func Processor(c *gin.Context) {
	input := c.Param("machid")

	if input == "" {
		c.Data(returnResponse("No VM ID specified", http.StatusBadRequest, "failure", "error"))
		return
	}

	if input == "all" {
		output, err := utilities.CommandLine(`Get-WmiObject -namespace 'root\virtualization\v2' -class Msvm_SummaryInformation | Select-Object -Property ElementName, InstanceID, NumberOfProcessors | ConvertTo-Json `)
		if err != nil {
			c.Data(returnResponse(err.Error(), http.StatusInternalServerError, "failure", "error"))
			return
		}
		c.Data(returnResponse(output, http.StatusOK, "success", "Processor info is displayed in data field."))
		return
	}

	if !utilities.IsValidUUID(input) {
		c.Data(returnResponse("Invalid VM ID specified", http.StatusBadRequest, "failure", "error"))
		return
	}

	output, err := utilities.CommandLine(`Get-VM -Id ` + input + ` | Get-VMProcessor | ConvertTo-Json`)

	if string(output) == "" {
		c.Data(returnResponse("VM Not Found", http.StatusNotFound, "failure", "error"))
		return
	}

	if err != nil {
		c.Data(returnResponse(err.Error(), http.StatusInternalServerError, "failure", "error"))
		return
	}

	c.Data(returnResponse(output, http.StatusOK, "success", "Processor info is displayed in data field."))
}
