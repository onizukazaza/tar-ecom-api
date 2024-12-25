package custom

// import (
// 	"sync"

// 	"github.com/go-playground/validator/v10"
// 	"github.com/gofiber/fiber/v2"
// )

// type (
// 	FiberRequest interface {
// 		Bind(obj any) error
// 	}

// 	customFiberRequest struct {
// 		ctx       *fiber.Ctx
// 		validator *validator.Validate
// 	}
// )

// var (
// 	once              sync.Once
// 	validatorInstance *validator.Validate
// )


// func NewCustomFiberRequest(fiberCtx *fiber.Ctx) FiberRequest {
// 	once.Do(func() {
// 		validatorInstance = validator.New()
// 	})

// 	return &customFiberRequest{
// 		ctx:       fiberCtx,
// 		validator: validatorInstance,
// 	}
// }

// func (r *customFiberRequest) Bind(obj any) error {
//     if err := r.ctx.QueryParser(obj); err == nil {
//         if err := r.validator.Struct(obj); err != nil {
//             return err
//         }
//         return nil
//     }

//     if err := r.ctx.BodyParser(obj); err != nil {
//         return err
//     }

//     // ตรวจสอบ id ใน path parameter หากมีฟิลด์ ID
//     if id := r.ctx.Params("id"); id != "" {
//         if objWithID, ok := obj.(interface{ SetID(string) }); ok {
//             objWithID.SetID(id)
//         }
//     }

//     if err := r.validator.Struct(obj); err != nil {
//         return err
//     }

//     return nil
// }
