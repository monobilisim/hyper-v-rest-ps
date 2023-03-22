package wmi

// Get-VM -Id E1D7DB80-97D0-4849-A150-107886C95555 | Select-Object -ExpandProperty VMId | Get-VHD | ConvertTo-Json
// And then parse the JSON output to JSON
func Image(VMName string) ([]byte, error) {
	ps := `Get-VM -Id ` + VMName + ` | Select-Object -ExpandProperty VMId | Get-VHD | ConvertTo-Json`
	return execPS(ps)
}
