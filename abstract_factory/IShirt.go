package abstract_factory

type IShirt interface {
	SetLogo(logo string)
	GetLogo() string
	SetSize(size int)
	GetSize() int
}

type Shirt struct {
	logo string
	size int
}

func (s *Shirt) SetLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) SetSize(size int) {
	s.size = size
}

func (s *Shirt) GetSize() int {
	return s.size
}

func (s *Shirt) GetLogo() string {
	return s.logo
}
