package sound

import (
	"fmt"
	"os/exec"
	"runtime"
)

func PlayWav(path string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("afplay", path)
	case "windows":
		psCommand := fmt.Sprintf(`(New-Object Media.SoundPlayer '%s').PlaySync()`, path)
		cmd = exec.Command("powershell", "-c", psCommand)
	case "linux":
		cmd = exec.Command("aplay", path)
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}

	return cmd.Run()
}
