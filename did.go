package telphin

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//
type Did struct {
	ID          uint32 `json:"id"`
	Name        string `json:"name"`
	Domain      string `json:"domain"`
	ExtensionId uint32 `json:"extension_id"`
	ClientId    uint32 `json:"client_id"`
}

//
func (oauth *OAuth) GetDidList() ([]Did, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", API+"/client/"+fmt.Sprint(oauth.User.ClientId)+"/did/", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+oauth.AccessToken)
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	did, err := ioutil.ReadAll(resp.Body)

	if err != nil || resp.StatusCode > 201 {
		log.Println(err)
		fmt.Println(resp.StatusCode)
		return nil, errors.New(string(did))
	}
	var Didresult []Did
	err = json.Unmarshal(did, &Didresult)
	if err != nil {
		log.Println(err)
	}
	return Didresult, nil
}

//
func (oauth *OAuth) GetDidById(didId uint64) (*Did, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", API+"/client/"+fmt.Sprint(oauth.User.ClientId)+"/did/"+fmt.Sprint(didId), nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+oauth.AccessToken)
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	did, err := ioutil.ReadAll(resp.Body)

	if err != nil || resp.StatusCode > 201 {
		log.Println(err)
		fmt.Println(resp.StatusCode)
		return nil, errors.New(string(did))
	}
	var Didresult *Did
	err = json.Unmarshal(did, &Didresult)
	if err != nil {
		log.Println(err)
	}
	return Didresult, nil
}

//
func (oauth *OAuth) UpdateDidById(didId uint64, ExtensionId uint32) (*Did, error) {
	client := &http.Client{}
	d := map[string]uint32{
		"extension_id": ExtensionId,
	}

	jsonPayload, _ := json.Marshal(d)
	req, err := http.NewRequest("PUT", API+"/client/"+fmt.Sprint(oauth.User.ClientId)+"/did/"+fmt.Sprint(didId), strings.NewReader(string(jsonPayload)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+oauth.AccessToken)
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	did, err := ioutil.ReadAll(resp.Body)

	if err != nil || resp.StatusCode > 201 {
		log.Println(err)
		fmt.Println(resp.StatusCode)
		return nil, errors.New(string(did))
	}
	var Didresult *Did
	err = json.Unmarshal(did, &Didresult)
	if err != nil {
		log.Println(err)
	}
	return Didresult, nil
}
