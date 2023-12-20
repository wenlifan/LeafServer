package internal

import (
	"fmt"
	"net/http"
)

func GetApList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "{\"SecretKey\":\"\",\"BannedDescId\":0,\"AccessPointInfo\":{\"APList\":[{\"APAddress\":\"192.168.61.155:3563\",\"Name\":\"default\",\"APDomain\":\"sso.digisky.com\",\"ID\":0}],\"LastAPId\":1},\"BannedCode\":0,\"Succeed\":true,\"ErrorDesc\":\"\"}")
}

func init() {
	http.HandleFunc("/auth_player", GetApList)
}
