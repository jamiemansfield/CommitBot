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

func GetGitlab(ctx *macaron.Context) {
    if (ctx.Req.Header.Get("X-Gitlab-Event") == "Push Hook") {
        body, _ := ioutil.ReadAll(ctx.Req.Body().ReadCloser())

        var res github.PushEvent
        json.Unmarshal(body, &res)

        branch := utils.GetBranchName(*res.Ref)

        modules.BOT.Privmsg("#" + modules.CONFIG.Section("IRC").Key("channel").String(),
            "[" + *res.Repo.Name + "] " + *res.Pusher.Name + " pushed " + strconv.Itoa(len(res.Commits)) + " commits to " + branch)

        for _, commit := range res.Commits {
            message := utils.GetShortCommitMessage(*commit.Message)
            id := utils.GetShortCommitID(*commit.ID)

            modules.BOT.Privmsg("#" + modules.CONFIG.Section("DISCORD").Key("channel").String(),
                *res.Repo.Name + "/" + branch + " " + id + ": " + message + " (By " + *commit.Author.Name + ")")
        }
    }
}