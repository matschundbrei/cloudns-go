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


/*
requests by HTTP get to url
params need to be send as 


GET is the thing to use, but unfortunately that means we'll need to 
1. encode the struct (see types.go) in json (to apply filter/renaming
2. reencode it back to have only the relevant slice
3. add it to the querystring, see https://stackoverflow.com/questions/30652577/go-doing-a-get-request-and-building-the-querystring

*/

func doRequest(url string, httptype string, params []string) (response, err) {
	return response, err := http.Request(
}


// get the needed vars from ARG, CFG, ENV (prio l2r) https://www.cloudns.net/wiki/article/45/
func (conf *Apiaccess) Logincheck() (response, err error) {}

func listRecs(conf *Apiaccess, searchstring string) (response, err error) {}

func listZones(conf *Apiaccess, searchstring string) (response, err error) {}

// CRUD functions for our structs in types.go
func (record **Recordset) Read(auth *Apiaccess) (response map{}, err error) {}

func (record **Recordset) Create(auth *Apiaccess) (err error) {}

func (record **Recordset) Update(auth *Apiaccess) (err error) {}

func (record **Recordset) Destroy(auth *Apiaccess) (err error) {}

func (zone **Zone) Read(auth *Apiaccess) (response something?, err error) {}

func (zone **Zone) Create(auth *Apiaccess) (err error) {}

func (zone **Zone) Update(auth *Apiaccess) (err error) {}

func (zone **Zone) Destroy(auth *Apiaccess) (err error) {}

