package main

import (
    "testing"
    "os"
)

func checkIfFileExists(filepath string) bool {
    info, err := os.Stat(filepath)
    if os.IsNotExist(err) {
        return false
    }

    if err != nil {
        return false
    }

    return !info.IsDir()
}

func TestFileExists(t* testing.T) {
    filepath := "main.go"

    if !checkIfFileExists(filepath) {
        t.Errorf("File %s does not exist", filepath)
    } else {
        t.Logf("File %s exists", filepath)
    }
}

func TestClearScreen(t *testing.T) {
    result := ClearScreen()
    if result == "" {
        t.Errorf("ClearScreen() returned an empty result")
    }
}
