package handlers

import (
	"bytes"
	"encoding/xml"

	"net/http"
	"strconv"

	"github.com/beevik/etree"
	"github.com/gin-gonic/gin"
	"github.com/user/test_template/dtos"
	"github.com/user/test_template/exutilities"
)

// ...

func HandleAdd(c *gin.Context) {
	// Parse the numbers to be added from the URL parameters
	numA, err := strconv.Atoi(c.Param("numA"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, exutilities.ErrorResponse{Message: err.Error()})
		return
	}

	numB, err := strconv.Atoi(c.Param("numB"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, exutilities.ErrorResponse{Message: err.Error()})
		return
	}

	// Create the SOAP envelope with the Add operation and parameters
	soapRequest := dtos.SOAPEnvelope{
		SoapEnv: "http://schemas.xmlsoap.org/soap/envelope/",
		Tem:     "http://tempuri.org/",
		Body: dtos.SOAPBody{
			Add: dtos.Add{
				IntA: numA,
				IntB: numB,
			},
		},
	}

	// Convert the SOAP request struct to XML
	requestXML, err := xml.MarshalIndent(soapRequest, "", "  ")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, exutilities.ErrorResponse{Message: err.Error()})
		return
	}

	// Send the SOAP request to the service
	resp, err := http.Post("http://www.dneonline.com/calculator.asmx", "text/xml; charset=utf-8", bytes.NewBuffer(requestXML))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, exutilities.ErrorResponse{Message: err.Error()})
		return
	}
	defer resp.Body.Close()

	// Read the SOAP response body
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	responseXML := buf.String()

	// Parse the SOAP response XML
	doc := etree.NewDocument()
	if err := doc.ReadFromString(responseXML); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, exutilities.ErrorResponse{Message: err.Error()})
		return
	}

	// Extract the result from the SOAP response
	resultElement := doc.FindElement(".//AddResult")
	if resultElement == nil {
		c.IndentedJSON(http.StatusInternalServerError, exutilities.ErrorResponse{Message: err.Error()})
		return
	}
	resultStr := resultElement.Text()
	result, err := strconv.Atoi(resultStr)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, exutilities.ErrorResponse{Message: err.Error()})
		return
	}
	addJSONResponse := dtos.AddJSONResponse{Result: result}

	c.IndentedJSON(http.StatusOK, addJSONResponse)
}
