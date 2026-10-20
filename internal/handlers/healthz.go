package handlers

import "net/http"

func Health(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	//explicitly setting status code of our own
	w.WriteHeader(http.StatusOK) //oder of header and write matters, if we write before setting header then it will not work because write will set the status code to 200 by default and then we cannot change it after that

	w.Write([]byte(`{"status":"ok"}`))
}
