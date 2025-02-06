package usecase

import (
	reflect "reflect"
	"testing"

	"github.com/Ayyasy123/pt-aka-tech-vision/internal/model"
	"github.com/Ayyasy123/pt-aka-tech-vision/internal/repository"
	"github.com/golang/mock/gomock"
)

func TestUserUseCase_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockRepo := NewMockUserRepository(ctrl)

	type fields struct {
		repo repository.UserRepository
	}
	type args struct {
		req model.CreateUserReq
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		mockFn   func()
		wantData *model.UserRes
		wantErr  bool
	}{
		{
			"Test create user success",
			fields{
				repo: mockRepo,
			},
			args{
				req: model.CreateUserReq{
					Name:     "Test name",
					Email:    "email@example.com",
					Password: "",
				},
			},
			func() {
				// cek email sudah ada atau belum
				mockRepo.EXPECT().IsEmailExists(gomock.Any()).Return(false, nil)

				// create user
				mockRepo.EXPECT().CreateUser(gomock.Any()).Return(nil)
			},
			&model.UserRes{
				ID:    0,
				Name:  "Test name",
				Email: "email@example.com",
			},
			false,
		},
		// {
		// 	"Test create user error",
		// 	fields{
		// 		repo: mockRepo,
		// 	},
		// 	args{
		// 		req: model.CreateUserReq{
		// 			Name:     "Test name",
		// 			Email:    "email@example.com",
		// 			Password: "Test password",
		// 		},
		// 	},
		// 	func() {
		// 		mockRepo.EXPECT().Create(gomock.Any()).Return(entity.Booking{
		// 			ID: 1,
		// 		}, errors.New("ERROR"))
		// 	},
		// 	entity.Booking{
		// 		ID: 1,
		// 	},
		// 	true,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			usecase := &userUseCase{
				userRepository: mockRepo,
			}
			gotData, err := usecase.CreateUser(&tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("CreateUser() = %+v, want %+v", gotData, tt.wantData)
			}
		})
	}
}
