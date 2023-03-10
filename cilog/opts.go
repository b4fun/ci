package cilog

type applyOpts[T any] interface {
	apply(*T)
}

type applyOptsFunc[T any] func(*T)

var _ applyOpts[any] = (applyOptsFunc[any])(nil)

// nolint: unused
func (f applyOptsFunc[T]) apply(t *T) {
	f(t)
}
