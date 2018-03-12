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

type EventResp struct {
	Url       string `json:"url"`
	EventType string `json:"event_type"`
	Method    string `json:"method"`
	ID        uint32 `json:"id"`
}

//
func (oauth *OAuth) AddEvent(extensionId uint32, callbackUrl string, method string, eventType string) (*EventResp, error) {
	AllowMethod := map[string]string{"POST": "POST", "GET": "GET"}
	if _, ok := AllowMethod[strings.ToUpper(method)]; !ok {
		return nil, errors.New("Method " + method + " Not allowed")
	}

	var eventResp *EventResp

	d := map[string]string{
		"url":        callbackUrl,
		"event_type": strings.ToLower(eventType),
		"method":     strings.ToUpper(method),
	}

	jsonPayload, _ := json.Marshal(d)
	client := &http.Client{}
	req, err := http.NewRequest(strings.ToUpper(method), API+"/extension/"+fmt.Sprint(extensionId)+"/event/", strings.NewReader(string(jsonPayload)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+oauth.AccessToken)
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	event, err := ioutil.ReadAll(resp.Body)

	if err != nil || resp.StatusCode > 201 {
		log.Println(err)
		fmt.Println(resp.StatusCode)
		return nil, errors.New(string(event))
	}
	fmt.Println(event)
	err = json.Unmarshal(event, &eventResp)
	if err != nil {
		log.Println(err)
	}

	return eventResp, nil
}

//
func (oauth *OAuth) GetEventList(extensionId uint32) ([]EventResp, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", API+"/extension/"+fmt.Sprint(extensionId)+"/event", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+oauth.AccessToken)
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	event, err := ioutil.ReadAll(resp.Body)

	if err != nil || resp.StatusCode > 201 {
		log.Println(err)
		fmt.Println(resp.StatusCode)
		return nil, errors.New(string(event))
	}
	var eventResp []EventResp
	err = json.Unmarshal(event, &eventResp)
	if err != nil {
		log.Println(err)
	}
	return eventResp, nil
}

//
func (oauth *OAuth) GetEventByID(extensionId uint32, eventId uint32) (*EventResp, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", API+"/extension/"+fmt.Sprint(extensionId)+"/event/"+fmt.Sprint(eventId), nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+oauth.AccessToken)
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	event, err := ioutil.ReadAll(resp.Body)

	if err != nil || resp.StatusCode > 201 {
		log.Println(err)
		fmt.Println(resp.StatusCode)
		return nil, errors.New(string(event))
	}
	var eventResp *EventResp
	err = json.Unmarshal(event, &eventResp)
	if err != nil {
		log.Println(err)
	}
	return eventResp, nil
}

//
func (oauth *OAuth) DeleteEventAll(extensionId uint32) error {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", API+"/extension/"+fmt.Sprint(extensionId)+"/event/", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+oauth.AccessToken)
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()

	event, err := ioutil.ReadAll(resp.Body)
	if err != nil || resp.StatusCode > 204 {
		log.Println(err)
		fmt.Println(resp.StatusCode)
		return err
	}
	var eventResp *EventResp
	err = json.Unmarshal(event, &eventResp)
	if err != nil {
		log.Println(err)
	}
	return nil
}

//
func (oauth *OAuth) DeleteEventByID(extensionId uint32, eventId uint32) error {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", API+"/extension/"+fmt.Sprint(extensionId)+"/event/"+fmt.Sprint(eventId), nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+oauth.AccessToken)
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()

	event, err := ioutil.ReadAll(resp.Body)
	if err != nil || resp.StatusCode > 204 {
		log.Println(err)
		fmt.Println(resp.StatusCode)
		return err
	}
	var eventResp *EventResp
	err = json.Unmarshal(event, &eventResp)
	if err != nil {
		log.Println(err)
	}
	return nil
}

//
func (oauth *OAuth) UpdateEventByID(extensionId uint32, callbackUrl string, method string, eventType string, eventId uint32) (*EventResp, error) {
	var eventResp *EventResp

	d := map[string]string{
		"url":        callbackUrl,
		"event_type": strings.ToLower(eventType),
		"method":     strings.ToUpper(method),
	}

	jsonPayload, _ := json.Marshal(d)
	client := &http.Client{}
	req, err := http.NewRequest("PUT", API+"/extension/"+fmt.Sprint(extensionId)+"/event/"+fmt.Sprint(eventId), strings.NewReader(string(jsonPayload)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+oauth.AccessToken)
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	event, err := ioutil.ReadAll(resp.Body)

	if err != nil || resp.StatusCode > 201 {
		log.Println(err)
		fmt.Println(resp.StatusCode)
		return nil, errors.New(string(event))
	}
	fmt.Println(event)
	err = json.Unmarshal(event, &eventResp)
	if err != nil {
		log.Println(err)
	}

	return eventResp, nil
}
