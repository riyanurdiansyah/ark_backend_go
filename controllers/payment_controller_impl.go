package controller

import (
	"ark_backend_go/helper"
	service "ark_backend_go/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentControllerImpl struct {
	PaymentService service.PaymentService
}

func NewPaymentController(paymentService service.PaymentService) PaymentController {
	return &PaymentControllerImpl{
		PaymentService: paymentService,
	}
}

// GetPaymentMethod implements PaymentController
func (controller *PaymentControllerImpl) GetPaymentMethod(c *gin.Context) {

	paymentResponse := controller.PaymentService.GetPaymentMethod()
	responses := helper.DefaultResponse{
		Code:    http.StatusOK,
		Message: "Payment method has been listed",
		Data:    paymentResponse,
		Status:  true,
	}
	c.JSON(http.StatusOK, responses)
}
