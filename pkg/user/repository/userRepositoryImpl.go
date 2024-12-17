package repository

type userRepositoryImpl struct {
	
} 

func NewUserRepositoryImpl() UserRepository {
	return &userRepositoryImpl{}
}