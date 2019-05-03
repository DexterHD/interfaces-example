package circus

import "github.com/DexterHD/interfaces-example/animal"

func Perform(a animal.Animal) string { return a.Speaks() }
