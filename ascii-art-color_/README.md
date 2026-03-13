# ASCII Color

## Overview

ASCII Color is a Go program that prints text as ASCII art and allows specific parts of the text to be displayed in color.
It uses a banner file containing ASCII representations of characters and applies ANSI color codes to selected substrings.

The program:

* Converts input text into ASCII art.
* Detects occurrences of a specified substring.
* Applies a chosen color to those characters in the ASCII output.

---

## Features

* Converts normal text into ASCII banner art.
* Supports coloring specific substrings within the text.
* Handles multi-line input.
* Uses ANSI escape codes to display colored output in the terminal.

---

## Project Structure

```
.
├── coloring
│   ├── asciicolor.go
│   └── sub2color.go
├
├── README.md
├── shadow.txt
├── standard.txt
├── thinkertoy.txt
└── main.go
```

### Description

* **main.go**
  Entry point of the program. Handles arguments and reads banner files.

* **coloring/asciicolor.go**
  Responsible for printing ASCII characters and applying colors.

* **coloring/sub2color.go**
  Determines which characters in the input string should be colored.

* **banner.txt, standard.txt, thinkertoy.txt**
  ASCII banner files containing representations of printable characters.

---

## How It Works

### 1. Detecting Characters to Color

The function `Sub2color()` scans the input string and creates a boolean slice.
Each position corresponds to a character in the string.

Example:

```
Input: "Hello World"
Substring: "lo"

Result:
[false false false true true false false false false false false]
```

`true` means that character should be colored.

---

### 2. Printing ASCII Characters

Each ASCII character in the banner file is **8 lines tall**.

The program:

1. Reads the banner file specified in `main.go`.
2. Locates the ASCII art for each character.
3. Prints it line by line.

Character location is calculated with:

```
index = (ASCII value - 32) * 9
```

---

### 3. Applying Color

If a character's position in `checklist` is `true`, the program prints it with a color code.

Example:

```
fmt.Print(colorCode + asciiLine + "\033[0m")
```

`\033[0m` resets the color so it doesn't affect the next characters.

---

## Supported Colors

| Color   | Code       |
| ------- | ---------- |
| Red     | `\033[31m` |
| Green   | `\033[32m` |
| Yellow  | `\033[33m` |
| Blue    | `\033[34m` |
| Default | `\033[0m`  |

---

## Usage

```
go run . <text> <substring> <color>
```

Example:

```
go run . "Hello World" "World" red
```

This will print **"World"** in red ASCII art.

---

## Example Output

Input:

```
Hello
```
Output (ASCII style ~ using shadow.txt):
```
 _   _          _ _       
| | | |   ___  | | | ___  
| |_| |  / _ \ | | |/ _ \ 
|  _  | |  __/ | | | (_) |
|_| |_|  \___| |_|_|\___/
     █
```



Output (ASCII style ~ using thinkertoy.txt):
```
o  o     o
|  |     |
O--O  o  |  o-o
|  |  |  |  | |
o  o  o  o  o-o
```

Output (ASCII style ~ using standard.txt):

```
 _   _      _ _
| | | | ___| | | ___
| |_| |/ _ \ | |/ _ \
|  _  |  __/ | | (_) |
|_| |_|\___|_|_|\___/
```

If a substring is specified, those characters will appear in color.

---

## Requirements

* Go 1.18+
* Terminal that supports ANSI color codes

---

## Notes

* ASCII banner files must follow the standard format used in `shadow.txt`, `standard.txt`and `thinkertoy` depending on what file is specified in `main.go`.
* Each character occupies 9 lines in the file (1 empty + 8 lines of ASCII).

---

## Author

Built using Go as a learning project to practice:

* String manipulation
* ASCII rendering
* Terminal color formatting
* Slice handling
