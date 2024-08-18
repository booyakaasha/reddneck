package cursor

// Direction направление курсора.
type Direction int

const (
	// DirectionUnknown невалидное значение.
	DirectionUnknown Direction = iota
	// DirectionNext направление вперед.
	DirectionNext
	// DirectionPrev направление назад.
	DirectionPrev
)

type Marshallable interface {
	Marshal() string
	Unmarshal(s string) error
}

// Cursor курсор.
type Cursor[T Marshallable] struct {
	Direction Direction
	Value     T
	Limit     int
}

type Group struct {
	Next string
	Prev string
}

// func NewGroup
