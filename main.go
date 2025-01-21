package main

import (
	"fmt"
	"os"
	"os/exec"
)

func makeCommit(days int) {
	if days < 1 {
		// Push the changes
		cmd := exec.Command("git", "push")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		return
	}

	// Create a commit message with the date
	date := fmt.Sprintf("%d days ago", days)
	data := fmt.Sprintf("%s\n", date)

	// Append the date to a file
	file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(data); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Stage the file
	exec.Command("git", "add", "data.txt").Run()

	// Commit the changes with a specific date
	cmd := exec.Command("git", "commit", "--date", date, "-m", "First Commit")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	// Recursive call for the previous day
	makeCommit(days - 1)
}

func main() {
	days := 365 // Number of days to commit
	makeCommit(days)
}
