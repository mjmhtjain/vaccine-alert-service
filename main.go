package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/mjmhtjain/vaccine-alert-service/src/handler"
	"github.com/mjmhtjain/vaccine-alert-service/src/util"
)

const (
	DEFAULT_ADDR string = ":80"
)

//go:embed staticData/*
var embededFiles embed.FS

func init() {
	log.Println("init embededFiles ...")
	util.EmbededFiles = embededFiles
}

func main() {
	r := handler.Router()
	log.Fatal(http.ListenAndServe(DEFAULT_ADDR, r))
}
