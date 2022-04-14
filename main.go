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

	return copyFile(".gitignore", filePath)
}

func copyFile(dstFilePath string, srcFilePath string) bool {

	_, err := os.Stat(srcFilePath)

	if os.IsNotExist(err) {
		fmt.Println("The source file does not exist.")
		return false
	}

	srcFile, err := os.Open(srcFilePath)
	defer srcFile.Close()
	if err != nil {
		fmt.Println("The source file could not open.")
		return false
	}

	dstFile, err := os.OpenFile(dstFilePath, os.O_CREATE|os.O_WRONLY, 0644)
	defer dstFile.Close()
	if err != nil {
		fmt.Println("The destination file could not create.")
		return false
	}

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		fmt.Println("Could not file copy.")
		return false
	}
	return true
}
