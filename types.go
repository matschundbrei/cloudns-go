package main

// Apiaccess ClouDNS API Credentials, see https://www.cloudns.net/wiki/article/42/
type Apiaccess struct {
	Authid       int    `json:"auth-id"`
	Authpassword string `json:"auth-password"`
}

// Apierr we should be able to json.Unmarshal this from an *resty.Response.Body(),
// if there was an API error, see https://www.cloudns.net/wiki/article/45/
type Apierr struct {
	Status string `json:"status"`
	Desc   string `json:"statusDescription"`
}

// Zone is the external representation of a zone
type Zone struct {
	Domain string   `json:"domain-name"`
	Ztype  string   `json:"zone-type"`
	Ns     []string `json:"ns,omitempty"`
}

// Record is the external representation of a record
type Record struct {
	ID     string `json:"id"`
	Domain string `json:"domain-name"`
	Host   string `json:"host"`
	Rtype  string `json:"record-type"`
	TTL    int    `json:"ttl"`
	Record string `json:"record"`
}

// Createrec create a record, see https://www.cloudns.net/wiki/article/58/
type Createrec struct {
	Authid         int    `json:"auth-id"`
	Authpassword   string `json:"auth-password"`
	Domain         string `json:"domain-name"`
	Rtype          string `json:"record-type"`
	TTL            int    `json:"ttl"`
	Host           string `json:"host"`
	Record         string `json:"record"`
	Priority       int    `json:"priority,omitempty"`
	Weight         int    `json:"weight,omitempty"`
	Port           int    `json:"port,omitempty"`
	Frame          int    `json:"frame,omitempty"`
	FrameTitle     string `json:"frame-title,omitempty"`
	FrameKeywords  string `json:"frame-keywords,omitempty"`
	FrameDesc      string `json:"frame-description,omitempty"`
	SavePath       int    `json:"save-path,omitempty"`
	RedirectType   int    `json:"redirect-type,omitempty"`
	Mail           string `json:"mail,omitempty"`
	Txt            string `json:"txt,omitempty"`
	Algorithm      string `json:"algorithm,omitempty"`
	Fptype         string `json:"fptype,omitempty"`
	Status         int    `json:"status,omitempty"`
	GeodnsLocation int    `json:"geodns-location,omitempty"`
	CaaFlag        int    `json:"caa_flag,omitempty"`
	CaaType        string `json:"caa_type,omitempty"`
	CaaValue       string `json:"caa_value,omitempty"`
	Tusage         string `json:"tlsa_usage,omitempty"`
	Tselector      string `json:"tlsa_selector,omitempty"`
	Tmatchtype     string `json:"tlsa_matching_type,omitempty"`
}

// Updaterec is the alternative record struct, used here https://www.cloudns.net/wiki/article/60/
type Updaterec struct {
	Authid         int    `json:"auth-id"`
	Authpassword   string `json:"auth-password"`
	Domain         string `json:"domain-name"`
	Rid            int    `json:"record-id"`
	TTL            int    `json:"ttl"`
	Host           string `json:"host"`
	Record         string `json:"record"`
	Priority       int    `json:"priority,omitempty"`
	Weight         int    `json:"weight,omitempty"`
	Port           int    `json:"port,omitempty"`
	Frame          int    `json:"frame,omitempty"`
	FrameTitle     string `json:"frame-title,omitempty"`
	FrameKeywords  string `json:"frame-keywords,omitempty"`
	FrameDesc      string `json:"frame-description,omitempty"`
	SavePath       int    `json:"save-path,omitempty"`
	RedirectType   int    `json:"redirect-type,omitempty"`
	Mail           string `json:"mail,omitempty"`
	Txt            string `json:"txt,omitempty"`
	Algorithm      string `json:"algorithm,omitempty"`
	Fptype         string `json:"fptype,omitempty"`
	Status         int    `json:"status,omitempty"`
	GeodnsLocation int    `json:"geodns-location,omitempty"`
	CaaFlag        int    `json:"caa_flag,omitempty"`
	CaaType        string `json:"caa_type,omitempty"`
	CaaValue       string `json:"caa_value,omitempty"`
	Tusage         string `json:"tlsa_usage,omitempty"`
	Tselector      string `json:"tlsa_selector,omitempty"`
	Tmatchtype     string `json:"tlsa_matching_type,omitempty"`
}

// Returnrec to Unmarshal returned records, see https://www.cloudns.net/wiki/article/57/
type Returnrec struct {
	ID       string `json:"id"`
	Rtype    string `json:"type"`
	Host     string `json:"host"`
	Record   string `json:"record"`
	Dynurl   int    `json:"dynamicurl_status"`
	Failover string `json:"failover"`
	TTL      string `json:"ttl"`
	Dtatus   int    `json:"status"`
}

// Createzone create a zone, see https://www.cloudns.net/wiki/article/48/
type Createzone struct {
	Authid       int      `json:"auth-id"`
	Authpassword string   `json:"auth-password"`
	Domain       string   `json:"domain-name"`
	Ztype        string   `json:"zone-type"`
	Ns           []string `json:"ns,omitempty"`
	Master       string   `json:"master-ip,omitempty"`
}

//Returnzone to Unmarchal returned zones, see https://www.cloudns.net/wiki/article/50/
type Returnzone struct {
	Domain string `json:"domain-name"`
	Ztype  string `json:"zone-type"`
	Master string `json:"master-ip,omitempty"` // ??? not sure if this was in, check in lsdom
}

//types lists available types from api, see https://www.cloudns.net/wiki/article/157/
type rectypes struct {
	Authid       int    `json:"auth-id"`
	Authpassword string `json:"auth-password"`
	Ztype        string `json:"zone-type"`
}

//reclist lists records for a domain, see https://www.cloudns.net/wiki/article/57/
type reclist struct {
	Authid       int    `json:"auth-id"`
	Authpassword string `json:"auth-password"`
	Domain       string `json:"domain-name"`
	Host         string `json:"host,omitempty"`
	Rtype        string `json:"type,omitempty"`
}

//zonelist lists zones, see https://www.cloudns.net/wiki/article/50/
type zonelist struct {
	Authid       int    `json:"auth-id"`
	Authpassword string `json:"auth-password"`
	Page         int    `json:"page"`
	Hits         int    `json:"rows-per-page"`
	Search       string `json:"search,omitempty"`
	Gid          int    `json:"group-id,omitempty"`
}

//zone update/destroy struct, see https://www.cloudns.net/wiki/article/135/ or https://www.cloudns.net/wiki/article/49/
type zupdate struct {
	Authid       int    `json:"auth-id"`
	Authpassword string `json:"auth-password"`
	Domain       string `json:"domain-name"`
}
