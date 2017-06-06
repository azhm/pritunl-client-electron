package main

import (
	"github.com/pritunl/pritunl-client-electron/service/command"
	"os"
	"path/filepath"
)

func main() {
	rootDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	cmd := command.Command("taskkill.exe", "/F", "/IM", "pritunl.exe")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	cmd = command.Command(filepath.Join(rootDir, "nssm.exe"),
		"remove", "pritunl", "confirm")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	cmd = command.Command(filepath.Join(rootDir, "nssm.exe"),
		"stop", "pritunl")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	cmd = command.Command("taskkill.exe", "/F", "/IM", "openvpn.exe")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	cmd = command.Command(filepath.Join(rootDir, "tuntap", "tuntap.exe"),
		"uninstall")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
