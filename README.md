# TODO-CLI

TODO-CLI es una aplicación de línea de comandos para gestionar tareas. Está construida utilizando el lenguaje de programación Go y sigue el patrón de arquitectura limpia (Clean Architecture/ Hexagonal Architecure).

## Arquitectura

La arquitectura del proyecto está organizada en las siguientes capas:

- **cmd**: Contiene el punto de entrada de la aplicación.
- **config**: Maneja las configuraciones específicas para diferentes entornos.
- **infrastructure**: Contiene la configuración y conexión a la base de datos.
- **internal/domain**: Define las entidades y lógica de negocio.
- **internal/ports**: Define las interfaces para los servicios y repositorios.
- **internal/repository**: Implementa las interfaces de repositorio, interactuando con la base de datos.
- **internal/usecase**: Implementa la lógica de los casos de uso específicos.

## Funcionalidades

- **Agregar Tarea**: Permite agregar una nueva tarea a la lista.
- **Obtener Todas las Tareas**: Permite obtener todas las tareas almacenadas.

## Instalación

1. Clona el repositorio:
    ```sh
    git clone https://github.com/tu-usuario/todo-cli.git
    cd todo-cli
    ```

2. Instala las dependencias:
    ```sh
    go mod tidy
    ```

## Uso

### Agregar una Tarea

```sh
go run cmd/main.go add "Nueva tarea" "Descripcion de la tarea"
go run cmd/main.go list 