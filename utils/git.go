package utils

import (
    "log"
    "strings"
    "net/url"
    "net/http"
)

func GetBranchName(ref string) string {
    return strings.Split(ref, "/")[2]
}

func GetShortCommitMessage(longMessage string) string {
    return strings.Split(longMessage, "\n")[0]
}

func GetShortCommitID(longId string) string {
    return longId[0:8]
}

func GetGitioUrl(longUrl string) string {
    response, err := http.PostForm("https://git.io", url.Values{ "url": { longUrl } })
    if err != nil {
        log.Println("Failed to shorten url", err)
        return longUrl
    }

    defer response.Body.Close()
    return response.Header.Get("Location")
}
