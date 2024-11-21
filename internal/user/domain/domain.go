package domain

type Domain struct {
}

func NewDomain() *Domain {
	return &Domain{}
}

func (d *Domain) Build() *Domain {
	// do something
	return d
}
