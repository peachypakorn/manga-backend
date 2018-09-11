package main

import (
	"net/http"
	log "github.com/Sirupsen/logrus"
)

func StoreCreate(w http.ResponseWriter, r *http.Request) {
	log.Debugln("StoreCreate")

	//if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
	//	log.Errorln("decode json error:", err)
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//defer r.Body.Close()


	JSONResponse(w, http.StatusCreated, ResponseNormal{
		"00",
		"create store success",
	})

}