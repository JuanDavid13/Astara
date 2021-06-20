# Astara

Astara is a web application which goal is to help people to manage their goals and tasks, giving them structure and control it tries to increase poeple performance.


## Features

- Light/dark mode toggle
- Task management
- Goal management
- Search 
- Area organization
- Profile configuration
- Short Cuts
- Dynamic pagination
- Dynamic profile update

  
## Tech Stack

**Client-side:** Vue-Cli, Vue-Router, Jquery, SplittingJS, Sass

**Server-side:** Golang (Go), Fiber, Mysql, Nginx

  
## API Reference

### Log in

#### Check User

```http
  POST /api/v1/login
```

| Parámetro | Tipo     | Descrición                |
| :-------- | :------- | :------------------------- |
| `user` | `string` | **Obligatorio.** Nombre o correo electrónico del usuario |

#### Check

```http
  POST /api/v1/login/check
```

| Parámetros | Tipo     | Descrición                       |
| :-------- | :------- | :-------------------------------- |
| `user`      | `string` | **Obligatorio.** Nombre o correo electrónico del usuario |
| `pass`      | `string` | **Obligatorio.** Contraseña del usuario |

#### Create User

```http
  POST /api/v1/login/create
```

| Parámetros | Tipo     | Descrición                       |
| :-------- | :------- | :-------------------------------- |
| `user`      | `string` | **Obligatorio.** Nombre o correo electrónico del usuario |
| `pass`      | `string` | **Obligatorio.** Contraseña del usuario |
| `email`     | `string` | Correo electrónico |

### Authentication

#### Authentication

```http
  GET /api/v1/auth/validate
```

| Parámetros | Tipo     | Descrición                       |
| :-------- | :------- | :-------------------------------- |
| `token`      | `string` | JWT de la cookie llamada "token" |

Comprueba que exista la cookie "token" en la cabezera de la llmada. Si no existe devolverá un estatus 401.
Si la cabecera existe comprobará su contenido. Si este ha sido alterado, devolverá un estado 401.
Si el estado no ha sido alterado, sumará 10 minutos al tiempo de expiración de la cookie,
seguidamente, guardará el contenido del token en el contexto de la petición y continuará con la ejecución.

#### Log Out

```http
  GET /api/v1/auth/logout
```

| Parámetros | Tipo     | Descrición                       |
| :-------- | :------- | :-------------------------------- |
| `token`      | `string` | JWT de la cookie llamada "token" |

Comprueba que exista la cookie "token" en la cabezera de la llmada. Si no existe devolverá un estatus 401.
Si la cabecera existe comprobará su contenido. Si este ha sido alterado, devolverá un estado 401.
Si el estado no ha sido alterado, establecerá el tiempo de expiración, aun tiempo inferior al presente. 
Devolverá un estatus 200 si todo ha ido bien, o un 400 si no.

### Usuario

Cualquier llamada del grupo "/user" tendrá como primer paso la validación del token.

### Task

#### Create Task

```http
  Post /api/v1/user/task/create
```

| Parámetros | Tipo     | Descrición                       |
| :-------- | :------- | :-------------------------------- |
| `slug`      | `string` | **Obligatorio**. Slug del area al que pertenece |
| `name`      | `string` | **Obligatorio**. Nombre de la tarea. |
| `deadline`      | `string` | **Obligatorio.** Fecha limiete de la tarea. |
| `dated`      | `string` | Cuanto tiene pensado realizar la tarea. |
| `id`      | `int` | **Obligatorio.** Id del objetivo al que pertenece. (-1 si no pertenece a ninguno) |

#### Delete Task

```http
  Post /api/v1/user/task/delete
```

| Parámetros | Tipo     | Descrición                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Obligatorio**. Identificador de la tarea a eliminar |

#### Edit Task

```http
  Post /api/v1/user/task/edit
```

