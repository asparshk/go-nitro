package lb

type Lbsipparameters struct {
  Addrportvip string `json:"addrportvip,omitempty"`
  Retrydur int `json:"retrydur,omitempty"`
  Rnatdstport int `json:"rnatdstport,omitempty"`
  Rnatsrcport int `json:"rnatsrcport,omitempty"`
  Sip503ratethreshold int `json:"sip503ratethreshold,omitempty"`
}