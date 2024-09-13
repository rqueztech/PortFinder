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
    ipprotocol
    linkprotocol
    physicalprotocol
)

const(
    wellknownport uint8 = 1 << iota
    registeredport
    otherport
    anyport
)

// constant variables holding the database queries
const (
    OPEN_PORTDATABASE = "../../db/portdatabase.db"
    QUERY_BY_PORTNUMBER = "SELECT * FROM ports WHERE port_number = ?;"
    QUERY_BY_PORTNUMBER_WELLKNOWN = "SELECT * FROM ports WHERE port_number = ? AND port_category = 'Well Known Ports'";
    QUERY_BY_PORTNUMBER_REGISTERED = "SELECT * FROM ports WHERE port_number = ? AND port_category = 'Registered Ports'";
    QUERY_BY_PORTNUMBER_DYNAMIC = "SELECT * FROM ports WHERE port_number = ? AND port_category = 'Dynamic, Private, or Ephemeral'";
    QUERY_BY_PORTNAME_ANY = "SELECT * FROM ports WHERE short_name LIKE ?;"
    QUERY_BY_PORTNAME_WELLKNOWN = "SELECT * FROM ports WHERE short_name LIKE ? AND port_category = 'Well Known Ports'";
    QUERY_BY_PORTNAME_REGISTERED = "SELECT * FROM ports WHERE short_name LIKE ? AND port_category = 'Registered Ports'";
    QUERY_BY_PORTNAME_DYNAMIC = "SELECT * FROM ports WHERE short_name LIKE ? AND port_category = 'Dynamic, Private, or Ephemeral'";
    QUERY_BY_WELLKNWONPORTNUMBER = "SELECT * FROM ports WHERE port_number = ?;"
    QUERY_BY_LINKPROTOCOL = "SELECT * FROM link_protocols WHERE protocol_name LIKE ?;"
    QUERY_BY_IPPROTOCOL = "SELECT * FROM ip_protocols WHERE keyword LIKE ?;"
    QUERY_BY_PHYSICALPROTOCOL = "SELECT * FROM physical_protocols WHERE protocol_name LIKE ?;"
)

// hashmap containing ANSII colors
var AnsiiColors = map[string]string {
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

func PaintAnsii(color string, text string) string {
    return AnsiiColors[color] + text + AnsiiColors["reset"]
}

func ClearScreen() string {
    fmt.Print("\033[H\033[2J")
    return "\033[H\033[2J"
}

func QueryIPTable(input string, querytoexecute string, verbosity bool) {
    db, err := sql.Open("sqlite3", OPEN_PORTDATABASE)
    if err != nil {
        log.Fatal(err)
    }

    defer db.Close()
    rows, err := db.Query(querytoexecute, input)
    if err != nil {
        log.Fatal(err)
    }

    for rows.Next() {
        var hex string
        var protocol_number string
        var keyword string
        var protocol string
        var refernces_rfc string

        err = rows.Scan(&hex, &protocol_number, &keyword, &protocol, &refernces_rfc)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("\n%s: %s", PaintAnsii("red", "Hex"), hex)
        fmt.Printf("\n%s: %s", PaintAnsii("green", "Protocol Number"), protocol_number)
        fmt.Printf("\n%s: %s", PaintAnsii("yellow", "Keyword"), keyword)
        fmt.Printf("\n%s: %s", PaintAnsii("blue", "Protocol"), protocol)
        fmt.Printf("\n%s: %s", PaintAnsii("magenta", "References RFC"), refernces_rfc)
        fmt.Println("\x1b[--------------------------------------------------")
    }
}

func QueryPortsTable(input string, querytoexecute string, verbosity bool) {
    db, err := sql.Open("sqlite3", OPEN_PORTDATABASE)
    if err != nil {
        log.Fatal(err)
    }
    
    defer db.Close()
    rows, err := db.Query(querytoexecute, input)

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
                fmt.Printf("\n%s: %s", PaintAnsii("red", "Protocol Name: "), short_name)
                fmt.Println("")
			}
            if (port_number != "") {
                fmt.Printf("\n%s: %s", PaintAnsii("green", "Port Number"), port_number)
			}
            if (tcp != "") {
                fmt.Printf("\n%s: %s", PaintAnsii("yellow", "TCP"), tcp)
			}
            if (udp != "") {
                fmt.Printf("\n%s: %s", PaintAnsii("blue", "UDP"), udp)
			}
            if (sctp != "") {
                fmt.Printf("\n%s: %s", PaintAnsii("magenta", "SCTP"), sctp)
			}
            if (dccp != "") {
				fmt.Printf("\n%s: %s", PaintAnsii("cyan", "DCCP"), dccp)
			}
            if (port_description != "") {
				fmt.Printf("\n%s %s", PaintAnsii("brightred", "Port Description"), port_description)
			}
            if (port_category != "") {
				fmt.Printf("\n%s %s", PaintAnsii("brightblue", "Port Category"), port_category)
			}
            fmt.Println("\x1b[--------------------------------------------------")
        } else {
            nonverbosereading := PaintAnsii("green", short_name)
            nonverbosereading += " : "
            nonverbosereading += PaintAnsii("red", port_number)
            nonverbosereading += " : "
            nonverbosereading += PaintAnsii("yellow", ("TCP (" + tcp + ")"))
            nonverbosereading += " : "
            nonverbosereading += PaintAnsii("blue", ("UDP (" + udp + ")"))
            fmt.Println(nonverbosereading)
        }
    }
}

