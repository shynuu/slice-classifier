package runtime

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// HandleUpdateAdmissionControl handle the change of the policies in the slice
func HandleUpdateAdmissionControl(c *gin.Context) {

	var adm ADM

	if err := c.BindJSON(&adm); err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Error retrieving parameters",
		})
	}

	err := runtime.UpdateAdmissionControl(&adm)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"error": fmt.Sprintf("Error Control Data on %s", RuntimeConfig.ClassifierName),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Control data success on %s", RuntimeConfig.ClassifierName),
	})

}

// HandleAdmissionControl handle the admission control
func HandleAdmissionControl(c *gin.Context) {

	var adm ADM

	if err := c.BindJSON(&adm); err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Error retrieving parameters",
		})
	}

	err := runtime.AdmissionControl(&adm)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"error": fmt.Sprintf("Error Control Data on %s", RuntimeConfig.ClassifierName),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Control data success on %s", RuntimeConfig.ClassifierName),
	})

}

// HandlePDU handles the PDU sessions
func HandlePDU(c *gin.Context) {

	var pdu PDU

	if err := c.BindJSON(&pdu); err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Error retrieving parameters",
		})
	}

	err := runtime.NewPDU(&pdu)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"error": fmt.Sprintf("Error Pipe Data on %s", RuntimeConfig.ClassifierName),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Pipe success on %s", RuntimeConfig.ClassifierName),
	})
}
