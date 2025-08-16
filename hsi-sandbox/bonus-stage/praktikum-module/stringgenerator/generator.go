package stringgenerator

import "rsc.io/quote"

func MencetakSalam(name string) string {
	pesanIslami := membuatSalamIslami(name)
	return pesanIslami + " " + quote.Hello()
}

func membuatSalamIslami(name string) string {
	return "Assalamualaykum " + name
}