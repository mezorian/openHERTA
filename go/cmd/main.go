package main

import (
	"sync"

	"github.com/mezorian/openHERTA/go/pkg/openHERTA"
)

func main() {
	// create a open HERTA struct
	var oH openHERTA.OpenHERTA

	// generate dummy data
	oH.CreateDummyData()

	// create wait group to be able to wait for
	// the http server which is running in a separate go routine
	wg := new(sync.WaitGroup)

	// start http server in a separate go routine
	// this will block until
	//  - either : the http server is panic-ing
	//  - or     : the http is gracefully with the Shutdown function
	oH.Run(wg)

	// wait for all go routines to finish
	wg.Wait()

}
