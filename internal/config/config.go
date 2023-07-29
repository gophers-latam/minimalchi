package config

// Conf contiene los datos de configuración de la api.
// Implemente los campos que su aplicación necesite y defina un medio para obtener
// esos datos, ya sea cargandolos desde un archivo o desde flags de la línea de comandos.
type Conf struct {
	// APPName contiene el nombre de su aplicación
	APPName string `json:"appName" yaml:"app_name"`

	// Port infica en que puerto poner a escuchar a esta api
	Port int `json:"port" yaml:"port"`
}

var cfg Conf

// NewConf devuelve una instancia de la configuración de la aplicación.
// Es un buen punto para cargar los datos de configuración si ud. así lo necesitase.
func NewConf() Conf {

	if cfg.APPName == "" {
		cfg = Conf{
			APPName: "Gophers Latam Minimal Chi Api",
			Port:    3000, // reemplace por su puerto deseado o carguelo desde un archivo o desde la línea de comandos
		}
	}

	return cfg
}
