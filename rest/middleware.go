package rest

import (
	"github.com/gin-gonic/gin"
)

// Middleware defines a type of data that represents
// a rest method middleware function.
type Middleware func(gin.HandlerFunc) gin.HandlerFunc
