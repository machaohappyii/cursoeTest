package controllers

import (
	"gin-user-api/models"
	"gin-user-api/services"
	"gin-user-api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Age:      req.Age,
	}

	if err := uc.userService.CreateUser(user); err != nil {
		utils.InternalServerErrorResponse(c, "Failed to create user")
		return
	}

	utils.CreatedResponse(c, user)
}

func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.userService.GetAllUsers()
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to get users")
		return
	}

	utils.SuccessResponse(c, users)
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid user ID")
		return
	}

	user, err := uc.userService.GetUserByID(uint(id))
	if err != nil {
		utils.NotFoundResponse(c, "User not found")
		return
	}

	utils.SuccessResponse(c, user)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid user ID")
		return
	}

	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	updateData := &models.User{
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
	}

	user, err := uc.userService.UpdateUser(uint(id), updateData)
	if err != nil {
		utils.NotFoundResponse(c, "User not found or update failed")
		return
	}

	utils.SuccessResponse(c, user)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid user ID")
		return
	}

	if err := uc.userService.DeleteUser(uint(id)); err != nil {
		utils.InternalServerErrorResponse(c, "Failed to delete user")
		return
	}

	utils.SuccessResponse(c, gin.H{"message": "User deleted successfully"})
}

func (uc *UserController) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	user, err := uc.userService.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		utils.UnauthorizedResponse(c, "Invalid email or password")
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Name)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to generate token")
		return
	}

	response := models.LoginResponse{
		Token: token,
		User:  *user,
	}

	utils.SuccessResponse(c, response)
}

func (uc *UserController) Logout(c *gin.Context) {
	// 从请求头中获取 JWT token
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		utils.BadRequestResponse(c, "Authorization header required")
		return
	}

	// 提取 token（去掉 "Bearer " 前缀）
	token := ""
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		token = authHeader[7:]
	} else {
		utils.BadRequestResponse(c, "Invalid authorization header format")
		return
	}

	// 验证 token 是否有效
	if _, err := utils.ValidateToken(token); err != nil {
		utils.UnauthorizedResponse(c, "Invalid token")
		return
	}

	// TODO: 在实际应用中，这里应该将 token 添加到黑名单或缓存中
	// 例如：utils.AddTokenToBlacklist(token)
	// 或者：redisClient.Set(ctx, "blacklist_"+token, "true", tokenExpirationTime)

	utils.SuccessResponse(c, gin.H{"message": "Logout successful"})
}
