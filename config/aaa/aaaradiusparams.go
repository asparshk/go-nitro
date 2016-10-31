package aaa

type Aaaradiusparams struct {
  Accounting string `json:"accounting,omitempty"`
  Authtimeout int `json:"authtimeout,omitempty"`
  Callingstationid string `json:"callingstationid,omitempty"`
  Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
  Groupauthname string `json:"groupauthname,omitempty"`
  Ipaddress string `json:"ipaddress,omitempty"`
  Ipattributetype int `json:"ipattributetype,omitempty"`
  Ipvendorid int `json:"ipvendorid,omitempty"`
  Passencoding string `json:"passencoding,omitempty"`
  Pwdattributetype int `json:"pwdattributetype,omitempty"`
  Pwdvendorid int `json:"pwdvendorid,omitempty"`
  Radattributetype int `json:"radattributetype,omitempty"`
  Radgroupseparator string `json:"radgroupseparator,omitempty"`
  Radgroupsprefix string `json:"radgroupsprefix,omitempty"`
  Radkey string `json:"radkey,omitempty"`
  Radnasid string `json:"radnasid,omitempty"`
  Radnasip string `json:"radnasip,omitempty"`
  Radvendorid int `json:"radvendorid,omitempty"`
  Serverip string `json:"serverip,omitempty"`
  Serverport int `json:"serverport,omitempty"`
}