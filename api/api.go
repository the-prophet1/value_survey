package api

import (
	"value-survey/internal/handler"
	"value-survey/internal/server"
	. "value-survey/pkg/errors"
)

func RegisterRouter(s *server.HttpServer) Error {
	balanceSheetHandler := handler.GetBalanceSheetHandler()
	server.POST(s, "/balance_sheet", balanceSheetHandler.Create)

	return nil
}
