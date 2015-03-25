package log
/*
This file represents implementation of logging to remote syslog server.
The task was started, but was dropped until sometime in the future. The reason is simple: there's no support for syslog in
standard library on Windows (current go version is 1.4.2!). I think this is necessary. Until syslog is suported for Windows
also, this module will stay empty and logging will support only console and file loggers, usually this is enough.

*/

/*
import (
"fmt"
    "net"
    "log"
)

// SyslogLogger
type SyslogLogger struct {

    //
    conn *net.Conn
}

// CreateSyslogLogger creates a standard syslog logger sending messages to remote address using UDP port 514.
// The 'addr' parameter is any string containing IP address (v4 or v6).
func CreateSyslogLogger(ipaddr string, prio Priority) (*Logger, error) {

    addr := fmt.Sprintf("%s:514", ipaddr) // Standard syslog UDP port is 514
    conn, err := net.Dial("udp", addr)
    if err != nil {
        return nil, fmt.Errorf("Address %q UDP port 514 cannot be opened.\n", addr)
    }
    format := "%s%s %s %s"
    return CreateLogger(conn, prio, "", format, log.Lshortfile), err
}

// CreateCustom SyslogLogger creates a syslog logger sending messages to remote address using customnetwork protocol (anything but
// standard syslog UDP port 514).
// The 'net' parameter is any value depicting the network layer protocol, such as "tcp" or "udp". For all possibilities check the
// standard library 'net' package documentation. The 'addr' parameter is any string containing IP address (v4 or v6), custom port
// can also be specified (examples: "12.23.34.45" or "12:23:34:56:1234". See examples in standard library 'net' package
// documentation for more info.
func CreateCustomSyslogLogger(netw, addr string, prio Priority) (*Logger, error) {

    conn, err := net.Dial(netw, addr)
    if err != nil {
        return nil, fmt.Errorf("Address %q %s cannot be opened.\n", addr, netw)
    }
    format := "%s%s %s %s"
    return CreateLogger(conn, prio, "", format, log.Lshortfile), err
}

*/
