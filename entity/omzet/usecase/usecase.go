package usecase

import (
	"majoo-test-case/entity/omzet"
	"majoo-test-case/entity/omzet/repository"
	"math"
	"time"
)

type OmzetUsecase interface {
	GetMerchantOmzet(req *omzet.MerchantOmzetRequest) (*omzet.MerchantOmzetResponse, error)
	GetOutletOmzet(req *omzet.OutletOmzetRequest) (*omzet.OutletOmzetResponse, error)
}

type omzetUsecase struct {
	omzetRepository repository.OmzetRepository
}

func NewOmzetUseCase(u repository.OmzetRepository) OmzetUsecase {
	return &omzetUsecase{
		omzetRepository: u,
	}
}

func (o *omzetUsecase) GetMerchantOmzet(req *omzet.MerchantOmzetRequest) (*omzet.MerchantOmzetResponse, error) {
	offset := (req.Page * req.DataPerPage) - req.DataPerPage
	omzetObj, err := o.omzetRepository.GetMerchantOmzet(req.MerchantID, req.DataPerPage, offset)
	if err != nil {
		return nil, err
	}
	dayStart := offset + 1
	dayEnd := dayStart + req.DataPerPage - 1

	omzetList := new(omzet.MerchantOmzetResponse)
	for i := dayStart; i <= dayEnd; i++ {
		omzet := omzet.MerchantOmzet{}
		date := time.Date(2021, time.November, i, 0, 0, 0, 0, time.Local)
		dateString := date.Format("2006-01-02")
		if val, ok := omzetObj[dateString]; ok {
			valInterface := val.(map[string]interface{})
			omzet.Omzet = valInterface["omzet"].(int)
			omzet.Date = dateString
		} else {
			omzet.Omzet = 0
			omzet.Date = dateString
		}
		omzetList.Data = append(omzetList.Data, omzet)
	}
	omzetList.MerchantName, err = o.omzetRepository.GetMerchantName(req.MerchantID)
	if err != nil {
		return nil, err
	}
	omzetList.Pagination = paginate(req.Page, req.DataPerPage)
	return omzetList, nil
}

func (o *omzetUsecase) GetOutletOmzet(req *omzet.OutletOmzetRequest) (*omzet.OutletOmzetResponse, error) {
	offset := (req.Page * req.DataPerPage) - req.DataPerPage
	omzetObj, err := o.omzetRepository.GetOutletOmzet(req.OutletID, req.DataPerPage, offset)
	if err != nil {
		return nil, err
	}
	dayStart := offset + 1
	dayEnd := dayStart + req.DataPerPage - 1

	omzetList := new(omzet.OutletOmzetResponse)
	for i := dayStart; i <= dayEnd; i++ {
		omzet := omzet.OutletOmzet{}
		date := time.Date(2021, time.November, i, 0, 0, 0, 0, time.Local)
		dateString := date.Format("2006-01-02")
		if val, ok := omzetObj[dateString]; ok {
			valInterface := val.(map[string]interface{})
			omzet.Omzet = valInterface["omzet"].(int)
			omzet.Date = dateString
		} else {
			omzet.Omzet = 0
			omzet.Date = dateString
		}
		omzetList.Data = append(omzetList.Data, omzet)
	}
	omzetList.MerchantName, omzetList.OutletName, err = o.omzetRepository.GetOutletName(req.OutletID)
	if err != nil {
		return nil, err
	}
	omzetList.Pagination = paginate(req.Page, req.DataPerPage)
	return omzetList, nil
}

func paginate(page int, dataPerPage int) omzet.Pagination {
	pagination := omzet.Pagination{}
	pagination.TotalData = 30
	pagination.CurrentPage = page
	pagination.DataPerPage = dataPerPage
	pagination.LastPage = int(math.Ceil(float64(30) / float64(dataPerPage)))
	return pagination
}
