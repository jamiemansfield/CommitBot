package modules

import (
    "github.com/fluffle/goirc/client"
    "crypto/tls"
    "strings"
)

var (
    BOT *client.Conn
)

func InitBot() {
    config := client.NewConfig(
        CONFIG.Section("IRC").Key("nickname").String(), // nick
        CONFIG.Section("IRC").Key("nickname").String(), // ident
        CONFIG.Section("IRC").Key("username").String(), // username
    )
    config.Server = CONFIG.Section("IRC").Key("server").String()
    config.SSL = CONFIG.Section("IRC").Key("ssl").MustBool(false)
    config.SSLConfig = &tls.Config{ServerName: strings.Split(CONFIG.Section("IRC").Key("server").String(), ":")[0]}
    if CONFIG.Section("IRC").HasKey("password") {
        config.Pass = CONFIG.Section("IRC").Key("password").String()
    }

    BOT = client.Client(config)
}
