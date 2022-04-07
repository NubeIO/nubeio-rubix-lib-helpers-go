package rest

/*
REST Setters
*/

func (s *Service) SetPath(path string) *Service {
	s.Path = path
	return s
}

func (s *Service) SetBody(body interface{}) *Service {
	if s.Options != nil {
		s.Options.Body = body
	}
	return s
}

func (s *Service) SetMethod(method string) *Service {
	s.Method = method
	return s
}
