package internal

import (
	"fmt"
	"github.com/zhanglifan/leaf_server/leaf/log"
	"net/http"
)

var ()

type Module struct {
}

func (m *Module) Run(closeSig chan bool) {
	log.Release("game web Run")
	//panic("implement me")
	log.Release("Web listener: %d", 11102)
	err := http.ListenAndServe(":11102", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}

func (m *Module) OnInit() {

}

func (m *Module) OnDestroy() {

}
