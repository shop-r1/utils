package tools

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shop-r1/utils/data"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	// Mime types definitions
	extraMimeTypes = map[string]string{
		".icon": "image-x-icon",
		".ttf":  "application/x-font-ttf",
		".woff": "application/x-font-woff",
		".eot":  "application/vnd.ms-fontobject",
		".svg":  "image/svg+xml",
		".html": "text/html; charset-utf-8",
	}

	// Paths that dont require database connection
	allowedPaths = map[string]bool{
		"/api/sessions":  true,
		"/api/info":      true,
		"/api/connect":   true,
		"/api/bookmarks": true,
		"/api/history":   true,
	}

	// List of characters replaced by javascript code to make queries url-safe.
	base64subs = map[string]string{
		"-": "+",
		"_": "/",
		".": "=",
	}
)

type Error struct {
	Message string `json:"error"`
}

func NewError(err error) Error {
	return Error{err.Error()}
}

// Returns a clean query without any comment statements
func CleanQuery(query string) string {
	lines := []string{}

	for _, line := range strings.Split(query, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "--") {
			continue
		}
		lines = append(lines, line)
	}

	return strings.TrimSpace(strings.Join(lines, "\n"))
}

func Desanitize64(query string) string {
	// Before feeding the string into decoded, we must "reconstruct" the base64 data.
	// Javascript replaces a few characters to be url-safe.
	for olds, news := range base64subs {
		query = strings.Replace(query, olds, news, -1)
	}

	return query
}

func GetSessionId(req *http.Request) string {
	id := req.Header.Get("x-session-id")
	if id == "" {
		id = req.URL.Query().Get("_session_id")
	}
	return id
}

func GetQueryParam(c *gin.Context, name string) string {
	result := ""
	q := c.Request.URL.Query()

	if len(q[name]) > 0 {
		result = q[name][0]
	}

	return result
}

func ParseIntFormValue(c *gin.Context, name string, defValue int) (int, error) {
	val := c.Request.FormValue(name)

	if val == "" {
		return defValue, nil
	}

	num, err := strconv.Atoi(val)
	if err != nil {
		return defValue, fmt.Errorf("%s must be a number", name)
	}

	if num < 1 && defValue != 0 {
		return defValue, fmt.Errorf("%s must be greater than 0", name)
	}

	return num, nil
}

func assetContentType(name string) string {
	ext := filepath.Ext(name)
	result := mime.TypeByExtension(ext)

	if result == "" {
		result = extraMimeTypes[ext]
	}

	if result == "" {
		result = "text/plain; charset=utf-8"
	}

	return result
}

func ServeStaticAsset(path string, c *gin.Context) {
	data, err := data.Asset(path)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.Data(200, assetContentType(path), data)
}

// Send a query result to client
func ServeResult(c *gin.Context, result interface{}, err interface{}) {
	if err == nil {
		successResponse(c, result)
	} else {
		badRequest(c, err)
	}
}

// Send successful response back to client
func successResponse(c *gin.Context, data interface{}) {
	c.JSON(200, data)
}

// Send an error response back to client
func errorResponse(c *gin.Context, status int, err interface{}) {
	var message interface{}

	switch v := err.(type) {
	case error:
		message = v.Error()
	case string:
		message = v
	default:
		message = v
	}

	c.AbortWithStatusJSON(status, gin.H{"status": status, "error": message})
}

// Send a bad request (http 400) back to client
func badRequest(c *gin.Context, err interface{}) {
	errorResponse(c, 400, err)
}
