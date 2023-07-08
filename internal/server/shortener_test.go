package server

import (
	"context"
	mockRepo "github.com/hbashift/url-shortener/internal/domain/repository/mock"
	"github.com/hbashift/url-shortener/internal/domain/repository/model"
	"github.com/hbashift/url-shortener/internal/service"
	pb "github.com/hbashift/url-shortener/pb"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func Test_shortenerServer_GetUrl(t *testing.T) {
	type mockBehavior func(s *mockRepo.MockRepository, url uint64)

	tests := []struct {
		name          string
		inputUrl      string
		urlID         uint64
		mockBehavior  mockBehavior
		expected      string
		expectedError error
	}{
		{
			name:     "test 1",
			inputUrl: "aaaaaaaaab",
			urlID:    1,
			mockBehavior: func(s *mockRepo.MockRepository, url uint64) {
				s.EXPECT().GetUrl(&model.Url{ShortUrl: "aaaaaaaaab"}).Return("http://localhost:8081/test1", nil)
			},
			expected:      "http://localhost:8081/test1",
			expectedError: nil,
		},
		{
			name:     "test 2: length must be 10 not less",
			inputUrl: "jjj",
			urlID:    uint64(0),
			mockBehavior: func(s *mockRepo.MockRepository, url uint64) {
				s.EXPECT().GetUrl(&model.Url{ShortUrl: "jjj"}).Times(0)
			},
			expected:      "",
			expectedError: status.Errorf(codes.InvalidArgument, "short url length must be 10"),
		},
		{
			name:     "test 3: length must be 10 not more",
			inputUrl: "jjjjjjjjjjjjjj",
			urlID:    uint64(0),
			mockBehavior: func(s *mockRepo.MockRepository, url uint64) {
				s.EXPECT().GetUrl(&model.Url{ShortUrl: "jjjjjjjjjjjjjj"}).Times(0)
			},
			expected:      "",
			expectedError: status.Errorf(codes.InvalidArgument, "short url length must be 10"),
		},
		{
			name:     "test 4",
			inputUrl: "aaaaaaaaac",
			urlID:    2,
			mockBehavior: func(s *mockRepo.MockRepository, url uint64) {
				s.EXPECT().GetUrl(&model.Url{ShortUrl: "aaaaaaaaac"}).Return("http://localhost:8081/test1", nil)
			},
			expected:      "http://localhost:8081/test1",
			expectedError: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)

			repo := mockRepo.NewMockRepository(c)
			s := service.NewShortenerService(repo)
			test.mockBehavior(repo, test.urlID)

			// Init Endpoint
			serv := NewShortenerServer(*s)

			// Make Request
			grpcServ := grpc.NewServer()
			pb.RegisterShortenerServer(grpcServ, serv)

			res, err := serv.GetUrl(context.Background(), &pb.ShortUrl{ShortUrl: test.inputUrl})
			assert.Equal(t, err, test.expectedError)
			assert.Equal(t, test.expected, res.GetLongUrl())
		})
	}
}

func Test_shortenerServer_PostUrl(t *testing.T) {
	type mockBehavior func(s *mockRepo.MockRepository, url string)
	tests := []struct {
		name          string
		inputUrl      string
		expected      string
		mockBehavior  mockBehavior
		expectedError error
	}{
		{
			name:     "test 1",
			inputUrl: "http://localhost:8081/test2",
			expected: "aaaaaaaaab",
			mockBehavior: func(s *mockRepo.MockRepository, url string) {
				s.
					EXPECT().
					PostUrl(gomock.Any()).
					Return("sdlkfsas", nil).Times(1)
			},
			expectedError: nil,
		},
		{
			name:     "test 2",
			inputUrl: "",
			expected: "something",
			mockBehavior: func(s *mockRepo.MockRepository, url string) {
				s.EXPECT().PostUrl(gomock.Any()).Return("", nil).Times(0)
			},
			expectedError: status.Errorf(codes.InvalidArgument, "url length must > 0"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)

			repo := mockRepo.NewMockRepository(c)
			s := service.NewShortenerService(repo)
			test.mockBehavior(repo, test.inputUrl)

			// Init Endpoint
			serv := NewShortenerServer(*s)

			// Make Request
			grpcServ := grpc.NewServer()
			pb.RegisterShortenerServer(grpcServ, serv)

			res, err := serv.PostUrl(context.Background(), &pb.LongUrl{LongUrl: test.inputUrl})
			assert.Equal(t, err, test.expectedError)
			assert.NotEqual(t, test.expected, res.GetShortUrl())
		})
	}
}
