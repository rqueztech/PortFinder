package main

import (
    "bufio"
    "os"
    "fmt"
    "strings"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

const(
    portnumber uint8 = 1 << iota
    protocolname
)

const(
    wellknownflat uint8 = 1 << iota
    ephemeralflat
    otherflat
)

// constant variables holding the database queries
const (
    OPEN_PORTDATABASE = "./db/portdatabase.db"
    QUERY_BY_PORTNUMBER = "SELECT * FROM ports WHERE port_number = ?;"
    QUERY_BY_PORTNUMBER_WELLKNOWN = "SELECT * FROM ports WHERE port_number = ? AND port_category = 'Well Known Ports';"
    QUERY_BY_PORTNAME = "SELECT * FROM ports WHERE short_name LIKE ?;"
    QUERY_BY_WELLKNWONPORTNUMBER = "SELECT * FROM ports WHERE port_number = ?;"
)

// hashmap containing ANSII colors
var ansiicolors = map[string]string {
    "red": "\x1b[31m",
    "green": "\x1b[32m",
    "yellow": "\x1b[33m",
    "blue": "\x1b[34m",
    "magenta": "\x1b[35m",
    "cyan": "\x1b[36m",
    "white": "\x1b[37m",
    "brightred": "\x1b[91m",
    "brightgreen": "\x1b[92m",
    "brightyellow": "\x1b[93m",
    "brightblue": "\x1b[94m",
    "brightmagenta": "\x1b[95m",
    "brightcyan": "\x1b[96m",
    "brightwhite": "\x1b[97m",
    "reset": "\x1b[0m",
}

func paintansii(color string, text string) string {
    return ansiicolors[color] + text + ansiicolors["reset"]
}

func clearscreen() {
    fmt.Print("\033[H\033[2J")
}

func queryByPortNumber(input string, querytoexecute string, verbosity bool) {
    db, err := sql.Open("sqlite3", OPEN_PORTDATABASE)
    if err != nil {
        log.Fatal(err)
    }

    defer db.Close()

    rows, err := db.Query(string(querytoexecute), input)
    if err != nil {
        log.Fatal(err)
    }

    for rows.Next() {
        var main_id string
        var short_name string
        var port_number string
        var tcp string
        var udp string
        var sctp string
        var dccp string
        var port_description string
        var port_category string


        err = rows.Scan(&main_id, &short_name, &port_number, &tcp, &udp, &sctp, &dccp, &port_description, &port_category)
        
        if err != nil {
            if err == sql.ErrNoRows {
                fmt.Println("No rows found")
            } else {
                log.Fatal(err)
            }
        }
        
        if verbosity == true {
            if (short_name != "") {
                fmt.Printf("\n%s: %s", paintansii("red", "Protocol Name: "), short_name)
			}
            if (port_number != "") {
                fmt.Printf("\n%s: %s", paintansii("green", "Port Number"), port_number)
			}
            if (tcp != "") {
                fmt.Printf("\n%s: %s", paintansii("yellow", "TCP"), tcp)
			}
            if (udp != "") {
                fmt.Printf("\n%s: %s", paintansii("blue", "UDP"), udp)
			}
            if (sctp != "") {
                fmt.Printf("\n%s: %s", paintansii("magenta", "SCTP"), sctp)
			}
            if (dccp != "") {
				fmt.Printf("\n%s: %s", paintansii("cyan", "DCCP"), dccp)
			}
            if (port_description != "") {
				fmt.Printf("\n%s %s", paintansii("brightred", "Port Description"), port_description)
			}
            if (port_category != "") {
				fmt.Printf("\n%s %s", paintansii("brightblue", "Port Category"), port_category)
			}
            fmt.Println("\x1b[--------------------------------------------------")
        } else {
            fmt.Printf("\x1b[35m%s\x1b[0m: %s TCP: %sUDP: %s\n", short_name, port_number, tcp, udp)
        }
    }
}

func checkForIllegalCharacters(input string) bool {
    // Define the set of illegal characters
    illegalCharacters := map[rune]bool{
        '!': true, '@': true, '#': true, '$': true, '%': true, '^': true, '&': true,
        '*': true, '(': true, ')': true, '_': true, '+': true, '=': true, '{': true,
        '}': true, '[': true, ']': true, '|': true, '\\': true, ':': true, ';': true,
        '"': true, '\'': true, '<': true, '>': true, ',': true, '.': true, '?': true,
        '/': true, '`': true, '~': true,
    }

    // Iterate over each character in the input string
    for _, character := range input {
        if _, found := illegalCharacters[character]; found {
            fmt.Println("Illegal characters found")
            return true
        }
    }

    return false
}

func main() {
    var protocolportmask uint8 = 0
    var portcategorymask uint8 = 0
    var verbose bool = true

    clearscreen()

    for true {
        protocolnamestring := "protocolname"
        portnumberstring := "portnumber"

        verbosestring := "verbose"
        nonverbosestring := "nonverbose"

        wellknownportsstring := "wellknownports"
        ephemeralportsstring := "ephemeralports"
        otherports := "otherports"

        if protocolportmask == portnumber {
            portnumberstring = paintansii("red", "portnumber") 
        } else if protocolportmask == protocolname {
            protocolnamestring = paintansii("red", "protocolname")
        } 

        if portcategorymask == wellknownflat {
            wellknownportsstring = paintansii("blue", "wellknownports")
        } else if portcategorymask == ephemeralflat {
            ephemeralportsstring = paintansii("blue", "ephemeralports")
        } else if portcategorymask == otherflat {
             otherports = paintansii("blue", "otherports")
        }

        if verbose == true {
            verbosestring = paintansii("green", "verbose")
        } else {
            nonverbosestring = paintansii("green", "nonverbose")
        }

        fmt.Printf("Enter Option: (%s, %s, clear)\n%s :: %s\n", portnumberstring, protocolnamestring, verbosestring, nonverbosestring)
        fmt.Printf("%s, %s, %s\n\n", wellknownportsstring, ephemeralportsstring, otherports)

        fmt.Printf("Enter Input: ")
        reader:= bufio.NewReader(os.Stdin)

        input, err := reader.ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }

        input = strings.TrimSpace(input)
        illegalcharactersfound := checkForIllegalCharacters(input)

        if illegalcharactersfound {
            fmt.Println("Illegal characters found")
        } else {
            switch input{
            case "portnumber":
                clearscreen()
                fmt.Println("PROTOCOL NAME")
                protocolportmask = portnumber
            case "protocolname":
                clearscreen()
                protocolportmask = protocolname
            case "wellknownports":
                clearscreen()
                portcategorymask = wellknownflat
            case "ephemeralports":
                clearscreen()
                portcategorymask = ephemeralflat
            case "otherports":
                clearscreen()
                portcategorymask = otherflat
            case "verbose":
                clearscreen()
                verbose = true
            case "nonverbose":
                clearscreen()
                verbose = false
            case "clear":
                clearscreen()
            default:
                if protocolportmask == portnumber && input != ""{
                    clearscreen()
                    queryByPortNumber(input, QUERY_BY_PORTNUMBER, verbose)
                } else if protocolportmask == protocolname {
                    clearscreen()
                    input =  "%" + input + "%"
                    queryByPortNumber(input, QUERY_BY_PORTNAME, verbose)
                }
            }
            
            input = ""
        }
    }
}

