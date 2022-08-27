package product

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (c *Service) List() []Product {
	return allProducts
}
