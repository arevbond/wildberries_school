package simple_factory

type Door struct {
	width  int
	height int
}

func New(width, height int) Door {
	return Door{
		width:  width,
		height: height,
	}
}
