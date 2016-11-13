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
    var args string

    // Ignore any command not sent to Rosie
    if requestor == r.User || strings.HasPrefix(command, fmt.Sprintf("%s: ", r.Nick)) {
        log.Printf("Received: %s", line.Raw)

        go func() {
            command, args = r.normaliseCommand(command)

            if r.isValid(requestor, command) {
                se := NewScriptEngine(*scriptDir)
                sc := ScriptContext{requestor, command, args}

                if resp, err := se.Run(command, sc); err != nil {
                    conn.Privmsg(requestor, fmt.Sprintf("An error occured: %q", err))
                } else {
                    conn.Privmsg(requestor, resp.String())
                }

            } else {
                conn.Privmsg(requestor, fmt.Sprintf("I'm sorry, I cannot run '%s' for you\n", command))
            }
        }()

    }
}

func (r *Router) normaliseCommand(cmd string) (string, string) {
    postString := strings.Replace(cmd, fmt.Sprintf("%s:", r.Nick), "", -1)

    textSplit := strings.SplitAfterN( strings.TrimSpace(postString), " ", 2)

    if len(textSplit) < 2 {
        return strings.TrimSpace(textSplit[0]), ""
    }

    return strings.TrimSpace(textSplit[0]), textSplit[1]
}

func (r *Router) isValid(requestor, cmd string) bool {
    for _, v := range SafeCommands {
        if v == cmd {
            return true
        }
    }
    return requestor == r.User
}
