package tuple

type T2[F0 any, F1 any] struct {
	F0 F0
	F1 F1
}

// First returns the first field of the tuple
func (t T2[F0, F1]) First() F0 {
	return t.F0
}

// Last returns the last field of the tuple
func (t T2[F0, F1]) Last() F1 {
	return t.F1
}
