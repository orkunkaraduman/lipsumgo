package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	promgrpc "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"lipsumgo/pkg/pb"
	"lipsumgo/pkg/server"
)

var (
	appName    string
	appVersion string
	appBuild   string
)

func main() {
	var httpListenAddr string
	var grpcListenAddr string
	flag.StringVar(&httpListenAddr, "http-listen", ":8080", "http listen address")
	flag.StringVar(&grpcListenAddr, "grpc-listen", ":9000", "grpc listen address")
	flag.Parse()

	log.SetOutput(os.Stdout)

	log.Printf("starting %s version %s build %s", appName, appVersion, appBuild)

	httpLis, err := net.Listen("tcp", httpListenAddr)
	if err != nil {
		log.Fatalf("http listen error: %v", err)
		return
	}
	//goland:noinspection GoUnhandledErrorResult
	defer httpLis.Close()
	log.Printf("listening http %s", httpListenAddr)

	grpcLis, err := net.Listen("tcp", grpcListenAddr)
	if err != nil {
		log.Fatalf("grpc listen error: %v", err)
		return
	}
	//goland:noinspection GoUnhandledErrorResult
	defer grpcLis.Close()
	log.Printf("listening grpc %s", grpcListenAddr)

	gwMux := runtime.NewServeMux()
	gwConn, _ := grpc.Dial(grpcListenAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//goland:noinspection GoUnhandledErrorResult
	defer gwConn.Close()

	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/debug/pprof/", pprof.Index)
	httpMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	httpMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	httpMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	httpMux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	httpMux.Handle("/metrics/", promhttp.Handler())
	httpMux.Handle("/api/", gwMux)
	httpHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if p := recover(); p != nil {
				log.Printf("panic in http handler %q: %v", r.RequestURI, p)
				return
			}
		}()
		httpMux.ServeHTTP(w, r)
	})
	httpSrv := &http.Server{
		Handler:           httpHandler,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       65 * time.Second,
		MaxHeaderBytes:    1 << 20,
		ErrorLog:          log.New(ioutil.Discard, "", log.LstdFlags),
	}

	grpcUnaryInterceptor := func(handlerCtx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (reply interface{}, err error) {
		defer func() {
			if p := recover(); p != nil {
				log.Printf("panic in grpc unary rpc %q: %v", info.FullMethod, p)
				return
			}
		}()
		reply, err = handler(handlerCtx, req)
		if err != nil {
			if s, ok := status.FromError(err); ok {
				log.Printf("grpc unary rpc error: method=%q code=%d desc=%q", info.FullMethod, s.Code(), s.Message())
			} else {
				log.Printf("grpc unary rpc custom error in %q: %v", info.FullMethod, err)
			}
		}
		return
	}
	grpcStreamInterceptor := func(srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler) (err error) {
		defer func() {
			if p := recover(); p != nil {
				log.Printf("panic in grpc stream rpc %q: %v", info.FullMethod, p)
				return
			}
		}()
		err = handler(srv, stream)
		if err != nil {
			if s, ok := status.FromError(err); ok {
				log.Printf("grpc stream rpc error: method=%q code=%d desc=%q", info.FullMethod, s.Code(), s.Message())
			} else {
				log.Printf("grpc stream rpc custom error in %q: %v", info.FullMethod, err)
			}
		}
		return
	}

	grpcSrv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(promgrpc.UnaryServerInterceptor, grpcUnaryInterceptor),
		grpc.ChainStreamInterceptor(promgrpc.StreamServerInterceptor, grpcStreamInterceptor))
	pb.RegisterApiServer(grpcSrv, &server.Api{})
	_ = pb.RegisterApiHandler(context.Background(), gwMux, gwConn)

	log.Printf("running %s", appName)

	go func() {
		if e := httpSrv.Serve(httpLis); e != nil && e != http.ErrServerClosed {
			log.Printf("http serve error: %v", e)
			return
		}
	}()

	go func() {
		if e := grpcSrv.Serve(grpcLis); e != nil && e != grpc.ErrServerStopped {
			log.Printf("grpc serve error: %v", e)
			return
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Kill, os.Interrupt, syscall.SIGTERM)
	<-ch

	var wg sync.WaitGroup

	termCtx, termCtxCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer termCtxCancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if e := httpSrv.Shutdown(termCtx); e != nil {
			log.Printf("http shutdown error: %v", e)
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		gracefulStoppedCh := make(chan struct{})
		go func() {
			grpcSrv.GracefulStop()
			close(gracefulStoppedCh)
		}()
		select {
		case <-termCtx.Done():
			grpcSrv.Stop()
			log.Printf("grpc shutdown error: %v", termCtx.Err())
			return
		case <-gracefulStoppedCh:
		}
	}()

	wg.Wait()

	log.Printf("stopped %s", appName)
}
