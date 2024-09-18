package utils

import (
    "fmt"
)

var Colors = map[string]string{
    "red":           "\x1b[31m",
    "green":         "\x1b[32m",
    "yellow":        "\x1b[33m",
    "blue":          "\x1b[34m",
    "magenta":       "\x1b[35m",
    "cyan":          "\x1b[36m",
    "white":         "\x1b[37m",
    "brightred":     "\x1b[91m",
    "brightgreen":   "\x1b[92m",
    "brightyellow":  "\x1b[93m",
    "brightblue":    "\x1b[94m",
    "brightmagenta": "\x1b[95m",
    "brightcyan":    "\x1b[96m",
    "brightwhite":   "\x1b[97m",
    "reset":         "\x1b[0m",
}

// Paint applies an ANSI color to a string
func Paint(color string, text string) string {
    return Colors[color] + text + Colors["reset"]
}

// ClearScreen clears the terminal screen
func ClearScreen() string {
    fmt.Print("\033[H\033[2J")
    return "\033[H\033[2J"
}

// CheckForIllegalCharacters checks for illegal characters in a string
func CheckForIllegalCharacters(input string) bool {
    illegalCharacters := map[rune]bool{
        '!': true, '@': true, '#': true, '$': true, '%': true, '^': true, '&': true,
        '*': true, '(': true, ')': true, '_': true, '+': true, '=': true, '{': true,
        '}': true, '[': true, ']': true, '|': true, '\\': true, ':': true, ';': true,
        '"': true, '\'': true, '<': true, '>': true, ',': true, '.': true, '?': true,
        '/': true, '`': true, '~': true, 
    }

    for _, character := range input {
        if _, found := illegalCharacters[character]; found {
            fmt.Println("Illegal characters found")
            return true
        }
    }

    return false
}
