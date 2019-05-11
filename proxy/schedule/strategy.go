package schedule

const (
	PollName   = "poll"
	IpHashName = "iphash"
	RandomName = "random"
)

var registry = make(map[string]Strategy)

type Strategy interface {
	Init()
	Choose(string, []string) string
}

func init() {
	registry[PollName] = new(Poll)
	registry[IpHashName] = new(IpHash)
	registry[RandomName] = new(Random)
}

func GetStrategy(name string) Strategy {
	return registry[name]
}
