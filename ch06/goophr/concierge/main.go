package main

import (
        "goophr/concierge/api"
        "goophr/concierge/common"
        "net/http"
)

func main(){
        common.Log("Adding API handles...")
        http.HandleFunc("/api/feeder", api.FeedHandler)

        common.Log("Starting feeder...")
        api.StartFeederSystem()

        common.Log("Starting Goophr Concierge server on port :8080...")
        http.ListenAndServe(":8080", nil)
}

