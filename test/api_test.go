package main

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/mezorian/openHERTA/go/pkg/openHERTA"
)

func TestAPI(t *testing.T) {
	var oH openHERTA.OpenHERTA

	wg := new(sync.WaitGroup)
	oH.Run(wg)
	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
	time.Sleep(10 * time.Second)
	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
	oH.Shutdown(wg)
	wg.Wait()
}

//
// data := url.Values{
// 	"firstName": {"John"},
// 	"lastName":  {"Doe"},
// 	"userID":    {"1234567"},
// 	"sessionID": {"98765"},
// 	"eventID":   {"75683"},
// }
//
// resp, err := http.PostForm("http://localhost:8080/updateGuestDetailsName", data)
//
// if err != nil {
// 	log.Fatal(err)
// }
//
// var res map[string]interface{}
// json.NewDecoder(resp.Body).Decode(&res)
// //fmt.Println(res["form"])
//
