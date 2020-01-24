package main

// things that you need to access the api
type Apiaccess struct {
	Authid       int    `json:"auth-id"`
	Authpassword string `json:"auth-password"`
}

// things in a records set by api
type Recordset struct {
	Rtype          string `json:"record-type"`
	Ttl            int
	Host           string
	Record         string
	Priority       int    `json:",omitempty"`
	Weight         int    `json:",omitempty"`
	Port           int    `json:",omitempty"`
	Frame          int    `json:",omitempty"`
	FrameTitle     string `json:"frame-title,omitempty"`
	FrameKeywords  string `json:"frame-keywords,omitempty"`
	FrameDesc      string `json:"frame-description,omitempty"`
	SavePath       int    `json:"save-path,omitempty"`
	RedirectType   int    `json:"redirect-type,omitempty"`
	Mail           string `json:",omitempty"`
	Txt            string `json:",omitempty"`
	Algorithm      string `json:",omitempty"`
	Fptype         string `json:",omitempty"`
	Status         int    `json:",omitempty"`
	GeodnsLocation int    `json:"geodns-location,omitempty"`
	CaaFlag        int    `json:"caa_flag,omitempty"`
	CaaType        string `json:"caa_type,omitempty"`
	CaaValue       string `json:"caa_value,omitempty"`
}

// things in a zone set by api
type Zone struct {
	Domain string   `json:"domain-name"`
	Ztype  string   `json:"zone-type"`
	Ns     []string `json:",omitempty"`
	Master string   `json:"master-ip,omitempty"`
}
