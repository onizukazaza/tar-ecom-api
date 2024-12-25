package custom

import (
	"fmt"
	"strings"
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


func NewCustomFiberRequest(fiberCtx *fiber.Ctx) FiberRequest {
	once.Do(func() {
		validatorInstance = validator.New()
	})

	return &customFiberRequest{
		ctx:       fiberCtx,
		validator: validatorInstance,
	}
}

// Bind binds and validates the incoming request data (Query Parameters, JSON Body, and Path Parameters)
func (r *customFiberRequest) Bind(obj any) error {
	// Parse Query Parameters
	if err := r.ctx.QueryParser(obj); err != nil {
		return fmt.Errorf("failed to parse query parameters: %w", err)
	}

	// Parse JSON Body
	if err := r.ctx.BodyParser(obj); err != nil {
		if strings.Contains(err.Error(), "Unprocessable Entity") {
			return fmt.Errorf("failed to parse JSON body: invalid JSON format")
		}
		return fmt.Errorf("failed to parse JSON body: %w", err)
	}

	// Parse Path Parameters
	if id := r.ctx.Params("id"); id != "" {
		if objWithID, ok := obj.(interface{ SetID(string) }); ok {
			objWithID.SetID(id)
		}
	}

	// Validate Struct
	if err := r.validator.Struct(obj); err != nil {
		return r.formatValidationError(err)
	}

	return nil
}


func (r *customFiberRequest) formatValidationError(err error) error {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var messages []string
		for _, e := range validationErrors {
			messages = append(messages, fmt.Sprintf("Field '%s' failed validation with tag '%s'", e.Field(), e.Tag()))
		}
		return fmt.Errorf("validation errors: %s", strings.Join(messages, ", "))
	}
	return err
}
