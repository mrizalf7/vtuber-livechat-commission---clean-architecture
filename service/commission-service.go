package service


import (
	"main/repository"
	"main/entity"
	"github.com/mashingan/smapping"
	"log"
	"fmt"
)


type CommissionService interface {
	InsertCommission(p entity.Commission) entity.Commission
	UpdateCommission(p entity.Commission) entity.Commission
	GetCommission() []entity.Commission
	DeleteCommission(p entity.Commission)
	FindCommissionById(id uint64) entity.Commission
}

type Commission struct {
	commissionRepository repository.Commission
}

func NewCommissionService(commissionRepo repository.Commission) CommissionService{
	return &Commission{
		commissionRepository : commissionRepo,
	}
}


func (service *Commission) InsertCommission(p entity.Commission) entity.Commission {
	Commission := entity.Commission{}
	err := smapping.FillStruct(&Commission, smapping.MapFields(&p))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.commissionRepository.InsertCommission(Commission)
	return res
}

func (service *Commission) UpdateCommission(p entity.Commission) entity.Commission {
	Commission := entity.Commission{}
	err := smapping.FillStruct(&Commission, smapping.MapFields(&p))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.commissionRepository.UpdateCommission(Commission)
	return res
}

func (service *Commission) DeleteCommission(p entity.Commission) {
	service.commissionRepository.DeleteCommission(p)
}

// func (service *CommissionService) DeleteCommissionByCategory(p entity.Commission,category_name string){
// 	service.CommissionRepository.DeleteCommissionByCategory(p,category_name)
// }

func (service *Commission) GetCommission() []entity.Commission {
	return service.commissionRepository.GetCommission()
}

func (service *Commission) FindCommissionById(id uint64) entity.Commission {
	fmt.Sprintf("%v", id) 
	return service.commissionRepository.FindCommissionById(id)
}


