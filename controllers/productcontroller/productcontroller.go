package productcontroller

import (
	"net/http"

	"github.com/kritimauludin/go-restapi-mux-gorm-mysql/helper"
	"github.com/kritimauludin/go-restapi-mux-gorm-mysql/models"
)

var ResponseJson = helper.ResponseJson
var ResponseError = helper.ResponseError

func Index(w http.ResponseWriter, r *http.Request)  {
	var products []models.Product

	if err :=  models.DB.Find(&products).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJson(w, http.StatusOK, products)
}
func Show(w http.ResponseWriter, r *http.Request)  {
	
}
func Create(w http.ResponseWriter, r *http.Request)  {
	
}
func Update(w http.ResponseWriter, r *http.Request)  {
	
}
func Delete(w http.ResponseWriter, r *http.Request)  {
	
}