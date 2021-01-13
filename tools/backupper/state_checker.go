package main

type StateChecker interface {
	Check() (ok bool, err error)
}
