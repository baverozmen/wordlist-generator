
# Wordlist Generator

This Go program generates various combinations of personal information such as names, surnames, favorite numbers, birth year, birth months, birth day, favorite colors, usernames, and hobbies. The generated combinations are saved into an output file.

## Features

- Takes inputs for names, surnames, favorite numbers, birth details, colors, usernames, and hobbies.
- Generates combinations with different formatting variations (uppercase, lowercase, capitalized).
- Combines inputs with special characters (e.g., ".", "_", "-", etc.).
- Creates combinations of names, surnames, hobbies, favorite numbers, and more.
- Saves the generated combinations to the specified output file.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/baverozmen/wordlist-generator.git
   cd wordlist-generator
   ```

2. Ensure that Go is installed on your system. You can download Go from [here](https://golang.org/dl/).

3. Install dependencies (if any):
   ```bash
   go mod tidy
   ```

## Usage

1. Run the program:
   ```bash
   go run main.go
   ```

2. The program will prompt you to enter the following details:
   - Names (comma-separated)
   - Surnames (comma-separated)
   - Favorite Numbers (comma-separated)
   - Birth Year
   - Birth Months (comma-separated)
   - Birth Day
   - Favorite Colors (comma-separated)
   - Usernames (comma-separated)
   - Hobbies (comma-separated)
   - Output file name (e.g., wordlist.txt)

3. The program will generate combinations based on the input and save them to the specified output file.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
