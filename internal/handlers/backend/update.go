package backend

import (
	"encoding/json"
	"log"
	"monitoring/internal/backend"
)

type Container struct {
	IPv4            string `json:"IPv4"`
	PingDuration    string `json:"PingDuration"`
	SuccessPingTime string `json:"SuccessPingTime"`
	IsSuccess       bool   `json:"IsSuccess"`
}

func (hnd *BackendHandler) UpdateContainers(jsonData []byte) {
	var resp Container
	err := json.Unmarshal(jsonData, &resp)
	if err != nil {
		log.Println("error unmarshaling JSON:", err)
		return
	}

	ctr := backend.Container{}
	ctr.IPv4 = resp.IPv4
	ctr.IsSuccess = resp.IsSuccess
	ctr.PingDuration = resp.PingDuration
	ctr.SuccessPingTime = resp.SuccessPingTime
	err = hnd.BackendRepo.UpdateContainer(ctr)
	if err != nil {
		log.Printf("error returned from method `UpdateContainer`, package `backend`: %v", err)
	}
}
