package handler

import (
	"net/http"

	"github.com/0zyyy/money_record/helper"
	"github.com/0zyyy/money_record/history"
	"github.com/gin-gonic/gin"
)

type HistoryHandler struct {
	historyService history.Service
}

func NewHistoryHandler(historyService history.Service) *HistoryHandler {
	return &HistoryHandler{historyService}
}

func (h *HistoryHandler) History(ctx *gin.Context) {
	var input history.Search

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		response := helper.ErrorResponse(err)
		ctx.JSON(http.StatusMethodNotAllowed, response)
		return
	}
}

func (h *HistoryHandler) Create(ctx *gin.Context) {
	var input history.NewHistoryInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		response := helper.ErrorResponse(err)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newHistory, err := h.historyService.Create(input)
	if err != nil {
		response := helper.ErrorResponse(err)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIResponse("Successfully created history", http.StatusOK, "success", newHistory)
	ctx.JSON(http.StatusOK, response)
}

func (h *HistoryHandler) Update(ctx *gin.Context) {
	var input history.NewHistoryInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ErrorResponse(err)
		errMsg := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to update history", http.StatusUnprocessableEntity, "fail", errMsg)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updated, err := h.historyService.Update(input)
	if err != nil {
		response := helper.ErrorResponse(err)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIResponse("Successfully updated history", http.StatusOK, "success", history.ResponseHistoryFormatter(updated))
	ctx.JSON(http.StatusOK, response)
}

func (h *HistoryHandler) Delete(ctx *gin.Context) {
	var input history.DeleteHistory

	err := ctx.ShouldBindJSON(input)
	if err != nil {
		response := helper.ErrorResponse(err)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	deletedHis, err := h.historyService.Delete(input.IDHistory)
	if err != nil {
		response := helper.ErrorResponse(err)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIResponse("Successfully deleted history", http.StatusOK, "success", deletedHis)
	ctx.JSON(http.StatusOK, response)
}

func (h *HistoryHandler) SearchHistory(ctx *gin.Context) {
	var input history.Search

	err := ctx.ShouldBindJSON(input)
	if err != nil {
		response := helper.ErrorResponse(err)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	data, err := h.historyService.SearchHistory(input.UserID, input.Date)
	if err != nil {
		response := helper.ErrorResponse(err)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIResponse("Successfully getting histories", http.StatusOK, "success", data)
	ctx.JSON(http.StatusOK, response)
}
