package carservice

type CreateCarRequest struct {
	Brand string
	Model string
	Year  int64
	Price float64
}
