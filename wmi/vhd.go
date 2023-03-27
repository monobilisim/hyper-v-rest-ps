package wmi

func Vhd(VMName string) ([]byte, error) {
	ps := `Get-VM -Id ` + VMName + ` | Select-Object -ExpandProperty VMId | Get-VHD | ConvertTo-Json`
	return execPS(ps)
}
