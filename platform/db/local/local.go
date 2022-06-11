package local

type Monitors map[string]interface{}
type CriticalResources []string

func LocalMonitor() Monitors {
	lm := make(Monitors)
	return lm
}

func LocalCriticalResources() CriticalResources {
	var cr []string
	return cr
}
