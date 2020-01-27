package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-resty/resty/v2"
)

// this is the communicated api endpoint
const (
	apiurl = "https://api.cloudns.net"
)

// this is how we will do requests
func apireq(path string, body interface{}) (*resty.Response, error) {
	fullurl := strings.Join([]string{apiurl, path}, "")
	client := resty.New()
	client.R().SetHeader("Content-Type", "application/json")
	client.R().SetHeader("Accept", "application/json")
	client.R().SetHeader("User-Agent", "github.com/sta-travel/cloudns-go")
	return client.R().SetBody(body).Post(fullurl)
}

// get the needed vars from ENV, ARG, CFG (prio l2r) https://www.cloudns.net/wiki/article/45/
func fetchconfig() Apiaccess {
	conf := Apiaccess{}
	// ENV
	// for convinience, we're using the same ones as go-acme/lego
	envid, err := strconv.ParseInt(os.Getenv("CLOUDNS_AUTH_ID"), 0, 32)
	if err != nil {
		spew.Printf("error converting %#+v to an int64\n", os.Getenv("CLOUDNS_AUTH_ID"))
	} else {
		conf.Authid = int(envid)
	}
	conf.Authpassword = os.Getenv("CLOUDNS_AUTH_PASSWORD")
	// look for args
	var argid int
	var argpw string
	flag.IntVar(&argid, "authid", 0, "Your ClouDNS HTTP API ID")
	flag.StringVar(&argpw, "authpw", "", "Your ClouDNS HTTP API Password")
	flag.Parse()
	if argid > 0 {
		conf.Authid = argid
	}
	if argpw != "" {
		conf.Authpassword = argpw
	}
	// todo config from file
	return conf
}

// Logincheck checks if the credentials work
func (c Apiaccess) Logincheck() (*resty.Response, error) {
	const path = "/dns/login.json"
	return apireq(path, c)
}

// Availablettl gets the currently available TTL values
func (c Apiaccess) Availablettl() (*resty.Response, error) {
	const path = "/dns/get-available-ttl.json"
	return apireq(path, c)
}

// Availabletype gets the currently available Record-Types
func (r rectypes) Availabletype() (*resty.Response, error) {
	const path = "/dns/get-available-record-types.json"
	return apireq(path, r)
}

// list records
func (r reclist) lsrec() (*resty.Response, error) {
	const path = "/dns/records.json"
	return apireq(path, r)
}

// list zones
func (z zonelist) lszone() (*resty.Response, error) {
	const path = "/dns/list-zones.json"
	return apireq(path, z)
}

/*
// CRUD functions for our structs in types.go
func (r Createrec) Read() (map[string]Returnrec, error) {
	// utilise list function ...

	resp, err := auth.lsrec(r.Domain, r.Host, record.Rtype)
	var unmres map[string]Returnrec
	var unmerr error
	if err == nil {
		unmerr = json.Unmarshal(resp.Body(), &unmres)
		if unmerr != nil {
			spew.Printf("wow, there's an Unmarshal error: %#+v \n %#+v \n", unmerr, string(resp.Body()))
			return unmres, unmerr
		}
	} else {
		spew.Printf("wow, there's a request error! %#+v \n%#+v \n", err, resp)
		return unmres, err
	}
	return unmres, unmerr
}

func (record Createrec) Create(auth *Apiaccess) (*resty.Response, error) {
	const path = "/dns/add-record.json"
}

func (record Createrec) Update(auth *Apiaccess) (*resty.Response, error) {
	const path = "/dns/mod-record.json"
}

func (record Createrec) Destroy(auth *Apiaccess) (*resty.Response, error) {
	const path = "/dns/delete-record.json"
}

func (zone Createzone) Read(auth *Apiaccess) (*resty.Response, error) {
	// utilise list function ...
}

func (zone Createzone) Create(auth *Apiaccess) (*resty.Response, error) {
	const path = "/dns/register.json"
}

func (zone Createzone) Update(auth *Apiaccess) (*resty.Response, error) {
	// not sure what this does ...
	// see https://www.cloudns.net/wiki/article/135/
	const path = "/dns/update-zone.json"
}

func (zone Createzone) Destroy(auth *Apiaccess) (*resty.Response, error) {
	const path = "/dns/delete.json"
}
*/

func main() {
	fmt.Println("we start now")
	foo := fetchconfig()

	spew.Printf("this is foo now: %#+v \n", foo)
	// bar := Createrec{
	// 	Domain: "sta.net",
	// 	Rtype:  "A",
	// 	Host:   "cr",
	// }
	//zap := Createzone{}
	//thingy, _ := bar.Read(foo)

	thingy := zonelist{
		Authid:       foo.Authid,
		Authpassword: foo.Authpassword,
		Page:         1,
		Hits:         100,
	}

	resp, _ := thingy.lszone()

	spew.Printf("here's what we got: %#+v \n", resp)

}
