package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/osscameroon/jobsika/pkg/models/v1beta"
	log "github.com/sirupsen/logrus"
)

//PostPay handler for POST
func PostPay(c *gin.Context) {
	contentType := c.Request.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		log.Error(errors.New("Wrong contentType"))
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "could not post rating"})
		return
	}

	query := v1beta.PayPostQuery{}
	if err := c.ShouldBind(&query); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "could not post rating"})
		return
	}

	err := query.Validate()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}

	//Record customer payment request

	//Create a new opencollective tier
	//The deletion should happen on the webhook or a day after created

	//Send the link to the newly created opencollective tier to back to the client
	c.JSON(http.StatusOK, gin.H{"message": "Payment request recorded"})
}
