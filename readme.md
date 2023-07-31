# 🗃 BeStock 
**_This is an enterprise application addressing the problem of inventory organisation, bringing solutions such as staff roles with access to clothing management, which will require user management, also providing statistics and graphs for both users and inventory, message alerts, management feedback for users with specific roles, chatbot for frequently asked questions, and other features that will be taken into account to achieve simplicity and convenience to inventory control._**

# 📖 Architecture
Let's start with the architecture of the project, it is a combination of Clean Architecture and DDD (Domain Driven Design). This means that the software is distributed in layers of different functionalities inherited by other layers.

First layer, it tries to communicate with the database called repository;
This is inherited to the business logic layer, this includes the business logic called services;
And this layer is going to inherit to the presentation, being the way in which the data is displayed (we are going to use API, but you can use containers, gRPC ..etc) in conjunction with the UI/UX, it is called API in this case.

### Aclaracion de arquitectura:
The repository layer cannot have references to layers above it, i.e. you should never import elements or functionalities from higher layers, in GO the error is known as circular dependency.

Let's continue through the folders (Packages) in the app "Settings" will carry the configuration of the application, the port where the REST API will be configured, the configuration of the database credentials ..others.
Another folder/package called "Internal", in Golang there are special properties, everything that refers to a sensitive state of the application goes inside this folder, Golang will hide the content of functionalities that this package brings, here are the different layers..
 
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
