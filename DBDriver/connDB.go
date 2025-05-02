package DBDriver

import (
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/lib/pq"
)

// DB es la variable global de conexión
var DB *sql.DB

// Inicializa la conexión a la base de datos MySQL
func InitDBMYSQL(GBD, user, pass, host, port, dbname string) {
    var err error
    DSN := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbname
    DB, err = sql.Open(GBD, DSN)
    if err != nil {
        log.Fatal("Error al conectar con la base de datos:", err)
    }

    // Verificar si la conexión es exitosa
    err = DB.Ping()
    if err != nil {
        log.Fatal("No se puede establecer la conexión con la base de datos:", err)
    }
}

// Inicializa la conexión a la base de datos externa (Supabase)
func InitDBSupabase(connStr string) {
    var err error
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Error al conectar con la base de datos:", err)
    }

    // Verificar si la conexión es exitosa
    err = DB.Ping()
    if err != nil {
        log.Fatal("No se puede establecer la conexión con la base de datos:", err)
    }
}
