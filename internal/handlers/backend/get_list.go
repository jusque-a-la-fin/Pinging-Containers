package backend

import (
	"encoding/json"
	"log"
	"monitoring/internal/handlers"
	"net/http"
)

func (hnd *BackendHandler) GetList(wrt http.ResponseWriter, rqt *http.Request) {
	if rqt.Method != http.MethodGet {
		errSend := handlers.SendBadReq(wrt)
		if errSend != nil {
			log.Printf("error while sending bad request message: %v\n", errSend)
		}
		return
	}

	ips, err := hnd.BackendRepo.GetList()
	if err != nil {
		log.Printf("error returned from function `GetList`, package `backend`: %v", err)
	}

	wrt.Header().Set("Content-Type", "application/json")
	wrt.WriteHeader(http.StatusOK)
	errJSON := json.NewEncoder(wrt).Encode(ips)
	if errJSON != nil {
		log.Printf("error while sending response: %v\n", errJSON)
	}
}
