package utils

// Variables
const (
	DBFILE = "bakes.db"
	PORT   = ":8080"
	ADDR   = "localhost" + PORT // TODO change when deploying
	TOKEN  = "supersecret"      // TODO change when deploying
)

// API consts
const VERSION = "v1"

// Endpoints
const (
	ROOT = "/"
	BAKE = "/" + VERSION + "/bake/"
)
