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
            "\x0303[" + *res.Repository.Name  + "]\x0F " + *res.UserName + " pushed \x02" + strconv.Itoa(len(res.Commits)) + "\x0F commits to " + branch)

        for _, commit := range res.Commits {
            message := utils.GetShortCommitMessage(*commit.Message)
            id := utils.GetShortCommitID(*commit.ID)

            modules.BOT.Privmsg("#" + modules.CONFIG.Section("IRC").Key("channel").String(),
                "\x0303" + *res.Repository.Name + "\x0F/\x0308" + branch + "\x0F " + id + ": \x0315" + message + "\x0F \x0313(By " + *commit.Author.Name + ")\x0F")
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
