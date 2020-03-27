package main

type bike struct {
	isWorking bool
	isDocked  bool
}

func newBike() bike {
	return bike{true, true}
}

func (b *bike) SetIsWorking(val bool) {
	b.isWorking = val
}

func (b *bike) SetIsDocked(val bool) {
	b.isDocked = val
}
