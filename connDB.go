
package utilToolsGo

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Inicializa la conexión a la base de datos
func InitDB(GBD, user, pass, host, port, dbname string) {
	// Conexión a la base de datos MySQL
	var err error
	DB, err = sql.Open(GBD, user+":"+pass+"@tcp("+host+":"+port+")/"+dbname)
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	// Verificar si la conexión es exitosa
	err = DB.Ping()
	if err != nil {
		log.Fatal("No se puede establecer la conexión con la base de datos:", err)
	}
}

// Inserta un nuevo usuario en una tabla específica
func InsertUser(table, usr, pass string) error {
	// Usamos el nombre de la tabla proporcionado por el usuario
	_, err := DB.Exec("INSERT INTO "+table+" (usr, pass) VALUES (?, ?)", usr, pass)
	return err
}