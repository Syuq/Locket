package v2

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	apiv2pb "github.com/Syuq/Locket/proto/gen/api/v2"
	"github.com/Syuq/Locket/server/profile"
	"github.com/Syuq/Locket/store"
)

type APIV2Service struct {
	apiv2pb.UnimplementedWorkspaceServiceServer
	apiv2pb.UnimplementedWorkspaceSettingServiceServer
	apiv2pb.UnimplementedAuthServiceServer
	apiv2pb.UnimplementedUserServiceServer
	apiv2pb.UnimplementedLocketServiceServer
	apiv2pb.UnimplementedResourceServiceServer
	apiv2pb.UnimplementedTagServiceServer
	apiv2pb.UnimplementedInboxServiceServer
	apiv2pb.UnimplementedActivityServiceServer
	apiv2pb.UnimplementedWebhookServiceServer
	apiv2pb.UnimplementedLinkServiceServer

	Secret  string
	Profile *profile.Profile
	Store   *store.Store

	grpcServer     *grpc.Server
	grpcServerPort int
}

func NewAPIV2Service(secret string, profile *profile.Profile, store *store.Store, grpcServerPort int) *APIV2Service {
	grpc.EnableTracing = true
	authProvider := NewGRPCAuthInterceptor(store, secret)
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			NewLoggerInterceptor().LoggerInterceptor,
			authProvider.AuthenticationInterceptor,
		),
	)
	apiv2Service := &APIV2Service{
		Secret:         secret,
		Profile:        profile,
		Store:          store,
		grpcServer:     grpcServer,
		grpcServerPort: grpcServerPort,
	}

	apiv2pb.RegisterWorkspaceServiceServer(grpcServer, apiv2Service)
	apiv2pb.RegisterWorkspaceSettingServiceServer(grpcServer, apiv2Service)
	apiv2pb.RegisterAuthServiceServer(grpcServer, apiv2Service)
	apiv2pb.RegisterUserServiceServer(grpcServer, apiv2Service)
	apiv2pb.RegisterLocketServiceServer(grpcServer, apiv2Service)
	apiv2pb.RegisterTagServiceServer(grpcServer, apiv2Service)
	apiv2pb.RegisterResourceServiceServer(grpcServer, apiv2Service)
	apiv2pb.RegisterInboxServiceServer(grpcServer, apiv2Service)
	apiv2pb.RegisterActivityServiceServer(grpcServer, apiv2Service)
	apiv2pb.RegisterWebhookServiceServer(grpcServer, apiv2Service)
	apiv2pb.RegisterLinkServiceServer(grpcServer, apiv2Service)
	reflection.Register(grpcServer)

	return apiv2Service
}

func (s *APIV2Service) GetGRPCServer() *grpc.Server {
	return s.grpcServer
}

// RegisterGateway registers the gRPC-Gateway with the given Echo instance.
func (s *APIV2Service) RegisterGateway(ctx context.Context, e *echo.Echo) error {
	// Create a client connection to the gRPC Server we just started.
	// This is where the gRPC-Gateway proxies the requests.
	conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf(":%d", s.grpcServerPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}

	gwMux := runtime.NewServeMux()
	if err := apiv2pb.RegisterWorkspaceServiceHandler(context.Background(), gwMux, conn); err != nil {
		return err
	}
	if err := apiv2pb.RegisterWorkspaceSettingServiceHandler(context.Background(), gwMux, conn); err != nil {
		return err
	}
	if err := apiv2pb.RegisterAuthServiceHandler(context.Background(), gwMux, conn); err != nil {
		return err
	}
	if err := apiv2pb.RegisterUserServiceHandler(context.Background(), gwMux, conn); err != nil {
		return err
	}
	if err := apiv2pb.RegisterLocketServiceHandler(context.Background(), gwMux, conn); err != nil {
		return err
	}
	if err := apiv2pb.RegisterTagServiceHandler(context.Background(), gwMux, conn); err != nil {
		return err
	}
	if err := apiv2pb.RegisterResourceServiceHandler(context.Background(), gwMux, conn); err != nil {
		return err
	}
	if err := apiv2pb.RegisterInboxServiceHandler(context.Background(), gwMux, conn); err != nil {
		return err
	}
	if err := apiv2pb.RegisterActivityServiceHandler(context.Background(), gwMux, conn); err != nil {
		return err
	}
	if err := apiv2pb.RegisterWebhookServiceHandler(context.Background(), gwMux, conn); err != nil {
		return err
	}
	if err := apiv2pb.RegisterLinkServiceHandler(context.Background(), gwMux, conn); err != nil {
		return err
	}
	e.Any("/api/v2/*", echo.WrapHandler(gwMux))

	// GRPC web proxy.
	options := []grpcweb.Option{
		grpcweb.WithCorsForRegisteredEndpointsOnly(false),
		grpcweb.WithOriginFunc(func(origin string) bool {
			return true
		}),
	}
	wrappedGrpc := grpcweb.WrapServer(s.grpcServer, options...)
	e.Any("/lockets.api.v2.*", echo.WrapHandler(wrappedGrpc))

	// Start gRPC server.
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Profile.Addr, s.grpcServerPort))
	if err != nil {
		return errors.Wrap(err, "failed to start gRPC server")
	}
	go func() {
		if err := s.grpcServer.Serve(listen); err != nil {
			slog.Error("failed to start gRPC server", err)
		}
	}()

	return nil
}
