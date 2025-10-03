package tuple

type T6[F0 any, F1 any, F2 any, F3 any, F4 any, F5 any] struct {
	F0 F0
	F1 F1
	F2 F2
	F3 F3
	F4 F4
	F5 F5
}

// NewT6 creates a new T6 tuple
func NewT6[F0 any, F1 any, F2 any, F3 any, F4 any, F5 any](f0 F0, f1 F1, f2 F2, f3 F3, f4 F4, f5 F5) T6[F0, F1, F2, F3, F4, F5] {
	return T6[F0, F1, F2, F3, F4, F5]{
		F0: f0,
		F1: f1,
		F2: f2,
		F3: f3,
		F4: f4,
		F5: f5,
	}
}

// First returns the first field of the tuple
func (t T6[F0, F1, F2, F3, F4, F5]) First() F0 {
	return t.F0
}

// Last returns the last field of the tuple
func (t T6[F0, F1, F2, F3, F4, F5]) Last() F5 {
	return t.F5
}

// CutLast returns a tuple with one fewer field (omitting the last field)
func (t T6[F0, F1, F2, F3, F4, F5]) CutLast() T5[F0, F1, F2, F3, F4] {
	return T5[F0, F1, F2, F3, F4]{
		F0: t.F0,
		F1: t.F1,
		F2: t.F2,
		F3: t.F3,
		F4: t.F4,
	}
}

// Get returns all tuple field values
func (t T6[F0, F1, F2, F3, F4, F5]) Get() (F0, F1, F2, F3, F4, F5) {
	return t.F0, t.F1, t.F2, t.F3, t.F4, t.F5
}
