package config

// web config
const (

	// WebConfig

	TEMPLATEPATH = "web/template"
	IMGPATH      = "web/img"
	SERVERPORT   = "80"

	// RPC var

	RPCSERVER = "127.0.0.1"
	RPCPORT   = "2333"
	MAXCONN   = 150
	INITCONN  = 150
)

// SUFFIX is file suffix check map for file upload
var SUFFIX = map[string]string{
	".jpg":  ".jpg",
	".jpeg": ".jpeg",
	".gif":  ".gif",
	".png":  ".png",
}
