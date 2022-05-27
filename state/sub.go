package state

func Sub() {
	classicImpl()
}

func classicImpl() {
	sw := NewSwitch()
	sw.On()
	sw.On()
	sw.Off()
	sw.Off()
}
