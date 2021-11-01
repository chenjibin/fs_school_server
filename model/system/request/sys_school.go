package request

type SchoolAdd struct {
	Name string `json:"name"`
	Address string `json:"address"`
	Principal string `json:"principal"`
	PrincipalMobile string `json:"principalMobile"`
	Type string `json:"type"`
	Logo string `json:"logo"`
}

type SchoolUpdate struct {
	ID uint `json:"ID"`
	Name string `json:"name"`
	Address string `json:"address"`
	Principal string `json:"principal"`
	PrincipalMobile string `json:"principalMobile"`
	Type string `json:"type"`
	Logo string `json:"logo"`
}
