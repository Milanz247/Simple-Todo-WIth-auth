package controllers

import (
	"Golang/config"
	"Golang/models"
	"Golang/utils"
	"encoding/json"
	"net/http"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {

    userID := utils.GetUserID(r)
	
    if userID == 0 {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    w.Header().Set("Content-Type", "application/json")

    db := config.InitDB()

    var todos []models.Todo

    result := db.Where("user_id = ?", userID).Find(&todos)
    if result.Error != nil {
        http.Error(w, "Error fetching todos", http.StatusInternalServerError)
        return
    }
	
    json.NewEncoder(w).Encode(todos)
}


func CreateTodos(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	db := config.InitDB()

	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	userID := utils.GetUserID(r)
	todo.UserID = userID

	result := db.Create(&todo)
	if result.Error != nil {
		http.Error(w, "Error inserting todo into database", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "Todo created successfully",
		"todo":    todo,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}


func UpdateTodos(w http.ResponseWriter, r *http.Request) {

	userID := utils.GetUserID(r)

    w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("ID")
	if id == "" {
        http.Error(w, "ID is required", http.StatusBadRequest)
        return
    }

	var todo models.Todo

	db := config.InitDB()

	  result := db.Where("id = ? AND user_id = ?", id, userID).First(&todo)
	  if result.Error != nil {
		  if result.RowsAffected == 0 {
			  http.Error(w, "Todo not found or not authorized", http.StatusNotFound)
		  } else {
			  http.Error(w, "Error fetching todo", http.StatusInternalServerError)
		  }
		  return
	  }

	var updatedData models.Todo
    if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
        http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
        return
    }

	todo.Title = updatedData.Title
    todo.Complete = updatedData.Complete

	saveResult := db.Save(&todo)
	if saveResult.Error != nil {
		http.Error(w, "Error updating todo", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "Todo updated successfully",
		"todo":    todo,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	    
}

		

func DeleteTodos(w http.ResponseWriter, r *http.Request) {

	userID := utils.GetUserID(r)
    if userID == 0 {
        http.Error(w, `{"error": "Unauthorized"}`, http.StatusUnauthorized)
        return
    }

	id := r.PathValue("ID")
    if id == "" {
        http.Error(w, `{"error": "ID is required"}`, http.StatusBadRequest)
        return
    }

    db  := config.InitDB()
    
    result := db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Todo{})
    if result.Error != nil {
        http.Error(w, `{"error": "Error deleting todo"}`, http.StatusInternalServerError)
        return
    }

    if result.RowsAffected == 0 {
        http.Error(w, `{"error": "Todo not found or unauthorized"}`, http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
        "message": "Todo deleted successfully",
    }

    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, `{"error": "JSON encoding failed"}`, http.StatusInternalServerError)
    }
}