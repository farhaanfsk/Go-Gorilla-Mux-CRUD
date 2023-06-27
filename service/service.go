package service

import (
	"fmt"

	M "github.com/farhaanfsk/Go-Gorilla-Mux-CRUD/models"
	"github.com/farhaanfsk/Go-Gorilla-Mux-CRUD/repository"
	"github.com/google/uuid"
)

type Service struct {
	empRepo repository.EmployeeRepo
}

var Employees []M.Employee

func (s *Service) GetEmployees() []M.Employee {
	return s.empRepo.GetAllEmployees()
}

func (s *Service) CreateEmployee(e M.Employee) M.Employee {
	emp := s.empRepo.Save(e)
	return emp
}

func (s *Service) GetEmployee(id string) M.Employee {
	uid, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("invalid uuid passed")
		return M.Employee{}
	}
	fmt.Println(uid)
	return s.empRepo.GetEmployeeById(uid)
}

func (s *Service) DeleteEmployee(id string) bool {
	uid, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("invalid uuid passed")
		return false
	}
	return s.empRepo.DeleteEmployeeById(uid)
}

func (s *Service) UpdateEmployee(e M.Employee) bool {
	return s.empRepo.UpdateEmployee(e)
}

func NewService(repo repository.EmployeeRepo) Service {
	return Service{repo}
}
