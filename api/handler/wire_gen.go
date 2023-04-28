// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package handler

import (
	"github.com/gin-gonic/gin"
	"go_web/api/usecase/auth"
	"go_web/api/usecase/user"
	"go_web/config"
	"go_web/domain/repository"
	"go_web/domain/service"
	"go_web/domain/service/implement"
	"go_web/infra/repository/mysql"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func initUserRepository(db *gorm.DB) repository.UserRepo {
	userRepo := mysql.NewUserRepo(db)
	return userRepo
}

func initUserService(db *gorm.DB) service.UserService {
	userRepo := mysql.NewUserRepo(db)
	userService := implement.NewUserService(userRepo)
	return userService
}

func InitAuthService(db *gorm.DB, cfg *config.Config) service.AuthService {
	userRepo := mysql.NewUserRepo(db)
	authService := implement.NewAuthService(cfg, userRepo)
	return authService
}

func initUcCreateUser(db *gorm.DB) user.CreateUser {
	userService := initUserService(db)
	userCreateUser := user.NewCreateUser(db, userService)
	return userCreateUser
}

func initUcUpdateUser(db *gorm.DB) user.UpdateUser {
	userService := initUserService(db)
	userUpdateUser := user.NewUpdateUser(db, userService)
	return userUpdateUser
}

func initUcDeleteUser(db *gorm.DB) user.DeleteUser {
	userService := initUserService(db)
	userDeleteUser := user.NewDeleteUser(db, userService)
	return userDeleteUser
}

func initUcGetUser(db *gorm.DB) user.GetUser {
	userService := initUserService(db)
	userGetUser := user.NewGetUser(userService)
	return userGetUser
}

func initUcLogin(db *gorm.DB, cfg *config.Config) auth.Login {
	authService := InitAuthService(db, cfg)
	userService := initUserService(db)
	login := auth.NewLogin(authService, userService)
	return login
}

func initUcAuthInit(db *gorm.DB, cfg *config.Config) auth.AuthInit {
	authService := InitAuthService(db, cfg)
	authInit := auth.NewAuthInit(authService)
	return authInit
}

func InitLoginHandler(db *gorm.DB, cfg *config.Config) func(ctx *gin.Context) {
	login := initUcLogin(db, cfg)
	v := doLogin(login, cfg)
	return v
}

func InitCreateUserHandler(db *gorm.DB, cfg *config.Config) func(ctx *gin.Context) {
	userCreateUser := initUcCreateUser(db)
	v := createUser(userCreateUser)
	return v
}

func InitUpdateUserHandler(db *gorm.DB, cfg *config.Config) func(ctx *gin.Context) {
	userUpdateUser := initUcUpdateUser(db)
	v := updateUser(userUpdateUser)
	return v
}

func InitDeleteUserHandler(db *gorm.DB, cfg *config.Config) func(ctx *gin.Context) {
	userDeleteUser := initUcDeleteUser(db)
	v := deleteUser(userDeleteUser)
	return v
}

func InitGetUserHandler(db *gorm.DB, cfg *config.Config) func(ctx *gin.Context) {
	userGetUser := initUcGetUser(db)
	v := getUser(userGetUser)
	return v
}

func InitAuthInitHandler(db *gorm.DB, cfg *config.Config) func(ctx *gin.Context) {
	authInit := initUcAuthInit(db, cfg)
	v := doAuthInit(authInit, cfg)
	return v
}

func initUcListUser(db *gorm.DB) user.ListUser {
	userService := initUserService(db)
	listUser := user.NewListUser(userService)
	return listUser
}
