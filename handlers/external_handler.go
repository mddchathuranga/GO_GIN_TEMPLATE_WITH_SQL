package handlers

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/user/test_template/exutilities"
)

func FetchJasonData(c *gin.Context) {
	idStr := c.Param("id")
	_, err := strconv.Atoi(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, exutilities.ErrorResponse{Message: err.Error()})
		return
	}
	apiUrl := "https://jsonplaceholder.typicode.com/todos/" + idStr
	response, err := http.Get(apiUrl)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, exutilities.ErrorResponse{Message: err.Error()})
		return
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		c.IndentedJSON(http.StatusInternalServerError, exutilities.ErrorResponse{Message: "API returned a non-200 status code"})
		return
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, exutilities.ErrorResponse{Message: err.Error()})
		return
	}
	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(data))

}
