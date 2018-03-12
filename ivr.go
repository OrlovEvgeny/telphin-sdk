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
type IVR struct {
	LifetimeExpAction      string `json:"lifetime_exp_action"`
	EntryContext           uint64 `json:"entry_context"`
	LifetimeExpTransferDst string `json:"lifetime_exp_transfer_dst"`
	VmEnabled              bool   `json:"vm_enabled"`
	VmGreeting             uint64 `json:"vm_greeting"`
	Lifetime               uint64 `json:"lifetime"`
	VmAttachFile           bool   `json:"vm_attach_file"`
	SleepTime              uint64 `json:"sleep_time"`
	VmMailto               string `json:"vm_mailto"`
}

//
func (oauth *OAuth) GetIVR(extensionId uint32) (*IVR, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", API+"/extension/"+fmt.Sprint(extensionId)+"/ivr/", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+oauth.AccessToken)
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	ivr, err := ioutil.ReadAll(resp.Body)

	if err != nil || resp.StatusCode > 201 {
		log.Println(err)
		fmt.Println(resp.StatusCode)
		return nil, errors.New(string(ivr))
	}
	var ivrConfig *IVR
	err = json.Unmarshal(ivr, &ivrConfig)
	if err != nil {
		log.Println(err)
	}
	return ivrConfig, nil
}

//
func (oauth *OAuth) UpdateIVR(extensionId uint32, object ...interface{}) (*IVR, error) {
	jsonPayload, err := json.Marshal(object)
	if err != nil {
		log.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("PUT", API+"/extension/"+fmt.Sprint(extensionId)+"/ivr/", strings.NewReader(string(jsonPayload)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+oauth.AccessToken)
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	ivr, err := ioutil.ReadAll(resp.Body)

	if err != nil || resp.StatusCode > 201 {
		log.Println(err)
		fmt.Println(resp.StatusCode)
		return nil, errors.New(string(ivr))
	}
	var ivrConfig *IVR
	err = json.Unmarshal(ivr, &ivrConfig)
	if err != nil {
		log.Println(err)
	}
	return ivrConfig, nil
}
