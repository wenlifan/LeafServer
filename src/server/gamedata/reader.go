package gamedata

import (
	"github.com/zhanglifan/leaf_server/leaf/log"
	"github.com/zhanglifan/leaf_server/leaf/recordfile"
	"reflect"
)

func readRf(st interface{}) *recordfile.RecordFile {
	rf, err := recordfile.New(st)
	if err != nil {
		log.Fatal("%v", err)
	}
	fn := reflect.TypeOf(st).Name() + ".txt"
	err = rf.Read("gamedata/" + fn)
	if err != nil {
		log.Fatal("%v: %v", fn, err)
	}

	return rf
}
