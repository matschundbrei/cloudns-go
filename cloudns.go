package main

import (
  "fmt"
  "os"
  "strings"
  "net/http"
  "encoding/json"
)

const (
  apiurl = "https://api.cloudns.net"
)

var (
  authid int
  authkey string
  apitimeout int = 60
  apiinterval int = 10
  defaultttl int = 3600
)


func main() {
  configmap = getConfig()
}

/* fetch settings from either
    - configuration file
    - command line variable
    - environment variable 
*/
func getConfig() (*Apiaccess) {
  return
}

func doRequest(url string, rtype string, params map[string]string, auth *Apiaccess) (response, err) {
  return
}

func (c *Recordset) Read() {}

func (c *Recordset) Create() {}

func (c *Recordset) Update() {}

func (c *Recordset) Destroy() {}

func (c *Zone) Read() {}

func (c *Zone) Create() {}

func (c *Zone) Update() {}

func (c *Zone) Destroy() {}


func checkrecord(rtype string, rvalue string) (valid bool) {}

func createZone(client{} *ApiClient, domain string) (response{} map[string]string) {}
func updateZone(client{} *ApiClient, domain string) (response{} map[string]string) {}
func readZone(client{} *ApiClient, domain string) (response{} map[string]string, zone *ZoneRecord) {}
func destroyZone(client{} *ApiClient, domain string) (response{} map[string]string) {}
func createRecord(client{} *ApiClient, domain string, rname string, rtype string, rvalue string) (response{} map[string]string) {}
func updateRecord(client{} *ApiClient, domain string, rname string, rtype string, rvalue string) (response{} map[string]string) {}
func readRecord(client{} *ApiClient, domain string, rname string, rtype string, rvalue string) (response{} map[string]string, record *Record) {}
func destroyRecord(client{} *ApiClient, domain string, rname string, rtype string, rvalue string) (response{} map[string]string) {}
