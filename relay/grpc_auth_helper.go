package relay

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"
	"strings"

	grpcotel "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc"
	"go.opentelemetry.io/otel/api/trace"
	"google.golang.org/api/idtoken"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	grpcMetadata "google.golang.org/grpc/metadata"
)

type GRPCAuthHelper interface {
	AuthorizeGCPContext(ctx context.Context, addr string) (context.Context, error)
	DialRPCService(ctx context.Context, rpcEndpoint string, tracer trace.Tracer) (context.Context, *grpc.ClientConn, error)
}

type grpcAuthHelper struct {
}

func NewGRPCAuthHelper() GRPCAuthHelper {
	return &grpcAuthHelper{}
}

func (g *grpcAuthHelper) AuthorizeGCPContext(ctx context.Context, addr string) (context.Context, error) {
	splitAddr := strings.Split(addr, ":")
	tokenSource, err := idtoken.NewTokenSource(ctx, fmt.Sprintf("https://%v", splitAddr[0]))
	if err != nil {
		return nil, fmt.Errorf("idtoken.NewTokenSource: %v", err)
	}
	token, err := tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("TokenSource.Token: %v", err)
	}

	ctx = grpcMetadata.NewOutgoingContext(ctx, grpcMetadata.MD{})
	ctx = grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token.AccessToken)
	return ctx, nil
}

func (g *grpcAuthHelper) DialRPCService(ctx context.Context, rpcEndpoint string, tracer trace.Tracer) (context.Context, *grpc.ClientConn, error) {
	ctx, err := AuthorizeGCPContext(ctx, rpcEndpoint)
	if err != nil {
		return ctx, nil, err
	}

	var opts []grpc.DialOption
	if os.Getenv("ENVIRONMENT") == "production" {
		creds := credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	interceptor := grpcotel.UnaryClientInterceptor(tracer)
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))
	conn, err := grpc.DialContext(ctx, rpcEndpoint, opts...)
	if err != nil {
		return ctx, conn, err
	}

	return ctx, conn, nil
}
