package fp

func noStopLoop() (ret func()) {
	ret = func() {
		println("init ret")
	}
	ret = nil
	return func() {
		println("return ret")
		ret()
	}
}
