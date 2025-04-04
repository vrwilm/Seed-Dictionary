package main

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"math/big"
	"os"
	"time"
)

// generateHTML creates an HTML file with the scrambled word list
func generateHTML(entries []struct {
	Position int
	Word     string
}, filename string,
) error {
	// Parse the HTML template
	tmpl, err := template.New("wordlist").Parse(htmlTemplate)
	if err != nil {
		return err
	}

	// Create the output file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Current timestamp
	timestamp := time.Now().Format("January 2, 2006 15:04:05")

	// Data for the template
	data := struct {
		Words []struct {
			Position int
			Word     string
		}
		Timestamp string
	}{
		Words:     entries,
		Timestamp: timestamp,
	}

	// Execute the template with data
	return tmpl.Execute(file, data)
}

// shuffleWordsCrypto randomly shuffles the BIP39 word list using cryptographically secure random numbers
func shuffleWordsCrypto() ([]struct {
	Position int
	Word     string
}, error,
) {
	// Create a copy of the word list
	wordList := make([]string, len(BIP39WordList))
	copy(wordList, BIP39WordList)

	// Fisher-Yates shuffle algorithm with crypto/rand
	for i := len(wordList) - 1; i > 0; i-- {
		// Generate cryptographically secure random number between 0 and i
		nBig, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return nil, fmt.Errorf("failed to generate secure random number: %v", err)
		}
		j := int(nBig.Int64())

		// Swap elements
		wordList[i], wordList[j] = wordList[j], wordList[i]
	}

	// Create entries with new positions
	entries := make([]struct {
		Position int
		Word     string
	}, len(wordList))

	// Assign new positions (1-based)
	for i, word := range wordList {
		entries[i] = struct {
			Position int
			Word     string
		}{
			Position: i + 1, // 1-based position
			Word:     word,
		}
	}

	return entries, nil
}

// HTML template for the output
const htmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Bitcoin Seed Phrase Word List</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            max-width: 1200px;
            margin: 0 auto;
            padding: 20px;
            line-height: 1.6;
        }
        h1 {
            text-align: center;
            color: #333;
            margin-bottom: 30px;
        }
        .info {
            background-color: #f5f5f5;
            padding: 15px;
            border-radius: 5px;
            margin-bottom: 20px;
            text-align: center;
        }
        .word-grid {
            display: grid;
            grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
            gap: 10px;
            margin-bottom: 30px;
        }
        .word-item {
            background-color: #f9f9f9;
            border: 1px solid #ddd;
            border-radius: 4px;
            padding: 10px;
            display: flex;
            align-items: center;
            transition: background-color 0.2s;
        }
        .word-item:hover {
            background-color: #e9e9e9;
        }
        .word-number {
            font-weight: bold;
            margin-right: 10px;
            min-width: 40px;
            text-align: right;
            color: #555;
        }
        .word-text {
            flex-grow: 1;
        }
        .timestamp {
            text-align: center;
            font-size: 0.8em;
            color: #777;
            margin-top: 30px;
        }
        @media print {
            body {
                padding: 0;
                max-width: none;
            }
            .word-grid {
                grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
            }
            .word-item {
                break-inside: avoid;
            }
            .no-print {
                display: none;
            }
        }
    </style>
</head>
<body>
    <div class="word-grid">
        {{range .Words}}
        <div class="word-item">
            <span class="word-number">{{.Position}}.</span>
            <span class="word-text">{{.Word}}</span>
        </div>
        {{end}}
    </div>
    
    <div class="timestamp">
        Generated on: {{.Timestamp}}
    </div>
</body>
</html>`

func main() {
	// Use the BIP39WordList from word.go
	fmt.Printf("Loaded %d seed words\n", len(BIP39WordList))

	// Shuffle the words using cryptographically secure random number generation
	shuffledWords, err := shuffleWordsCrypto()
	if err != nil {
		fmt.Printf("Error shuffling words: %v\n", err)
		return
	}

	// Generate HTML output
	err = generateHTML(shuffledWords, "scrambled_seed_words.html")
	if err != nil {
		fmt.Printf("Error generating HTML: %v\n", err)
		return
	}

	fmt.Println("Scrambled words saved to scrambled_seed_words.html")
	fmt.Println("You can open this file in your browser and print it to save as PDF")
}
