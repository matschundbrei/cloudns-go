// Package cloudns private api functions
package cloudns

import (
	"encoding/json"
	"strings"

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

// Apierr we should be able to json.Unmarshal this from an *resty.Response.Body(),
// if there was an API error, see https://www.cloudns.net/wiki/article/45/
type apierr struct {
	status string `json:"status"`
	desc   string `json:"statusDescription"`
}

// this function will check a byte array for the error message from ClouDNS
// it's a little backwards, but it works pretty slick
func checkapierr(d []byte) (string, bool) {
	var status apierr
	err := json.Unmarshal(d, &status)
	if err == nil && status.status != "Success" {
		return status.desc, true
	}
	return "", false
}

// logincheck checks if the credentials work, see https://www.cloudns.net/wiki/article/45/
func (c Apiaccess) logincheck() (*resty.Response, error) {
	const path = "/dns/login.json"
	return apireq(path, c)
}

// availablettl gets the currently available TTL values, see https://www.cloudns.net/wiki/article/153/
func (c Apiaccess) availablettl() (*resty.Response, error) {
	const path = "/dns/get-available-ttl.json"
	return apireq(path, c)
}

//types lists available types from api, see https://www.cloudns.net/wiki/article/157/
type rectypes struct {
	authid       int    `json:"auth-id"`
	authpassword string `json:"auth-password"`
	ztype        string `json:"zone-type"`
}

// Availabletype gets the currently available Record-Types
func (r rectypes) availabletype() (*resty.Response, error) {
	const path = "/dns/get-available-record-types.json"
	return apireq(path, r)
}

//reclist lists records for a domain, see https://www.cloudns.net/wiki/article/57/
type reclist struct {
	authid       int    `json:"auth-id"`
	authpassword string `json:"auth-password"`
	domain       string `json:"domain-name"`
	host         string `json:"host,omitempty"`
	rtype        string `json:"type,omitempty"`
}

// list records
func (r reclist) lsrec() (*resty.Response, error) {
	const path = "/dns/records.json"
	return apireq(path, r)
}

//zonelist struct to lists zones, see https://www.cloudns.net/wiki/article/50/
type zonelist struct {
	authid       int    `json:"auth-id"`
	authpassword string `json:"auth-password"`
	page         int    `json:"page"`
	hits         int    `json:"rows-per-page"`
	search       string `json:"search,omitempty"`
	gid          int    `json:"group-id,omitempty"`
}

// list zones
func (z zonelist) lszone() (*resty.Response, error) {
	const path = "/dns/list-zones.json"
	return apireq(path, z)
}

// createrec stuct to create a record, see https://www.cloudns.net/wiki/article/58/
type createrec struct {
	authid         int    `json:"auth-id"`
	authpassword   string `json:"auth-password"`
	domain         string `json:"domain-name"`
	rtype          string `json:"record-type"`
	ttl            int    `json:"ttl"`
	host           string `json:"host"`
	record         string `json:"record"`
	priority       int    `json:"priority,omitempty"`
	weight         int    `json:"weight,omitempty"`
	port           int    `json:"port,omitempty"`
	frame          int    `json:"frame,omitempty"`
	frameTitle     string `json:"frame-title,omitempty"`
	frameKeywords  string `json:"frame-keywords,omitempty"`
	frameDesc      string `json:"frame-description,omitempty"`
	savePath       int    `json:"save-path,omitempty"`
	redirectType   int    `json:"redirect-type,omitempty"`
	mail           string `json:"mail,omitempty"`
	txt            string `json:"txt,omitempty"`
	algorithm      string `json:"algorithm,omitempty"`
	fptype         string `json:"fptype,omitempty"`
	status         int    `json:"status,omitempty"`
	geodnsLocation int    `json:"geodns-location,omitempty"`
	caaFlag        int    `json:"caa_flag,omitempty"`
	caaType        string `json:"caa_type,omitempty"`
	caaValue       string `json:"caa_value,omitempty"`
	tusage         string `json:"tlsa_usage,omitempty"`
	tselector      string `json:"tlsa_selector,omitempty"`
	tmatchtype     string `json:"tlsa_matching_type,omitempty"`
}

// Read returns the created records (map[string]Returnrec in response)
func (r createrec) read() (*resty.Response, error) {
	// this should give us a list containing this exact record
	listrec := reclist{
		authid:       r.authid,
		authpassword: r.authpassword,
		rtype:        r.rtype,
		host:         r.host,
	}
	return listrec.lsrec()
}

// create actually creates a record
func (r createrec) create() (*resty.Response, error) {
	const path = "/dns/add-record.json"
	return apireq(path, r)
}

// updaterec is the alternative record struct, used here https://www.cloudns.net/wiki/article/60/
type updaterec struct {
	authid         int    `json:"auth-id"`
	authpassword   string `json:"auth-password"`
	domain         string `json:"domain-name"`
	rid            int    `json:"record-id"`
	ttl            int    `json:"ttl"`
	host           string `json:"host"`
	record         string `json:"record"`
	priority       int    `json:"priority,omitempty"`
	weight         int    `json:"weight,omitempty"`
	port           int    `json:"port,omitempty"`
	frame          int    `json:"frame,omitempty"`
	frameTitle     string `json:"frame-title,omitempty"`
	frameKeywords  string `json:"frame-keywords,omitempty"`
	frameDesc      string `json:"frame-description,omitempty"`
	savePath       int    `json:"save-path,omitempty"`
	redirectType   int    `json:"redirect-type,omitempty"`
	mail           string `json:"mail,omitempty"`
	txt            string `json:"txt,omitempty"`
	algorithm      string `json:"algorithm,omitempty"`
	fptype         string `json:"fptype,omitempty"`
	status         int    `json:"status,omitempty"`
	geodnsLocation int    `json:"geodns-location,omitempty"`
	caaFlag        int    `json:"caa_flag,omitempty"`
	caaType        string `json:"caa_type,omitempty"`
	caaValue       string `json:"caa_value,omitempty"`
	tusage         string `json:"tlsa_usage,omitempty"`
	tselector      string `json:"tlsa_selector,omitempty"`
	tmatchtype     string `json:"tlsa_matching_type,omitempty"`
}

// Update updates an existing record
func (r updaterec) update() (*resty.Response, error) {
	const path = "/dns/mod-record.json"
	return apireq(path, r)
}

// Destroy destroys the record
func (r updaterec) destroy() (*resty.Response, error) {
	const path = "/dns/delete-record.json"
	return apireq(path, r)
}

// createzone creates a zone, see https://www.cloudns.net/wiki/article/48/
type createzone struct {
	authid       int      `json:"auth-id"`
	authpassword string   `json:"auth-password"`
	domain       string   `json:"domain-name"`
	ztype        string   `json:"zone-type"`
	ns           []string `json:"ns,omitempty"`
	master       string   `json:"master-ip,omitempty"`
}

// read should return the exact zone from the list
func (z createzone) read() (*resty.Response, error) {
	listzone := zonelist{
		authid:       z.authid,
		authpassword: z.authpassword,
		page:         1,
		hits:         10,
		search:       z.domain,
	}
	return listzone.lszone()
}

// create registers a new DNS zone
func (z createzone) create() (*resty.Response, error) {
	const path = "/dns/register.json"
	return apireq(path, z)
}

//zone update/destroy struct, see https://www.cloudns.net/wiki/article/135/ or https://www.cloudns.net/wiki/article/49/
type zupdate struct {
	authid       int    `json:"auth-id"`
	authpassword string `json:"auth-password"`
	domain       string `json:"domain-name"`
}

// Update in this context does not make much sense, but we implement it anyway
func (z createzone) update() (*resty.Response, error) {
	// not sure what this does ...
	// see https://www.cloudns.net/wiki/article/135/
	const path = "/dns/update-zone.json"
	up := zupdate{
		authid:       z.authid,
		authpassword: z.authpassword,
		domain:       z.domain,
	}
	return apireq(path, up)
}

// Destroy removes a zone
func (z createzone) Destroy(auth *Apiaccess) (*resty.Response, error) {
	const path = "/dns/delete.json"
	rm := zupdate{
		authid:       z.authid,
		authpassword: z.authpassword,
		domain:       z.domain,
	}
	return apireq(path, rm)
}
