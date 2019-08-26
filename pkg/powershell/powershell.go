package powershell

import (
	"encoding/base64"
	"os/exec"

	"golang.org/x/text/encoding/unicode"
)

func newEncodedPSScript(script string) (string, error) {
	uni := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	encoded, err := uni.NewEncoder().String(script)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString([]byte(encoded)), nil
}

// RunCommand will encoded and run a powershell script.
func RunCommand(script string) error {
	powershellLocation, _ := exec.LookPath("powershell.exe")
	encoded, _ := newEncodedPSScript(script)
	c := exec.Command(powershellLocation, "-nop", "-exec", "bypass", "-EncodedCommand", encoded)
	return c.Run()
}
