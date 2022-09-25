package controller

import "github.com/gin-gonic/gin"

type PaymentController interface {
	GetPaymentMethod(c *gin.Context)
}
