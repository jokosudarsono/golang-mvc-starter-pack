package controllers

import (
	"net/http"
	"todo/core/response"
)

type DefaultController struct{}

func (ctrl *DefaultController) PageNotFound(w http.ResponseWriter, r *http.Request) {
	payload := map[string]interface{}{
		"status":      "errors",
		"status_code": "not_found",
		"message":     "route not found!",
	}

	response.JSON(w, http.StatusNotFound, payload)
}
