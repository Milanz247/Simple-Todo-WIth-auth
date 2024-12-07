package utils

import "net/http"


func GetUserID(r *http.Request) uint {

	userID := r.Context().Value("user_id")
	if userID == nil {
		return 0 // Or handle error
	}
	
	return userID.(uint)
}
