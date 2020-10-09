package config

func GetConfiguration(strategy string) Configuration {
	var c Configuration

	switch strategy {
	case "in_config":
		c.DbUser = DB_USER
		c.DbPassword = DB_PASSWORD
		c.DbHost = DB_HOST
		c.DbPort = DB_PORT
		c.DbDatabase = DB_DATABASE
		c.ApiPort = API_PORT

	case "external":
		// TODO: get values from external provider since these are supposed to be secrets values
		c.DbUser = "TODO"
		c.DbPassword = "TODO"
		c.DbHost = "TODO"
		c.DbPort = 3306
		c.DbDatabase = "TODO"
		c.ApiPort = API_PORT

	default:
		panic("Estrategy not allowed bro")
	}
	return c
}
