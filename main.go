package main

import (
    "flag"
    "fmt"
    "log"

    irc "github.com/fluffle/goirc/client"
)

var (
    connectionString *string
    router Router
)


func init() {
    router.Channel = flag.String("c", "#control-room", "Channel with which to connect")
    connectionString = flag.String("h", "localhost:6667", "IRC daemon to connect to")
    router.Nick = flag.String("n", "rosie", "Nick with which to connect")
    router.User = flag.String("u", "jspc", "User to listen to")
}

func main() {
    flag.Parse()

    log.Printf("Connecting to %s as %s. This bot lives in %s and listens to %s",
        *connectionString,
        *router.Nick,
        *router.Channel,
        *router.User)

    cfg := irc.NewConfig(*router.Nick)
    cfg.Server = *connectionString
    cfg.NewNick = func(n string) string { return n + "^" }
    c := irc.Client(cfg)

    c.HandleFunc(irc.CONNECTED, router.ConnectToChannel)
    c.HandleFunc(irc.PRIVMSG, router.Route)

    quit := make(chan bool)
    c.HandleFunc(irc.DISCONNECTED, func(conn *irc.Conn, line *irc.Line) {
        quit <- true
    })


    if err := c.Connect(); err != nil {
        fmt.Printf("Connection error: %s\n", err.Error())
    }

    <-quit
}
