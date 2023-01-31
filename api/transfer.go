package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/techschool/simplebank/db/sqlc"
)

type CreateTransferRequest struct {
	FromAccountID int64 `json:"from_account_id" binding:"required"`
	ToAccountID int64 `json:"to_account_id" binding:"required"`
	Amount int64 `json:"amount" binding:"required,gt=0"`
	Currency string `json:"currency" binding:"required", oneof=USD EUR VND`
}

func (server *Server) CreateTransfer(ctx *gin.Context){
	var req CreateTransferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID: req.ToAccountID,
		Amount: req.Amount,	
	}

	if !server.IsValidAccount(ctx, req.FromAccountID, req.Currency){
		return
	}

	if !server.IsValidAccount(ctx, req.ToAccountID, req.Currency){
		return
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// IsValidAccount check if account is exist or account's currency is match with transfer's currency
func (server *Server) IsValidAccount(ctx *gin.Context, accountID int64, currency string) bool {
	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account with ID '%d' currency mismatch: %s vs %s", accountID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}

	return true
}