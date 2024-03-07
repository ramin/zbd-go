package zebedee

type Environment int

const (
	Production Environment = iota
	Sandbox
	Public
)

var endpoints = map[Environment]string{
	Production: "https://api.zebedee.io/v0",
	Sandbox:    "https://sandbox-api.zebedee.io/v0",
	Public:     "https://api.zebedee.io/public/v1",
}

func (e Environment) String() string {
	switch e {
	case Production:
		return "production"
	case Sandbox:
		return "sandbox"
	case Public:
		return "public"
	default:
		return "unknown"
	}
}
