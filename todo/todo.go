package todo

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title string `json:"text"`
	Done  bool
}

func (Todo) TableName() string {
	return "todos"
}

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) CreateTodoHandler(c *gin.Context) {
	var todo Todo
	if err := c.Bind(&todo); err != nil {
		log.Println(err)
		return
	}

	if err := h.db.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) ListTodoHandler(c *gin.Context) {
	bearer := c.Request.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(bearer, "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("AllYourBase"), nil
	})
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["aud"])
	}

	var todos []Todo
	if err := h.db.Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, todos)
}
