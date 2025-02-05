package factory

type Gun struct {
	name  string
	power int
}

func (gun *Gun) SetName(name string) {
	gun.name = name
}

func (gun *Gun) GetName() string {
	return gun.name
}

func (gun *Gun) SetPower(power int) {
	gun.power = power
}

func (gun *Gun) GetPower() int {
	return gun.power
}
