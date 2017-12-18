package repositories

type Repository interface {
	Get() (interface{}, error)
	Save(data interface{}) error
	DeleteById(id interface{}) error
}
