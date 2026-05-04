# ASCII Art Generator

A simple command-line program written in Go that converts text into ASCII art using banner files.

## Description

This project reads normal text input and prints the matching ASCII art representation in the terminal.

It can also save the generated ASCII art into a file using an output flag.

## Features

- Convert text into ASCII art
- Supports banner files
- Output directly to the terminal
- Save output to a file using a flag

## Usage

```bash
go run . [OPTION] [STRING] [BANNER]

## Text Alignment

This project supports aligning ASCII art output in the terminal.

Available alignment options:

- `left` — aligns output to the left (default)
- `center` — centers the ASCII art
- `right` — aligns the ASCII art to the right
- `justify` — distributes spacing between words so the text stretches across the terminal width

### Usage

```bash
go run . --align=<mode> [STRING] [BANNER]