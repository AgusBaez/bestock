# 🗃 BeStock 
**_Se trata de un aplicación de uso empresarial abordando la problematica de la organizacion del inventario, traer soluciones como roles del personal con acceso a la administracion de la indumentaria,lo que necesitaremos una administracion de usuarios, tambien proporcionara estadisticas y graficos tanto de usuarios como de inventario, avisos mediante mensajes, comentarios de la administracion para los usuarios con roles especificos, chatbot para preguntas frecuentes, y otras funcionalidades que se tendran encuenta para lograr simplicidad y comodidad al control del inventario._**

# 📖 About
Empecemos por la arquitectura del proyecto, se trata de una convicacion entre Clean Arquitectur junto con arquitectura DDD(diseño guiado por el dominio). Quiere decir que el software esta distribuido en capas de distintas funcionalidades heredadas por otras capas

Primera capa, trata de comunicarse con la base de datos llamada repositorio;
Esta se hereda a la capa de logica de negocios, esta incluye la logica del negocio llamada services;
Y esta capa se va a heredar a la presnetacion, siendo la forma en que se muestran los datos (vamos a utilizar API, pero se puede utilizar enventos, gRPC ..etc) en conjunto con el UI/UX, lleva el nombre de API en este caso.

Aclaracion de arquitectura:
La capa de repositorio no puede tener referencias a capas superiores a la misma, es decir jamas se deberan importar elementos o funcionalidades de capas superiores, en GO el error es conocido como dependecia circular

Sigamos por las carpetas(Packages) en la app "Settings" llevara la configuracion de la aplicaion, el puerto en donde se comucara el REST API, la configuracion a las credenciales de la base de datos ..otros.

## Otra carpeta/paquete llamada "Internal", En Golang existen propiedades especiales, todo lo que se refiera a un estado sensible de la aplicacion va dentro de esa carpeta, Golang va a ocultar el contenido de funcionalidades que trae este paquete, aca se encuentran las diferetes capas.

---

## 🛠️ Built with the packages: 

- Embed: - Package embed provides access to files embedded in the running Go program.

- Yaml V.3: - The yaml package enables Go programs to comfortably encode and decode YAML values

- Fx Uber: - Fx is a dependency injection system for Go and provide a framework that makes it easy to build applications out of reusable, composable modules

- log: - Package log implements a simple logging package. It defines a type, Logger, with methods for formatting output. That logger writes to standard error and prints the date and time of each logged message

## Agus Baez 👨🏼‍💻
_Hello! 🖖🏼 this proyect is for my practice whit Golang so i start to the a simple application API, and how to use this language when working on a project❣_
**Thanks for reading**

### ⚠ License 

[MIT](https://choosealicense.com/licenses/mit/) [![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)
