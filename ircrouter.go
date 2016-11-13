package main

import (
    "fmt"
    "log"
    "strings"

    irc "github.com/fluffle/goirc/client"
)

var SafeCommands []string = []string{"uptime", "echo"}

type Router struct {
    Channel, Nick, User string
}

func (r *Router) ConnectToChannel(conn *irc.Conn, line *irc.Line) {
    conn.Join(r.Channel)
    log.Printf("Connected to %s", r.Channel)

    conn.Privmsg(r.Channel, fmt.Sprintf("Hello! I am %s, your friendly command bot!\n%s\n", r.Nick))

}

func (r *Router) Route(conn *irc.Conn, line *irc.Line) {
    var requestor = line.Target()
    var command = line.Text()

    var resp string

    // Ignore any command not sent to Rosie
    if requestor == r.User || strings.HasPrefix(command, fmt.Sprintf("%s: ", r.Nick)) {
        log.Printf("Received: %s", line.Raw)

        command = r.normaliseCommand(command)

        if r.isValid(requestor, command) {
            switch command {
            default:
                resp = command
            }

            conn.Privmsg(requestor, resp)
        } else {
            conn.Privmsg(requestor, fmt.Sprintf("I'm sorry, I cannot run '%s' for you\n", command))
        }

    }
}

func (r *Router) normaliseCommand(cmd string) string {
    postString := strings.Replace(cmd, fmt.Sprintf("%s:", r.Nick), "", -1)
    return strings.TrimSpace(postString)
}

func (r *Router) isValid(requestor, cmd string) bool {
    for _, v := range SafeCommands {
        if v == cmd {
            return true
        }
    }
    return requestor == r.User
}
