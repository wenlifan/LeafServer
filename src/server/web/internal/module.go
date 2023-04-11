package internal

import (
	"fmt"
	"net/http"
)

var ()

type Module struct {
}

func (m *Module) Run(closeSig chan bool) {
	//TODO implement me
	panic("implement me")
}

func (m *Module) OnInit() {
	err := http.ListenAndServe(":11102", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}

func (m *Module) OnDestroy() {

}
