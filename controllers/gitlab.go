package controllers

import (
    "strconv"
    "io/ioutil"
    "encoding/json"
    "github.com/jamierocks/CommitBot/utils"
    "github.com/jamierocks/CommitBot/modules"
    "gopkg.in/macaron.v1"
)

func GetGitlab(ctx *macaron.Context) {
    if (ctx.Req.Header.Get("X-Gitlab-Event") == "Push Hook") {
        body, _ := ioutil.ReadAll(ctx.Req.Body().ReadCloser())

        var res PushEvent
        json.Unmarshal(body, &res)

        branch := utils.GetBranchName(*res.Ref)

        modules.BOT.Privmsg("#" + modules.CONFIG.Section("IRC").Key("channel").String(),
            "[" + *res.Repository.Name + "] " + *res.UserName + " pushed " + strconv.Itoa(len(res.Commits)) + " commits to " + branch)

        for _, commit := range res.Commits {
            message := utils.GetShortCommitMessage(*commit.Message)
            id := utils.GetShortCommitID(*commit.ID)

            modules.BOT.Privmsg("#" + modules.CONFIG.Section("DISCORD").Key("channel").String(),
                *res.Repository.Name + "/" + branch + " " + id + ": " + message + " (By " + *commit.Author.Name + ")")
        }
    }
}

type PushEvent struct {
    UserName *string `json:"user_name"`
    Ref *string `json:"ref"`
    Commits []Commit `json:"commits"`
    Repository Repository `json:"repository"`
}

type Repository struct {
    Name *string `json:"name"`
}

type Commit struct {
    ID *string `json:"id"`
    Message *string `json:"message"`
    Author CommitAuthor `json:"author"`
}

type CommitAuthor struct {
    Name *string `json:"name"`
}
