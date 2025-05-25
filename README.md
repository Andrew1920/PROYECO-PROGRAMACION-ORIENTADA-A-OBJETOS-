
# 🛒 Sistema de Gestión de E-commerce en Go

Este proyecto implementa un sistema de gestión de e-commerce básico desarrollado en el lenguaje de programación Go. El objetivo es permitir a los usuarios realizar operaciones como visualizar productos, gestionar usuarios, procesar pedidos y manejar la lógica de una tienda en línea.

## 🚀 Objetivo del Proyecto
Desarrollar un sistema modular, funcional y escalable para gestionar una tienda online utilizando Go y sus herramientas modernas, aplicando principios de programación funcional en su diseño.

## 📦 Estructura del Proyecto

```
ecomerce-go/
├── main.go              # Punto de entrada de la aplicación
├── handlers/            # Controladores para manejar peticiones HTTP
├── models/              # Estructuras de datos para productos, usuarios, pedidos
├── routes/              # Configuración de rutas y endpoints
├── utils/               # Funciones auxiliares y utilitarias
└── go.mod               # Archivo de gestión de dependencias
```

## 📚 Módulos Implementados

- **Gestión de Productos**: Crear, listar, actualizar y eliminar productos.
- **Gestión de Usuarios**: Registro, listado y autenticación.
- **Pedidos**: Registro y visualización de órdenes de compra.
- **Rutas RESTful**: Estructura modular con gorilla/mux para facilitar el enrutamiento.

## 🔧 Paquetes Utilizados

- `net/http`: Servidor HTTP estándar.
- `encoding/json`: Manejo de datos en formato JSON.
- `github.com/gorilla/mux`: Router HTTP robusto y flexible.

## 📈 Alcance del Proyecto

Este sistema busca simular el funcionamiento básico de un backend para un sitio e-commerce, incluyendo la gestión de productos, usuarios y pedidos. Está diseñado para servir como base extensible para sistemas reales.

## 📌 Funcionalidades Futuras

- Conexión a base de datos PostgreSQL o MongoDB.
- Autenticación de usuarios con JWT.
- Interfaz gráfica en HTML/CSS/JS.
- Integración con pasarelas de pago (Stripe/PayPal).
- Administración de inventario y promociones.

## ✅ Ejecución del Proyecto

Para ejecutar el proyecto, asegúrate de tener Go instalado y luego:

```bash
git clone https://github.com/tu-usuario/ecomerce-go.git
cd ecomerce-go
go mod tidy
go run main.go
```

