package factory

type Musket struct {
	Gun
}

func NewMusket() IGun {
	return &Musket{
		Gun: Gun{
			name:  "musket",
			power: 10,
		},
	}
}
