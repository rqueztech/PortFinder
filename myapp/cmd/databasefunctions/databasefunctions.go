package databasefunctions

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "fmt"
    "log"
    "PortProgram/myapp/cmd/utils"
)

// Constants for database queries
const (
    OPEN_PORTDATABASE = "../../db/portdatabase.db"
    QUERY_BY_PORTNUMBER = "SELECT protocol_name, port_number, tcp, udp, sctp, dccp, protocol_description FROM ports WHERE port_number = ?;"
    QUERY_BY_PORTNUMBER_WELLKNOWN = "SELECT protocol_name, port_number, tcp, udp, sctp, dccp, protocol_description FROM ports WHERE port_number = ? AND port_category = 'Well Known Ports';"
    QUERY_BY_PORTNUMBER_REGISTERED = "SELECT protocol_name, port_number, tcp, udp, sctp, dccp, protocol_description FROM ports WHERE port_number = ? AND port_category = 'Registered Ports';"
    QUERY_BY_PORTNUMBER_DYNAMIC = "SELECT protocol_name, port_number, tcp, udp, sctp, dccp, protocol_description FROM ports FROM ports WHERE port_number = ? AND port_category = 'Dynamic, Private, or Ephemeral';"
    QUERY_BY_PORTNAME_ANY = "SELECT protocol_name, port_number, tcp, udp, sctp, dccp, protocol_description FROM ports WHERE protocol_name LIKE ? OR protocol_description LIKE ?;"
    QUERY_BY_PORTNAME_WELLKNOWN = "SELECT protocol_name, port_number, tcp, udp, sctp, dccp, protocol_description FROM ports WHERE (protocol_name LIKE ? OR protocol_description LIKE ?) AND port_category = 'Well Known Ports';"
    QUERY_BY_PORTNAME_REGISTERED = "SELECT protocol_name, port_number, tcp, udp, sctp, dccp, protocol_description FROM ports WHERE (protocol_name LIKE ? OR protocol_description LIKE ?) AND port_category = 'Registered Ports';"
    QUERY_BY_PORTNAME_DYNAMIC = "SELECT protocol_name, port_number, tcp, udp, sctp, dccp, protocol_description FROM ports WHERE (protocol_name LIKE ? OR protocol_description LIKE ?) AND port_category = 'Dynamic, Private, or Ephemeral';"
    QUERY_ALL_LINKPROTOCOL = "SELECT * FROM link_protocols;"
    QUERY_BY_LINKPROTOCOL = "SELECT * FROM link_protocols WHERE protocol_abbreviation LIKE ? OR protocol_name LIKE ? OR protocol_description LIKE ?;"
    QUERY_ALL_IPPROTOCOL = "SELECT * FROM ip_protocols;"
    QUERY_BY_IPPROTOCOL = "SELECT * FROM ip_protocols WHERE protocol_abbreviation LIKE ? OR protocol_description LIKE ?;"
    QUERY_ALL_PHYSICALPROTOCOL = "SELECT * FROM physical_protocol;"
    QUERY_BY_PHYSICALPROTOCOL = "SELECT * FROM physical_protocol WHERE protocol_name LIKE ? OR protocol_description LIKE ?;"
    QUERY_ALL_NETWORK_METHODS = "SELECT * FROM networking_methods;"
    QUERY_NETWORK_METHODS = "SELECT * FROM networking_methods WHERE method_abbreviation LIKE ? OR method_name LIKE ?;"
)

// TableRow holds information for a table row
type TableRow struct {
    Columns []interface{}
    Headers []string
}

func OpenDatabase(dbname string) *sql.DB {
    db, err := sql.Open("sqlite3", dbname)
    if err != nil {
        log.Fatal(err)
    }
    return db
}

func CloseDatabase(db *sql.DB) {
    db.Close()
}

// QueryPortsTable executes a query and processes the results
func QueryPortsTable(querytoexecute string, args []interface{}, titles []string, verbosity bool) {
    db := OpenDatabase(OPEN_PORTDATABASE)

    rows, err := db.Query(querytoexecute, args...)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    columns, err := rows.Columns()
    if err != nil {
        log.Fatal(err)
    }

    columnPointers := make([]interface{}, len(columns))
    for i := range columnPointers {
        var col string
        columnPointers[i] = &col
    }

    var tableRow TableRow
    tableRow.Headers = titles
    tableRow.Columns = columnPointers

    for rows.Next() {
        err = rows.Scan(tableRow.Columns...)
        if err != nil {
            log.Fatal(err)
        }

        if verbosity {
            for i, header := range tableRow.Headers {
                value := *(tableRow.Columns[i].(*string))
                fmt.Printf("\n%s: %s", utils.Paint("blue", header), value)
            }
            fmt.Println("\x1b[--------------------------------------------------")
        } else {
            nonverbosereading := ""
            for i, header := range tableRow.Headers {
                value := *(tableRow.Columns[i].(*string))
                nonverbosereading += utils.Paint("green", header) + ": " + utils.Paint("red", value) + " "
            }
            fmt.Println(nonverbosereading)
        }
        fmt.Println("")
    }
}
