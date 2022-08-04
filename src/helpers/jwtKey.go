package helpers

func JwtKey() []byte {
	var jwtKey = []byte("secret_key")
	return jwtKey
}
