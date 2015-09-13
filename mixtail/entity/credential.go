package entity


type Credential struct {
	User string
	Pass string
	Identity string
}


func NewCredential() (c *Credential) {
	c = new(Credential)
	return
}


func (c *Credential) IsPasswordAuth() bool {
	return (c.Identity == "")
}

