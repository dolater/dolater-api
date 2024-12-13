// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	AppCheckScopes   = "AppCheck.Scopes"
	AuthBearerScopes = "AuthBearer.Scopes"
)

// Defines values for TaskPoolType.
const (
	Active   TaskPoolType = "active"
	Archived TaskPoolType = "archived"
	Bin      TaskPoolType = "bin"
	Pending  TaskPoolType = "pending"
)

// CreateTaskInput defines model for CreateTaskInput.
type CreateTaskInput struct {
	Url string `json:"url"`
}

// Error defines model for Error.
type Error struct {
	Message *string `json:"message,omitempty"`
}

// FollowStatus defines model for FollowStatus.
type FollowStatus struct {
	From      User      `json:"from"`
	Timestamp time.Time `json:"timestamp"`
	To        User      `json:"to"`
}

// Notification defines model for Notification.
type Notification struct {
	Body      *string            `json:"body,omitempty"`
	CreatedAt time.Time          `json:"createdAt"`
	Id        openapi_types.UUID `json:"id"`
	Title     string             `json:"title"`
	Url       *string            `json:"url,omitempty"`
}

// Task defines model for Task.
type Task struct {
	ArchivedAt  *time.Time         `json:"archivedAt,omitempty"`
	CompletedAt *time.Time         `json:"completedAt,omitempty"`
	CreatedAt   time.Time          `json:"createdAt"`
	Id          openapi_types.UUID `json:"id"`
	Owner       User               `json:"owner"`
	Pool        TaskPool           `json:"pool"`
	RemovedAt   *time.Time         `json:"removedAt,omitempty"`
	Url         string             `json:"url"`
}

// TaskPool defines model for TaskPool.
type TaskPool struct {
	Id    openapi_types.UUID `json:"id"`
	Owner *User              `json:"owner,omitempty"`
	Type  TaskPoolType       `json:"type"`
}

// TaskPoolType defines model for TaskPool.Type.
type TaskPoolType string

// UpdateTaskInput defines model for UpdateTaskInput.
type UpdateTaskInput struct {
	ArchivedAt  *time.Time          `json:"archivedAt,omitempty"`
	CompletedAt *time.Time          `json:"completedAt,omitempty"`
	PoolId      *openapi_types.UUID `json:"poolId,omitempty"`
	RemovedAt   *time.Time          `json:"removedAt,omitempty"`
	Url         *string             `json:"url,omitempty"`
}

// UpdateUserInput defines model for UpdateUserInput.
type UpdateUserInput struct {
	DisplayName *string `json:"displayName,omitempty"`
	PhotoURL    *string `json:"photoURL,omitempty"`
}

// User defines model for User.
type User struct {
	DisplayName string `json:"displayName"`
	Id          string `json:"id"`
	PhotoURL    string `json:"photoURL"`
}

// Id defines model for id.
type Id = openapi_types.UUID

// PoolId defines model for poolId.
type PoolId = openapi_types.UUID

// Uid defines model for uid.
type Uid = string

// FollowStatusOK defines model for FollowStatusOK.
type FollowStatusOK = FollowStatus

// InternalServerError defines model for InternalServerError.
type InternalServerError = Error

// MethodNotAllowed defines model for MethodNotAllowed.
type MethodNotAllowed = Error

// NotFound defines model for NotFound.
type NotFound = Error

// TaskOK defines model for TaskOK.
type TaskOK = Task

// TaskPoolOK defines model for TaskPoolOK.
type TaskPoolOK = TaskPool

// Unauthorized defines model for Unauthorized.
type Unauthorized = Error

// UserOK defines model for UserOK.
type UserOK = User

// UsersOK defines model for UsersOK.
type UsersOK = []User

// UpsertFCMToken defines model for UpsertFCMToken.
type UpsertFCMToken struct {
	Token string `json:"token"`
}

// UpsertFCMTokenJSONBody defines parameters for UpsertFCMToken.
type UpsertFCMTokenJSONBody struct {
	Token string `json:"token"`
}

// GetTasksParams defines parameters for GetTasks.
type GetTasksParams struct {
	// PoolId Pool ID
	PoolId *PoolId `form:"poolId,omitempty" json:"poolId,omitempty"`
}

// UpsertFCMTokenJSONRequestBody defines body for UpsertFCMToken for application/json ContentType.
type UpsertFCMTokenJSONRequestBody UpsertFCMTokenJSONBody

// CreateTaskJSONRequestBody defines body for CreateTask for application/json ContentType.
type CreateTaskJSONRequestBody = CreateTaskInput

// UpdateTaskJSONRequestBody defines body for UpdateTask for application/json ContentType.
type UpdateTaskJSONRequestBody = UpdateTaskInput

// UpdateTaskForciblyJSONRequestBody defines body for UpdateTaskForcibly for application/json ContentType.
type UpdateTaskForciblyJSONRequestBody = UpdateTaskInput

