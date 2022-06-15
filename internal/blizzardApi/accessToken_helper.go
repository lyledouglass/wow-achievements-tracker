package blizzardApi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func GetAccessToken() string {
	type blizzOauthData struct {
		AccessToken string `json:"access_token"`
	}

	clientId := os.Getenv("BLIZZ_CLIENT_ID")
	clientSecret := os.Getenv("BLIZZ_CLIENT_SECRET")
	params := url.Values{}
	params.Add("grant_type", `client_credentials`)
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", "http://us.battle.net/oauth/token", body)
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
	return data.AccessToken
}
