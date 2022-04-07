package certfile

import _ "embed"

//go:embed server.crt
var Cert []byte

//go:embed server.key
var Key []byte

// https://www.cnblogs.com/logchen/p/16030515.html
//go:embed badcert.der
var BadCert []byte
