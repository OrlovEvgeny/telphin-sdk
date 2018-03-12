package telphin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//
type Extension struct {
	Status             string `json:"status"`
	Domain             string `json:"domain"`
	CreateDate         string `json:"create_date"`
	Name               string `json:"name"`
	RFCid              bool   `json:"rfc_public_caller_id_number"`
	ExtensionGroupId   uint32 `json:"extension_group_id"`
	PublicCallerId     uint32 `json:"public_caller_id_number"`
	Label              string `json:"label"`
	CallerId           uint32 `json:"caller_id_name"`
	ClientId           uint32 `json:"client_id"`
	FromCallerId       bool   `json:"from_public_caller_id_number"`
	ExtraParams        string `json:"extra_params"`
	DialRuleLimit      string `json:"dial_rule_limit"`
	DialRuleId         string `json:"dial_rule_id"`
	AniRFC3325         bool   `json:"ani_rfc3325"`
	Type               string `json:"type"`
	ID                 uint32 `json:"id"`
	DidAsTransCallerId uint32 `json:"did_as_transfer_caller_id"`
	ExtensionNumber    string
}

//
func (oauth *OAuth) ExtensionList(ClientId uint32, setExtensionNumber ...bool) ([]Extension, error) {
	var extensn []Extension
	client := &http.Client{}
	req, err := http.NewRequest("GET", API+"/client/"+fmt.Sprint(ClientId)+"/extension/", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+oauth.AccessToken)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	ext, err := ioutil.ReadAll(resp.Body)

	if err != nil || resp.StatusCode > 201 {
		log.Println(err)
		log.Println("response body: ", string(ext))
		log.Fatal("Status code: ", resp.StatusCode)
	}
	err = json.Unmarshal(ext, &extensn)
	if err != nil {
		log.Println(err)
	}
	if len(setExtensionNumber) > 0 && setExtensionNumber[0] {
		for k, v := range extensn {
			s := strings.Split(v.Name, "*")
			extensn[k].ExtensionNumber = s[1]
		}
	}

	return extensn, nil
}
