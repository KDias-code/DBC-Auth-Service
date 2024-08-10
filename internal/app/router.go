package app

func (s *server) router() {
	s.app.Get("/healthz", s.handlers.HealthCheck)
	s.app.Post("/sign-up", s.handlers.SignUp)
	s.app.Post("/sign-in", s.handlers.SignIn)
}
