package service

import(

	"fmt"
	"log"
	"github.com/mashingan/smapping"
	"main/entity"
	"main/repository"
	"main/dto"

)

type PriceListService interface {
	InsertPriceList(p dto.PriceListCreateDTO) entity.PriceList
	UpdatePriceList(p dto.PriceListUpdateDTO) entity.PriceList
	GetPriceList() []entity.PriceList
	DeletePriceList(p entity.PriceList)
	FindPriceListById(id uint64) entity.PriceList
	DeletePriceListByCategory(p entity.PriceList)
	// IsAllowedToEdit(userID string, bookID uint64) bool
}

type priceListService struct {
	priceListRepository repository.PriceList
}

//NewPriceListService .....
func NewPriceListService(priceListRepo repository.PriceList) PriceListService {
	return &priceListService{
		priceListRepository: priceListRepo,
	}
}
//priceList
func (service *priceListService) InsertPriceList(p dto.PriceListCreateDTO) entity.PriceList {
	priceList := entity.PriceList{}
	err := smapping.FillStruct(&priceList, smapping.MapFields(&p))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.priceListRepository.InsertPriceList(priceList)
	return res
}

func (service *priceListService) UpdatePriceList(p dto.PriceListUpdateDTO) entity.PriceList {
	priceList := entity.PriceList{}
	err := smapping.FillStruct(&priceList, smapping.MapFields(&p))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.priceListRepository.UpdatePriceList(priceList)
	return res
}

func (service *priceListService) DeletePriceList(p entity.PriceList) {
	service.priceListRepository.DeletePriceList(p)
}

// func (service *priceListService) DeletePriceListByCategory(p entity.PriceList,category_name string){
// 	service.priceListRepository.DeletePriceListByCategory(p,category_name)
// }
func (service *priceListService) DeletePriceListByCategory(p entity.PriceList){
	service.priceListRepository.DeletePriceListByCategory(p)
}

func (service *priceListService) GetPriceList() []entity.PriceList {
	return service.priceListRepository.GetPriceList()
}

func (service *priceListService) FindPriceListById(id uint64) entity.PriceList {
	fmt.Sprintf("%v", id) 
	return service.priceListRepository.FindPriceListById(id)
}

// func (service *bookService) IsAllowedToEdit(userID string, bookID uint64) bool {
// 	b := service.bookRepository.FindBookByID(bookID)
// 	id := fmt.Sprintf("%v", b.UserID)
// 	return userID == id
// }
