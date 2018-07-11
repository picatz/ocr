# OCR
> Ocular character recognition command-line utility.

This tool is built using the [`Tesseract OCR Engine`](https://github.com/tesseract-ocr/tesseract).

# Installation

Be sure to [install tesseract](https://github.com/tesseract-ocr/tesseract/wiki#installation).

```shell
$ go get github.com/picatz/ocr
```

# Usage
```shell
# download example image to test
$ ocr extract ~/Downloads/example.jpeg
```

# Help Menu
```shell
Usage:
  ocr [command]

Available Commands:
  extract     Extract recognizable characters from a given image.
  help        Help about any command

Flags:
  -h, --help   help for ocr

Use "ocr [command] --help" for more information about a command.
```