// UpdateUserJSONRequestBody defines body for UpdateUser for application/json ContentType.
type UpdateUserJSONRequestBody = UpdateUserInput

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get Notifications
	// (GET /notifications)
	GetNotifications(c *gin.Context)
	// Upsert FCM Token
	// (PATCH /notifications/fcmToken)
	UpsertFCMToken(c *gin.Context)
	// Get Pools
	// (GET /pools)
	GetPools(c *gin.Context)
	// Get Pool
	// (GET /pools/{id})
	GetPool(c *gin.Context, id Id)
	// Get Tasks
	// (GET /tasks)
	GetTasks(c *gin.Context, params GetTasksParams)
	// Create Task
	// (POST /tasks)
	CreateTask(c *gin.Context)
	// Delete Task
	// (DELETE /tasks/{id})
	DeleteTask(c *gin.Context, id Id)
	// Get Task
	// (GET /tasks/{id})
	GetTask(c *gin.Context, id Id)
	// Update Task
	// (PATCH /tasks/{id})
	UpdateTask(c *gin.Context, id Id)
	// Update Task Forcibly
	// (PUT /tasks/{id})
	UpdateTaskForcibly(c *gin.Context, id Id)
	// Get Users
	// (GET /users)
	GetUsers(c *gin.Context)
	// Create User
	// (POST /users)
	CreateUser(c *gin.Context)
	// Delete User
	// (DELETE /users/{uid})
	DeleteUser(c *gin.Context, uid Uid)
	// Get User
	// (GET /users/{uid})
	GetUser(c *gin.Context, uid Uid)
	// Update User
	// (PATCH /users/{uid})
	UpdateUser(c *gin.Context, uid Uid)
	// Unfollow User
	// (DELETE /users/{uid}/follow)
	UnfollowUser(c *gin.Context, uid Uid)
	// Follow User
	// (POST /users/{uid}/follow)
	FollowUser(c *gin.Context, uid Uid)
	// Get Followers
	// (GET /users/{uid}/followers)
	GetFollowers(c *gin.Context, uid Uid)
	// Get Followings
	// (GET /users/{uid}/followings)
	GetFollowings(c *gin.Context, uid Uid)
	// Get Friends
	// (GET /users/{uid}/friends)
	GetFriends(c *gin.Context, uid Uid)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetNotifications operation middleware
