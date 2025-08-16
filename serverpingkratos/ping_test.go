package serverpingkratos_test

import (
	"context"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/uuid"
	"github.com/orzkratos/pingkratos/clientpingkratos"
	"github.com/orzkratos/pingkratos/serverpingkratos"
	"github.com/orzkratos/zapkratos"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/must"
	"github.com/yyle88/rese"
	"github.com/yyle88/zaplog"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	httpPort string = "28000"
	grpcPort string = "28001"
)

func TestMain(m *testing.M) {
	// Create logger to show ping request logs
	// 创建 logger 以显示 ping 请求日志
	zapKratos := zapkratos.NewZapKratos(zaplog.LOGGER, zapkratos.NewOptions())

	// Create HTTP server with dynamic port
	// 使用动态端口创建 HTTP 服务器
	httpSrv := http.NewServer(
		http.Address(":"+httpPort),
		http.Middleware(
			recovery.Recovery(),
			logging.Server(zapKratos.GetLogger("HTTP")),
		),
	)

	// Create gRPC server with dynamic port
	// 使用动态端口创建 gRPC 服务器
	grpcSrv := grpc.NewServer(
		grpc.Address(":"+grpcPort),
		grpc.Middleware(
			recovery.Recovery(),
			logging.Server(zapKratos.GetLogger("GRPC")),
		),
	)

	// Create ping service
	// 创建 ping 服务
	pingService := serverpingkratos.NewPingService(zapKratos.GetLogger("PING"))
	clientpingkratos.RegisterPingHTTPServer(httpSrv, pingService)
	clientpingkratos.RegisterPingServer(grpcSrv, pingService)

	app := kratos.New(
		kratos.Name("test-ping-kratos"),
		kratos.Server(httpSrv, grpcSrv),
	)

	// Start server in background
	// 后台启动服务器
	go func() {
		must.Done(app.Run())
	}()
	defer rese.F0(app.Stop)

	// Wait for server to start
	// 等待服务器启动
	time.Sleep(time.Millisecond * 200)

	m.Run()
}

func TestPingService_Ping_HTTP(t *testing.T) {
	// Create HTTP client
	// 创建 HTTP 客户端
	conn := rese.P1(http.NewClient(
		context.Background(),
		http.WithMiddleware(recovery.Recovery()),
		http.WithEndpoint("127.0.0.1:"+httpPort),
	))
	defer rese.F0(conn.Close)

	// Test ping service via HTTP
	// 通过 HTTP 测试 ping 服务
	pingClient := clientpingkratos.NewPingHTTPClient(conn)

	// Generate random test message for unique testing
	// 生成随机测试消息确保测试唯一性
	testMessage := uuid.New().String()

	resp, err := pingClient.Ping(context.Background(), wrapperspb.String(testMessage))
	require.NoError(t, err)
	require.Equal(t, testMessage, resp.GetValue())
}

func TestPingService_Ping_gRPC(t *testing.T) {
	// Create gRPC client
	// 创建 gRPC 客户端
	conn := rese.P1(grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:"+grpcPort),
		grpc.WithMiddleware(recovery.Recovery()),
	))
	defer rese.F0(conn.Close)

	// Test ping service via gRPC
	// 通过 gRPC 测试 ping 服务
	pingClient := clientpingkratos.NewPingClient(conn)

	testMessage := uuid.New().String()

	resp, err := pingClient.Ping(context.Background(), wrapperspb.String(testMessage))
	require.NoError(t, err)
	require.Equal(t, testMessage, resp.GetValue())
}

// NOTE: Add test for empty message handling
// NOTE: 暂不测试空字符串的
