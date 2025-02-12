package backend

import (
	"encoding/json"
	"log"
	"monitoring/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func (hnd *BackendHandler) GetContainer(wrt http.ResponseWriter, rqt *http.Request) {
	if rqt.Method != http.MethodGet {
		errSend := handlers.SendBadReq(wrt)
		if errSend != nil {
			log.Printf("error while sending bad request message: %v\n", errSend)
		}
		return
	}

	ipv4 := mux.Vars(rqt)["id"]
	ips, err := hnd.BackendRepo.GetContainer(ipv4)
	if err != nil {
		log.Printf("error returned from function `GetContainer`, package `backend`: %v", err)
	}

	wrt.Header().Set("Content-Type", "application/json")
	wrt.WriteHeader(http.StatusOK)
	errJSON := json.NewEncoder(wrt).Encode(ips)
	if errJSON != nil {
		log.Printf("error while sending response: %v\n", errJSON)
	}
}
