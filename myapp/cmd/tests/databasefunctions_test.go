package tests

import (
    "database/sql"
    "testing"
    "PortProgram/myapp/cmd/databasefunctions"
    "log"
    _ "github.com/mattn/go-sqlite3"
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

func TestOpenDatabase(t *testing.T) {
    dbConnector := databasefunctions.OpenDatabase(OPEN_PORTDATABASE)

    if dbConnector == nil {
        t.Errorf("Error: Database connection failed.")
    }

    dbConnector.Close()
}
