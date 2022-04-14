package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Please specify an argument.")
		fmt.Println("For exsample, mygit [language]")
		return
	}
	lang := args[0]

	// Create git repository
	msg, err := exec.Command("git", "init").CombinedOutput()
	fmt.Println("Create git repository.")
	fmt.Printf("git message:\n%s", msg)

	if err != nil {
		fmt.Printf("git error:\n%s", err)
		return
	}

	if !copyIgnorefile(lang) {
		fmt.Println("Failed to copy ignore file")
		return
	}
	fmt.Println("Repository created.")
}

func copyIgnorefile(lang string) bool {
	exePath, _ := os.Executable()
	exeDirPath := filepath.Dir(exePath)
	filePath := fmt.Sprintf("%s/ignores/%s/.gitignore", exeDirPath, lang)

	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		fmt.Println(".gitignore does not exist.")
		return false
	}

	srcFile, err := os.Open(filePath)
	defer srcFile.Close()
	if err != nil {
		fmt.Println(".gitignore could not open.")
		return false
	}

	dstFile, err := os.OpenFile(".gitignore", os.O_CREATE| os.O_WRONLY, 0644)
	defer dstFile.Close()
	if err != nil {
		fmt.Println(".gitignore could not create.")
		return false
	}

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		fmt.Println(".gitignore could not copy.")
		return false
	}
	return true
}
