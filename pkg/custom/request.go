package custom

import (
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	FiberRequest interface {
		Bind(obj any) error
	}

	customFiberRequest struct {
		ctx       *fiber.Ctx
		validator *validator.Validate
	}
)

var (
	once              sync.Once
	validatorInstance *validator.Validate
)

// NewCustomFiberRequest creates a custom request handler for Fiber
func NewCustomFiberRequest(fiberCtx *fiber.Ctx) FiberRequest {
	once.Do(func() {
		validatorInstance = validator.New()
	})

	return &customFiberRequest{
		ctx:       fiberCtx,
		validator: validatorInstance,
	}
}

// Bind binds and validates the incoming request data (both Query Parameters and JSON Body)
func (r *customFiberRequest) Bind(obj any) error {
	// Try to parse from Query Parameters first
	if err := r.ctx.QueryParser(obj); err == nil {
		if err := r.validator.Struct(obj); err != nil {
			return err
		}
		return nil
	}

	// Fallback to parsing JSON Body
	if err := r.ctx.BodyParser(obj); err != nil {
		return err
	}

	// Validate the struct
	if err := r.validator.Struct(obj); err != nil {
		return err
	}

	return nil
}
