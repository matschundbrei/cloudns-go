// Package cloudns public structs/functions
package cloudns

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

// List returns all records from a zone
func (z Zone) List(a *Apiaccess) ([]Record, error) {
	var err error = nil
	r := Record{}
	ra := []Record{r}
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
