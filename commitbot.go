package main

import (
    "fmt"
    "github.com/fluffle/goirc/client"
    "github.com/jamierocks/CommitBot/modules"
    "github.com/jamierocks/CommitBot/controllers"
    "gopkg.in/macaron.v1"
)

func main() {
    // macaron
    s := macaron.Classic()

    // init all the stuff
    modules.InitConfig()
    modules.InitBot()

    // Routes
    s.Post("/commit", controllers.GetGithub)

    // Tell client to connect.
    if err := modules.BOT.Connect(); err != nil {
        fmt.Printf("Connection error: %s\n", err.Error())
    }

    // Ensure bot is in the channel
    modules.BOT.HandleFunc(client.CONNECTED,
        func(conn *client.Conn, line *client.Line) {
            conn.Join("#" + modules.CONFIG.Section("IRC").Key("channel").String())
        })

    // Run webserver
    s.Run()
}
