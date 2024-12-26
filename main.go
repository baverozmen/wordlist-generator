package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	names := getInput(reader, "enter your names (comma-separated, e.g., Ali,Ahmet): ")
	surnames := getInput(reader, "enter your surnames (comma-separated, e.g., Yılmaz,Öztürk): ")
	favNumbers := getInput(reader, "enter your favorite numbers (comma-separated, e.g., 76,30,45): ")
	birthYear := getInput(reader, "enter your birth year (e.g., 1990): ")
	birthMonths := getInput(reader, "enter your birth months (comma-separated, e.g., 1,12): ")
	birthDay := getInput(reader, "enter your birth day (e.g., 15): ")
	colors := getInput(reader, "enter your favorite colors (comma-separated, e.g., blue,red): ")
	usernames := getInput(reader, "enter your usernames (comma-separated): ")
	hobbies := getInput(reader, "enter your hobbies (comma-separated, e.g., basketball,csgo): ")
	filename := getInput(reader, "enter output file name (e.g., wordlist.txt): ")

	combinations := generateCombinations(names, surnames, favNumbers, birthYear, birthMonths, birthDay, colors, usernames, hobbies)

	err := writeToFile(filename, combinations)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	} else {
		fmt.Printf("Wordlist has been saved to '%s'\n", filename)
	}
}

func getInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func generateCombinations(names, surnames, favNumbers string, year string, months string, day string, colors, usernames, hobbies string) []string {
	var combinations []string

	getVariants := func(str string) []string {
		return []string{
			str,
			strings.ToUpper(str),
			strings.ToLower(str),
			strings.ToUpper(string(str[0])) + strings.ToLower(str[1:]),
		}
	}

	specialChars := []string{".", "_", "-", "+", "*", "!", "?"}

	for _, name := range strings.Split(names, ",") {
		for _, surname := range strings.Split(surnames, ",") {
			for _, color := range strings.Split(colors, ",") {
				for _, username := range strings.Split(usernames, ",") {
					nameVariants := getVariants(name)
					surnameVariants := getVariants(surname)
					colorVariants := getVariants(color)
					usernameVariants := getVariants(username)
					hobbyVariants := strings.Split(hobbies, ",")
					favNumVariants := strings.Split(favNumbers, ",")
					monthVariants := strings.Split(months, ",")

					for _, nameVariant := range nameVariants {
						for _, surnameVariant := range surnameVariants {
							for _, colorVariant := range colorVariants {
								for _, usernameVariant := range usernameVariants {
									combinations = append(combinations,
										nameVariant, surnameVariant,
										nameVariant+surnameVariant, surnameVariant+nameVariant,
										nameVariant+colorVariant, surnameVariant+colorVariant,
									)

									for _, specialChar := range specialChars {
										combinations = append(combinations,
											nameVariant+specialChar+surnameVariant,
											surnameVariant+specialChar+nameVariant,
											colorVariant+specialChar+usernameVariant,
											usernameVariant+specialChar+day,
										)
									}

									for _, hobby := range hobbyVariants {
										for _, favNum := range favNumVariants {
											combinations = append(combinations,
												nameVariant+hobby+favNum,
												surnameVariant+hobby+favNum,
												usernameVariant+hobby+favNum,
											)
										}
									}

									for _, month := range monthVariants {
										combinations = append(combinations,
											nameVariant+year+month+day,
											usernameVariant+month+day,
											nameVariant+month+surnameVariant,
											surnameVariant+month+nameVariant,
										)
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return combinations
}

func writeToFile(filename string, data []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range data {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}
