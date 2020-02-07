package types

//
// JWK is used with AWS Cognito's well known
// public rsa keys
//
type JWK struct {
	Al  string `json:"al"`
	E   string `json:"e"`
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	N   string `json:"n"`
	Use string `json:"use"`
}

//
// Keys is the return from the AWS Well known keys
// route.
//
type Keys struct {
	Keys []JWK `jwon:"keys"`
}
