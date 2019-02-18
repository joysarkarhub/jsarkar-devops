package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joysarkarhub/jsarkar-devops/model"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

var appPort = ":9090"
var projectName = "BranchMessenger DevOps App"

type DockerData struct {
	DockerHostname    string
	DockerContainerID string
	DockerCreated     string
	DockerMacAddress  string
	DockerIPAddress   string
	DockerImage       string
}

func init() {
	model.Setenv()
	model.SetIP()
	model.SetMac()
	model.FetchContainerID()
	model.SetenvContainerID()
}

func getEnv(key, fetchDefault string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fetchDefault
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response, err := getDockerJson()
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)

	log.Infof("200; %s => %s; Host => %s, From => %s; User-Agent => %s",
		r.Method,
		r.URL.Path,
		r.Host,
		r.RemoteAddr,
		r.Header.Get("User-Agent"))

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func NotFound(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNotFound)

	log.Infof("404; %s => %s; Host => %s, From => %s; User-Agent => %s",
		r.Method,
		r.URL.Path,
		r.Host,
		r.RemoteAddr,
		r.Header.Get("User-Agent"))

	w.Write([]byte("DevOps Challenge - Page does not Exist!"))
}

func main() {
	log.Infof("Starting %s on %s", projectName, appPort)
	router := httprouter.New()
	router.GET("/", Index)
	router.NotFound = http.HandlerFunc(NotFound)
	log.Fatal(http.ListenAndServe(appPort, router))
}

func getDockerJson() ([]byte, error) {

	data := DockerData{
		DockerHostname:    getEnv("DockerHostname", fmt.Sprintf("Couldn't locate Hostname")),
		DockerContainerID: getEnv("DockerContainerID", fmt.Sprintf("ID not found!")),
		DockerCreated:     getEnv("DockerCreated", fmt.Sprintf("Couldn't locate time created")),
		DockerMacAddress:  getEnv("DockerMacAddress", fmt.Sprintf("Couldn't locate MacAddress")),
		DockerIPAddress:   getEnv("DockerIPAddress", fmt.Sprintf("Couldn't locate IPAddress")),
		DockerImage:       getEnv("DockerImage", fmt.Sprintf("Couldn't locate Image")),
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		panic(err)
	}

	return jsonData, nil
}
