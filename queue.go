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
type Queue struct {
	Sound                     uint64 `json:"sound"`
	TierRulesApply            bool   `json:"tier_rules_apply"`
	ExitKey                   string `json:"exit_key"`
	TimeoutTransferDst        string `json:"timeout_transfer_dst"`
	MaxWaitTime               uint64 `json:"max_wait_time"`
	RecordEnabled             bool   `json:"record_enabled"`
	MaxWaitTimeWithNoAgent    uint64 `json:"max_wait_time_with_no_agent"`
	Strategy                  string `json:"strategy"`
	TierRuleWaitMultiplyLevel bool   `json:"tier_rule_wait_multiply_level"`
	ExtensionID               uint64 `json:"extension_id"`
	AnnounceFrequency         uint64 `json:"announce_frequency"`
	ExitTransferDst           string `json:"exit_transfer_dst"`
	SoundType                 string `json:"sound_type"`
	TierRuleWaitSecond        uint64 `json:"tier_rule_wait_second"`
	AnnounceSound             uint64 `json:"announce_sound"`
	StartSound                uint64 `json:"start_sound"`
}

//
func (oauth *OAuth) GetQueue(extensionId uint32) (*Queue, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", API+"/extension/"+fmt.Sprint(extensionId)+"/queue/", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+oauth.AccessToken)
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	queue, err := ioutil.ReadAll(resp.Body)

	if err != nil || resp.StatusCode > 201 {
		log.Println(err)
		fmt.Println(resp.StatusCode)
		return nil, errors.New(string(queue))
	}
	var queueConfig *Queue
	err = json.Unmarshal(queue, &queueConfig)
	if err != nil {
		log.Println(err)
	}
	return queueConfig, nil
}

//
func (oauth *OAuth) UpdateQueue(extensionId uint32, object ...interface{}) (*IVR, error) {
	jsonPayload, err := json.Marshal(object)
	if err != nil {
		log.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("PUT", API+"/extension/"+fmt.Sprint(extensionId)+"/queue/", strings.NewReader(string(jsonPayload)))
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
