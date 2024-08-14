package problem

import "github.com/gin-gonic/gin"

type ProblemApi struct{}

func (api *ProblemApi) GetProblemList(c *gin.Context) {
	ProblemService.GetProblemList(c)
}
