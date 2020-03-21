package main

type bike struct {
	isWorking bool
	isDocked  bool
}

func newBike() bike {
	return bike{true, true}
}

func (b *bike) release() {
	b.isDocked = false
}

func (b *bike) dock() {
	b.isDocked = true
}
