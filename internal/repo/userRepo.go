package repo

import (
	"ecommercestore/internal/conf"
	"ecommercestore/internal/database"
	"ecommercestore/internal/helper"
	"ecommercestore/internal/helper/jwt"
	"ecommercestore/internal/models"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user *models.User) error {
	data, _ := GetUser(map[string]interface{}{"email": user.Email})
	if data != nil {
		log.Info().Str("email", user.Email).Msg("Error email has already existed")
		return models.ErrEmailExisted
	}

	// logic
	salt := helper.GenerateSalt(50)
	toHash := user.Password + salt
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(toHash), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("Error hashing password")
		return err
	}
	user.Password = string(hashPassword)
	user.Salt = salt

	// create user
	err = database.DB.Create(user).Error
	if err != nil {
		log.Error().Err(err).Msg("Error creating new user")
		return helper.ErrCannotCreateEntity(helper.EntityUser, err)
	}
	return nil
}

func AuthenticateUser(sv *models.UserSignIn) (string, error) {
	user, err := GetUser(map[string]interface{}{"email": sv.Email})
	if err != nil {
		log.Error().Err(err).Str("email", sv.Email).Msg("Error fetching user for authentication")
		return "", models.ErrEmailOrPasswordInvalid
	}

	// logic
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(sv.Password+user.Salt)); err != nil {
		log.Error().Err(err).Msg("Error comparing hash and password")
		return "", models.ErrEmailOrPasswordInvalid
	}

	tokenProvider := jwt.NewJWTProvider(conf.NewConfig().JwtSecret)
	token := tokenProvider.GenerateToken(user.Id, helper.Expiry)

	return token, nil
}

func GetUser(condition map[string]interface{}) (*models.User, error) {
	var user models.User
	if err := database.DB.Where(condition).First(&user).Error; err != nil {
		log.Error().Err(err).Msg("Error fetching user")
		return nil, helper.ErrCannotGetEntity(helper.EntityUser, err)
	}
	return &user, nil
}
func GetListUsers(
	paging *helper.Paging,
) ([]models.User, error) {
	var users []models.User
	db := database.DB.Table(models.User{}.TableName())

	if err := db.Count(&paging.Total).Error; err != nil {
		log.Error().Err(err).Msg("Error counting user")
		return nil, helper.ErrDb(err)
	}

	db = db.Offset(paging.Limit * (paging.Page - 1))
	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&users).Error; err != nil {
		log.Error().Err(err).Msg("Error listing user")
		return nil, helper.ErrCannotListEntity(helper.EntityUser, err)
	}
	return users, nil
}

func UpdateInfoUser(id int, infoUser *models.InfoUser) error {
	err := database.DB.Where("id = ?", id).Updates(infoUser).Error
	if err != nil { //Duplicate entry 'toida@mail.com' for key 'users.email'
		log.Error().Err(err).Msg("Error updating user")
		return helper.ErrCannotUpdateEntity(helper.EntityUser, err)
	}
	return nil
}

func UpdatePasswordUser(user *models.User, passwordUser *models.UserPassword) error {
	// verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordUser.Password+user.Salt)); err != nil {
		return models.ErrPasswordInvalid
	}
	// get password
	toHash := passwordUser.NewPassword + user.Salt
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(toHash), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("Error hashing password for update password user")
		return err
	}
	password := string(hashPassword)

	// update password
	err = database.DB.Table(models.User{}.TableName()).Where("id = ?", user.Id).Updates(map[string]interface{}{"password": password}).Error
	if err != nil {
		log.Error().Err(err).Msg("Error updating new password user")
		return helper.ErrCannotCreateEntity("PASSWORD", err)
	}
	return nil
}
