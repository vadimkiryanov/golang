package pkgmain

import (
	"log"
	"main-mode/pkg/20_nethttp/coincap"
	"time"
)

func Main() {
	coincapClient, err := coincap.NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
		return
	}

	data, err := coincapClient.GetAssets(true)

	if err != nil {
		log.Fatal(err)
		return
	}

	for _, v := range data {
		v.GetInfo()
	}
}
