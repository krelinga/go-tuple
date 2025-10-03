package tuple

type T5[F0 any, F1 any, F2 any, F3 any, F4 any] struct {
	F0 F0
	F1 F1
	F2 F2
	F3 F3
	F4 F4
}

// First returns the first field of the tuple
func (t T5[F0, F1, F2, F3, F4]) First() F0 {
	return t.F0
}

// Last returns the last field of the tuple
func (t T5[F0, F1, F2, F3, F4]) Last() F4 {
	return t.F4
}
