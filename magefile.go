// +build mage

package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

func Test() error {
	out, err := exec.Command("go", "test", "./...").Output()
	fmt.Println(string(out))
	if err != nil {
		return err
	}
	return nil
}

// SqlMigrateDown execute sql-migrate down
func SqlMigrateDown() error {
	out, err := exec.Command("sql-migrate", "down", "-config=./resources/db/dbconfig.yml").Output()
	fmt.Println(string(out))
	if err != nil {
		return err
	}
	return nil
}

// SqlMigrateUp execute sql-migrate up
func SqlMigrateUp() error {
	out, err := exec.Command("sql-migrate", "up", "-config=./resources/db/dbconfig.yml").Output()
	fmt.Println(string(out))
	if err != nil {
		return err
	}
	return nil
}

// SqlMigrateNew execute sql-migrate new. option is maigration file name.
func SqlMigrateNew() error {

	out, err := exec.Command("sql-migrate", "new", "-config=./resources/db/dbconfig.yml").Output()
	fmt.Println(string(out))
	if err != nil {
		return err
	}
	return nil
}

// DropDataAccessModel drop data access model
func DropDataAccessModel() {
	os.RemoveAll("./app/infrastructures/sqlboiler")
	//os.Mkdir("./app/infrastructures/sqlboiler", 0777)
}

// CreateDataAccessModel create data access model
func CreateDataAccessModel() error {
	out, err := exec.Command("sqlboiler", "--wipe", "--output", "./app/infrastructures/sqlboiler", "--pkgname", "sqlboiler", "psql").Output()
	fmt.Println(string(out))
	if err != nil {
		return err
	}
	return nil
}

// Run execute program
func Run() error {
	// mg.Deps(Build)

	os.Chdir("./build")
	defer os.Chdir("../")

	out, err := exec.Command("./regulus_linux_amd64.bin").Output()
	fmt.Println(string(out))
	if err != nil {
		return err
	}
	return nil

}

// BuildJs build javascript
func BuildJs() error {

	os.Chdir("./frontend")
	defer os.Chdir("../")

	fmt.Println("Building frontend")
	cmd := exec.Command("npx", "ng", "build", "--deploy-url=/resources/")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// A build step that requires additional params, or platform specific steps for example
func Build() error {

	mg.Deps(Clean)
	mg.Deps(BuildJs)

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

	fmt.Println("copy start")
	// resourcesコピー
	separator := getSeparator()
	rootDir := "." + separator + "resources"
	targetDir := "." + separator + "build" + separator + "resources"
	copy(rootDir, targetDir)

	// frontendコピー
	rootDir = "." + separator + "frontend" + separator + "dist"
	targetDir = "." + separator + "build" + separator + "resources"
	copy(rootDir, targetDir)

	fmt.Println("build finished !")

	return nil
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("build")
}

func resetEnv() {
	os.Setenv("CGO_ENABLED", "0")
	os.Setenv("GOARCH", "amd64")
	os.Setenv("GOOS", "linux")
}

func getSeparator() string {
	return string(os.PathSeparator)
}

func copy(rootDir string, targetDir string) error {

	separator := getSeparator()
	files := []string{}

	// ディレクトリ内のファイルを取得
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		rel, err := filepath.Rel(rootDir, path)
		files = append(files, rel)
		return nil
	})

	// ディレクトリを作成
	for _, file := range files {
		if fInfo, _ := os.Stat(rootDir + separator + file); fInfo.IsDir() == true {
			os.MkdirAll(targetDir+separator+file, 0777)
		}
	}

	// コピー実行
	for _, file := range files {
		if fInfo, _ := os.Stat(rootDir + separator + file); fInfo.IsDir() == false {
			src, err := os.Open(rootDir + separator + file)
			if err != nil {
				panic(err)
			}

			dst, err := os.Create(targetDir + separator + file)
			if err != nil {
				panic(err)
			}

			_, err = io.Copy(dst, src)
			if err != nil {
				panic(err)
			}
			src.Close()
			dst.Close()
		}
	}

	return err
}
