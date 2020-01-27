package main

/*
Good morning, Jan, this is Jan from yesterday, I have a few tasks for you:

- make all current functions/structs private
- make public CRUD functions for those higher lvl structs (Record, Zone), that use the private ones
- declare api error type for {"status":"Failed","statusDescription":".... (implemented already: checkapierr)


the public crud should *always* return the same struct (_self?) and errors if any

create: this is only complex for records, since ClouDNS allows for duplicate records of the same type in the same host
so, we need to grapple the ID from the create request and add it to the struct returning

read: needs to be done by listing (either recs or zones)

update: can only be done for records (domains will err)

destroy: functions already present

we need also:
- an init function for the module read this: https://blog.golang.org/using-go-modules
- a proper cli for the binary this looks promising: https://github.com/urfave/cli/blob/master/docs/v2/manual.md

*/

import (
	"encoding/json"
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

// this function will check a byte array for the error message from ClouDNS
// it's a little backwards, but it works pretty slick
func checkapierr(d []byte) (string, bool) {
	var status Apierr
	err := json.Unmarshal(d, &status)
	if err == nil && status.Status != "Success" {
		return status.Desc, true
	}
	return "", false
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

// CRUD functions for our structs in types.go

// Createrec.Read() returns the created records (map[string]Returnrec in response)
func (r Createrec) Read() (*resty.Response, error) {
	// this should give us a list containing this exact record
	listrec := reclist{
		Authid:       r.Authid,
		Authpassword: r.Authpassword,
		Rtype:        r.Rtype,
		Host:         r.Host,
	}
	return listrec.lsrec()
}

// Create actually creates a record
func (r Createrec) Create() (*resty.Response, error) {
	const path = "/dns/add-record.json"
	return apireq(path, r)
}

// Update updates an existing record
func (r Updaterec) Update() (*resty.Response, error) {
	const path = "/dns/mod-record.json"
	return apireq(path, r)
}

// Destroy destroys the record
func (r Updaterec) Destroy() (*resty.Response, error) {
	const path = "/dns/delete-record.json"
	return apireq(path, r)
}

// Read should return the exact zone from the list
func (z Createzone) Read() (*resty.Response, error) {
	listzone := zonelist{
		Authid:       z.Authid,
		Authpassword: z.Authpassword,
		Page:         1,
		Hits:         10,
		Search:       z.Domain,
	}
	return listzone.lszone()
}

// Create registers a new DNS zone
func (z Createzone) Create() (*resty.Response, error) {
	const path = "/dns/register.json"
	return apireq(path, z)
}

// Update in this context does not make much sense, but we implement it anyway
func (z Createzone) Update() (*resty.Response, error) {
	// not sure what this does ...
	// see https://www.cloudns.net/wiki/article/135/
	const path = "/dns/update-zone.json"
	up := zupdate{
		Authid:       z.Authid,
		Authpassword: z.Authpassword,
		Domain:       z.Domain,
	}
	return apireq(path, up)
}

// Destroy removes a zone
func (z Createzone) Destroy(auth *Apiaccess) (*resty.Response, error) {
	const path = "/dns/delete.json"
	rm := zupdate{
		Authid:       z.Authid,
		Authpassword: z.Authpassword,
		Domain:       z.Domain,
	}
	return apireq(path, rm)
}

func main() {
	fmt.Println("we start now")
	auth := fetchconfig()

	spew.Printf("this is auth now: %#+v \n", auth)
	testzone := Createzone{
		Authid:       auth.Authid,
		Authpassword: auth.Authpassword,
		Ztype:        "master",
		Domain:       "testdomain.xxx",
	}

	zcres, zcerr := testzone.Create()
	spew.Printf("Raw output of zone creation: %#+v", zcres)
	if zcerr != nil {
		fmt.Println("HTTP error creating zone: ", zcerr.Error())
	}
	zcar, zcab := checkapierr(zcres.Body())
	if zcab {
		fmt.Println("error creating zone: ", zcar)
	}

	testrecord := Createrec{
		Authid:       auth.Authid,
		Authpassword: auth.Authpassword,
		Domain:       testzone.Domain,
		TTL:          3600,
		Host:         "foo",
		Rtype:        "AAAA",
		Record:       "::1",
	}

	rcres, rcerr := testrecord.Create()
	spew.Printf("Raw output of record creation: %#+v", rcres)
	if rcerr != nil {
		fmt.Println("HTTP error creating record: ", rcerr.Error())
	}
	rcar, rcab := checkapierr(rcres.Body())
	if rcab {
		fmt.Println("error creating record: ", rcar)
	}
}
