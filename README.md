# ccwc

`ccwc` is a command-line utility written in Go that replicates the basic functionality of the Unix `wc` (word count) command.

This is a Go study project based on the [Coding Challenges](https://codingchallenges.fyi) exercises, particularly [this one](https://codingchallenges.fyi/challenges/challenge-wc)

## Features

- **Bytes**: Counts the number of bytes in a file.
- **Lines**: Counts the number of lines in a file.
- **Words**: Counts the number of words in a file.
- **Chars**: Counts the number of words in a file.
- **Supports multiple files**: Processes multiple files in one command.
- **Standard Input**: Reads from standard input if no file is provided.


## Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/friarhob/ccwc.git
   cd ccwc
   ```

1. **Build the executable:**
   ```bash
   go build ccwc.go
   ```

## Usage
   ```bash
   ccwc [options] [filepaths]
   ```

### Options
- `-c`, `--bytes`: Print the number of bytes.
- `-l`, `--lines`: Print the number of lines.
- `-w`, `--words`: Print the number of words.
- `-m`, `--chars`: Print the number of chars (considers encoding, may differ from `-c`).
- `-h`, `--help`: Display help message.

### Examples

1. **Count bytes only:**
   ```bash
   ccwc -c example.txt
   ```

1. **Count lines only:**
   ```bash
   ccwc -l example.txt
   ```

1. **Count words only:**
   ```bash
   ccwc -w example.txt
   ```

1. **Count chars only:**
   ```bash
   ccwc -m example.txt
   ```

1. **Count bytes and chars:**
   ```bash
   ccwc -c -m example.txt
   ```

1. **Count lines, words, and bytes in a file:**
   ```bash
   ccwc example.txt
   ```

1. **Count for multiple files:**
   ```bash
   ccwc file1.txt file2.txt
   ```

1. **Use with standard input:**
   ```bash
   echo "Hello World" | ccwc
   ```

_*the output order will always be lines, words, bytes and chars_

## Exit Codes

Code | Description | Obs.
:-:|---|---
0 | No errors | also returns when `--help` is called
1 | Error reading from stdin |
2 | Error reading from file |

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

