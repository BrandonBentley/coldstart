package health

import (
	"context"
	"errors"
	"testing"

	"github.com/BrandonBentley/coldstart/client/health/mock_health"
	mock_healthe "github.com/BrandonBentley/coldstart/entity/health/mock_health"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCheck_Success(t *testing.T) {
	srv, mocks := initMocks(t)

	mocks.HealthClient.EXPECT().Ping(gomock.Any()).Return(nil)
	mocks.HealthEntity.EXPECT().Ping(gomock.Any()).Return(nil)

	ctx := context.Background()

	resp, err := srv.Check(ctx)

	assert.NoError(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, "ok", resp.Status)
	}
}

func TestCheck_HealthEntity_Error(t *testing.T) {
	srv, mocks := initMocks(t)

	expectedError := errors.New("test error")

	mocks.HealthClient.EXPECT().Ping(gomock.Any()).Return(nil)
	mocks.HealthEntity.EXPECT().Ping(gomock.Any()).Return(expectedError)

	ctx := context.Background()

	resp, err := srv.Check(ctx)

	assert.EqualError(t, err, expectedError.Error())
	assert.Nil(t, resp)
}

func TestCheck_HealthClient_Error(t *testing.T) {
	srv, mocks := initMocks(t)

	expectedError := errors.New("test error")

	mocks.HealthClient.EXPECT().Ping(gomock.Any()).Return(expectedError)

	ctx := context.Background()

	resp, err := srv.Check(ctx)

	assert.EqualError(t, err, expectedError.Error())
	assert.Nil(t, resp)
}

type mockComponents struct {
	Ctrl         *gomock.Controller
	HealthClient *mock_health.MockClient
	HealthEntity *mock_healthe.MockEntity
}

func initMocks(t *testing.T) (*Service, mockComponents) {
	ctrl := gomock.NewController(t)

	mocks := mockComponents{
		Ctrl:         ctrl,
		HealthClient: mock_health.NewMockClient(ctrl),
		HealthEntity: mock_healthe.NewMockEntity(ctrl),
	}

	srv := NewService(ServiceParams{})

	srv.healthClient = mocks.HealthClient
	srv.healthEntity = mocks.HealthEntity

	return srv, mocks
}
