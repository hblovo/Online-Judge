package problem

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type ProblemService struct{}

func (apiService *ProblemService) GetProblemList(c *gin.Context) string {
	fmt.Println("Get ProblemList")
	return "Success"
}