func (siw *ServerInterfaceWrapper) GetNotifications(c *gin.Context) {

	c.Set(AppCheckScopes, []string{})

	c.Set(AuthBearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetNotifications(c)
}

// UpsertFCMToken operation middleware
func (siw *ServerInterfaceWrapper) UpsertFCMToken(c *gin.Context) {

	c.Set(AppCheckScopes, []string{})

	c.Set(AuthBearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpsertFCMToken(c)
}

// GetPools operation middleware
func (siw *ServerInterfaceWrapper) GetPools(c *gin.Context) {

	c.Set(AppCheckScopes, []string{})

	c.Set(AuthBearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetPools(c)
}

// GetPool operation middleware
func (siw *ServerInterfaceWrapper) GetPool(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id Id

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(AppCheckScopes, []string{})

	c.Set(AuthBearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetPool(c, id)
}

// GetTasks operation middleware
func (siw *ServerInterfaceWrapper) GetTasks(c *gin.Context) {

	var err error

	c.Set(AppCheckScopes, []string{})

	c.Set(AuthBearerScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetTasksParams

	// ------------- Optional query parameter "poolId" -------------

	err = runtime.BindQueryParameter("form", true, false, "poolId", c.Request.URL.Query(), &params.PoolId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter poolId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetTasks(c, params)
}

// CreateTask operation middleware
func (siw *ServerInterfaceWrapper) CreateTask(c *gin.Context) {

	c.Set(AppCheckScopes, []string{})

	c.Set(AuthBearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateTask(c)
}

// DeleteTask operation middleware
func (siw *ServerInterfaceWrapper) DeleteTask(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id Id

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(AppCheckScopes, []string{})

	c.Set(AuthBearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteTask(c, id)
}

// GetTask operation middleware
func (siw *ServerInterfaceWrapper) GetTask(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id Id

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(AppCheckScopes, []string{})

	c.Set(AuthBearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetTask(c, id)
}

// UpdateTask operation middleware
func (siw *ServerInterfaceWrapper) UpdateTask(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id Id

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(AppCheckScopes, []string{})

	c.Set(AuthBearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateTask(c, id)
}

// UpdateTaskForcibly operation middleware
func (siw *ServerInterfaceWrapper) UpdateTaskForcibly(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id Id

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(AppCheckScopes, []string{})

	c.Set(AuthBearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateTaskForcibly(c, id)
}

// GetUsers operation middleware
func (siw *ServerInterfaceWrapper) GetUsers(c *gin.Context) {

	c.Set(AppCheckScopes, []string{})

	c.Set(AuthBearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUsers(c)
}

// CreateUser operation middleware
func (siw *ServerInterfaceWrapper) CreateUser(c *gin.Context) {

	c.Set(AppCheckScopes, []string{})

	c.Set(AuthBearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateUser(c)
}

// DeleteUser operation middleware
func (siw *ServerInterfaceWrapper) DeleteUser(c *gin.Context) {

	var err error

	// ------------- Path parameter "uid" -------------
	var uid Uid

	err = runtime.BindStyledParameterWithOptions("simple", "uid", c.Param("uid"), &uid, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter uid: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(AppCheckScopes, []string{})

	c.Set(AuthBearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteUser(c, uid)
}

// GetUser operation middleware
func (siw *ServerInterfaceWrapper) GetUser(c *gin.Context) {

	var err error

	// ------------- Path parameter "uid" -------------
	var uid Uid

	err = runtime.BindStyledParameterWithOptions("simple", "uid", c.Param("uid"), &uid, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter uid: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(AppCheckScopes, []string{})

	c.Set(AuthBearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUser(c, uid)
}

// UpdateUser operation middleware
func (siw *ServerInterfaceWrapper) UpdateUser(c *gin.Context) {

	var err error

	// ------------- Path parameter "uid" -------------
	var uid Uid

	err = runtime.BindStyledParameterWithOptions("simple", "uid", c.Param("uid"), &uid, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter uid: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(AppCheckScopes, []string{})

	c.Set(AuthBearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateUser(c, uid)
}

// UnfollowUser operation middleware
func (siw *ServerInterfaceWrapper) UnfollowUser(c *gin.Context) {

	var err error

	// ------------- Path parameter "uid" -------------
	var uid Uid

	err = runtime.BindStyledParameterWithOptions("simple", "uid", c.Param("uid"), &uid, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter uid: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(AppCheckScopes, []string{})

	c.Set(AuthBearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UnfollowUser(c, uid)
}

// FollowUser operation middleware
func (siw *ServerInterfaceWrapper) FollowUser(c *gin.Context) {

	var err error

	// ------------- Path parameter "uid" -------------
	var uid Uid

	err = runtime.BindStyledParameterWithOptions("simple", "uid", c.Param("uid"), &uid, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter uid: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(AppCheckScopes, []string{})

	c.Set(AuthBearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.FollowUser(c, uid)
}

// GetFollowers operation middleware
func (siw *ServerInterfaceWrapper) GetFollowers(c *gin.Context) {

	var err error

	// ------------- Path parameter "uid" -------------
	var uid Uid

	err = runtime.BindStyledParameterWithOptions("simple", "uid", c.Param("uid"), &uid, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter uid: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(AppCheckScopes, []string{})

	c.Set(AuthBearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetFollowers(c, uid)
}

// GetFollowings operation middleware
func (siw *ServerInterfaceWrapper) GetFollowings(c *gin.Context) {

	var err error

	// ------------- Path parameter "uid" -------------
	var uid Uid

	err = runtime.BindStyledParameterWithOptions("simple", "uid", c.Param("uid"), &uid, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter uid: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(AppCheckScopes, []string{})

	c.Set(AuthBearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetFollowings(c, uid)
}

// GetFriends operation middleware
func (siw *ServerInterfaceWrapper) GetFriends(c *gin.Context) {

	var err error

	// ------------- Path parameter "uid" -------------
	var uid Uid

	err = runtime.BindStyledParameterWithOptions("simple", "uid", c.Param("uid"), &uid, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter uid: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(AppCheckScopes, []string{})

	c.Set(AuthBearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetFriends(c, uid)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/notifications", wrapper.GetNotifications)
	router.PATCH(options.BaseURL+"/notifications/fcmToken", wrapper.UpsertFCMToken)
	router.GET(options.BaseURL+"/pools", wrapper.GetPools)
	router.GET(options.BaseURL+"/pools/:id", wrapper.GetPool)
	router.GET(options.BaseURL+"/tasks", wrapper.GetTasks)
	router.POST(options.BaseURL+"/tasks", wrapper.CreateTask)
	router.DELETE(options.BaseURL+"/tasks/:id", wrapper.DeleteTask)
	router.GET(options.BaseURL+"/tasks/:id", wrapper.GetTask)
	router.PATCH(options.BaseURL+"/tasks/:id", wrapper.UpdateTask)
	router.PUT(options.BaseURL+"/tasks/:id", wrapper.UpdateTaskForcibly)
	router.GET(options.BaseURL+"/users", wrapper.GetUsers)
	router.POST(options.BaseURL+"/users", wrapper.CreateUser)
	router.DELETE(options.BaseURL+"/users/:uid", wrapper.DeleteUser)
	router.GET(options.BaseURL+"/users/:uid", wrapper.GetUser)
	router.PATCH(options.BaseURL+"/users/:uid", wrapper.UpdateUser)
	router.DELETE(options.BaseURL+"/users/:uid/follow", wrapper.UnfollowUser)
	router.POST(options.BaseURL+"/users/:uid/follow", wrapper.FollowUser)
	router.GET(options.BaseURL+"/users/:uid/followers", wrapper.GetFollowers)
	router.GET(options.BaseURL+"/users/:uid/followings", wrapper.GetFollowings)
	router.GET(options.BaseURL+"/users/:uid/friends", wrapper.GetFriends)
}
