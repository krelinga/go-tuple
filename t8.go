package tuple

type T8[F0 any, F1 any, F2 any, F3 any, F4 any, F5 any, F6 any, F7 any] struct {
	F0 F0
	F1 F1
	F2 F2
	F3 F3
	F4 F4
	F5 F5
	F6 F6
	F7 F7
}

// NewT8 creates a new T8 tuple
func NewT8[F0 any, F1 any, F2 any, F3 any, F4 any, F5 any, F6 any, F7 any](f0 F0, f1 F1, f2 F2, f3 F3, f4 F4, f5 F5, f6 F6, f7 F7) T8[F0, F1, F2, F3, F4, F5, F6, F7] {
	return T8[F0, F1, F2, F3, F4, F5, F6, F7]{
		F0: f0,
		F1: f1,
		F2: f2,
		F3: f3,
		F4: f4,
		F5: f5,
		F6: f6,
		F7: f7,
	}
}

// First returns the first field of the tuple
func (t T8[F0, F1, F2, F3, F4, F5, F6, F7]) First() F0 {
	return t.F0
}

// Last returns the last field of the tuple
func (t T8[F0, F1, F2, F3, F4, F5, F6, F7]) Last() F7 {
	return t.F7
}
