package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/time/rate"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

const defaultPort = ":8081"

var buildtime string = "-"

func main() {
	_, err := os.Create("/tmp/live")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove("/tmp/live")

	err = godotenv.Load()
	if err != nil {
		log.Println("not found .env file")
	}

	port := defaultPort
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	limiter := rate.NewLimiter(5, 5)

	r := gin.Default()
	r.GET("/x", func(c *gin.Context) {
		if !limiter.Allow() {
			c.Status(http.StatusTooManyRequests)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"buildtime": buildtime,
		})
	})
	r.GET("/healthz", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	r.POST("/ping/:id", pingPongHandler)
	r.POST("/logins", loginHandler)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	log.Println("listening on", port)

	go func() {
		srv.ListenAndServe() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	}()

	<-ctx.Done()
	stop()
	fmt.Println("shutting down gracefully, press Ctrl+C again to force")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(timeoutCtx); err != nil {
		fmt.Println(err)
	}
}

type Credetial struct {
	Account  string
	Password string
}

func loginHandler(c *gin.Context) {
	var cred Credetial
	if err := c.Bind(&cred); err != nil {
		return
	}

	mySigningKey := []byte("AllYourBase")

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute).Unix(),
		Audience:  cred.Account,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": ss,
	})
}

type Person struct {
	Name string `json:"name" binding:"required"`
}

func pingPongHandler(c *gin.Context) {
	var person Person

	if err := c.ShouldBind(&person); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": person.Name + " " + c.Param("id"),
	})
}
