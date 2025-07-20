package response

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Generic Response Format
type APIResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"error,omitempty"`
}

type PaginatedResponse struct {
	SearchBy string      `json:"search_by,omitempty"`
	OrderBy  string      `json:"order_by,omitempty"`
	OrderDir string      `json:"order_dir,omitempty"`
	Page     int32       `json:"page"`
	PerPage  int32       `json:"per_page"`
	Total    int64       `json:"total"`
	Results  interface{} `json:"results"`
}

func PaginatedSuccess(c *gin.Context, statusCode int, message string, paginated PaginatedResponse) {
	c.JSON(statusCode, APIResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       paginated,
	})
}

// Success Response Helper
func Success(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, APIResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	})
}

// Error Response Helper
func Error(c *gin.Context, statusCode int, message string, err interface{}) {
	c.JSON(statusCode, APIResponse{
		StatusCode: statusCode,
		Message:    message,
		Error:      err,
	})
}

// Validation Error Response
func ValidationError(c *gin.Context, err interface{}) {
	c.JSON(http.StatusUnprocessableEntity, APIResponse{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    "Validation failed",
		Error:      err,
	})
}

// Unauthorized Error Response
func UnauthorizedError(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, APIResponse{
		StatusCode: http.StatusUnauthorized,
		Message:    message,
	})
}

// Forbidden Error Response
func ForbiddenError(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, APIResponse{
		StatusCode: http.StatusForbidden,
		Message:    message,
	})
}

// Not Found Error Response
func NotFoundError(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, APIResponse{
		StatusCode: http.StatusNotFound,
		Message:    message,
	})
}

// Bad Request Error Response
func BadRequestError(c *gin.Context, message string, err interface{}) {
	c.JSON(http.StatusBadRequest, APIResponse{
		StatusCode: http.StatusBadRequest,
		Message:    message,
		Error:      err,
	})
}

// Internal Server Error Response
func InternalServerError(c *gin.Context, message string, err interface{}) {
	c.JSON(http.StatusInternalServerError, APIResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
		Error:      err,
	})
}

type PaginationInput struct {
	SearchBy *string `json:"search_by"` // e.g., name or title
	OrderBy  *string `json:"order_by"`  // e.g., name, created_at
	OrderDir *string `json:"order_dir"` // asc or desc
	Page     *int32  `json:"page"`      // page number (1-based)
	PerPage  *int32  `json:"per_page"`  // items per page
}

func (p *PaginationInput) SetDefaults() {
	lowerOrderDir := strings.ToLower(*p.OrderDir)
	if lowerOrderDir != "asc" && lowerOrderDir != "desc" {
		defaultOrderDir := "desc"
		p.OrderDir = &defaultOrderDir
	} else {
		*p.OrderDir = lowerOrderDir
	}

	// Default for Page
	if p.Page == nil || *p.Page <= 0 {
		defaultPage := int32(1)
		p.Page = &defaultPage
	}

	// Default for PerPage
	if p.PerPage == nil || *p.PerPage <= 0 {
		defaultPerPage := int32(10)
		p.PerPage = &defaultPerPage
	}
}

func (p *PaginationInput) Offset() int32 {
	return (*p.Page - 1) * *p.PerPage
}
