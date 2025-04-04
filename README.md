# BIP39 Seed Phrase Word List Generator

A simple Go application that generates a randomized list of all BIP39 seed words (used for Bitcoin and other cryptocurrency wallets) and displays them in an HTML format for easy reference, printing, or saving as PDF.

## Overview

This tool randomizes the standard 2048 BIP39 seed words used in cryptocurrency wallets and presents them in a clean, grid-based HTML layout. This can be useful for various purposes:

- Creating educational materials about cryptocurrency wallet seed phrases
- Generating reference sheets of all possible seed words
- Testing or educational purposes where you need a complete list of seed words in random order

## Features

- Presents all 2048 BIP39 seed words in a randomized order
- Clean, responsive HTML output with a grid layout
- Print-friendly design (for saving as PDF)
- Time-stamped generation for record-keeping
- Words are numbered for easy reference

## Usage

### Requirements

- Go programming environment (1.16 or higher recommended)

### Running the Application

1. Clone this repository
2. Run the application using Go:

```bash
go run main.go word.go
```

3. The application will generate an HTML file named `scrambled_seed_words.html` in the current directory
4. Open this file in any modern web browser
5. You can print the page or save it as PDF using your browser's print functionality

## File Structure

- `main.go`: Contains the main application logic, HTML template, and functions for shuffling words and generating HTML output
- `word.go`: Contains the BIP39 word list and helper functions for word lookups

## HTML Output

The generated HTML file includes:

- A responsive grid layout that adapts to different screen sizes
- Clear presentation of each word with its assigned number
- Timestamp showing when the list was generated
- Print-friendly styling that automatically adjusts for printing

## Security Considerations

This tool is intended for educational and reference purposes only. Please note:

- **DO NOT** use this tool to generate actual seed phrases for cryptocurrency wallets
- For actual cryptocurrency wallet creation, always use trusted wallet software that generates truly random seed phrases
- The randomization in this tool is for display purposes only and should not be considered cryptographically secure

## License

[MIT License](LICENSE)

## Acknowledgments

- The BIP39 word list is derived from the Bitcoin Improvement Proposal 39, which standardized mnemonic codes for generating deterministic keys

