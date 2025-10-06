package entity

import "errors"

var (
	// Errores 404 - Recursos no encontrados
	ErrNotFound            = errors.New("resource not found")
	ErrJobNotFound         = errors.New("job not found")
	ErrApplicationNotFound = errors.New("application not found")
	ErrUserNotFound        = errors.New("user not found")

	// Errores 403 - Autorización
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")

	// Errores 400 - Validación
	ErrInvalidInput = errors.New("invalid input")
	ErrInvalidID    = errors.New("invalid ID format")
	ErrMissingField = errors.New("required field missing")

	// Errores 409 - Conflictos
	ErrAlreadyExists = errors.New("resource already exists")
	ErrConflict      = errors.New("operation conflicts with current state")

	// Errores 500 - Internos
	ErrInternal = errors.New("internal server error")
	ErrDatabase = errors.New("database error")
	ErrExternal = errors.New("external service error")
)
