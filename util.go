package nzym

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

// Fatal prints s.
func Fatal(s string) {
	fmt.Println(s)
	os.Exit(1)
}

// GetConfig returns the unmarshaled contents of config.yml
func GetConfig(nzymrcPath string) map[string]string {
	// Read contents of config file.
	config, err := ioutil.ReadFile(nzymrcPath)
	if err != nil {
		Fatal("nzym: unable to locate \"nzym.yml\".")
	}

	// Unmarshal the config file.
	var nzymAlises map[string]string
	err = yaml.Unmarshal([]byte(config), &nzymAlises)
	if err != nil {
		Fatal("nzym: unable to read \"nzym.yml\".")
	}
	return nzymAlises
}

// WriteStringToFile writes the string s to file
// located at filepath.
func WriteStringToFile(filepath, s string) error {
	fo, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer fo.Close()

	_, err = io.Copy(fo, strings.NewReader(s))
	if err != nil {
		return err
	}
	return nil
}

// Clean removes excessive whitespace.
func Clean(s string) string {
	space := regexp.MustCompile(`\s+`)
	return strings.TrimSpace(space.ReplaceAllString(s, " "))
}

// GetConfirmation tries 3 times to receive yes
// or no for the given prompt.
func GetConfirmation(prompt, retry string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	for x := 0; x < 3; x++ {
		input, _ := reader.ReadString('\n')
		if Clean(input) == "yes" {
			return true
		} else if Clean(input) == "no" {
			return false
		}
		fmt.Printf(retry)
	}
	fmt.Println()
	return false
}
