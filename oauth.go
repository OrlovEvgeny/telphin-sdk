package telphin

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

//
type OAuth struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   uint32 `json:"expires_in"`
	Scope       string `json:"scope"`
	User        *User
}

//
type Trusted struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type PasswordAuth struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

//
type User struct {
	DealerID         string `json:"dealer_id"`
	IsAdmin          bool   `json:"admin"`
	Access           string `json:"access"`
	ExtensionId      uint32 `json:"extension_id"`
	ClientId         uint32 `json:"client_id"`
	ExtensionGroupID uint32 `json:"extension_group_id"`
	Timezone         string `json:"timezone"`
	Login            string `json:"login"`
	Id               uint32 `json:"id"`
}

//
var oauth *OAuth

func (p *PasswordAuth) PassowrdAuth() (*OAuth, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", API_URL+"/oauth/token", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	query := req.URL.Query()
	query.Add("grant_type", p.GrantType)
	query.Add("client_id", p.ClientID)
	query.Add("client_secret", p.ClientSecret)
	query.Add("username", p.Username)
	query.Add("password", p.Password)
	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	hosts, err := ioutil.ReadAll(resp.Body)

	if err != nil || resp.StatusCode > 201 {
		log.Fatal(err)
		return nil, errors.New(string(hosts))
	}

	err = json.Unmarshal(hosts, &oauth)
	if err != nil {
		log.Println(err)
	}
	oauth.getUser()

	go HeartBeat(p.PassowrdAuth)

	return oauth, nil
}

func (t *Trusted) TrustedAuth() (*OAuth, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", API_URL+"/oauth/token", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	query := req.URL.Query()
	query.Add("grant_type", t.GrantType)
	query.Add("client_id", t.ClientID)
	query.Add("client_secret", t.ClientSecret)
	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	hosts, err := ioutil.ReadAll(resp.Body)

	if err != nil || resp.StatusCode > 201 {
		log.Fatal(err)
		return nil, errors.New(string(hosts))
	}

	err = json.Unmarshal(hosts, &oauth)
	if err != nil {
		log.Println(err)
	}
	oauth.getUser()

	go HeartBeat(t.TrustedAuth)

	return oauth, nil
}

//
func (oauth *OAuth) getUser() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", API+"/user", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+oauth.AccessToken)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	me, err := ioutil.ReadAll(resp.Body)

	if err != nil || resp.StatusCode > 201 {
		log.Fatal(err)
	}
	err = json.Unmarshal(me, &oauth.User)
	if err != nil {
		log.Println(err)
	}
}
