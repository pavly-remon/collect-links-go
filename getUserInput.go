package main

import "fmt"

func GetUserInput(msg string) (string, error) {
	var input string
	fmt.Println(msg)
	_, err := fmt.Scan(&input)
	if err != nil {
		return "", err
	}

	return input, nil
}
