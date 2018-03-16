package telphin

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"strings"
	"time"
)

//
type Storage struct {
	Url string `json:"record_url"`
}

//
type RecordList struct {
	ToUsername            string `json:"to_username"`
	SourceNumber          string `json:"source_number"`
	Result                string `json:"result"`
	Duration              uint32 `json:"duration"`
	ToDomain              string `json:"to_domain"`
	DidNumber             string `json:"did_number"`
	HangupTimeGMT         string `json:"hangup_time_gmt"`
	RecordFileSize        uint64 `json:"record_file_size"`
	FromUsername          string `json:"from_username"`
	Application           string `json:"application"`
	StartTimeGMT          string `json:"start_time_gmt"`
	ExtNumberReg          string `json:"ext_number_reg"`
	HangupCause           string `json:"hangup_cause"`
	ExtensionGroupOwnerID uint64 `json:"extension_group_owner_id"`
	InitTimeGMT           string `json:"init_time_gmt"`
	DestDomain            string `json:"dest_domain"`
	RecordUUID            string `json:"record_uuid"`
	FromDomain            string `json:"from_domain"`
	SourceDomain          string `json:"source_domain"`
	ExtensionType         string `json:"extension_type"`
	DidDomain             string `json:"did_domain"`
	FromScreenName        string `json:"from_screen_name"`
	ExtensionName         string `json:"extension_name"`
	Flow                  string `json:"flow"`
	DestNumber            string `json:"dest_number"`
	ExtensionId           uint32 `json:"extension_id"`
	CallUUID              string `json:"call_uuid"`
	ClientOwnerID         uint64 `json:"client_owner_id"`
}

//
func (oauth *OAuth) DownloadRecord(recordUUID string, path string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", API+"/client/"+fmt.Sprint(oauth.User.ClientId)+"/record/"+fmt.Sprint(recordUUID), nil)
	req.Header.Add("Authorization", "Bearer "+oauth.AccessToken)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	if err != nil || resp.StatusCode > 201 {
		log.Println(err)
		fmt.Println(resp.StatusCode)
	}

	_, params, _ := mime.ParseMediaType(resp.Header.Get("Content-Disposition"))
	var fileName = params["filename"]

	output, err := os.Create(path + fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return "Error while creating", err
	}
	defer output.Close()
	n, err := io.Copy(output, resp.Body)
	if err != nil {
		fmt.Println("Error while downloading", fileName, "-", err)
		return "Error while downloading", err
	}
	fmt.Println(n, "bytes downloaded.")
	return path + fileName, nil
}

//
func (oauth *OAuth) GetRecordStorageUrl(recordUUID string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", API+"/client/"+fmt.Sprint(oauth.User.ClientId)+"/record/"+fmt.Sprint(recordUUID)+"/storage_url/", nil)
	req.Header.Add("Authorization", "Bearer "+oauth.AccessToken)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	url, err := ioutil.ReadAll(resp.Body)

	if err != nil || resp.StatusCode > 201 {
		log.Println(err)
		fmt.Println(resp.StatusCode)
	}
	var storage *Storage
	err = json.Unmarshal(url, &storage)
	if err != nil {
		log.Println(err)
	}

	return storage.Url
}

//
func (oauth *OAuth) GetRecordList(extensionId uint32, startDate time.Time, endDate time.Time, order string) ([]RecordList, error) {
	AllowOrder := map[string]string{"asc": "asc", "desc": "desc"}
	if _, ok := AllowOrder[strings.ToLower(order)]; !ok {
		return nil, errors.New("order to " + order + " Not allowed")
	}

	startDateStr := startDate.Format(DATE_FORMAT)
	endDateStr := endDate.Format(DATE_FORMAT)

	client := &http.Client{}
	req, err := http.NewRequest("GET", API+"/extension/"+fmt.Sprint(extensionId)+"/record/", nil)
	req.Header.Add("Authorization", "Bearer "+oauth.AccessToken)
	queryStr := req.URL.Query()
	queryStr.Add("start_datetime", startDateStr)
	queryStr.Add("end_datetime", endDateStr)
	queryStr.Add("order", strings.ToLower(order))
	req.URL.RawQuery = queryStr.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	records, err := ioutil.ReadAll(resp.Body)

	if err != nil || resp.StatusCode > 201 {
		log.Println(err)
		fmt.Println(resp.StatusCode)
		return nil, errors.New(string(records))
	}

	var recordList []RecordList
	err = json.Unmarshal(records, &recordList)
	if err != nil {
		log.Println(err)
	}

	return recordList, nil
}
