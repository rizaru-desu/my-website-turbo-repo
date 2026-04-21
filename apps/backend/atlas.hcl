env "local" {
  src = "ent://internal/infrastructure/persistence/ent/schema"
  dev = "postgres://postgres:postgres@localhost:5432/mydb?sslmode=disable"

  migration {
    dir = "file://internal/infrastructure/persistence/migrations"
  }
}
