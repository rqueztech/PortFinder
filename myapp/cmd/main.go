package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "PortProgram/myapp/cmd/databasefunctions"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("1. Major Searches")
		fmt.Println("2. Minor Searches")
		chosemode, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		chosemode = strings.TrimSpace(chosemode)

		if chosemode == "1" {
            d, err := databasefunctions.NewDatabaseConnection(databasefunctions.OPEN_PORTDATABASE)
            if err != nil {
                log.Fatal(err)
            }
			d.MajorSearches()
		} else if chosemode == "2" {
			fmt.Println("Chosen two")
			_, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("nil")
			}
		}
	}
}

