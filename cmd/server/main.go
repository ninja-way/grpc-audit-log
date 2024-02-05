package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"log"
	"time"

	"github.com/ninja-way/grpc-audit-log/internal/config"
	"github.com/ninja-way/grpc-audit-log/internal/repository"
	"github.com/ninja-way/grpc-audit-log/internal/server"
	"github.com/ninja-way/grpc-audit-log/internal/service"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.Password, cfg.DB.DBName, cfg.DB.SSLMode)

	db, err := pgx.Connect(ctx, psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(ctx)

	auditRepo := repository.NewAudit(db)
	auditService := service.NewAudit(auditRepo)
	auditSrv := server.NewAuditServer(auditService)
	srv := server.New(auditSrv)

	//if auditRepo.Insert(ctx, audit.LogItem{
	//	Entity:    "test",
	//	Action:    "test",
	//	EntityID:  1,
	//	Timestamp: time.Now(),
	//}); err != nil {
	//	log.Fatal(err)
	//}

	fmt.Println("SERVER STARTED", time.Now())

	if err := srv.ListenAndServe(cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
