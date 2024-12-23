package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jmoiron/sqlx"
	"github.com/onizukazaza/tar-ecom-api/config"
)

type fiberServer struct {
	app  *fiber.App
	db   *sqlx.DB
	conf *config.Config
}

var (
	server *fiberServer
	once   sync.Once
)

func NewFiberServer(conf *config.Config, db *sqlx.DB) *fiberServer {
	// Initialize Fiber application
	fiberApp := fiber.New(fiber.Config{
		BodyLimit: conf.Server.BodyLimit,
	})

	once.Do(func() {
		server = &fiberServer{
			app:  fiberApp,
			db:   db,
			conf: conf,
		}
	})

	return server
}

// Start initializes middleware, routes, and starts the server
func (s *fiberServer) Start() {
	s.initMiddlewares()
	s.initRoutes()
	s.initAdminRouter()
	s.initProductManagingRouter()
	s.app.Use(getCORSMiddleware(s.conf.Server.AllowOrigins))
	s.app.Use(getTimeoutMiddleware(s.conf.Server.Timeout))
	
	// Graceful shutdown
	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)
	go s.gracefullyShutdown(quitCh)
	
	// Start server
	s.httpListening()
}

func (s *fiberServer) initMiddlewares() {
	
	customLogger := logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path} - ${latency}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
	})

	s.app.Use(recover.New())

	s.app.Use(customLogger)
}

func (s *fiberServer) initRoutes() {
	// Healthcheck route
	s.app.Get("/v1/healthcheck", s.healthCheck)
	s.app.Get("/panic", func(c *fiber.Ctx) error {
		panic("This is a test panic")
	})


}

func (s *fiberServer) httpListening() {
	addr := fmt.Sprintf(":%d", s.conf.Server.Port)
	log.Printf("[INFO] Server is starting on %s", addr)

	if err := s.app.Listen(addr); err != nil && err != http.ErrServerClosed {
		log.Fatalf("[ERROR] Failed to start server: %v", err)
	}
}

func (s *fiberServer) gracefullyShutdown(quitCh <-chan os.Signal) {
	<-quitCh
	log.Println("[INFO] Gracefully shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.app.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("[ERROR] Error shutting down: %s", err)
	}
	log.Println("[INFO] Server stopped successfully.")
}

// Routes Handlers
func (s *fiberServer) healthCheck(c *fiber.Ctx) error {
	// Database health check with sqlx
	var result string
	err := s.db.Get(&result, "SELECT 'UP' as status")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": "DOWN",
			"error":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{"status": result, "timestamp": time.Now().Format(time.RFC3339)})
}

// Custom Middleware
func getTimeoutMiddleware(timeout time.Duration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), timeout)
		defer cancel()

		c.SetUserContext(ctx)
		done := make(chan error, 1)
		go func() {
			done <- c.Next()
		}()

		select {
		case <-ctx.Done():
			return c.Status(http.StatusRequestTimeout).JSON(fiber.Map{
				"error": "Request timed out",
			})
		case err := <-done:
			return err
		}
	}
}

func getCORSMiddleware(allowOrigins []string) fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: strings.Join(allowOrigins, ","),
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	})
}
