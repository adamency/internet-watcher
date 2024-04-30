package main

import (
  "context"
  "encoding/json"
  "fmt"
  "net/http"
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

func GetUserData() string {
  client := NewPatreonClient(oauth2.NoContext, GetApiToken())
  
  user, err := client.FetchUser()
  if err != nil {
    fmt.Println("error ", err.Error())
  }

  jsonUserData, err := json.MarshalIndent(user.Data, "", "  ")

  if err != nil {
    fmt.Println("error jsonuserdata: ", err)
  }

  return string(jsonUserData)
}

// Function to handle requests to /user
func userHandler(w http.ResponseWriter, r *http.Request) {

  userJSON := GetUserData()

  w.Header().Set("Content-Type", "application/json")

  w.Write([]byte(userJSON))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/text")
  w.Write([]byte("shlm"))
}

func main() {
  http.HandleFunc("/", rootHandler)
  http.HandleFunc("/user", userHandler)

  // Start the HTTP server
  http.ListenAndServe(":8796", nil)
}
