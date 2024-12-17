package abstract_factory

type Adidas struct{}

type AdidasShoe struct {
	Shoe
}

type AdidasShirt struct {
	Shirt
}

func (a *Adidas) MakeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 24,
		},
	}
}

func (a *Adidas) MakeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "adidas",
			size: 12,
		},
	}
}
