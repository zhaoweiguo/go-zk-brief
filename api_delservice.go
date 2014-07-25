package main

import(
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/samuel/go-zookeeper/zk"
)

func delservice(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(DEFAULT_MIN_MEMORY)

	keys := r.Form["key"]
	destNames := r.Form["destName"]
	zkidcs := r.Form["zkidc"]

	input := fmt.Sprintf("keys:%v, zkidcs:%v, destNames:%v", keys, zkidcs, destNames)
	api := "delservice"
	defer handleError(w, input, api)


	// 参数检验
	checkParams(keys, destNames, zkidcs)
	// 判断key是否正确
	checkKeys(keys[0])

	c, _, err := zk.Connect([]string{ZKHOST[zkidcs[0]]}, ZKTIMEOUT)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	zkServerPath := ZKPATH + "/" + destNames[0]

//	fmt.Println(zkServerPath)
	err = c.Delete(zkServerPath, -1)

	if err != nil {
		panic(err)
	}
	rtnNormal := &RtnNormal {
		Code : 1,
	}
	rtnJson, _ := json.Marshal(rtnNormal)
	rtnStr := string(rtnJson)
	fmt.Fprintf(w, rtnStr)

	apilog(input, api, rtnStr)   // 日志记录

}

