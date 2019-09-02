package ci

// Travis struct
type Travis struct{}

// Get returns necessary information from Circle
func (c Travis) Get() Info {

	//TODO: retrieve info
	return Info{
		Sha1:   "",
		Branch: "",
		CI:     "travis",
	}
}

// NewTravis is a constructor for Circle
func NewTravis() Travis {
	return Travis{}
}
