package customer

type Service struct {
	repository Repository
}

func (s *Service) GetById(uuid string) (Model, error) {
	return s.repository.GetById(uuid)
}

func (s *Service) GetAll() {
	s.repository.GetAll()
}

func (s *Service) Update() {
	s.repository.Update()
}

func (s *Service) Delete() {
	s.repository.Delete()
}

func (s *Service) Create() {
	s.repository.Create()
}
