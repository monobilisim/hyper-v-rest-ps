package hyperv

import (
	"wmi-rest/utilities"

	"github.com/gin-gonic/gin"
)

func Network(c *gin.Context) {
	input := c.Param("machid")
	output, err := utilities.CommandLine(`Get-VM -Id ` + input + ` | Get-VMNetworkAdapter | ConvertTo-Json`)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.Data(200, "application/json", []byte(output))
}
