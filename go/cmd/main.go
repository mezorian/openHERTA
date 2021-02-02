package main

import (
	"sync"

	"github.com/mezorian/openHERTA/go/pkg/openHERTA"
)

func main() {
	var oH openHERTA.OpenHERTA

	wg := new(sync.WaitGroup)
	oH.Run(wg)
	wg.Wait()

}
