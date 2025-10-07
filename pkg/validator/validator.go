package customvalidator

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func New() *validator.Validate {
	v := validator.New()

	v.RegisterCustomTypeFunc(func(field reflect.Value) interface{} {
		if val, ok := field.Interface().(decimal.Decimal); ok {
			return val.String()
		}
		return nil
	}, decimal.Decimal{})

	if err := v.RegisterValidation("decimal_gt0", func(fl validator.FieldLevel) bool {
		s, ok := fl.Field().Interface().(string)
		if !ok {
			return false
		}
		val, err := decimal.NewFromString(s)
		if err != nil {
			return false
		}
		return val.GreaterThan(decimal.Zero)
	}); err != nil {
		panic(err)
	}

	if err := v.RegisterValidation("uuid_not_nil", func(fl validator.FieldLevel) bool {
		if u, ok := fl.Field().Interface().(uuid.UUID); ok {
			return u != uuid.Nil
		}
		return false
	}); err != nil {
		panic(err)
	}

	return v
}
