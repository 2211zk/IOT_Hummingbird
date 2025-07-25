package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	userpb "kratos/api/user/v1"
	"kratos/internal/pkg"
	"kratos/internal/service"

	httpx "github.com/go-kratos/kratos/v2/transport/http"
)

func RegisterUserRoutes(srv *httpx.Server, userSvc *service.UserService) {
	srv.HandleFunc("/user/register", userRegisterHandler(userSvc))
	srv.HandleFunc("/user/login", userLoginHandler(userSvc))
	srv.HandleFunc("/user/get", userGetHandler(userSvc))
	srv.HandleFunc("/user/list", userListHandler(userSvc))
	srv.HandleFunc("/user/update", userUpdateHandler(userSvc))
	srv.HandleFunc("/user/delete", userDeleteHandler(userSvc))
}

func userRegisterHandler(svc *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req userpb.RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		reply, err := svc.Register(context.Background(), &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reply)
	}
}

func userLoginHandler(svc *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req userpb.LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		reply, err := svc.Login(context.Background(), &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reply)
	}
}

func userGetHandler(svc *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := checkJWT(r); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// 从查询参数获取用户ID
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "missing id parameter", http.StatusBadRequest)
			return
		}

		id, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			http.Error(w, "invalid id parameter", http.StatusBadRequest)
			return
		}

		req := &userpb.GetUserRequest{Id: int32(id)}
		reply, err := svc.GetUser(context.Background(), req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reply)
	}
}

func userListHandler(svc *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := checkJWT(r); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// 从查询参数获取分页信息
		pageStr := r.URL.Query().Get("page")
		pageSizeStr := r.URL.Query().Get("page_size")

		page := int32(1)
		pageSize := int32(10)

		if pageStr != "" {
			if p, err := strconv.ParseInt(pageStr, 10, 32); err == nil {
				page = int32(p)
			}
		}

		if pageSizeStr != "" {
			if ps, err := strconv.ParseInt(pageSizeStr, 10, 32); err == nil {
				pageSize = int32(ps)
			}
		}

		req := &userpb.ListUserRequest{Page: page, PageSize: pageSize}
		reply, err := svc.ListUser(context.Background(), req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reply)
	}
}

func userUpdateHandler(svc *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := checkJWT(r); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		var req userpb.UpdateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		reply, err := svc.UpdateUser(context.Background(), &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reply)
	}
}

func userDeleteHandler(svc *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := checkJWT(r); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// 从查询参数获取用户ID
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "missing id parameter", http.StatusBadRequest)
			return
		}

		id, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			http.Error(w, "invalid id parameter", http.StatusBadRequest)
			return
		}

		req := &userpb.DeleteUserRequest{Id: int32(id)}
		reply, err := svc.DeleteUser(context.Background(), req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reply)
	}
}

func checkJWT(r *http.Request) error {
	auth := r.Header.Get("Authorization")
	if !strings.HasPrefix(auth, "Bearer ") {
		return http.ErrNoCookie // 使用标准错误，调用方会处理HTTP状态码
	}
	tokenStr := strings.TrimPrefix(auth, "Bearer ")
	_, err := pkg.ParseToken(tokenStr)
	return err
}
