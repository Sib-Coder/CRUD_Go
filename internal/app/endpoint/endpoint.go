package endpoint

type Endpoint struct {
	s Service
}

func New() *Endpoint {
	return &Endpoint{
		s: s,
	}
}
