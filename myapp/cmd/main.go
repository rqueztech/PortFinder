package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "PortProgram/myapp/cmd/utils"
    "PortProgram/myapp/cmd/databasefunctions"

    _ "github.com/mattn/go-sqlite3"
)

const (
    portnumber uint8 = 1 << iota
    protocolname
    ipprotocol
    linkprotocol
    physicalprotocol
    methodname
)

const (
    wellknownport uint8 = 1 << iota
    registeredport
    otherport
    anyport
)


func majorSearches() {
	var searchPortNumberOrProtocolName uint8 = 1
    var wellKnownorRegisteredorEphemeralPort uint8 = 1
    var verbose bool = true
    ClearScreen := utils.ClearScreen
    ClearScreen()
	ipprotocoltitles := []string{"Hex", "Protocol Number", "Protocol Name: ", "Protocol Description", "References RFC"}
	linkprotocoltitles := []string {"Protocol Abbreviation", "Protocol Name", "Protocol Description", "Frequency"}
	physicalprotocoltitles := []string{"Abbrebiation", "Description", "Long Name", "Commonality"}
	methodprotocoltitles := []string{"Method Abbreviation", "Method Name", "Description"}
	reader := bufio.NewReader(os.Stdin)

    for {
        protocolnamestring := "protocolname"
        portnumberstring := "portnumber"
        ipprotocolstring := "ip"
        linkprotocolstring := "link"
        physicalprotocolstring := "physical"
        methodnamestring := "method"

        verbosestring := "verbose"
        nonverbosestring := "nonverbose"

        wellknownportsstring := "wellknownports"
        registeredportstring := "registeredports"
        anyportstring := "anyport"
        otherports := "otherports"

        if searchPortNumberOrProtocolName == portnumber {
            portnumberstring = utils.Paint("red", "portnumber")
        } else if searchPortNumberOrProtocolName == protocolname {
            protocolnamestring = utils.Paint("red", "protocolname")
        } else if searchPortNumberOrProtocolName == ipprotocol {
            ipprotocolstring = utils.Paint("red", "ip")
        } else if searchPortNumberOrProtocolName == linkprotocol {
            linkprotocolstring = utils.Paint("red", "link")
        } else if searchPortNumberOrProtocolName == physicalprotocol {
            physicalprotocolstring = utils.Paint("red", "physical")
        } else if searchPortNumberOrProtocolName == methodname {
            methodnamestring = utils.Paint("red", "method")
        }

        if wellKnownorRegisteredorEphemeralPort == wellknownport {
            wellknownportsstring = utils.Paint("blue", "wellknownports")
        } else if wellKnownorRegisteredorEphemeralPort == registeredport {
            registeredportstring = utils.Paint("blue", "registeredports")
        } else if wellKnownorRegisteredorEphemeralPort == otherport {
            otherports = utils.Paint("blue", "otherports")
        } else if wellKnownorRegisteredorEphemeralPort == anyport {
            anyportstring = utils.Paint("blue", "anyport")
        }

        if verbose {
            verbosestring = utils.Paint("green", "verbose")
        } else {
            nonverbosestring = utils.Paint("green", "nonverbose")
        }

        fmt.Printf("Enter Option: (%s, %s, %s, %s, %s, %s clear)\n%s :: %s\n", portnumberstring, protocolnamestring, ipprotocolstring, linkprotocolstring, physicalprotocolstring, methodnamestring, verbosestring, nonverbosestring)
        fmt.Printf("%s, %s, %s, %s\n\n", wellknownportsstring, registeredportstring, otherports, anyportstring)

        fmt.Printf("Enter Input: ")

        input, err := reader.ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }

        input = strings.TrimSpace(input)
        illegalcharactersfound := utils.CheckForIllegalCharacters(input)

        
        if illegalcharactersfound {
            fmt.Println("Illegal characters found")
        } else {
            switch input {
            case "portnumber":
                searchPortNumberOrProtocolName = portnumber
            case "protocolname":
                searchPortNumberOrProtocolName = protocolname
            case "wellknownports":
                wellKnownorRegisteredorEphemeralPort = wellknownport
            case "registeredports":
                wellKnownorRegisteredorEphemeralPort = registeredport
            case "otherports":
                wellKnownorRegisteredorEphemeralPort = otherport
            case "anyport":
                wellKnownorRegisteredorEphemeralPort = anyport
            case "ip":
                searchPortNumberOrProtocolName = ipprotocol
            case "link":
                searchPortNumberOrProtocolName = linkprotocol
            case "physical":
                searchPortNumberOrProtocolName = physicalprotocol
            case "method":
                searchPortNumberOrProtocolName = methodname
            case "verbose":
                verbose = true
            case "nonverbose":
                verbose = false
            case "clear":
                ClearScreen()
            default:
                portprotocoltitles := []string{"Protocol Name", "Port Number", "TCP", "UDP", "SCTP", "RDP", "Description"}
                if searchPortNumberOrProtocolName == portnumber && input != "ip" && input != "link" && input != "physical" {
                    databasefunctions.QueryPortsTable(databasefunctions.QUERY_BY_PORTNUMBER, []interface{}{input}, portprotocoltitles, verbose)
                } else if searchPortNumberOrProtocolName == protocolname && input != "ip" && input != "link" && input != "physical" {
                    input = "%" + input + "%"
                    
                    switch wellKnownorRegisteredorEphemeralPort {
                    case wellknownport:
                        databasefunctions.QueryPortsTable(databasefunctions.QUERY_BY_PORTNAME_WELLKNOWN, []interface{}{input, input}, portprotocoltitles, verbose)
                    case registeredport:
                        databasefunctions.QueryPortsTable(databasefunctions.QUERY_BY_PORTNAME_REGISTERED, []interface{}{input, input}, portprotocoltitles, verbose)
                    case otherport:
                        databasefunctions.QueryPortsTable(databasefunctions.QUERY_BY_PORTNAME_DYNAMIC, []interface{}{input, input}, portprotocoltitles, verbose)
                    case anyport:
                        databasefunctions.QueryPortsTable(databasefunctions.QUERY_BY_PORTNAME_ANY, []interface{}{input, input}, portprotocoltitles, verbose)
                    }
                } else {
                    fmt.Println("Input: ", searchPortNumberOrProtocolName, input)
                    switch searchPortNumberOrProtocolName {
                    case ipprotocol:
                        if input == "" {
                            databasefunctions.QueryPortsTable(databasefunctions.QUERY_ALL_IPPROTOCOL, []interface{}{input}, ipprotocoltitles, verbose)
                        }
                        input = "%" + input + "%"
                        databasefunctions.QueryPortsTable(databasefunctions.QUERY_BY_IPPROTOCOL, []interface{}{input, input}, ipprotocoltitles, verbose)
                    case linkprotocol:
                        fmt.Println("Link Protocol")
                        if input == "" {
                            databasefunctions.QueryPortsTable(databasefunctions.QUERY_ALL_LINKPROTOCOL, []interface{}{input}, linkprotocoltitles, verbose)
                        }
                        input = "%" + input + "%"
                        databasefunctions.QueryPortsTable(databasefunctions.QUERY_BY_LINKPROTOCOL, []interface{}{input, input, input}, linkprotocoltitles, verbose)
                    case physicalprotocol:
                        fmt.Println("Physical Protocol")
                        if input == "" {
                            databasefunctions.QueryPortsTable(databasefunctions.QUERY_ALL_PHYSICALPROTOCOL, []interface{}{input}, physicalprotocoltitles, verbose)
                        }
                        input = "%" + input + "%"
                        databasefunctions.QueryPortsTable(databasefunctions.QUERY_BY_PHYSICALPROTOCOL, []interface{}{input, input}, physicalprotocoltitles, verbose)
                    case methodname:
                        fmt.Println("Method")
                        input = "%" + input + "%"
                        if input == "" {
                            databasefunctions.QueryPortsTable(databasefunctions.QUERY_ALL_NETWORK_METHODS, []interface{}{input}, methodprotocoltitles, verbose)
                        }
                        databasefunctions.QueryPortsTable(databasefunctions.QUERY_NETWORK_METHODS, []interface{}{input, input}, methodprotocoltitles, verbose)
                    }
                }
            }

            fmt.Print("Press enter to continue...")
            _, err := reader.ReadString('\n')
            if err != nil {
                log.Fatal(err)
            }
            ClearScreen()
        }
	}
}


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
			majorSearches()
		} else if chosemode == "2" {
			fmt.Println("Chosen two")
			_, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("nil")
			}
		}
	}
}

