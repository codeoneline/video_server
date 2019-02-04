package main

import (
	"github.com/ethereum/go-ethereum/log"
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, sc int, errMsg string) {
	log.Error(errMsg)
	w.WriteHeader(sc)
	io.WriteString(w, errMsg)
}