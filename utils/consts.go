package utils

// Variables
const (
	DBFILE    = "bakes.db"
	PORT      = ":8080"
	ADDR      = "localhost" + PORT // TODO change when deploying
	TOKEN     = "supersecret"      // TODO change when deploying
	TOKEN_ENV = "BAKERY_TOKEN"     // Environmental variable
	// TODO Set ENV variable with "export BAKERY_TOKEN='supersecret'
)

// API consts
const VERSION = "v1"

// Endpoints
const (
	ROOT      = "/"
	BAKE      = "/bake"
	DASHBOARD = "/dashboard"
	SEED      = "/seed"
)
