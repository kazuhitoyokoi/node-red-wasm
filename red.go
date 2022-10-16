package main
import ("fmt"; "encoding/json"; "time")

func execute (nodeId string) {
    var flowJson = []byte(`[{"id":"9677e77d09b17122","type":"inject","z":"979ccb452915f125","name":"","props":[{"p":"payload"},{"p":"topic","vt":"str"}],"repeat":"1","crontab":"","once":false,"onceDelay":0.1,"topic":"","payload":"","payloadType":"date","x":240,"y":160,"wires":[["3552c8ecdf006278"]]},{"id":"3552c8ecdf006278","type":"debug","z":"979ccb452915f125","name":"debug 1","active":true,"tosidebar":true,"console":true,"tostatus":false,"complete":"payload","targetType":"msg","statusVal":"","statusType":"auto","x":450,"y":160,"wires":[]}]`)
    var flowItems []map[string]interface {}
    json.Unmarshal(flowJson, &flowItems)
    for i := 0; i < len(flowItems); i++ {
        var nodeType = flowItems[i]["type"]
        if nodeType == "inject" {
            fmt.Println("hoge injecthoge")
            go func() {
            ticker := time.NewTicker(time.Second)
            defer ticker.Stop()
            done := make(chan bool)
            go func() {
                time.Sleep(10 * time.Second)
                done <- true
            }()
            for {
                select {
                case <-done:
                    fmt.Println("Done!")
                    return
                case t := <-ticker.C:
                    fmt.Println("Current time: ", t)
                }
            }
            }()
        } else if nodeType == "debug" {
            fmt.Println("hoge debughoge")
        }
        var nodeWires = flowItems[i]["wires"].([]interface {})
        for j := 0; j < len(nodeWires); j++ {
            var nodeWire = nodeWires[j].([]interface {})
            fmt.Println(nodeWire[0])
        }
    }
}

func main () {
    execute("")
    time.Sleep(24 * time.Hour)
}
