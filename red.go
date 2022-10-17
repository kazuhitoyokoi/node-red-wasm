package main
import ("fmt"; "encoding/json"; "time"; "net/http")

type handler struct {
    n int
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "testtest")
}

func execute (nodeId string) {
    var flowJson = []byte(flows())
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
    http.Handle("/", new(handler))
    http.ListenAndServe(":1880", nil)
}
