// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

// reset environment
func resetEnv() {
	os.Setenv("CGO_ENABLED", "0")
	os.Setenv("GOARCH", "amd64")
	os.Setenv("GOOS", "linux")
}

// A build step that requires additional params, or platform specific steps for example
func Build() error {

	mg.Deps(Clean)

	defer resetEnv()

	os.Setenv("CGO_ENABLED", "0")
	os.Setenv("GOARCH", "amd64")

	fmt.Println("Building for linux")
	os.Setenv("GOOS", "linux")
	cmdLinux := exec.Command("go", "build", "-o", "build/regulus_linux_amd64.bin", ".")
	if err := cmdLinux.Run(); err != nil {
		return err
	}

	fmt.Println("Building for windows")
	os.Setenv("GOOS", "windows")
	cmdWindows := exec.Command("go", "build", "-o", "build/regulus_windows_amd64.exe", ".")
	if err := cmdWindows.Run(); err != nil {
		return err
	}

	fmt.Println("Building for darwin")
	os.Setenv("GOOS", "darwin")
	cmdDarwin := exec.Command("go", "build", "-o", "build/regulus_darwin_amd64.app", ".")
	if err := cmdDarwin.Run(); err != nil {
		return err
	}

	fmt.Println("build finished !")
	return nil
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("build")
}
