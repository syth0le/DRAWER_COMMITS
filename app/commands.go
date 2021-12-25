package main

import (
	"fmt"
	"os/exec"
)

func runGitStatus() {
	command := "git status"
	execute(command)
	writeToLogFile(command)
	//fmt.Println(command)
}

func runGitAdd() {
	command := "git add ."
	writeToLogFile(command)
	//fmt.Println(command)
}

func runGitCommit(number int) {
	dateTime := getDateTimeCommit("s")
	message := getRandomCommitMessage(number)
	command := "git commit -m " + message + " --date=" + dateTime
	writeToLogFile(command)
	//fmt.Println(command)
}

func runGitPush() {
	command := "git push"
	writeToLogFile(command)
	//fmt.Println(command)
}

func execute(command string) {
	out, err := exec.Command("cmd", "/C", command).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	fmt.Println(output)
}

func runCommands() {
	for globalAccumulator < 10 {
		number := getNumber()
		writeChangesToFile(number)
		runGitStatus()
		runGitAdd()
		runGitCommit(number)
		runGitPush()
	}
	cleanLogFile(false)
}
