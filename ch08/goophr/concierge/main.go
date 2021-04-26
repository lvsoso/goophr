package main

import (
        "fmt"
        "concierge/api"
        "concierge/common"
        "net/http"
        "os"
)

func main(){
        common.Log("Adding API handles...")
        http.HandleFunc("/api/feeder", api.FeedHandler)
        http.HandleFunc("/api/query", api.QueryHandler)

        common.Log("Starting feeder...")
        api.StartFeederSystem()

        port := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
        common.Log(fmt.Sprintf("Starting Goophr Concierge server on port %s...", port))
        http.ListenAndServe(port, nil)
}

