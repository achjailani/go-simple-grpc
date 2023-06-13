package event

// register is a method initialize listener & event names
func (dpc *Dispatcher) register(listener Listener, names ...Name) {
	for _, name := range names {
		if _, ok := dpc.events[name]; ok {
			continue
		}

		dpc.events[name] = listener
	}
}
