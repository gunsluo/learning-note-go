package main

type Collection interface {
	Add(node interface{}) bool
	Del(node interface{}) bool
	List()
}
