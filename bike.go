package main

type bike struct {
	isWorking bool
	isDocked  bool
}

func newBike() bike {
	return bike{true, true}
}

func (b *bike) SetIsWorking(bln bool) {
	b.isWorking = bln
}

func (b *bike) SetIsDocked(bln bool) {
	b.isDocked = bln
}
