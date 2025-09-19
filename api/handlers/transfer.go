package handlers

import (
	"net/http"

	db "github.com/Ali-Hasan-Khan/go-bankify/db/sqlc"
	"github.com/gin-gonic/gin"
)

type TransferHandler struct {
	store *db.Store
}

func NewTransferHandler(store *db.Store) *TransferHandler {
	return &TransferHandler{store: store}
}

type createTransferRequest struct {
	FromAccountID int32 `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int32 `json:"to_account_id" binding:"required,min=1"`
	Amount        int64 `json:"amount" binding:"required,min=0"`
}

func (handler *TransferHandler) CreateTransfer(ctx *gin.Context) {
	var req createTransferRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: int64(req.FromAccountID),
		ToAccountID:   int64(req.ToAccountID),
		Amount:        req.Amount,
	}

	transfer, err := handler.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, transfer)
}

type listTransfersRequest struct {
	FromAccountID int32 `form:"from_account_id" binding:"required,min=1"`
	ToAccountID   int32 `form:"to_account_id" binding:"required,min=1"`
	PageID        int32 `form:"page_id" binding:"required,min=1"`
	PageSize      int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (handler *TransferHandler) ListTransfers(ctx *gin.Context) {
	var req listTransfersRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListTransferParams{
		FromAccountID: int64(req.FromAccountID),
		ToAccountID:   int64(req.ToAccountID),
		Limit:         int64(req.PageSize),
		Offset:        int64(req.PageID-1) * int64(req.PageSize),
	}

	transfers, err := handler.store.ListTransfer(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, transfers)
}
