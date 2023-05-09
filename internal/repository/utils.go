package repository

func CreateOffset(page, limit int) int {
	return (page - 1) * limit
}
