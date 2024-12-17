package abstract_factory

type Nike struct{}

type NickShoe struct {
	Shoe
}

type NickShirt struct {
	Shirt
}

func (n *Nike) MakeShoe() IShoe {
	return &NickShoe{
		Shoe: Shoe{
			logo: "Nike",
			size: 42,
		},
	}
}

func (n *Nike) MakeShirt() IShirt {
	return &NickShirt{
		Shirt: Shirt{
			logo: "Nike",
			size: 14,
		},
	}
}
