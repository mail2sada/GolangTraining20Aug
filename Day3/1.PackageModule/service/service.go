package service

type Service struct {
	ServiceName        string
	ServiceID          int
	ServiceDescription string
}

func (s Service) SetServiceName(name string) Service {
	s.ServiceName = name
	return s
}

func (s Service) SetServiceID(id int) Service {
	s.ServiceID = id
	return s
}

func (s Service) SetServiceDescription(des string) Service {
	s.ServiceDescription = des
	return s
}

func (s Service) GetServiceDetails() (name string, id int, description string) {
	return s.ServiceName, s.ServiceID, s.ServiceDescription
}
