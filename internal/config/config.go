package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Path     string `json:"path"`
	FilePath string `json:"filePath"`
	Method   string `json:"method"`
	Type     string `json:"type"`
	Response struct {
		Status int         `json:"status"`
		Body   interface{} `json:"body"`
	} `json:"response"`
}

// RoutesConfig holds the entire configuration for routes
type RoutesConfig struct {
	Routes []Route `json:"routes"`
}

// LoadConfig reads and parses all JSON files from the specified directory
func LoadConfig(dirPath string) (*RoutesConfig, error) {
	absPath, err := filepath.Abs(dirPath)
	if err != nil {
		return nil, err
	}

	var routesConfig RoutesConfig

	err = filepath.Walk(absPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			file, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			var tempConfig RoutesConfig
			if err := json.Unmarshal(file, &tempConfig); err != nil {
				return err
			}
			routesConfig.Routes = append(routesConfig.Routes, tempConfig.Routes...)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &routesConfig, nil
}

// CreateHandler generates a Gin handler function based on the provided status and body or file path
func CreateHandler(status int, body interface{}, filePath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if body != nil {
			c.JSON(status, body)
		} else if filePath != "" {
			c.File(filePath)
		} else {
			c.Status(status)
		}
	}
}

// RegisterRoutes dynamically registers routes based on the configuration
func RegisterRoutes(r *gin.Engine, config *RoutesConfig) {
	for _, route := range config.Routes {
		var handler gin.HandlerFunc
		if route.Response.Body != nil {
			handler = CreateHandler(route.Response.Status, route.Response.Body, "")
		} else {
			// Assume the file path is relative to the static directory
			filePath := filepath.Join("static", route.FilePath)
			handler = CreateHandler(route.Response.Status, nil, filePath)
		}
		switch route.Method {
		case "GET":
			r.GET(route.Path, handler)
		case "POST":
			r.POST(route.Path, handler)
		// Add more cases here for other HTTP methods if needed
		default:
			log.Printf("Unsupported method %s for path %s", route.Method, route.Path)
		}
	}
}
