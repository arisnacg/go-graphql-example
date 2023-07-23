package graph

import (
	"net/http"

	"github.com/arisnacg/go-graphql-example/todo"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

type ReqParams struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func TodoGraphRouter(c *gin.Context) {
	var reqObj ReqParams
	if err := c.ShouldBindJSON(&reqObj); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := graphql.Do(graphql.Params{
		Context:        c,
		Schema:         todo.TodoRootSchema,
		RequestString:  reqObj.Query,
		VariableValues: reqObj.Variables,
		OperationName:  reqObj.Operation,
	})

	if len(result.Errors) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Errors,
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
