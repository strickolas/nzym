package nzym

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

func getConfig() map[string]string {
	// Read contents of config file.
	config, err := ioutil.ReadFile("nzym.yml")
	if err != nil {
		log.Fatalf("nzym: unable to locate \"nzym.yml\".")
	}

	// Unmarshal the config file.
	var nzymAlises map[string]string
	err = yaml.Unmarshal([]byte(config), &nzymAlises)
	if err != nil {
		log.Fatalf("nzym: unable to read \"nzym.yml\".")
	}
	return nzymAlises
}

func writeStringToFile(filepath, s string) error {
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

func clean(s string) string {
	space := regexp.MustCompile(`\s+`)
	return strings.TrimSpace(space.ReplaceAllString(s, " "))
}

func getConfirmation(prompt, retry string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	for x := 0; x < 3; x++ {
		input, _ := reader.ReadString('\n')
		if clean(input) == "yes" {
			return true
		} else if clean(input) == "no" {
			return false
		}
		fmt.Printf(retry)
	}
	fmt.Println()
	return false
}
