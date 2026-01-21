package client

type Response[T any] struct {
	Data *T
	StatusCode int
	Status string
}