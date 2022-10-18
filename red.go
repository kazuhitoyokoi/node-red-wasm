package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"syscall/js" // wasm
	"time"
)

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "testtest")
}

func execute(nodeId string, msg string) {
	var flowJson = []byte(flows())
	var flowItems []map[string]interface{}
	json.Unmarshal(flowJson, &flowItems)
	for i := 0; i < len(flowItems); i++ {
		var nodeType = flowItems[i]["type"]
		var currentNodeId = flowItems[i]["id"]
		var nodeWires = flowItems[i]["wires"].([]interface{})
		if nodeType == "inject" && nodeId == "" {
			go func() {
				ticker := time.NewTicker(time.Second)
				defer ticker.Stop()
				done := make(chan bool)
				go func() {
					time.Sleep(24 * time.Hour)
					done <- true
				}()
				for {
					select {
					case <-done:
						return
					case t := <-ticker.C:
						var msg2 = `{"payload":` + strconv.FormatInt(t.UnixMilli(), 10) + "}"
						for j := 0; j < len(nodeWires); j++ {
							var nodeWire = nodeWires[j].([]interface{})
							execute(nodeWire[0].(string), msg2)
						}
					}
				}
			}()
		} else if nodeType == "debug" && nodeId == currentNodeId {
			var msgData interface{}
			json.Unmarshal([]byte(msg), &msgData)
			var output = msgData.(map[string]interface{})["payload"]
			var text = ""
			if reflect.TypeOf(output).Kind() == reflect.Float64 {
				text = strconv.FormatFloat(output.(float64), 'f', -1, 64)
			} else if reflect.TypeOf(output).Kind() == reflect.Map {
				jsonData, _ := json.Marshal(output)
				text = string(jsonData)
			} else {
				text = output.(string)
			}
			fmt.Println(text)
			js.Global().Get("debug").Invoke(text) // wasm
		} else if nodeType == "http request" && nodeId == currentNodeId {
			var url = flowItems[i]["url"]
			resp, _ := http.Get(url.(string))
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			var msg2 = `{"payload":` + string(body) + "}"
			for j := 0; j < len(nodeWires); j++ {
				var nodeWire = nodeWires[j].([]interface{})
				execute(nodeWire[0].(string), msg2)
			}
		}
	}
}

func main() {
	execute("", "{}")
	http.Handle("/", new(handler))
	http.ListenAndServe(":1880", nil)
}
