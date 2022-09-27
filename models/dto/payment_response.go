package dto

import (
	"ark_backend_go/graph/model"
	"ark_backend_go/models/entity"
)

func ToPaymentMethodResponseDTO(payment *entity.PaymentMethod) *model.PaymentMethodGql {
	var status bool
	if payment.Status == 1 {
		status = true
	} else {
		status = false
	}
	return &model.PaymentMethodGql{
		ID:          payment.ID,
		Value:       payment.Value,
		Chanel:      payment.Chanel,
		Code:        payment.Code,
		Description: payment.Description,
		Image:       payment.Image,
		Limit:       payment.Limit,
		Status:      status,
		Tipe:        payment.Tipe,
		Title:       payment.Title,
		TitleType:   payment.TitleType,
	}
}

func ToListPaymentMethodResponseDTO(promo []*entity.PaymentMethod) []*PaymentMethodDTO {
	var listTemp = []*PaymentMethodDTO{}
	for _, data := range promo {
		listTemp = append(listTemp, &PaymentMethodDTO{
			ID:          data.ID,
			Value:       data.Value,
			Chanel:      data.Chanel,
			Code:        data.Code,
			Description: data.Description,
			Image:       data.Image,
			Limit:       data.Limit,
			Tipe:        data.Tipe,
			Title:       data.Title,
			TitleType:   data.TitleType,
		})
	}
	return listTemp
}

func ToListPaymentMethodResponseGQL(promo []*entity.PaymentMethod) []*model.PaymentMethodGql {
	var listTemp = []*model.PaymentMethodGql{}
	for _, data := range promo {
		var cvtStatus bool
		if data.Status == 1 {
			cvtStatus = true
		} else {
			cvtStatus = false
		}
		listTemp = append(listTemp, &model.PaymentMethodGql{
			ID:          data.ID,
			Value:       data.Value,
			Chanel:      data.Chanel,
			Code:        data.Code,
			Description: data.Description,
			Image:       data.Image,
			Limit:       data.Limit,
			Status:      cvtStatus,
			Tipe:        data.Tipe,
			Title:       data.Title,
			TitleType:   data.TitleType,
		})
	}
	return listTemp
}
