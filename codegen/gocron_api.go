// Package codegen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.2 DO NOT EDIT.
package codegen

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// Defines values for NotificationType.
const (
	Feishu   NotificationType = "feishu"
	Telegram NotificationType = "telegram"
)

// Cron defines model for Cron.
type Cron struct {
	Command string `json:"command"`
	Id      *int   `json:"id,omitempty"`
	Name    string `json:"name"`
	Notify  int    `json:"notify"`
	Time    string `json:"time"`
}

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// Log defines model for Log.
type Log struct {
	CronId string  `json:"cron_id"`
	Id     *int    `json:"id,omitempty"`
	Log    string  `json:"log"`
	Name   string  `json:"name"`
	Time   *string `json:"time,omitempty"`
}

// Notification defines model for Notification.
type Notification struct {
	Content *string           `json:"content,omitempty"`
	Id      *int              `json:"id,omitempty"`
	Name    *string           `json:"name,omitempty"`
	Title   *string           `json:"title,omitempty"`
	Type    *NotificationType `json:"type,omitempty"`
	Url     *string           `json:"url,omitempty"`
}

// NotificationType defines model for Notification.Type.
type NotificationType string

// OK defines model for OK.
type OK struct {
	Message *string `json:"message,omitempty"`
}

// BadRequest defines model for BadRequest.
type BadRequest = Error

// GetCronJobsOK defines model for GetCronJobsOK.
type GetCronJobsOK struct {
	Data    *[]Cron `json:"data,omitempty"`
	Message *string `json:"message,omitempty"`
}

// GetLogByIdOK defines model for GetLogByIdOK.
type GetLogByIdOK struct {
	Data *[]Log `json:"data,omitempty"`
}

// GetNotificationsOK defines model for GetNotificationsOK.
type GetNotificationsOK struct {
	Data *[]Notification `json:"data,omitempty"`
}

// OKRequest defines model for OKRequest.
type OKRequest = OK

// CreateCronJobJSONRequestBody defines body for CreateCronJob for application/json ContentType.
type CreateCronJobJSONRequestBody = Cron

// CreateNotificationJSONRequestBody defines body for CreateNotification for application/json ContentType.
type CreateNotificationJSONRequestBody = Notification

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all cron jobs
	// (GET /cron)
	GetCronJobs(ctx echo.Context) error
	// Create a cron job
	// (POST /cron)
	CreateCronJob(ctx echo.Context) error
	// Delete a cron job
	// (DELETE /cron/{id})
	DeleteCronJob(ctx echo.Context, id int) error
	// Get log by id
	// (GET /log/{id})
	GetLogById(ctx echo.Context, id int64) error
	// Get all notifications
	// (GET /notify)
	GetNotifications(ctx echo.Context) error
	// Create a notification
	// (POST /notify)
	CreateNotification(ctx echo.Context) error
	// Delete a notification
	// (DELETE /notify/{id})
	DeleteNotification(ctx echo.Context, id int) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetCronJobs converts echo context to params.
func (w *ServerInterfaceWrapper) GetCronJobs(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetCronJobs(ctx)
	return err
}

// CreateCronJob converts echo context to params.
func (w *ServerInterfaceWrapper) CreateCronJob(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateCronJob(ctx)
	return err
}

// DeleteCronJob converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteCronJob(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteCronJob(ctx, id)
	return err
}

// GetLogById converts echo context to params.
func (w *ServerInterfaceWrapper) GetLogById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetLogById(ctx, id)
	return err
}

// GetNotifications converts echo context to params.
func (w *ServerInterfaceWrapper) GetNotifications(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetNotifications(ctx)
	return err
}

// CreateNotification converts echo context to params.
func (w *ServerInterfaceWrapper) CreateNotification(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateNotification(ctx)
	return err
}

// DeleteNotification converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteNotification(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteNotification(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/cron", wrapper.GetCronJobs)
	router.POST(baseURL+"/cron", wrapper.CreateCronJob)
	router.DELETE(baseURL+"/cron/:id", wrapper.DeleteCronJob)
	router.GET(baseURL+"/log/:id", wrapper.GetLogById)
	router.GET(baseURL+"/notify", wrapper.GetNotifications)
	router.POST(baseURL+"/notify", wrapper.CreateNotification)
	router.DELETE(baseURL+"/notify/:id", wrapper.DeleteNotification)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"LWT1bf2/AAAA///pp/ZCyxoAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
