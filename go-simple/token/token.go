package token

//https://www.sohamkamani.com/golang/2019-01-01-jwt-authentication/
import (
	"fmt"
	//...
	// import the jwt-go library

	"time"

	"github.com/dgrijalva/jwt-go"
	//...
)

// Create the JWT key used to create the signature
var jwtKey = []byte("La_camara_de_los_secretos_ha_sido_abierta_Enemigos_del_heredero_temed")

// Se crea el objeto credential que tiene usuario y contraseña
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Se crea el objeto claims que sera codificado en jwt.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Create the Signin handler
func GetToken(user string) (respuesta string, err bool) {
	fmt.Println("Creamos el token")

	// Añadimos el tiempo de valided del token
	expirationTime := time.Now().Add(5 * time.Minute)

	// Creamos nuestro estructura de token con el username y el tiempo de expiracion
	claims := &Claims{
		Username: user,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declaramos el token con el correspondiente algoritmo, elegimos el tipo de firma y la estructura
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Pasamos nuestro token a string, que es lo que enviaremos en el header
	tokenString, errortoken := token.SignedString(jwtKey)

	//Si hay un error en la creacion del token devolvemos un error
	if errortoken != nil {
		// If there is an error in creating the JWT return an internal server error
		return "", true
	}

	return tokenString, false

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	//json.NewEncoder(w).Encode(&tokenString)
}
