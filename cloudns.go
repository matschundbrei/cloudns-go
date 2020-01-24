package main

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-resty/resty/v2"
	//"os"
	"strconv"
	"strings"
)

const (
	apiurl = "https://api.cloudns.net"
)

/*
requests by HTTP get to url
params need to be send as


GET is the thing to use, but unfortunately that means we'll need to
1. encode the struct (see types.go) in json (to apply filter/renaming
2. reencode it back to have only the relevant slice
3. add it to the querystring, see https://stackoverflow.com/questions/30652577/go-doing-a-get-request-and-building-the-querystring

*/

func doreq(path string, params map[string]string) (response *resty.Response, err error) {
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
		spew.Printf("current vals:\nk: %#+v\nv: %#+v\n", k, v)
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
					spew.Printf("error handling type for key %#+v, heres the var: %#+v", newk, ut)
				}
			}
		default:
			spew.Printf("error handling type for key %#+v, heres the var: %#+v", k, vt)
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

/*
// get the needed vars from ARG, CFG, ENV (prio l2r) https://www.cloudns.net/wiki/article/45/
func fetchconfig() (conf Apiaccess) {

}

func (conf Apiaccess) Logincheck() (response resty.Response, err error) {
	const path = "/dns/login.json"
}

func listRecs(conf Apiaccess, searchstring string) (response resty.Response, err error) {
	const path = "/dns/records.json"
}

func listZones(conf Apiaccess, searchstring string) (response resty.Response, err error) {
	const path = "/dns/list-zones.json"
}

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
	foo := Apiaccess{
		Authid:       4711,
		Authpassword: "example-password",
	}

	spew.Printf("this is foo now: %#+v \n", foo)
	//bar := Recordset{}
	//zap := Zone{}

	thisparms := foo.mkparams()
	fmt.Println(thisparms)
	spew.Printf("params? %#+v", thisparms)

}
