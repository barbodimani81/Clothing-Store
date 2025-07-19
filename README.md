# 🛍️ B2C eCommerce Backend (Go + Gin + PostgreSQL)

A production-ready backend for a simple B2C online store — built with **Go**, **Gin**, **GORM**, **PostgreSQL**, and **Docker**.  
Includes full authentication, cart/order flow, and a clean structure for scaling.

---

## 🚀 Features

- ✅ User registration & login with JWT auth
- 🛒 Add/remove/view cart items
- 📦 Checkout flow: create orders from cart
- 📜 View order history
- 🧱 Product management (admin-ready)
- 🐘 PostgreSQL with GORM ORM
- 🐳 Dockerized app with `docker-compose`
- 🧱 Clean modular structure

---

## 📁 Project Structure

store/
├── config/ # Database connection and environment config
│ └── db.go
├── controllers/ # Route handler logic
│ ├── auth.go
│ ├── cart.go
│ ├── order.go
│ └── product.go
├── docker/ # Dockerfile lives here
│ └── Dockerfile
├── middleware/ # JWT auth middleware
├── models/ # GORM models (User, Product, CartItem, Order, etc.)
├── requests/ # Input validation structs
├── routes/ # All route registrations
├── .env # Environment variables for DB config
├── docker-compose.yml # Compose file for Go + PostgreSQL
├── go.mod / go.sum # Go dependencies
├── main.go # App entry point

---

## ⚙️ Getting Started

### ✅ Requirements

- Docker + Docker Compose installed
- Go 1.21+ (optional for local dev only)

---

### 📄 .env File

Create a `.env` file in the root:

```env
DB_HOST=db
DB_USER=root
DB_PASSWORD=password
DB_NAME=mydb
DB_PORT=5432


📡 API Endpoints

🔐 Authentication
Method	Endpoint	Description
POST	/register	Register a new user
POST	/login	Login and receive a JWT
GET	/me	Get current authenticated user info
🛒 Cart
Method	Endpoint	Description
GET	/cart	View all items in user's cart
POST	/cart/add	Add product to cart
POST	/cart/remove/:id	Remove product from cart by ID
📦 Orders
Method	Endpoint	Description
POST	/order/checkout	Create an order from cart items
GET	/order/history	View all previous user orders
🧱 Products
Method	Endpoint	Description
GET	/products	Get list of all products
POST	/products	Create a new product
🔐 Admin (optional/future)
Method	Endpoint	Description
GET	/admin/users	List all users
GET	/admin/orders	View all orders in system
POST	/admin/products	Add a new product (admin only)
You can insert this section under a ## 📡 API Endpoints heading in your README.md. Let me know if you'd like to include example request/response bodies as well.