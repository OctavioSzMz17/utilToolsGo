package DBDriver

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// Estructura de usuario
type User struct {
	ID       int
	Username string
	Password string
}

// Crear un nuevo usuario con contraseña encriptada
/*func CreateUser(table, username, password string) error {
	// Encriptar la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error encriptando contraseña: %v", err)
	}

	query := fmt.Sprintf("INSERT INTO %s (username, password) VALUES (?, ?)", table)
	_, err = DB.Exec(query, username, hashedPassword)
	if err != nil {
		return fmt.Errorf("error creando usuario: %v", err)
	}
	return nil
}*/

	func CreateUser(table []string, username, password string) error {
		// Encriptar la contraseña
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("error encriptando contraseña: %v", err)
		}

		query := fmt.Sprintf("INSERT INTO %s (%s, %s) VALUES (?, ?)", table[0], table[1], table[2])
		_, err = DB.Exec(query, username, hashedPassword)
		if err != nil {
			return fmt.Errorf("error creando usuario: %v", err)
		}
		return nil
	}











// Obtener usuario por nombre de usuario
func GetUserByUsername(table, username string) (User, error) {
	var user User
	query := fmt.Sprintf("SELECT id, username, password FROM %s WHERE username = ?", table)
	row := DB.QueryRow(query, username)

	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("usuario no encontrado")
		}
		return user, fmt.Errorf("error obteniendo usuario: %v", err)
	}
	return user, nil
}

// Verificar si las credenciales del usuario son correctas
func AuthenticateUser(table, username, password string) (bool, error) {
	// Obtener el usuario
	user, err := GetUserByUsername(table, username)
	if err != nil {
		return false, err
	}

	// Comparar la contraseña ingresada con la almacenada en la base de datos
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, fmt.Errorf("contraseña incorrecta")
		}
		return false, fmt.Errorf("error al autenticar usuario: %v", err)
	}

	return true, nil
}

// Obtener usuario por ID
func GetUserByID(table string, id int) (User, error) {
	var user User
	query := fmt.Sprintf("SELECT id, username, password FROM %s WHERE id = ?", table)
	row := DB.QueryRow(query, id)

	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("usuario no encontrado")
		}
		return user, fmt.Errorf("error obteniendo usuario: %v", err)
	}
	return user, nil
}

// Actualizar información de usuario
func UpdateUser(table string, id int, username, password string) error {
	// Encriptar la nueva contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error encriptando contraseña: %v", err)
	}

	query := fmt.Sprintf("UPDATE %s SET username = ?, password = ? WHERE id = ?", table)
	_, err = DB.Exec(query, username, hashedPassword, id)
	if err != nil {
		return fmt.Errorf("error actualizando usuario: %v", err)
	}
	return nil
}

// Eliminar usuario
func DeleteUser(table string, id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", table)
	_, err := DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error eliminando usuario: %v", err)
	}
	return nil
}

// Obtener todos los usuarios (solo para administración)
func GetAllUsers(table string) ([]User, error) {
	var users []User
	query := fmt.Sprintf("SELECT id, username, password FROM %s", table)
	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo usuarios: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			return nil, fmt.Errorf("error leyendo los datos del usuario: %v", err)
		}
		users = append(users, user)
	}

	return users, nil
}

// Actualizar solo la contraseña de un usuario
func UpdatePassword(table string, id int, newPassword string) error {
	// Encriptar la nueva contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error encriptando nueva contraseña: %v", err)
	}

	query := fmt.Sprintf("UPDATE %s SET password = ? WHERE id = ?", table)
	_, err = DB.Exec(query, hashedPassword, id)
	if err != nil {
		return fmt.Errorf("error actualizando contraseña: %v", err)
	}
	return nil
}
