package main

type bicycle interface {
	release()
	dock()
}

type bike struct {
	isWorking bool
	isDocked  bool
}
