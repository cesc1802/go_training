package services

import (
	"context"
	"database/sql"
	"encoding/json"

	mysql "github.com/cesc1802/go_training/internal/storages/mysql"
	"log"
	"net/http"
	"time"

	"github.com/cesc1802/go_training/internal/storages"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// ToDoServiceMySQL implement HTTP server
type ToDoServiceMySQL struct {
	JWTKey string
	Store  *mysql.MySQLDB
}

func (s *ToDoServiceMySQL) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "*")

	if req.Method == http.MethodOptions {
		resp.WriteHeader(http.StatusOK)
		return
	}

	switch req.URL.Path {
	case "/login":
		s.getAuthToken(resp, req)
		return
	case "/tasks":
		var ok bool
		req, ok = s.validToken(req)
		if !ok {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}

		switch req.Method {
		case http.MethodGet:
			s.listTasks(resp, req)
		case http.MethodPost:
			s.addTask(resp, req)
		}
		return
	}
}

func (s *ToDoServiceMySQL) getAuthToken(resp http.ResponseWriter, req *http.Request) {
	id := valueMySQL(req, "user_id")
	password := valueMySQL(req, "password")
	if !s.Store.ValidateUser(req.Context(), id, password) {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "incorrect user_id/pwd",
		})
		return
	}
	resp.Header().Set("Content-Type", "application/json")

	token, err := s.createToken(id.String)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]string{
		"data": token,
	})
}

func (s *ToDoServiceMySQL) listTasks(resp http.ResponseWriter, req *http.Request) {
	id, _ := userIDFromCtxMySQL(req.Context())
	tasks, err := s.Store.RetrieveTasks(
		req.Context(),
		sql.NullString{
			String: id,
			Valid:  true,
		},
		valueMySQL(req, "created_date"),
	)

	resp.Header().Set("Content-Type", "application/json")

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string][]*storages.Task{
		"data": tasks,
	})
}

func (s *ToDoServiceMySQL) addTask(resp http.ResponseWriter, req *http.Request) {
	t := &storages.Task{}
	err := json.NewDecoder(req.Body).Decode(t)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	now := time.Now()
	userID, _ := userIDFromCtxMySQL(req.Context())
	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedDate = now.Format("2006-01-02")

	maxTask := 9

	numberTasks, errTasks := s.numberTaskInDay(req)
	if errTasks != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	// check tasks
	if len(numberTasks) > maxTask {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(map[string]string{
			"errors": "Can't add tasks today",
		})
		return
	}

	resp.Header().Set("Content-Type", "application/json")

	err = s.Store.AddTask(req.Context(), t)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]*storages.Task{
		"data": t,
	})
}

func valueMySQL(req *http.Request, p string) sql.NullString {
	return sql.NullString{
		String: req.FormValue(p),
		Valid:  true,
	}
}

func (s *ToDoServiceMySQL) createToken(id string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(s.JWTKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *ToDoServiceMySQL) validToken(req *http.Request) (*http.Request, bool) {
	token := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(s.JWTKey), nil
	})
	if err != nil {
		log.Println(err)
		return req, false
	}

	if !t.Valid {
		return req, false
	}

	id, ok := claims["user_id"].(string)
	if !ok {
		return req, false
	}

	req = req.WithContext(context.WithValue(req.Context(), userAuthKeyMySQL(0), id))
	return req, true
}

type userAuthKeyMySQL int8

func userIDFromCtxMySQL(ctx context.Context) (string, bool) {
	v := ctx.Value(userAuthKeyMySQL(0))
	id, ok := v.(string)
	return id, ok
}

func (s *ToDoServiceMySQL) numberTaskInDay(req *http.Request) ([]string, error) {
	id, _ := userIDFromCtxMySQL(req.Context())
	date := valueMySQL(req, "created_date")
	tasks, err := s.Store.RetrieveTasks(
		req.Context(),
		sql.NullString{
			String: id,
			Valid:  true,
		},
		date,
	)
	var listTasks []string
	for i := range tasks {
		listTasks = append(listTasks, tasks[i].CreatedDate)
	}
	return listTasks, err
}
