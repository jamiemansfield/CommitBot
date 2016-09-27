package controllers

import (
    "strconv"
    "io/ioutil"
    "encoding/json"
    "github.com/google/go-github/github"
    "github.com/jamierocks/CommitBot/utils"
    "github.com/jamierocks/CommitBot/modules"
    "gopkg.in/macaron.v1"
)

func GetGithub(ctx *macaron.Context) {
    if (ctx.Req.Header.Get("X-GitHub-Event") == "push") {
        body, _ := ioutil.ReadAll(ctx.Req.Body().ReadCloser())

        var res github.PushEvent
        json.Unmarshal(body, &res)

        branch := utils.GetBranchName(*res.Ref)
        compare := utils.GetGitioUrl(*res.Compare)

        modules.BOT.Privmsg("#" + modules.CONFIG.Section("IRC").Key("channel").String(),
            "\x0313[" + *res.Repo.Name + "]\x0F " + *res.Pusher.Name + " pushed \x02" + strconv.Itoa(len(res.Commits)) + "\x0F commits to " + branch + " " + compare)

        for _, commit := range res.Commits {
            message := utils.GetShortCommitMessage(*commit.Message)
            id := utils.GetShortCommitID(*commit.ID)

            modules.BOT.Privmsg("#" + modules.CONFIG.Section("IRC").Key("channel").String(),
                "\x0303" + *res.Repo.Name + "\x0F/\x0307" + branch + "\x0F " + id + ": \x0315" + message + "\x0F \x0313(By " + *commit.Author.Name + ")\x0F")
        }
    }
}
