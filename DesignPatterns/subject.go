package main

type Subject interface {
	register(observer Observer)
	removeObserver(observer Observer)
	notifyAll()
}
