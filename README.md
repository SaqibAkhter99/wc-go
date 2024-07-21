# gowc

`gowc` is a command-line tool written in Go for counting lines, words, bytes, and characters in a text file. It is similar to the Unix `wc` command.

## Features

- Count lines in a file
- Count words in a file
- Count bytes in a file
- Count characters in a file

## Installation

1. **Clone the repository**:
    ```sh
    git clone https://github.com/yourusername/wc-go.git
    ```

2. **Navigate to the project directory**:
    ```sh
    cd wc-go
    ```

3. **Build the binary**:
    ```sh
    go build -o gowc main.go
    ```

4. **Move the binary to a directory in your PATH** (optional, for easier usage):
    ```sh
    sudo mv gowc /usr/local/bin/
    ```

## Usage

You can run `gowc` directly from the command line. Here are the available flags:

- `-c`: Count bytes
- `-l`: Count lines
- `-w`: Count words
- `-m`: Count characters

### Examples

1. **Count bytes in a file**:
    ```sh
    gowc -c example.txt
    ```

2. **Count lines in a file**:
    ```sh
    gowc -l example.txt
    ```

3. **Count words in a file**:
    ```sh
    gowc -w example.txt
    ```

4. **Count characters in a file**:
    ```sh
    gowc -m example.txt
    ```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any inquiries or issues, please contact saqibakhtar88@gmail.com.
