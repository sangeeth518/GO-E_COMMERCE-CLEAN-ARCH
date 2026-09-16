package usecase

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	mockhelper "github.com/sangeeth518/go-Ecommerce/pkg/mock/mockhelper"
	mockrepo "github.com/sangeeth518/go-Ecommerce/pkg/mock/mockrepo"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/models"
	"github.com/stretchr/testify/assert"
)

func TestUserSignup_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockrepo := mockrepo.NewMockUserRepo(ctrl)
	mockhelper := mockhelper.NewMockHelper(ctrl)

	usecase := NewUserUsecase(mockrepo, mockhelper)

	user := models.UserDetails{
		Name:            "Sangeeth",
		Email:           "sangeeth@gmail.com",
		Password:        "123456",
		ConfirmPassword: "123456",
		Phone:           "9999999999",
	}

	userResponse := models.UserDetailsResponse{
		Id:    1,
		Name:  "Sangeeth",
		Email: "sangeeth@gmail.com",
		Phone: "9999999999",
	}

	mockrepo.EXPECT().CheckUserAvailability(user.Email).Return(false)
	mockhelper.EXPECT().PasswordHashing(user.Password).Return("hashedpassowrd", nil)

	expectedUser := user
	expectedUser.Password = "hashedpassowrd"

	mockrepo.EXPECT().UserSignup(expectedUser).Return(userResponse, nil)
	mockhelper.EXPECT().GenerateTokenClient(userResponse).Return("token123", nil)

	result, err := usecase.UserSignup(user)
	assert.NoError(t, err)
	assert.Equal(t, "sangeeth@gmail.com", result.User.Email)
	assert.Equal(t, "token123", result.Token)
	assert.NotEmpty(t, result.Token)

}
func TestUserSignup_UserAlreadyExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakeRepo := mockrepo.NewMockUserRepo(ctrl)
	fakeHelper := mockhelper.NewMockHelper(ctrl)

	usecase := NewUserUsecase(fakeRepo, fakeHelper)

	input := models.UserDetails{
		Name:            "Sangeeth",
		Email:           "existing@gmail.com",
		Phone:           "9876543210",
		Password:        "password123",
		ConfirmPassword: "password123",
	}

	fakeRepo.EXPECT().
		CheckUserAvailability("existing@gmail.com").
		Return(true) // user already exists

	result, err := usecase.UserSignup(input)

	assert.Error(t, err)
	assert.Equal(t, "user already exist", err.Error())
	assert.Empty(t, result.Token)
}

func TestUserSignup_PasswordMismatch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakeRepo := mockrepo.NewMockUserRepo(ctrl)
	fakeHelper := mockhelper.NewMockHelper(ctrl)

	usecase := NewUserUsecase(fakeRepo, fakeHelper)

	input := models.UserDetails{
		Name:            "Sangeeth",
		Email:           "sangeeth@gmail.com",
		Phone:           "9876543210",
		Password:        "password123",
		ConfirmPassword: "wrong123", // different
	}

	fakeRepo.EXPECT().
		CheckUserAvailability("sangeeth@gmail.com").
		Return(false)

	result, err := usecase.UserSignup(input)

	assert.Error(t, err)
	assert.Equal(t, "password dosen't match", err.Error())
	assert.Empty(t, result.Token)
}
func TestAddAdress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mockrepo.NewMockUserRepo(ctrl)
	mockHlper := mockhelper.NewMockHelper(ctrl)

	usecase := NewUserUsecase(mockRepo, mockHlper)

	userId := 1
	adress := models.AddAdress{
		Name:      "Sangeeth",
		HouseName: "ABC House",
		Street:    "Main Street",
		City:      "Pathanamthitta",
		State:     "Kerala",
		Phone:     "9999999999",
		Pin:       "689645",
	}

	mockRepo.EXPECT().AddAdress(userId, adress).Return(nil)

	err := usecase.AddAdress(userId, adress)

	if err != nil {
		t.Errorf(" expected no errpr %v", err)
	}

}

func TestAddAdress_error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockrepo := mockrepo.NewMockUserRepo(ctrl)
	mockhelper := mockhelper.NewMockHelper(ctrl)

	usecase := NewUserUsecase(mockrepo, mockhelper)
	userId := 1
	adress := models.AddAdress{
		Name:      "Sangeeth",
		HouseName: "ABC House",
		Street:    "Main Street",
		City:      "Pathanamthitta",
		State:     "Kerala",
		Phone:     "9999999999",
		Pin:       "689645",
	}

	mockrepo.EXPECT().AddAdress(userId, adress).Return(errors.New("internal error"))
	err := usecase.AddAdress(userId, adress)
	if err == nil {
		t.Errorf("Expected Error , got nil")
	}

}
