package backend

import (
	"encoding/json"
	"log"
	"monitoring/internal/handlers"
	"net/http"
)

func (hnd *BackendHandler) GetLogs(wrt http.ResponseWriter, rqt *http.Request) {
	if rqt.Method != http.MethodGet {
		errSend := handlers.SendBadReq(wrt)
		if errSend != nil {
			log.Printf("error while sending bad request message: %v\n", errSend)
		}
		return
	}

	cns, err := hnd.BackendRepo.GetLogs()
	if err != nil {
		log.Printf("error returned from method `GetLogs`, package `backend`: %v", err)
	}

	wrt.Header().Set("Content-Type", "application/json")
	wrt.WriteHeader(http.StatusOK)
	errJSON := json.NewEncoder(wrt).Encode(cns)
	if errJSON != nil {
		log.Printf("error while sending response: %v\n", errJSON)
	}
}
