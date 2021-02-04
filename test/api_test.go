package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
	"testing"

	"github.com/mezorian/openHERTA/go/pkg/openHERTA"
	"github.com/stretchr/testify/assert"
)

func TestAPI(t *testing.T) {
	// start the back-end in a separate go routine
	var oH openHERTA.OpenHERTA
	wg := new(sync.WaitGroup)
	oH.CreateDummyData()
	oH.Run(wg)

	// prepare form-data
	data := url.Values{
		"firstName": {"John"},
		"lastName":  {"Doe"},
		"userID":    {"1234567"},
		"sessionID": {"98765"},
		"eventID":   {"75683"},
	}

	// execute http post request
	resp, err := http.PostForm("http://localhost:8080/updateGuestDetailsName", data)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, oH.Events["75683"].Name, "Lorem ipsum Grillabend")

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Println(res["form"])

	//assert.Equal(t, "", "", "")

	// shutdown the backend
	oH.Shutdown()
	wg.Wait()
}
