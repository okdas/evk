package main

import (
	"errors"
	"os/exec"
	"strings"
)

func getEnvListByBucket(bucket string) ([]string, error) {
	security := exec.Command("security", "dump-keychain")
	grep := exec.Command("grep", "-E", "acct.*evk_"+bucket)
	cut := exec.Command("cut", "-d", "=", "-f2")
	sed := exec.Command("sed", "s/\"//g;s/evk_"+bucket+"_//")
	output, err := pipedCommands(security, grep, cut, sed)
	if err != nil {
		return []string{}, errors.New("Can't execute environment variables list search.")
	}

	result := strings.Split(string(output), "\n")
	if result[len(result)-1] == "" {
		result = result[:len(result)-1]
	}

	return result, nil
}

func pipedCommands(commands ...*exec.Cmd) ([]byte, error) {
	for i, command := range commands[:len(commands)-1] {
		out, err := command.StdoutPipe()
		if err != nil {
			return nil, err
		}
		command.Start()
		commands[i+1].Stdin = out
	}
	final, err := commands[len(commands)-1].Output()
	if err != nil {
		return nil, err
	}
	return final, nil
}