func CheckForIllegalCharacters(input string) bool {
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
    var searchPortNumberOrProtocolName uint8 = 1
    var wellKnownorRegisteredorEphemeralPort uint8 = 1
    var verbose bool = true

    ClearScreen()

    for true {
        protocolnamestring := "protocolname"
        portnumberstring := "portnumber"
        ipprotocolstring := "ip"
        linkprotocolstring := "link"
        physicalprotocolstring := "physical"

        verbosestring := "verbose"
        nonverbosestring := "nonverbose"

        wellknownportsstring := "wellknownports"
        registeredportstring := "registeredports"
        anyportstring := "anyport"
        otherports := "otherports"


        if searchPortNumberOrProtocolName == portnumber {
            portnumberstring = PaintAnsii("red", "portnumber") 
        } else if searchPortNumberOrProtocolName == protocolname {
            protocolnamestring = PaintAnsii("red", "protocolname")
        } else if searchPortNumberOrProtocolName == ipprotocol {
            ipprotocolstring = PaintAnsii("red", "ip")
        } else if searchPortNumberOrProtocolName == linkprotocol {
            linkprotocolstring = PaintAnsii("red", "link")
        } else if searchPortNumberOrProtocolName == physicalprotocol {
            physicalprotocolstring = PaintAnsii("red", "physical")
        }

        if wellKnownorRegisteredorEphemeralPort == wellknownport {
            wellknownportsstring = PaintAnsii("blue", "wellknownports")
        } else if wellKnownorRegisteredorEphemeralPort == registeredport {
            registeredportstring = PaintAnsii("blue", "registeredports")
        } else if wellKnownorRegisteredorEphemeralPort == otherport {
             otherports = PaintAnsii("blue", "otherports")
        } else if wellKnownorRegisteredorEphemeralPort == anyport {
            anyportstring = PaintAnsii("blue", "anyport")
        }

        if verbose == true {
            verbosestring = PaintAnsii("green", "verbose")
        } else {
            nonverbosestring = PaintAnsii("green", "nonverbose")
        }

        fmt.Printf("Enter Option: (%s, %s, %s, %s, %s clear)\n%s :: %s\n", portnumberstring, protocolnamestring, ipprotocolstring, linkprotocolstring, physicalprotocolstring, verbosestring, nonverbosestring)
        fmt.Printf("%s, %s, %s, %s\n\n", wellknownportsstring, registeredportstring, otherports, anyportstring)

        fmt.Printf("Enter Input: ")
        reader:= bufio.NewReader(os.Stdin)

        input, err := reader.ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }

        input = strings.TrimSpace(input)
        illegalcharactersfound := CheckForIllegalCharacters(input)

        if illegalcharactersfound {
            fmt.Println("Illegal characters found")
        } else {
            switch input{
            case "portnumber":
                //ClearScreen()
                searchPortNumberOrProtocolName = portnumber
            case "protocolname":
                //ClearScreen()
                searchPortNumberOrProtocolName = protocolname
            case "wellknownports":
                //ClearScreen()
                wellKnownorRegisteredorEphemeralPort = wellknownport
            case "registeredports":
                //ClearScreen()
                wellKnownorRegisteredorEphemeralPort = registeredport
            case "otherports":
                //ClearScreen()
                wellKnownorRegisteredorEphemeralPort = otherport
            case "anyport":
                wellKnownorRegisteredorEphemeralPort = anyport
            case "ip":
                searchPortNumberOrProtocolName = ipprotocol
            case "link":
                searchPortNumberOrProtocolName = linkprotocol
            case "physical":
                searchPortNumberOrProtocolName = physicalprotocol
            case "verbose":
                //ClearScreen()
                verbose = true
            case "nonverbose":
                //ClearScreen()
                verbose = false
            case "clear":
                ClearScreen()
            default:
                if (searchPortNumberOrProtocolName == portnumber) {
                    QueryPortsTable(input, QUERY_BY_PORTNUMBER, verbose)
                } else if searchPortNumberOrProtocolName == protocolname {
                    input = "%" + input + "%"
                    if (wellKnownorRegisteredorEphemeralPort == wellknownport) {
                        //ClearScreen()
                        QueryPortsTable(input, QUERY_BY_PORTNAME_WELLKNOWN, verbose)
                    }
                    if (wellKnownorRegisteredorEphemeralPort == registeredport) {
                        //ClearScreen()
                        QueryPortsTable(input, QUERY_BY_PORTNAME_REGISTERED, verbose)
                    }
                    if (wellKnownorRegisteredorEphemeralPort == otherport) {
                        //ClearScreen()
                        QueryPortsTable(input, QUERY_BY_PORTNAME_DYNAMIC, verbose)
                    }

                    if (wellKnownorRegisteredorEphemeralPort == anyport) {
                        //ClearScreen()
                        QueryPortsTable(input, QUERY_BY_PORTNAME_ANY, verbose)
                    }
                }
            }
            input = ""

            fmt.Print("Press enter to continue...")
            _, err := reader.ReadString('\n')
            if err != nil {
                log.Fatal(err)
            }
            ClearScreen()
        }
    }
}

