# Minimal Chi

Un template mínimo recargado con el router [Chi](https://go-chi.io/#/) para ser usado como base para construir apis REST.

Minimal Chi no intenta decidir por ud. que clase de proyecto realizar ni que dependencias importar. Ofrece simplemente una estructura de directorios base, componible y expandible para ser usada en cualquier proyecto api REST, manteniendo algunos elementos idiomáticos, con el agregado del router chi. 


## Prueba rápida

Descargue el repositorio e instale las dependencias:

```bash
$ make deps
```

Levante la api

```bash
$ make run
```

Compruebe el funcionamiento pinchando en elndpoint de ejemplo (por defecto el puerto inicial es el 3000, lo que puede ser cambiado en *internal/config/config.go*)

```bash
$ curl --location --request GET 'http://localhost:3000/v1/ping'
```

Debería ver la siguiente respuesta:

```json
{"code":200,"msg":"Bienvenido a la plantilla minimal Chi de Gophers Latam"}
```

## Api

En el paquete *api/v1*, en el archivo *routes.go puede registrar las rutas de su aplicación, escribiendo los handlers en *api/v1/handlers* según los vaya necesitando.

Puede revisar el endpoint de ejemplo `Ping` que se incluye.

## Visualización de stats

Se incluyen los manejadores para poder revisar stats de la api como el consumo de memoria y otros usando el paquete [statsviz](https://github.com/arl/statsviz). Con su aplicación ejecutandose, solo dirijase en su navegador favorito a http://localhost:3000/debug/statsviz.

Si desea desactivar esta feature, solo comente las líneas donde se le referencia en el archivo *api/v1/routes.go* y ejecute `$ make deps`.

![image](https://github.com/gophers-latam/minimalchi/assets/20423399/855d1ae3-4836-4507-ae3a-73b8ede4ced6)


## Configuración

En el paquete internal/config puede, si gusta, agregar campos al struct `Conf` para cargar las configuraciones de su aplicación, usando la función `NewConf`.

```Go
// Conf contiene los datos de configuración de la api.
// Implemente los campos que su aplicación necesite y defina un medio para obtener
// esos datos, ya sea cargandolos desde un archivo o desde flags de la línea de comandos.
type Conf struct {
	// APPName contiene el nombre de su aplicación
	APPName string `json:"appName" yaml:"app_name"`

	// Port infica en que puerto poner a escuchar a esta api
	Port int `json:"port" yaml:"port"`
}
```

De esta forma ud. decide si cargarlas desde un archivo, desde la línea de comandos o desde otros medios.
