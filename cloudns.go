// Package cloudns public structs/functions
package cloudns

import "encoding/json"

import "github.com/davecgh/go-spew/spew"

import "strconv"

// Apiaccess ClouDNS API Credentials, see https://www.cloudns.net/wiki/article/42/
type Apiaccess struct {
	Authid       int    `json:"auth-id,omitempty"`
	Subauthid    int    `json:"sub-auth-id,omitempty"`
	Authpassword string `json:"auth-password"`
}

// Zone is the external representation of a zone
// check the ...zone types in api.go for details
type Zone struct {
	Domain string   `json:"domain-name"`
	Ztype  string   `json:"zone-type"`
	Ns     []string `json:"ns,omitempty"`
}

// Listzones returns all zones (max: 100)
func (a Apiaccess) Listzones() ([]Zone, error) {
	zls := zonelist{
		Authid:       a.Authid,
		Subauthid:    a.Subauthid,
		Authpassword: a.Authpassword,
		Page:         1,
		Hits:         100,
	}
	resp, err := zls.lszone()
	var rz []Zone
	if err == nil {
		var intrz []retzone
		err2 := json.Unmarshal(resp.Body(), &intrz)
		for _, zn := range intrz {
			tmpzn := Zone{
				Domain: zn.Domain,
				Ztype:  zn.Ztype,
			}
			rz = append(rz, tmpzn)
		}
		return rz, err2
	}
	return rz, err
}

// List returns all records from a zone
func (z Zone) List(a *Apiaccess) ([]Record, error) {
	var ra []Record
	rls := reclist{
		Authid:       a.Authid,
		Subauthid:    a.Subauthid,
		Authpassword: a.Authpassword,
		Domain:       z.Domain,
	}
	resp, err := rls.lsrec()
	if err == nil {
		var ratmp map[string]retrec
		spew.Println(resp)
		err2 := json.Unmarshal(resp.Body(), &ratmp)
		for _, rec := range ratmp {
			tmpttl, _ := strconv.Atoi(rec.TTL)
			rectmp := Record{
				Domain: z.Domain,
				ID:     rec.ID,
				Rtype:  rec.Rtype,
				Host:   rec.Host,
				TTL:    tmpttl,
				Record: rec.Record,
			}
			ra = append(ra, rectmp)
		}
		return ra, err2
	}
	return ra, err
}

// Create a new zone
func (z Zone) Create(a *Apiaccess) (Zone, error) {
	var err error = nil

	return z, err
}

// Read a zone
func (z Zone) Read(a *Apiaccess) (Zone, error) {
	var err error = nil
	return z, err
}

// Update a zone [dummy]
func (z Zone) Update(a *Apiaccess) (Zone, error) {
	var err error = nil
	return z, err
}

// Destroy a zone
func (z Zone) Destroy(a *Apiaccess) (Zone, error) {
	var err error = nil
	return z, err
}

// Record is the external representation of a record
// check the ...record types in api.go for details
type Record struct {
	ID     string `json:"id"`
	Domain string `json:"domain-name"`
	Host   string `json:"host"`
	Rtype  string `json:"record-type"`
	TTL    int    `json:"ttl"`
	Record string `json:"record"`
}

// Create a new record
func (r Record) Create(a *Apiaccess) (Record, error) {
	var err error = nil
	return r, err
}

// Read a record
func (r Record) Read(a *Apiaccess) (Record, error) {
	var err error = nil
	return r, err
}

// Update a record
func (r Record) Update(a *Apiaccess) (Record, error) {
	var err error = nil
	return r, err
}

// Destroy a record
func (r Record) Destroy(a *Apiaccess) (Record, error) {
	var err error = nil
	return r, err
}
