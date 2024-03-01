package controller

import (
	"beneficiary-validate/model"
	"beneficiary-validate/validation"

	"github.com/gin-gonic/gin"
)

func (s *Server) validateBeneficiary(c *gin.Context) {
	var req model.BeneficiaryValidateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	config, err := validation.LoadConfigFromFile(req.Currency)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if !validation.ApplyValidationRules(config.GeneralRules, &req) {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	if !validation.ApplyConditionalRules(config.ConditionalRules, &req) {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	c.JSON(200, gin.H{"message": "valid request"})

}
