package main

import (
    "context"
    "log"
    "net/http"

    "github.com/grpc-custom/graphql-gateway/runtime"
    "google.golang.org/grpc"

    "app/book" // protocで生成されたファイルをimport
    "app/user/pb"
)

func main() {
    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    mux, err := runtime.NewServeMux()
    if err != nil {
        log.Fatal(err)
    }
    opts := []grpc.DialOption{
        grpc.WithInsecure(),
    }
    // gRPCのUserServiceにアクセスするための定義を追加
    err = user.RegisterUserServiceFromEndpoint(ctx, mux, "localhost:9001", opts)
    if err != nil {
        log.Fatal(err)
    }
    // gRPCのBookServiceにアクセスするための定義を追加
    err = book.RegisterBookServiceFromEndpoint(ctx, mux, "localhost:9002", opts)
    if err != nil {
        log.Fatal(err)
    }
    err = http.ListenAndServe(":8080", mux)
    if err != nil {
        log.Fatal(err)
    }
}