package errors

import "fmt"

type BoatNotFoundError struct {
	StateNumber string
}
type BoatAlreadyOnWaterError struct {
	StateNumber string
}

type BoatAlreadyOnBaseError struct {
	StateNumber string
}

type ValidationError struct {
	Field string
}

// Ошибка при неверном значении при входе на сайт
func (e *BoatNotFoundError) Error() string {
	return fmt.Sprintf("Лодка с госномером %s не найдена", e.StateNumber)
}

// Ошибка при попытке создать заявку на спуск, когда статус на воде
func (e *BoatAlreadyOnWaterError) Error() string {
	return fmt.Sprintf("Лодка %s и так находится на воде!", e.StateNumber)
}

// Аналогичная ошибка только с подьемом
func (e *BoatAlreadyOnBaseError) Error() string {
	return fmt.Sprintf("Лодка %s и так находится на базе!", e.StateNumber)
}

// Ошибка при создании заявки(не указано какое-либо поле)
func (e *ValidationError) Error() string {
	return fmt.Sprintf("Поле %s, обязательно должно быть заполнено", e.Field)
}
