package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/divslingerx/sitewatcher/pkg/clearscreen"
	"github.com/divslingerx/sitewatcher/pkg/output"
)

func readURLsFromFile(filePath string) ([]string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Create a scanner and read the file line by line
	scanner := bufio.NewScanner(file)
	var urls []string
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	// If there's an error reading the file, return the error
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return urls, nil
}

func main() {
	// Define the path to the file containing the URLs
	filePath := "urls.txt"

	// Read the URLs from the file
	urls, err := readURLsFromFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Sort the URLs alphabetically
	sort.Strings(urls)

	// wait to ensure the file is read before clearing the screen
	clearscreen.Exec()
	fmt.Println("Checking sites...")

	// Infinite loop
	for {

		// Generate the output for the URLs and print it
		output := output.Generate(urls)
		//Clear the console after every iteration
		//clearscreen.Exec()
		fmt.Print(output)

		// Sleep for 30 seconds before starting the next loop iteration
		time.Sleep(5 * time.Minute)
	}
}
