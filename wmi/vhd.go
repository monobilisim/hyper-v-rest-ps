package wmi

func Vhd(VMId string) ([]byte, error) {
	ps := `Get-VHD -Id ` + VMId + ` | ConvertTo-Json`
	output, err := execPS(ps)
	return output, err
}
