package api

import (
	"github.com/hbashift/url-shortener/internal/domain/repository/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/hbashift/url-shortener/pb"
)

// HttpClient - interface for gRPC gateway
type HttpClient interface {
	GetShortUrl(ctx *gin.Context)
	PostLongUrl(ctx *gin.Context)
}

// NewHttpClient - creates new HttpClient
func NewHttpClient(client pb.ShortenerClient) HttpClient {
	return &httpClient{client: client}
}

// RunHttpClient - runs gRPC gateway proxy with gin.Engine
func RunHttpClient(addr, port string) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to grpc server: %v", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	grpcClient := pb.NewShortenerClient(conn)
	client := NewHttpClient(grpcClient)

	r := gin.Default()
	r.GET("/v1/url/:shortUrl", client.GetShortUrl)
	r.POST("/v1/url", client.PostLongUrl)

	err = r.Run(port)
	if err != nil {
		log.Fatalf("could not run http client: %v", err)
	}
}

// httpClient - HttpClient implementation
type httpClient struct {
	client pb.ShortenerClient
}

// GetShortUrl - handler for proto.ShortenerClient GetUrl procedure
func (h *httpClient) GetShortUrl(ctx *gin.Context) {
	shortUrl := ctx.Param("shortUrl")
	url := pb.ShortUrl{ShortUrl: shortUrl}

	res, err := h.client.GetUrl(ctx, &url)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &model.LongUrl{LongUrl: res.LongUrl})
}

// PostLongUrl - handler for proto.ShortenerClient PostUrl procedure
func (h *httpClient) PostLongUrl(ctx *gin.Context) {
	var longUrl *model.LongUrl

	err := ctx.ShouldBind(&longUrl)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	url := &pb.LongUrl{LongUrl: longUrl.LongUrl}
	res, err := h.client.PostUrl(ctx, url)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, &model.ShortUrl{ShortUrl: res.GetShortUrl()})
}
