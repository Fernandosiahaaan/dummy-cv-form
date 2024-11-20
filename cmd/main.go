package main

import (
	"context"
	"dummy-cv-form/infrastructure/redis"
	"dummy-cv-form/internal/handler"
	"dummy-cv-form/internal/repository"
	"dummy-cv-form/internal/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("=== APP START ===")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	repo, err := repository.NewRepository(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer repo.Close()
	fmt.Println("🔥 Init Repository...")

	redis, err := redis.NewReddisClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer redis.Close()
	fmt.Println("🔥 Init Redis...")

	svc := service.NewService(service.ServiceParam{
		Ctx:   ctx,
		Redis: redis,
		Repo:  repo,
	})
	defer svc.Close()
	fmt.Println("🔥 Init Service...")

	handler := handler.NewHandler(&handler.ParamHandler{
		Ctx:     ctx,
		Redis:   redis,
		Repo:    repo,
		Service: svc,
	})
	defer handler.Close()
	fmt.Println("🔥 Init Handler...")

	portHttp := os.Getenv("PORT_HTTP")
	localHost := fmt.Sprintf("0.0.0.0:%s", portHttp)
	routing := router(handler)

	fmt.Printf("🌐 %s\n", localHost)
	err = http.ListenAndServe(localHost, routing)
	if err != nil {
		log.Fatalf("Could not listen and serve %s. err = %v", localHost, err)
	}
}
