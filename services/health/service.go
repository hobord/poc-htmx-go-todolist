package health

type service struct {
	checkers []func() error
}

func NewService(checkers ...func() error) Service {
	return &service{
		checkers: checkers,
	}
}

func (s *service) Health() error {
	for _, checker := range s.checkers {
		if err := checker(); err != nil {
			return err
		}
	}

	return nil
}

func (s *service) AddChecker(checker func() error) {
	s.checkers = append(s.checkers, checker)
}
