package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"

	httpHandler "github.com/rmndvngrpslhr/app-ticket/go-client/internal/events/infra/http"
	"github.com/rmndvngrpslhr/app-ticket/go-client/internal/events/infra/repository"
	"github.com/rmndvngrpslhr/app-ticket/go-client/internal/events/infra/service"
	"github.com/rmndvngrpslhr/app-ticket/go-client/internal/events/usecase"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Events API
// @version 1.0
// @description This is a server for managing events
// @host localhost:8080
// @BasePath /
func main() {
	// DataBase configs
	db, err := sql.Open("mysql", "test_user:test_password@tcp(golang-mysql:3306)/test_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Repository
	eventRepo, err := repository.NewMysqlEventRepository(db)
	if err != nil {
		log.Fatal(err)
	}

	// Specific base URLs to each partner
	partnerBaseURls := map[int]string{
		1: "http://host.docker.internal:8000/partner1",
		2: "http://host.docker.internal:8000/partner2",
	}
	partnerFactory := service.NewPartnerFactory(partnerBaseURls)

	listEventsUseCase := usecase.NewListEventsUseCase(eventRepo)
	listSpotsUseCase := usecase.NewListSpotsUseCase(eventRepo)
	getEventUseCase := usecase.NewGetEventUseCase(eventRepo)
	createEventUseCase := usecase.NewCreateEventUseCase(eventRepo)
	createSpotsUseCase := usecase.NewCreateSpotsUseCase(eventRepo)
	buyTickersUseCase := usecase.NewBuyTicketsUseCase(eventRepo, partnerFactory)

	// HTTP Handlers
	eventsHandler := httpHandler.NewEventsHandler(
		buyTickersUseCase,
		createEventUseCase,
		createSpotsUseCase,
		getEventUseCase,
		listEventsUseCase,
		listSpotsUseCase,
	)

	r := http.NewServeMux()
	r.HandleFunc("GET /docs", httpSwagger.WrapHandler)
	r.HandleFunc("GET /events", eventsHandler.ListEvents)
	r.HandleFunc("GET /events/{eventID}", eventsHandler.GetEvents)
	r.HandleFunc("GET /events/{eventID}/spots", eventsHandler.ListSpots)
	r.HandleFunc("POST /events", eventsHandler.CreateEvent)
	r.HandleFunc("POST /checkout", eventsHandler.BuyTickets)
	r.HandleFunc("POST /events/{eventID}/spots", eventsHandler.CreateSpots)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Channel to listen OS signals
	idleConnsClosed := make(chan struct{})
	go func() {
		sigInt := make(chan os.Signal, 1)
		signal.Notify(sigInt, syscall.SIGINT, syscall.SIGTERM)
		<-sigInt

		// After interruption signal is given, initialize graceful shutdown
		log.Println("Recebido o sinal de interrupção, inicializando o 'graceful shutdown'...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Erro ao executar 'graceful shutdown: %v\n", err)
		}
		close((idleConnsClosed))
	}()

	// Initializing HTTP server
	log.Printf("\nservidor HTTP rodando na porta %s", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("erro ao iniciar servidor HTTP: %v\n", err)
	}

	<-idleConnsClosed
	log.Println("Servidor HTTP finalizado")
}
