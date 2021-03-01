package appliance

import "errors"

type Appliance interface {
	Start()
	GetPurpose()
}

const (
	MICROWAVE = iota
	STOVE
	FRIDGE
)

func GetAppliance(i int) (Appliance, error) {
	switch i {
	case MICROWAVE:
		return new(Microwave), nil
	case STOVE:
		return new(Stove), nil
	case FRIDGE:
		return new(Fridge), nil
	default:
		return nil, errors.New("Invalid Appliance choice")
	}
}
