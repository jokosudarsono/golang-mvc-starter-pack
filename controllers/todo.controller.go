package controllers

import (
	"net/http"
	"strconv"
	ct "todo/constants"
	"todo/core/response"
	"todo/models"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

type TodoController struct{}

var todoModel = models.TodoModel{}

func (ctrl *TodoController) GetTodos(w http.ResponseWriter, r *http.Request) {

	data, err := todoModel.GetTodos(context.Get(r, "user_id").(int64))
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err.Error())
	}

	results := map[string]interface{}{
		"status":  "success",
		"message": "Todos Found",
		"results": data,
	}

	response.JSON(w, http.StatusOK, results)
}

func (ctrl *TodoController) DetailTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := context.Get(r, "user_id").(int64)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err.Error())
		return
	}

	data, err := todoModel.DetailTodo(id, userId)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err.Error())
		return
	}

	results := map[string]interface{}{
		"status":  "success",
		"message": "Todo Found",
		"results": data,
	}

	response.JSON(w, http.StatusOK, results)
}

func (ctrl *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	userId := context.Get(r, "user_id").(int64)
	title := r.FormValue("title")
	description := r.FormValue("description")
	startDate := r.FormValue("start_date")
	endDate := r.FormValue("end_date")
	status := ct.StDraft

	if title == "" {
		response.Err(w, http.StatusBadRequest, "Title cannot be empty!")
		return
	}

	if startDate == "" {
		response.Err(w, http.StatusBadRequest, "Start Date cannot be empty!")
		return
	}

	if endDate == "" {
		response.Err(w, http.StatusBadRequest, "End Date cannot be empty!")
		return
	}

	data, err := todoModel.CreateTodo(
		userId,
		title,
		description,
		startDate,
		endDate,
		status,
	)

	if err != nil {
		response.Err(w, http.StatusInternalServerError, err.Error())
		return
	}

	results := map[string]interface{}{
		"status":  "success",
		"message": "Todo Created",
		"results": data,
	}

	response.JSON(w, http.StatusOK, results)
}

func (ctrl *TodoController) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := context.Get(r, "user_id").(int64)
	title := r.FormValue("title")
	description := r.FormValue("description")
	startDate := r.FormValue("start_date")
	endDate := r.FormValue("end_date")
	status := r.FormValue("status")

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err.Error())
		return
	}

	if title == "" {
		response.Err(w, http.StatusBadRequest, "Title cannot be empty!")
		return
	}

	if startDate == "" {
		response.Err(w, http.StatusBadRequest, "Start Date cannot be empty!")
		return
	}

	if endDate == "" {
		response.Err(w, http.StatusBadRequest, "End Date cannot be empty!")
		return
	}

	if status == "" {
		response.Err(w, http.StatusBadRequest, "Status cannot be empty!")
		return
	}

	data, err := todoModel.UpdateTodo(
		id,
		userId,
		title,
		description,
		startDate,
		endDate,
		status,
	)

	if err != nil {
		response.Err(w, http.StatusInternalServerError, err.Error())
		return
	}

	results := map[string]interface{}{
		"status":  "success",
		"message": "Todo Updated",
		"results": data,
	}

	response.JSON(w, http.StatusOK, results)
}

func (ctrl *TodoController) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := context.Get(r, "user_id").(int64)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err.Error())
		return
	}

	data, err := todoModel.UpdateStatus(id, userId, vars["status"])
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err.Error())
		return
	}

	results := map[string]interface{}{
		"status":  "success",
		"message": "Todo Updated",
		"results": data,
	}

	response.JSON(w, http.StatusOK, results)
}

func (ctrl *TodoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := context.Get(r, "user_id").(int64)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err.Error())
		return
	}

	data, err := todoModel.DetailTodo(id, userId)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err.Error())
		return
	}

	results := map[string]interface{}{
		"status":  "success",
		"message": "Todo Deleted",
		"results": data,
	}

	response.JSON(w, http.StatusOK, results)
}
