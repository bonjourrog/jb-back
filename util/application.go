package util

import "errors"

var (
	applicationStatus = map[string]bool{
		"Received":  true, // Aplicación recibida / pendiente
		"Viewed":    true, // Visto / en revisión
		"InProcess": true, // En proceso / entrevista agendada
		"Rejected":  true, // Rechazado / no seleccionado
		"Accepted":  true, // Aceptado / contratado
		"Cancelled": true, // Cancelado / retirado
		"OnHold":    true, // En espera / pausado
	}
)

func VerifyApplicationStatus(schedule string) error {
	if !applicationStatus[schedule] {
		return errors.New("invalid status")
	}
	return nil
}
