package http

import (
	"go-crud-learn/domain"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAppDataHandlerGetListAppData(t *testing.T) {
	type fields struct {
		appDataUsecase domain.AppDataUsecase
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &appDataHandler{
				appDataUsecase: tt.fields.appDataUsecase,
			}
			h.GetListAppData(tt.args.c)
		})
	}
}
