package demo

//type Repository interface {
//	SaveUser(user UserDto) (UserEntity, error)
//	FindAllUsers() ([]UserEntity, error)
//}
//
//type RepositoryImpl struct {
//	db *gorm.DB
//}
//
//func NewRepository(ds *db.DataSource) *RepositoryImpl {
//	if ds == nil {
//		log.Panic("DB connection can't be nil")
//	}
//	return &RepositoryImpl{
//		db: ds.Connection,
//	}
//}
//
//func (repo *RepositoryImpl) SaveUser(user UserDto) (UserEntity, error) {
//	id := uuid.New().String()
//	userEntity := UserEntity{Id: id, Email: user.Email}
//	result := repo.db.Clauses(clause.Returning{}).Select("Id", "Email").Create(&userEntity)
//	if result.Error != nil {
//		return userEntity, fmt.Errorf("SaveUser: %w", result.Error)
//	}
//	return userEntity, nil
//}
//
//func (repo *RepositoryImpl) FindAllUsers() ([]UserEntity, error) {
//	var users []UserEntity
//	err := repo.db.Find(&users).Error
//	if err != nil {
//		return users, fmt.Errorf("FindByBookIds: %w", err)
//	}
//	return users, nil
//}
