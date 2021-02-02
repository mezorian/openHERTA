package main

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/mezorian/openHERTA/go/pkg/openHERTA"
	"github.com/stretchr/testify/assert"
)

func TestAPI(t *testing.T) {
	var oH openHERTA.OpenHERTA

	wg := new(sync.WaitGroup)
	oH.Run(wg)
	log.Printf("Hello World")
	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())

	time.Sleep(10 * time.Second)

	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
	assert.Equal(t, "", "", "")
	oH.Shutdown()
	wg.Wait()
}
