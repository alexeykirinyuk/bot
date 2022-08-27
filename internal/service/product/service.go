package product

import "errors"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List(offset int, limit int) (products []Product, done bool) {
	done = false

	if offset > len(allProducts) {
		offset = 0
		done = true
	}

	right := offset + limit
	if right > len(allProducts) {
		right = len(allProducts)
		done = true
	}

	products = allProducts[offset:right]
	return
}

func (s *Service) Get(prodID int) (*Product, error) {
	for _, prod := range allProducts {
		if prod.ID == prodID {
			return &prod, nil
		}
	}

	return nil, errors.New("product not found")
}

func (s *Service) Add(title string) {
	allProducts = append(allProducts, Product{ID: allProducts[len(allProducts)-1].ID + 1, Title: title})
}

func (s *Service) Update(prodID int, title string) error {
	var product *Product
	for _, prod := range allProducts {
		if prod.ID == prodID {
			product = &prod
		}
	}

	if product == nil {
		return errors.New("product not found")
	}

	product.Title = title
	return nil
}

func (s *Service) Rm(prodID int) error {
	index := -1
	for idx, prod := range allProducts {
		if prod.ID == prodID {
			index = idx
		}
	}

	if index == -1 {
		return errors.New("product not found")
	}

	allProducts = append(allProducts[:index], allProducts[index+1:]...)
	return nil
}
