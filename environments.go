package zebedee

type Environment int

const (
	Production Environment = iota
	Sandbox
)

var endpoints = map[Environment]string{
	Production: "https://api.zebedee.io/v0",
	Sandbox:    "https://sandbox-api.zebedee.io/v0",
}

func (e Environment) String() string {
	switch e {
	case Production:
		return "production"
	case Sandbox:
		return "sandbox"
	default:
		return "unknown"
	}
}
