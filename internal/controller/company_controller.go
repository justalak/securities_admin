package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"securities_admin/internal/service"
	"strconv"
)

type ICompanyController interface {
	GetCompanies(ctx *gin.Context)
}

type CompanyController struct {
	companyService service.ICompanyService
}

func NewCompanyController(companyService service.ICompanyService) *CompanyController {
	return &CompanyController{companyService: companyService}
}

func (c *CompanyController) GetCompanies(ctx *gin.Context) {
	query := ctx.Param(CompanyQuery)

	page, err := strconv.Atoi(ctx.Query(PageQuery))
	if err != nil {
		page = 1
		log.Printf("Failed to convert logType %s", err)
	}

	logs, err := c.companyService.GetCompanies(page, query)
	if err != nil {
		ctx.JSON(HTTPCodeServerError, gin.H{
			"message": "Get Companies failed",
		})
	}

	ctx.JSON(HTTPCodeOK, gin.H{
		"data": logs,
	})
}
