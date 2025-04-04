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

This tool uses cryptographically secure random number generation (`crypto/rand`) to shuffle the word list, making it safer than using the standard random number generator. However:

- While the randomization is cryptographically secure, this tool is still primarily intended for educational and reference purposes
- For actual cryptocurrency wallet creation, it's recommended to use established wallet software that follows all BIP39 specifications completely
- This tool only randomizes the display order of all 2048 words, it doesn't generate actual wallet seed phrases (which would be a subset of typically 12 or 24 words)

## License

[MIT License](LICENSE)

## Acknowledgments

- The BIP39 word list is derived from the Bitcoin Improvement Proposal 39, which standardized mnemonic codes for generating deterministic keys
