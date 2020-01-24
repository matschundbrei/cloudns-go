package main

// things that you need to access the api
type Apiaccess struct {
	url          *string `json:"-"`
	authid       *int    `json:"auth-id"`
	authpassword *string `json:"auth-password"`
	apitimeout   *int    `json:"-"`
	apiinterval  *int    `json:"-"`
	defaultttl   *int    `json:"-"`
}

// things in a records set by api
type Recordset struct {
	rtype          *string `json:"record-type"`
	ttl            int
	host           string
	record         string
	priority       *int    `json:",omitempty"`
	weight         *int    `json:",omitempty"`
	port           *int    `json:",omitempty"`
	frame          *int    `json:",omitempty"`
	frameTitle     *string `json:"frame-title,omitempty"`
	frameKeywords  *string `json:"frame-keywords,omitempty"`
	frameDesc      *string `json:"frame-description,omitempty"`
	savePath       *int    `json:"save-path,omitempty"`
	redirectType   *int    `json:"redirect-type,omitempty"`
	mail           *string `json:",omitempty"`
	txt            *string `json:",omitempty"`
	algorithm      *string `json:",omitempty"`
	fptype         *string `json:",omitempty"`
	status         *int    `json:",omitempty"`
	geodnsLocation *int    `json:"geodns-location,omitempty"`
	caaFlag        *int    `json:"caa_flag,omitempty"`
	caaType        *string `json:"caa_type,omitempty"`
	caaValue       *string `json:"caa_value,omitempty"`
}

// things in a zone set by api
type Zone struct {
	domain *string  `json:"domain-name"`
	ztype  *string  `json:"zone-type"`
	ns     []string `json:,omitempty`
	master *string  `json:"master-ip,omitempty"`
}
