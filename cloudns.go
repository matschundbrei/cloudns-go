package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-resty/resty/v2"
	"os"
	"strconv"
	"strings"
)

// this is the communicated api endpoint
const (
	apiurl = "https://api.cloudns.net"
)

// generic request handler, up until now, it seems GET is sufficient
func doreq(path string, params map[string]string) (*resty.Response, error) {
	fullurl := strings.Join([]string{apiurl, path}, "")
	client := resty.New()
	return client.R().SetQueryParams(params).SetHeader("Accept", "application/json").Get(fullurl)
}

// a very fugly thing to flatten json to a map
// this is mostly stolen from https://blog.golang.org/json-and-go
// and then Knuth and I banged our heads against it until it worked
func flattenjson(b []byte) map[string]string {
	var tmp interface{}
	err2 := json.Unmarshal(b, &tmp)
	if err2 != nil {
		spew.Printf("error in json.Unmarshal: %#+v \n", err2)
	}
	m := tmp.(map[string]interface{})
	var params map[string]string
	params = make(map[string]string)
	for k, v := range m {
		switch vt := v.(type) {
		case string:
			params[k] = vt
		case int:
			params[k] = strconv.Itoa(vt)
		case bool:
			params[k] = strconv.FormatBool(vt)
		case float64: // this case is exceptionally shitty, but it works here
			var flt float64 = vt
			params[k] = strconv.Itoa(int(flt))
		case []interface{}: // if the layer is an array
			for i, u := range vt {
				numstr := strconv.Itoa(i)
				newk := strings.Join([]string{k, "[", numstr, "]"}, "")
				switch ut := u.(type) {
				case string:
					params[newk] = ut
				case int:
					params[newk] = strconv.Itoa(ut)
				case bool:
					params[newk] = strconv.FormatBool(ut)
				case float64:
					var fltu float64 = ut
					params[newk] = strconv.Itoa(int(fltu))
				default:
					spew.Printf("error handling type for key %#+v, heres the value: %#+v\n", newk, ut)
				}
			}
		default:
			spew.Printf("error handling type for key %#+v, heres the value: %#+v\n", k, vt)
		}
	}
	return params
}

// helper functions to create params as flattened map
func (conf Apiaccess) mkparams() map[string]string {
	b, err := json.Marshal(conf)
	if err != nil {
		spew.Printf("error in first json.Marshal: %#+v \n", err)
	}
	return flattenjson(b)
}

func (record Recordset) mkparams(conf Apiaccess) map[string]string {
	params := conf.mkparams()
	b, err := json.Marshal(record)
	if err != nil {
		spew.Printf("error in first json.Marshal: %#+v \n", err)
	}
	tmpmap := flattenjson(b)
	for k, v := range tmpmap {
		params[k] = v
	}
	return params
}

func (zone Zone) mkparams(conf Apiaccess) map[string]string {
	params := conf.mkparams()
	b, err := json.Marshal(zone)
	if err != nil {
		spew.Printf("error in first json.Marshal: %#+v \n", err)
	}
	tmpmap := flattenjson(b)
	for k, v := range tmpmap {
		params[k] = v
	}
	return params
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

// check if the credentials work
func (conf Apiaccess) Logincheck() (*resty.Response, error) {
	const path = "/dns/login.json"
	params := conf.mkparams()
	return doreq(path, params)
}

// get the available values for TTL
func (conf Apiaccess) Availablettl() (*resty.Response, error) {
	const path = "/dns/get-available-ttl.json"
	params := conf.mkparams()
	return doreq(path, params)
}

func (conf Apiaccess) Availabletype() (*resty.Response, error) {
	const path = "/dns/get-available-record-types.json"
	params := conf.mkparams()
	params["zone-type"] = "domain"
	return doreq(path, params)
}

func (conf Apiaccess) lsrec(domain string) (*resty.Response, error) {
	const path = "/dns/records.json"
	params := conf.mkparams()
	params["domain-name"] = domain
	return doreq(path, params)
}

func (conf Apiaccess) lszone(searchstring string) (*resty.Response, error) {
	const path = "/dns/list-zones.json"
	params := conf.mkparams()
	if searchstring != "" {
		params["search"] = searchstring
	}
	//TODO:
	//this needs to recurse through pages
	//currently we just take a limit of 100 domains into account
	params["page"] = "1"
	params["rows-per-page"] = "100"
	return doreq(path, params)
}

/*

// CRUD functions for our structs in types.go
func (record Recordset) Read(auth *Apiaccess) (response resty.Response, err error) {
	// utilise list function ...
}

func (record Recordset) Create(auth *Apiaccess) (err error) {
	const path = "/dns/add-record.json"
}

func (record Recordset) Update(auth *Apiaccess) (err error) {
	const path = "/dns/mod-record.json"
}

func (record Recordset) Destroy(auth *Apiaccess) (err error) {
	const path = "/dns/delete-record.json"
}

func (zone Zone) Read(auth *Apiaccess) (response resty.Response, err error) {
	// utilise list function ...
}

func (zone Zone) Create(auth *Apiaccess) (err error) {
	const path = "/dns/register.json"
}

func (zone Zone) Update(auth *Apiaccess) (err error) {
	// not sure what this does ...
	// see https://www.cloudns.net/wiki/article/135/
	const path = "/dns/update-zone.json"
}

func (zone Zone) Destroy(auth *Apiaccess) (err error) {
	const path = "/dns/delete.json"
}
*/

func main() {
	fmt.Println("we start now")
	foo := fetchconfig()

	spew.Printf("this is foo now: %#+v \n", foo)
	//bar := Recordset{}
	//zap := Zone{}

	req, err := foo.Logincheck()
	if err == nil {
		spew.Printf("#1 Logincheck: API says: %#+v \n", req)
	}
	req2, err2 := foo.Availablettl()
	if err2 == nil {
		spew.Printf("#2 Available TTLS: API says: %#+v \n", req2)
	}
	req3, err3 := foo.Availabletype()
	if err3 == nil {
		spew.Printf("#3 Available Types: API says: %#+v \n", req3)
	}
	req4, err4 := foo.lsrec("sta.net")
	if err4 == nil {
		spew.Printf("#4 Listing Records for Domain 'sta.net': %#+v \n", req4)
	}
	req5, err5 := foo.lszone("")
	if err5 == nil {
		spew.Printf("#5 Listing Domains with empty searchstring: %#+v \n", req5)
	}
}
