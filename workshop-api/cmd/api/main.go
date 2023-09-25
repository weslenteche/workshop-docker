package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/weslenteche/workshop-api/config"
	"github.com/weslenteche/workshop-api/internal"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	envs := config.LoadEnvs()
	srv := internal.NewServer()
	srv.Init()
	r := gin.Default()
	internal.Handlers(r, srv)

	httpServer := &http.Server{
		Addr:    ":" + envs.APIPort,
		Handler: r,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Recebido sinal de término. Encerrando...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		fmt.Printf("Erro ao encerrar o servidor: %v\n", err)
	}

	fmt.Println("Servidor encerrado com sucesso.")
}
