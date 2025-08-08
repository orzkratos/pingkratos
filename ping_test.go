package pingkratos_test

import (
	"context"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/orzkratos/pingkratos"
	"github.com/orzkratos/pingkratos/pingpong"
	"github.com/orzkratos/zapkratos"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/must"
	"github.com/yyle88/rese"
	"github.com/yyle88/zaplog"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestMain(m *testing.M) {
	zapKratos := zapkratos.NewZapKratos(zaplog.LOGGER, zapkratos.NewOptions())
	slog := zapKratos.SubZap()
	slog.LOG.Debug("starting ping-kratos tests")

	// cp from https://github.com/go-kratos/examples/blob/61daed1ec4d5a94d689bc8fab9bc960c6af73ead/helloworld/server/main.go#L42
	httpSrv := http.NewServer(
		http.Address(":28000"),
		http.Middleware(
			recovery.Recovery(),
			logging.Server(zapKratos.GetLogger("HTTP")),
		),
	)

	// cp from https://github.com/go-kratos/examples/blob/61daed1ec4d5a94d689bc8fab9bc960c6af73ead/helloworld/server/main.go#L48
	grpcSrv := grpc.NewServer(
		grpc.Address(":28001"),
		grpc.Middleware(
			recovery.Recovery(),
			logging.Server(zapKratos.GetLogger("GRPC")),
		),
	)

	pingService := pingkratos.NewPingService(zapKratos.GetLogger("PING"))
	pingpong.RegisterPingHTTPServer(httpSrv, pingService)
	pingpong.RegisterPingServer(grpcSrv, pingService)

	app := kratos.New(
		kratos.Name("demo-ping-kratos"),
		kratos.Server(
			httpSrv,
			grpcSrv,
		),
	)

	slog.LOG.Debug("starting ping-kratos tests")
	go func() {
		must.Done(app.Run())
	}()
	defer rese.F0(app.Stop)

	time.Sleep(time.Millisecond * 100)
	m.Run()
	slog.LOG.Debug("complete ping-kratos tests")
}

func TestPingService_Ping_Http(t *testing.T) {
	// cp from https://github.com/go-kratos/examples/blob/61daed1ec4d5a94d689bc8fab9bc960c6af73ead/helloworld/client/main.go#L20
	conn, err := http.NewClient(
		context.Background(),
		http.WithMiddleware(
			recovery.Recovery(),
		),
		http.WithEndpoint("127.0.0.1:28000"),
	)
	require.NoError(t, err)
	defer rese.F0(conn.Close)

	pingClient := pingpong.NewPingHTTPClient(conn)
	const msg = "message"
	resp, err := pingClient.Ping(context.Background(), wrapperspb.String(msg))
	require.NoError(t, err)
	require.Equal(t, msg, resp.GetValue())
}

func TestPingService_Ping_Grpc(t *testing.T) {
	// cp from https://github.com/go-kratos/examples/blob/61daed1ec4d5a94d689bc8fab9bc960c6af73ead/helloworld/client/main.go#L49
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:28001"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	require.NoError(t, err)
	defer rese.F0(conn.Close)

	pingClient := pingpong.NewPingClient(conn)
	const msg = "message"
	resp, err := pingClient.Ping(context.Background(), wrapperspb.String(msg))
	require.NoError(t, err)
	require.Equal(t, msg, resp.GetValue())
}
