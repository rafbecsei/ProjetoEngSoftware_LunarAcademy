package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	gocqlastra "github.com/datastax/gocql-astra"
	"github.com/gocql/gocql"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Função para testar conexão com o Supabase (PostgreSQL)
func testSupabase() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Erro ao carregar arquivo .env")
	}

	USUARIO := os.Getenv("USUARIO")
	PASSWORD := os.Getenv("PASSWORD")
	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")
	DBNAME := os.Getenv("DBNAME")


	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", USUARIO, PASSWORD, HOST, PORT, DBNAME)

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Fatalf("Falha ao conectar ao Supabase: %v", err)
	}
	defer conn.Close(context.Background())

	fmt.Println("Conectado ao SupaBase com sucesso!")
}

func testMongo() {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Println("MONGO_URI não configurada, pulando conexão MongoDB")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Falha ao conectar ao MongoDB: %v", err)
	}

	// Testa a conexão
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Falha ao testar conexão com MongoDB: %v", err)
	}

	fmt.Println("Conectado ao MongoDB com sucesso!")
	// Lembre-se de desconectar no final
	defer client.Disconnect(ctx)
}

// Função para testar conexão com o Cassandra (Astra DB)
func testCassandra() {
	if os.Getenv("ASTRA_DB_ID") == "" || os.Getenv("APPLICATION_TOKEN") == "" {
		log.Println("Variáveis de ambiente do Astra DB não configuradas, pulando conexão Cassandra.")
		return
	}

	cluster, err := gocqlastra.NewClusterFromURL(
		"https://api.astra.datastax.com",
		os.Getenv("ASTRA_DB_ID"),
		os.Getenv("APPLICATION_TOKEN"),
		10*time.Second,
	)
	if err != nil {
		log.Fatalf("Erro ao carregar cluster Astra: %v", err)
	}

	cluster.Timeout = 30 * time.Second
	session, err := gocql.NewSession(*cluster)
	if err != nil {
		log.Fatalf("Erro ao conectar ao Cassandra: %v", err)
	}
	defer session.Close()

	fmt.Println("Conectado ao Cassandra com sucesso!")

	// var version string
	// iter := session.Query("SELECT release_version FROM system.local").Iter()
	// for iter.Scan(&version) {
	// 	fmt.Println("Versão Cassandra:", version)
	// }
	// if err = iter.Close(); err != nil {
	// 	log.Printf("Erro ao rodar query Cassandra: %v", err)
	// }
}

func main() {
	// Carrega as variáveis do .env
	godotenv.Load()

	fmt.Println("Iniciando conexões com os bancos...")

	testSupabase()
	testMongo()
	testCassandra()

	fmt.Println("Conexões encerradas")
}
