package controller_test

import (
	"github.com/paulojr83/Go-Expert/Labs/temperatura-por-CEP/internal/api/wep/controller"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestFindCep_InvalidCep(t *testing.T) {
	// Criar um novo contexto Gin
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	// Definir o CEP inválido nos parâmetros
	c.Params = append(c.Params, gin.Param{Key: "cep", Value: "0115300"})

	// Executar o controlador
	controller.FindCep(c)

	// Verificar o código de status HTTP
	assert.Equal(t, http.StatusUnprocessableEntity, c.Writer.Status())
}

func TestFindCep_StatusInternalServerError(t *testing.T) {
	// Criar um novo contexto Gin
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = append(c.Params, gin.Param{Key: "cep", Value: "01153000"})

	// Executar o controlador
	controller.FindCep(c)

	// Verificar o código de status HTTP
	assert.Equal(t, http.StatusInternalServerError, c.Writer.Status())
}
