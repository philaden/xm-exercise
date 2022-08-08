package repositories

import (
	"testing"

	"github.com/golang/mock/gomock"
	mocks "github.com/philaden/xm-go-challenge/src/application/mocks"
	"github.com/stretchr/testify/require"
)

func Test_GetUsers(t *testing.T) {

	users := mocks.GetMockUsersDto()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepository := mocks.NewMockIUserRepository(ctrl)
	userRepository.EXPECT().GetUsers().Return(users, nil)

	response, err := userRepository.GetUsers()
	require.NoError(t, err)
	require.NotEmpty(t, response)
	require.Equal(t, response, users)
}

func Test_CreateUser_With_Valid_Params(t *testing.T) {
	testId := "69ff4588-392a-4378-80ef-05fa21f53871"

	data := mocks.CreateMockUserPayload()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepository := mocks.NewMockIUserRepository(ctrl)
	userRepository.EXPECT().CreateUser(data.FirstName, data.LastName, data.PhoneNumber, data.Email, data.Password).Return(testId, nil)

	response, err := userRepository.CreateUser(data.FirstName, data.LastName, data.PhoneNumber, data.Email, data.Password)
	require.NoError(t, err)
	require.Equal(t, response, testId)
}
