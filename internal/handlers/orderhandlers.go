package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/KozlovNikolai/CMDorders/internal/model"
	"github.com/KozlovNikolai/CMDorders/internal/store"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OrderHandler struct {
	logger *zap.Logger
	repo   store.IOrderRepository
}

func NewOrderHandler(logger *zap.Logger, repo store.IOrderRepository) *OrderHandler {
	return &OrderHandler{logger: logger, repo: repo}
}

// DeleteOrder(ctx context.Context, id uint64) error

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		h.logger.Error("Error binding JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.repo.CreateOrder(context.Background(), order)
	if err != nil {
		h.logger.Error("Error creating order", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	order.ID = id
	c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	employer, err := h.repo.GetOrderByID(context.Background(), uint64(id))
	if err != nil {
		h.logger.Error("Error getting order", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, employer)
}

func (h *OrderHandler) GetOrdersByPatientID(c *gin.Context) {
	patient_id, _ := strconv.Atoi(c.Param("patient_id"))
	is_active, _ := strconv.Atoi(c.Param("is_active"))

	orders, err := h.repo.GetOrdersByPatientID(context.Background(), uint64(patient_id), int8(is_active))
	if err != nil {
		h.logger.Error(fmt.Sprintf("Error getting orders where is_active=%d, patient_id=%d", is_active, patient_id), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) GetAllOrdersList(c *gin.Context) {
	is_active, _ := strconv.Atoi(c.Param("is_active"))
	orders, err := h.repo.GetAllOrdersList(context.Background(), int8(is_active))
	if err != nil {
		h.logger.Error(fmt.Sprintf("Error getting all orders where is_active=%d", is_active), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		h.logger.Error("Error binding JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.repo.UpdateOrder(context.Background(), order)
	if err != nil {
		h.logger.Error("Error updating order", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Updated successfully"})
}

func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.repo.DeleteOrder(context.Background(), uint64(id))
	if err != nil {
		h.logger.Error(fmt.Sprintf("Error deleting order whith id=%d", id), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Deleted successfully"})
}