| Parámetros | Tipo     | Descrición                       |
| :--------- | :------- | :-------------------------------- |
| `name`     | `string` | **Obligatorio**. Nombre de la tarea. |
| `deadline` | `string` | **Obligatorio**. Fecha límite. |
| `dated`    | `string` | **Obligatorio**. Cuando tiene pensado hacerla. |
| `task_id`  | `int`    | **Obligatorio**. Identificador de la tarea. |

#### Get task of goals

```http
  Post /api/v1/user/task/goal-tasks
```

| Parámetros | Tipo     | Descrición                       |
| :--------- | :------- | :-------------------------------- |
| `id`       | `int`    | **Obligatorio**. Identificador del objetivo. |

#### Check a task as done

```http
  Post /api/v1/user/task/check
```

| Parámetros | Tipo     | Descrición                       |
| :--------- | :------- | :-------------------------------- |
| `id`       | `int`    | **Obligatorio**. Identificador de la tarea. |

### Objetivos

Este grupo está fuera del grupo "/user", pero las llamadas a este grupo también realizan la validación del token.

#### Create

```http
  Post /api/v1/goal/create
```

| Parámetros | Tipo     | Descrición                       |
| :--------- | :------- | :-------------------------------- |
| `slug`       | `string`    | **Obligatorio**. Slug del area al que pertenece. |
| `name`       | `string`    | **Obligatorio**. Nombre del objetivo. |
| `deadline`   | `string`    | **Obligatorio**. Fecha límite del objetivo. |
| `description`| `string`    | **Obligatorio**. Descripción del objetivo. |

#### Delete

```http
  Post /api/v1/goal/delete
```

| Parámetros | Tipo     | Descrición                       |
| :--------- | :------- | :-------------------------------- |
| `id`       | `int`    | **Obligatorio**. Identificador de la tarea. |


#### Edit

```http
  Post /api/v1/goal/edit
```

| Parámetros | Tipo     | Descrición                       |
| :--------- | :------- | :-------------------------------- |
| `name`       | `string`    | **Obligatorio**. Nombre del objetivo. |
| `deadline`   | `string`    | **Obligatorio**. Fecha límite del objetivo. |
| `description`| `string`    | **Obligatorio**. Descripción del objetivo. |
| `goal_id`    | `int`    | **Obligatorio**. Id del objetivo. |

### Profile

#### Get info

```http
  GET /api/v1/user/profile/info
```

#### Edit

```http
  Post /api/v1/user/profile/update
```

| Parámetros | Tipo     | Descrición                       |
| :--------- | :------- | :-------------------------------- |
| `name`       | `string`    | **Obligatorio**. Nombre del objetivo. |
| `deadline`   | `string`    | **Obligatorio**. Fecha límite del objetivo. |
| `description`| `string`    | **Obligatorio**. Descripción del objetivo. |
| `goal_id`    | `int`    | **Obligatorio**. Id del objetivo. |

#### Check password

```http
  Post /api/v1/user/profile/checkpass
```

| Parámetros | Tipo     | Descrición                       |
| :--------- | :------- | :-------------------------------- |
| `password`       | `string`    | **Obligatorio**. Contraseña a comprobar. |

#### Update password

```http
  Post /api/v1/user/profile/changePass
```

| Parámetros | Tipo     | Descrición                       |
| :--------- | :------- | :-------------------------------- |
| `pass`     | `string`    | **Obligatorio**. Nueva contraseña. |

### Area

Aunque este grupo no pertence al grupo "/user", también realiza una validación del (jwt) token.

#### Get all areas

```http
  GET /api/v1/areas
```

#### Check if area correspond

```http
  Post /api/v1/areas/correspond
```

| Parámetros | Tipo     | Descrición                       |
| :--------- | :------- | :-------------------------------- |
| `slug`       | `string`    | **Obligatorio**. Slug del area a comprobar. |

#### Create

```http
  Post /api/v1/areas/create
```

| Parámetros | Tipo     | Descrición                       |
| :--------- | :------- | :-------------------------------- |
| `name`       | `string`    | **Obligatorio**. Nombre del area. |

