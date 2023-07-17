package hyperv

import (
	"wmi-rest/utilities"

	"github.com/gin-gonic/gin"
)

func VHD(c *gin.Context) {
	input := c.Param("machid")
	output, err := utilities.CommandLine(`Get-VHD -Id ` + input + ` | ConvertTo-Json`)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.Data(200, "application/json", output)
}
