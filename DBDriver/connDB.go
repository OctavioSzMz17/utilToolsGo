package DBDriver

import (
    "fmt"
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/lib/pq"
)

// DB es la variable global de conexión
var DBConn *sql.DB
var DBType string
// Inicializa la conexión a la base de datos MySQL
func InitDBMYSQL(GBD, user, pass, host, port, dbname string) {
	var err error
	DSN := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbname
	DBConn, err = sql.Open(GBD, DSN)
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}
	DBType = "mysql"

	err = DBConn.Ping()
	if err != nil {
		log.Fatal("No se puede establecer la conexión con la base de datos:", err)
	}
}


// Inicializa la conexión a la base de datos externa (Supabase)
func InitDBSupabase(connStr string) {
    var err error
    DBConn, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Error al conectar con la base de datos:", err)
    }

    DBType = "postgres"

    // Verificar si la conexión es exitosa
    err = DBConn.Ping()
    if err != nil {
        log.Fatal("No se puede establecer la conexión con la base de datos:", err)
    }
}

// Retorna el marcador de posición correcto
func getPlaceholder(i int) string {
	if DBType == "postgres" {
		return fmt.Sprintf("$%d", i)
	}
	return "?"
}
