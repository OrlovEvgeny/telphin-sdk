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


# License:

[MIT](LICENSE)
