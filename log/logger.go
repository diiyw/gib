package log

func Print(types ...Type) {
	for _, t := range types {
		t()
	}
}

func Info(types ...Type) {
	for _, t := range types {
		t()
	}
}

func Error(types ...Type) {
	for _, t := range types {
		t()
	}
}

func Warning(types ...Type) {
	for _, t := range types {
		t()
	}
}
