package main
//Se debe de exportar las librerias necesarias deben de realizar 
// go get "github.com/go-sql-driver/mysql"
// go get "github.com/gin-gonic/gin"


import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"



"github.com/gin-gonic/gin"
_ "github.com/go-sql-driver/mysql"
	
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/bdaplicaciones")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}

	type Correo struct {
		Cliente_idCliente int
		Correo            string
	}
	type Telefono struct {
		Cliente_idCliente int
		Telefono          string
	}

	type Telefonor struct {
		Restaurante_idRestaurante int
		Telefono                  string
	}

	type Cliente struct {
		IDCliente int
		Nombre    string
		Apellido  string
		Pass      string
		Tel       Telefono
		Cor       Correo
	}

	type Restaurante struct {
		ID            int
		Nit           string
		Razon_social  string
		Contacto      string
		Call_y_numero string
		Barrio        string
		Ciudad        string
		Pais          string
		Tel           Telefonor
	}



	type Reserva struct {
		ID             int
		FechaCreacion  string
		FechaReserva   string
		Hora           string
		NumeroPersonas int
		Cli Cliente
		Rest Restaurante

	}

	router := gin.Default()


//Acá ingresar el contenido de los metodos que daran el servicio API





	router.Run(":3000")
}
