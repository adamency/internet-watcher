package main

import (
  "context"
  "encoding/json"
  "fmt"
  "os"
  "strings"

  "golang.org/x/oauth2"
  patreon "gopkg.in/mxpv/patreon-go.v1"
)

func GetApiToken() string {
  bitContent, err := os.ReadFile("token.txt")
  if err != nil {
    panic(err)
  }
  
  content := string(bitContent)
  content = strings.TrimSuffix(content, "\n")

  return content
}

func NewPatreonClient(ctx context.Context, token string) *patreon.Client {
  ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
  tc := oauth2.NewClient(ctx, ts)

  client := patreon.NewClient(tc)
  return client
}

func main() {
  client := NewPatreonClient(oauth2.NoContext, GetApiToken())
  
  user, err := client.FetchUser()
  if err != nil {
    fmt.Println("error ", err.Error())
  }

  jsonUserData, err := json.MarshalIndent(user.Data, "", "  ")

  if err != nil {
    fmt.Println("error jsonuserdata: ", err)
  }

  fmt.Print(string(jsonUserData))
}
