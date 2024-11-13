env "development" {
    url = "postgres://${env.DB_USER}:${env.DB_PASSWORD}@${env.DB_HOST}:${env.DB_PORT}/${env.DB_NAME}?sslmode=disable"

    migrations {
        dir = "file://internal/database/migrations"
        format = "atlas"
    }
}