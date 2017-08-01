package routing

import (
	
	"net/http"
	//"time"
	
)

func WriteContent(w http.ResponseWriter, AText []byte) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)	
	w.Write( AText)
}