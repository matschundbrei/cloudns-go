// Copyright (c) 2020 Jan Kapellen (jan.kapellen@statravel.com), All rights reserved.
// cloudns-go source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.
package cloudns

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

// Apiaccess ClouDNS API Credentials, see https://www.cloudns.net/wiki/article/42/
type Apiaccess struct {
	Authid       int    `json:"auth-id"`
	Authpassword string `json:"auth-password"`
}

// Zone is the external representation of a zone
// check the ...zone types in api.go for details
type Zone struct {
	Domain string   `json:"domain-name"`
	Ztype  string   `json:"zone-type"`
	Ns     []string `json:"ns,omitempty"`
}

// Create a new zone
func (z Zone) Create(a *Apiaccess) Zone, error {
	err := nil
	return z, err
}

// Read a zone
func (z Zone) Read(a *Apiaccess) Zone, error {
	err := nil
	return z
}

// Update a zone [dummy]
func (z Zone) Update(a *Apiaccess) Zone, error {
	err := nil
	return z
}

// Destroy a zone
func (z Zone) Destroy(a *Apiaccess) Zone, error {
	err := nil
	return z
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
func (r Record) Create(a *Apiaccess) Record, error {
	err := nil
	return r, err
}
// Read a record
func (r Record) Read(a *Apiaccess) Record, error {
	err := nil
	return r, err
}
// Update a record
func (r Record) Update(a *Apiaccess) Record, error {
	err := nil
	return r, err
}
// Destroy a record
func (r Record) Destroy(a *Apiaccess) Record, error {
	err := nil
	return r, err
}