#### Delete

```http
  Post /api/v1/areas/delete
```

| Parámetros | Tipo     | Descrición                       |
| :--------- | :------- | :-------------------------------- |
| `slug`     | `string`    | **Obligatorio**. Slug del area a eliminar. |

#### Remove target

```http
  Post /api/v1/areas/remove-target
```

| Parámetros | Tipo     | Descrición                       |
| :--------- | :------- | :-------------------------------- |
| `slug`     | `string`    | **Obligatorio**. Slug del area a eliminar. |

#### Update area name

```http
  Post /api/v1/areas/edit
```

| Parámetros | Tipo     | Descrición                       |
| :--------- | :------- | :-------------------------------- |
| `area`     | `string`    | **Obligatorio**. Nombre del area a editar. |
| `name`     | `string`    | **Obligatorio**. Nuevo nombre. |

#### Get main (dashboard) goals

```http
  GET /api/v1/areas/main/goals
```

#### Get main (dashboard) tasks

```http
  GET /api/v1/areas/main/tasks
```

#### Get paginated goals

```http
  GET /api/v1/areas/:slug/paginated-goals/:size/:paginated
```

| Parámetros | Tipo     | Descrición                       |
| :--------- | :------- | :-------------------------------- |
| `slug`     | `string` | **Obligatorio**. Nombre del area. |
| `size`     | `int`    | **Obligatorio**. Número de objetivos ya cargados. |
| `paginated`| `bool`   | **Obligatorio**. Si es 'true' devolverá la siguiente tanda de resultados, si es 'false', devolverá los mismos actualizados. |



## Run Locally

Crea la carpeta donde quieras tener el projecto

```bash
    $ mkdir your-path/astara 
```

Mueve a la nueva carpeta

```bash
    $ cd astara/
```

Clona el proyecto

```bash
  $ git clone https://github.com/JuanDavid13/Astara
```

Instala Golang


> Link a la página de Golang desde donde podrás encontrar los archivos de instalación para
> tu sistema operativo. [https://golang.org/dl/](https://golang.org/dl/).

> Links para la configuración de Golang una vez instalado. [https://golang.org/doc/install](https://golang.org/doc/install), [https://golang.org/doc/install/source](https://golang.org/doc/install/source)


Instala Vue-Cli

```bash
  $ npm install -g vue@vue-cli
```

Ve a la carpeta "astara/"

```bash
  $ cd astara/
```

Instala las dependencias

```bash
  $ npm install
```

Abre una nueva terminal y dirigete a la carpeta de la api

```bash
  $ cd ../api
```

En la terminal de la api, ejecuta el archivo "astara"

```bash
    $ ./astara
```

En la terminal del frontend "astara/", ejecuta el comando
```bash
    $ npm run serve
```
  
## Documentation

[Documentación](https://github.com/JuanDavid13/Astara)

  
## FAQ

#### Es Astara responsive

Sí, la mayoría de elementos se adaptan bien a un formato movil.

#### ¿Cómo está desplegado Astara?

Astara está montado sobre un servidor Nginx, comprado en DigitalOcean. 
Se puede acceder a la página a través de [astarapp.site](http://astarapp.site/).
El servidor contiene instalado [Git](https://git-scm.com/), [Golang](https://golang.org/), y mysql-server.

Para más información revisa la [documentación](https://githug.com/JuanDavid13/Astara)

#### ¿Qué va a pasar con Astara?

El repositorio de Astara estará visible durante 2 semanas más, hasta 4 de Julio,
a partir de entonces el repositorio será privado y se seguirá desarrollando
la aplicación de forma independiente.

Como explico en la documentación, Astara es un projecto personal que me gustaría seguir
desarrollando, pronto recibirá un cambio de look y nombre, implementaré algunas 
funcionalidades más e incluso tengo pensado hacer una versión de escritorio con C# o C++.

  
## Author

[@JuanDavid13](https://www.github.com/JuanDavid13)

  
