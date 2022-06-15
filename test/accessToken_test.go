package test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
)

func TestGetAccessToken(t *testing.T) {
	t.Skip()
	type blizzOauthData struct {
		AccessToken string `json:"access_token"`
	}
	// env vars must be set before runtime
	clientId := os.Getenv("BLIZZ_CLIENT_ID")
	clientSecret := os.Getenv("BLIZZ_CLIENT_SECRET")
	params := url.Values{}
	params.Add("grant_type", `client_credentials`)
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", "https://us.battle.net/oauth/token", body)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(clientId, clientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data blizzOauthData
	jsonErr := json.Unmarshal(respData, &data)
	if jsonErr != nil {
		log.Fatal(err)
	}
	fmt.Println(data.AccessToken)
}
