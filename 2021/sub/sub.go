package sub

type Sub struct {
	Position Position
	Aim      int
}

func (s *Sub) Move(c Command) {
	switch c.Direction {
	case Forward:
		s.Position.X += c.Units
		s.Position.Y += (c.Units * s.Aim)
	case Up:
		s.Aim -= c.Units
	case Down:
		s.Aim += c.Units
	}
}

func (s *Sub) ExecuteCommands(cmds []Command) {
	for _, c := range cmds {
		s.Move(c)
	}
}
