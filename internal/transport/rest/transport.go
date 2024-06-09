package rest

import (
	"encoding/json"
	"net/http"

	"github.com/nurhidaylma/lover-app.git/internal/endpoint"
	"github.com/nurhidaylma/lover-app.git/internal/model"
)

type RESTServer struct {
	endpoints endpoint.LoverEndpoint
}

func NewRESTServer(endpoints endpoint.LoverEndpoint) *RESTServer {
	return &RESTServer{endpoints: endpoints}
}

func (s *RESTServer) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "failed to decode request body", http.StatusBadRequest)
		return
	}

	_, err = s.endpoints.SignUpEndpoint(user)
	if err != nil {
		http.Error(w, "failed to sign up: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
	})
}

func (s *RESTServer) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "failed to decode request body", http.StatusBadRequest)
		return
	}

	_, err = s.endpoints.LoginEndpoint(user.Email, user.Password)
	if err != nil {
		http.Error(w, "failed to login: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
	})
}

func (s *RESTServer) SetProfileHandler(w http.ResponseWriter, r *http.Request) {
	var profile model.Profile
	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		http.Error(w, "failed to decode request body", http.StatusBadRequest)
		return
	}

	_, err = s.endpoints.SetProfileEndpoint(r.Context(), profile)
	if err != nil {
		http.Error(w, "failed to login: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
	})
}

func (s *RESTServer) SwipeHanlder(w http.ResponseWriter, r *http.Request) {
	var swipe model.Swipe
	err := json.NewDecoder(r.Body).Decode(&swipe)
	if err != nil {
		http.Error(w, "failed to decode request body", http.StatusBadRequest)
		return
	}

	_, err = s.endpoints.SwipeEndpoint(r.Context(), swipe)
	if err != nil {
		http.Error(w, "failed to login: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
	})
}

func (s *RESTServer) UpgradeToPremiumHandler(w http.ResponseWriter, r *http.Request) {
	var purchase model.UserPurchase
	err := json.NewDecoder(r.Body).Decode(&purchase)
	if err != nil {
		http.Error(w, "failed to decode request body", http.StatusBadRequest)
		return
	}

	_, err = s.endpoints.UpgradeToPremiumEndpoint(r.Context(), purchase)
	if err != nil {
		http.Error(w, "failed to login: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
	})
}
