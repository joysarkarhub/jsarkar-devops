package model

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
)

// Setenv some env variable
func Setenv() {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	os.Setenv("DockerHostname", hostname)
}

// SetIP current ip as env var
func SetIP() {
	name, err := os.Hostname()
	if err != nil {
		fmt.Printf("Oops: %v\n", err)
		return
	}

	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Printf("Oops: %v\n", err)
		return
	}

	for _, a := range addrs {
		os.Setenv("DockerIPAddress", a)
	}
}

// SetMac to fetch current MAC Address
func SetMac() {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	for _, i := range interfaces {
		if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
			addr := i.HardwareAddr.String()
			os.Setenv("DockerMacAddress", addr)
		}
	}
}

// FetchContainerID to get container id to env
func FetchContainerID() {

	cmdStr := "/go/bin/script.sh"
	cmd := exec.Command("/bin/sh", cmdStr)
	_, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}
}

// SetenvContainerID read from file and set env var
func SetenvContainerID() {
	dat, err := ioutil.ReadFile("/go/bin/containerid")
	if err != nil {
		log.Fatal(err)
	}
	os.Setenv("DockerContainerID", string(dat))
}
