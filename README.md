# Telphin-SDK

[![Go Report Card](https://goreportcard.com/badge/github.com/OrlovEvgeny/telphin-sdk)](https://goreportcard.com/report/github.com/OrlovEvgeny/telphin-sdk)
[![GoDoc](https://godoc.org/github.com/OrlovEvgeny/telphin-sdk?status.svg)](https://godoc.org/github.com/OrlovEvgeny/telphin-sdk)

**golang lib api [Telphin billing](https://ringme-confluence.atlassian.net/wiki/spaces/RAL/pages)**

# Installation

```bash
~ $ go get -u gopkg.in/OrlovEvgeny/telphin-sdk.v1
```

Example Trusted auth

````golang
func main() {
        trusted := &telphin.Trusted {
    		GrantType:    "client_credentials",
    		ClientID:     <your client_id>,
    		ClientSecret: <your client_secret>,
    	}
    
    	api, err := trusted.TrustedAuth()
    	if err != nil {
    		fmt.Println(err)
    	}
    
    	fmt.Println("You telphin access token: ", api.AccessToken)
    	fmt.Println("You telphin token expires in: ", api.ExpiresIn)
    	fmt.Println("You telphin Login: ", api.User.Login)
    	fmt.Println("You telphin Clinet ID: ", api.User.ClientId)
    	fmt.Println("You telphin is Admin: ", api.User.IsAdmin)
}
	
````


Example get Extension list
````golang

    exten, err := api.ExtensionList(api.User.ClientId, true)
    if err != nil {
        fmt.Println(err)
    }
    
	for _, e := range exten {
		fmt.Printf("Extension id: %d, Extension name: %s \n", e.ID, e.Name)

	}
````

Example download records

````golang
    
    exten, err := api.ExtensionList(api.User.ClientId, true)
    if err != nil {
       fmt.Println(err)
    }
    
    t := time.Now()
	rounded := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	recordList, err := api.GetRecordList(exten[0].ID, rounded, time.Now(), "asc")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(recordList))
	for _, v := range recordList {
		api.DownloadRecord(v.RecordUUID, "./downloads/")
	}
````


**All methods:**
* *[AddEvent](https://godoc.org/github.com/OrlovEvgeny/telphin-sdk#OAuth.AddEvent)*
* *[DeleteEventAll](https://godoc.org/github.com/OrlovEvgeny/telphin-sdk#OAuth.DeleteEventAll)*
* *[DeleteEventByID](https://godoc.org/github.com/OrlovEvgeny/telphin-sdk#OAuth.DeleteEventByID)*
* *[DownloadRecord](https://godoc.org/github.com/OrlovEvgeny/telphin-sdk#OAuth.DownloadRecord)*
* *[ExtensionList](https://godoc.org/github.com/OrlovEvgeny/telphin-sdk#OAuth.ExtensionList)*
* *[GetDidById](https://godoc.org/github.com/OrlovEvgeny/telphin-sdk#OAuth.GetDidById)*
* *[GetDidList](https://godoc.org/github.com/OrlovEvgeny/telphin-sdk#OAuth.GetDidList)*
* *[GetEventByID](https://godoc.org/github.com/OrlovEvgeny/telphin-sdk#OAuth.GetEventByID)*
* *[GetEventList](https://godoc.org/github.com/OrlovEvgeny/telphin-sdk#OAuth.GetEventList)*
* *[GetIVR](https://godoc.org/github.com/OrlovEvgeny/telphin-sdk#OAuth.GetIVR)*
* *[UpdateIVR](https://godoc.org/github.com/OrlovEvgeny/telphin-sdk#OAuth.UpdateIVR)*
* *[GetQueue](https://godoc.org/github.com/OrlovEvgeny/telphin-sdk#OAuth.GetQueue)*
* *[GetRecordList](https://godoc.org/github.com/OrlovEvgeny/telphin-sdk#OAuth.GetRecordList)*
* *[GetRecordStorageUrl](https://godoc.org/github.com/OrlovEvgeny/telphin-sdk#OAuth.GetRecordStorageUrl)*
* *[UpdateDidById](https://godoc.org/github.com/OrlovEvgeny/telphin-sdk#OAuth.UpdateDidById)*
* *[UpdateEventByID](https://godoc.org/github.com/OrlovEvgeny/telphin-sdk#OAuth.UpdateEventByID)*
* *[UpdateQueue](https://godoc.org/github.com/OrlovEvgeny/telphin-sdk#OAuth.UpdateQueue)*


# License:

[MIT](LICENSE)
