package gogperf

// GPerfHelper 用于封装gperftools
type GPerfHelper interface {
	Start(name string)
	Started() bool
	Stop(name string)
}
