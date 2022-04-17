package facade

func Sub() {
	c := NewConsole()
	println(c)
	u := c.GetCharacterAt(1)
	println(u)
}
