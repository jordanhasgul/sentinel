package constraints

// Boolean is a constraint that permits any boolean type.
type Boolean interface {
	~bool
}

// Unsigned is a constraint that permits any unsigned integer type.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Signed is a constraint that permits any signed integer type.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Integer is a constraint that permits any integer type.
type Integer interface {
	Unsigned | Signed
}

// Float is a constraint that permits any floating-point type.
type Float interface {
	~float32 | ~float64
}

// Real is a constraint that permits any real type.
type Real interface {
	Integer | Float
}

// Complex is a constraint that permits any complex type.
type Complex interface {
	~complex64 | ~complex128
}

// String is a constraint that permits any string type.
type String interface {
	~string
}

// Map is a constraint that permits any map type.
type Map[K comparable, V any] interface {
	~map[K]V
}

// Slice is a constraint that permits any slice type.
type Slice[T any] interface {
	~[]T
}

// Equated is a constraint that permits types which can be equated.
type Equated interface {
	comparable
}

// Ordered is a constraint that permits types which can be ordered.
type Ordered interface {
	Real | String
}
