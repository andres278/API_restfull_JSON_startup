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
	

	// GET a reserva detail
	router.GET("/reserva/:id", func(c *gin.Context) {
		var (
			reserva  Reserva
			reservas []Reserva
		)
		id := c.Param("id")
		rows, err := db.Query("SELECT *FROM reservas r INNER JOIN cliente c ON (r.cliente_idCliente = c.idCliente) WHERE r.idReservas = ?;", id)
	
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&reserva.ID, &reserva.FechaCreacion, &reserva.FechaReserva, &reserva.Hora, &reserva.NumeroPersonas, &reserva.Cli.IDCliente, &reserva.Rest.ID, &reserva.Cli.IDCliente, &reserva.Cli.Nombre,&reserva.Cli.Apellido,&reserva.Cli.Pass)

			reservas = append(reservas, reserva)

			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": reservas,
			"count":  len(reservas),
		})
	})


// POST new reserva details
	router.POST("/reserva", func(c *gin.Context) {
		var buffer bytes.Buffer
		
		fechaCreacion := c.PostForm("fechaCreacion")
		fechaReserva := c.PostForm("fechaReserva")
		hora := c.PostForm("hora")
		numeroPersonas := c.PostForm("numeroPersonas")
		clienteIDCliente := c.PostForm("cliente_idCliente")
		restauranteid := c.PostForm("restaurante_id")
		
		stmt, err := db.Prepare("insert into reservas ( fechaCreacion, fechaReserva, hora, numeroPersonas, cliente_idCliente, restaurante_id) values(?,?,?,?,?,?);")
	
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(fechaCreacion, fechaReserva, hora, numeroPersonas, clienteIDCliente, restauranteid)

		if err != nil {
			fmt.Print(err.Error())
		}

		// Fastest way to append strings
		buffer.WriteString(fechaReserva)
		buffer.WriteString(" ")
		buffer.WriteString(hora)
		buffer.WriteString(" ")
		buffer.WriteString(numeroPersonas)
		defer stmt.Close()
		dato := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf(" %s successfully created", dato),
		})
	})


  // GET a reservas 
 	router.GET("/reservas", func(c *gin.Context) {
  		var (
   			reserva  Reserva
   			reservas []Reserva
  		)
  		rows, err := db.Query("SELECT *FROM reservas r INNER JOIN cliente c ON (r.cliente_idCliente = c.idCliente)")
 
  		if err != nil {
   			fmt.Print(err.Error())
  		}
  		for rows.Next() {
   			err = rows.Scan(&reserva.ID, &reserva.FechaCreacion, &reserva.FechaReserva, &reserva.Hora, &reserva.NumeroPersonas, &reserva.Cli.IDCliente, &reserva.Rest.ID, &reserva.Cli.IDCliente, &reserva.Cli.Nombre,&reserva.Cli.Apellido,&reserva.Cli.Pass)

   		reservas = append(reservas, reserva)

   			if err != nil {
    			fmt.Print(err.Error())
   			}
  		}
  		defer rows.Close()
  		c.JSON(http.StatusOK, gin.H{
   			"result": reservas,
   			"count":  len(reservas),
  		})
 	})


// GET all clientes
	router.GET("/clientes", func(c *gin.Context) {
		var (
			cliente  Cliente
			clientes []Cliente
		)
		rows, err := db.Query("select  *from cliente c, cli_telefono t, cli_correo e where c.idCliente = t.cliente_idCliente AND e.cliente_idCliente = c.idCliente")
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&cliente.IDCliente, &cliente.Nombre, &cliente.Apellido, &cliente.Pass, &cliente.Tel.Cliente_idCliente, &cliente.Tel.Telefono, &cliente.Cor.Cliente_idCliente, &cliente.Cor.Correo)
			clientes = append(clientes, cliente)

			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": clientes,
			"count":  len(clientes),
		})
	})



	// GET a cliente detail
	router.GET("/cliente/:id", func(c *gin.Context) {
		var (
			cliente  Cliente
			clientes []Cliente
		)
		id := c.Param("id")
		rows, err := db.Query("SELECT *FROM cliente c INNER JOIN cli_telefono t ON (c.idCliente = t.cliente_idCliente) INNER JOIN cli_correo e ON (c.idCliente = e.cliente_idCliente) WHERE c.idCliente = ?;", id)
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&cliente.IDCliente, &cliente.Nombre, &cliente.Apellido, &cliente.Pass, &cliente.Tel.Cliente_idCliente, &cliente.Tel.Telefono, &cliente.Cor.Cliente_idCliente, &cliente.Cor.Correo)
			clientes = append(clientes, cliente)

			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": clientes,
			"count":  len(clientes),
		})
	})


// POST new cliente details
	router.POST("/cliente", func(c *gin.Context) {
		var buffer bytes.Buffer
		nombre := c.PostForm("nombre")
		apellido := c.PostForm("apellido")
		pass := c.PostForm("pass")
		stmt, err := db.Prepare("insert into cliente (nombre, apellido, pass) values(?,?,?);")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(nombre, apellido, pass)

		if err != nil {
			fmt.Print(err.Error())
		}

		// Fastest way to append strings
		buffer.WriteString(nombre)
		buffer.WriteString(" ")
		buffer.WriteString(apellido)
		buffer.WriteString(" ")
		buffer.WriteString(pass)
		defer stmt.Close()
		dato := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf(" %s successfully created", dato),
		})
	})



	router.Run(":3000")
}
