package tester

import (
	"os/exec"
	"log"
	"fmt"
	"strings"
	"bytes"
)

func CalculatePRI(script string, optimalPath string, myPath string) {

	cmd := exec.Command("python3", script, optimalPath, myPath)

	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Println(err)
	}
	for _, str := range strings.Split(out.String(), "\n") {
		fmt.Printf("%s\n", str  )
	}
}
