
# E-Commerce Platform

A full-featured, scalable e-commerce web application designed to provide a seamless shopping experience for users and comprehensive management capabilities for administrators.

---

## Table of Contents

- [Project Overview](#project-overview)  
- [Key Features](#key-features)  
- [Technology Stack](#technology-stack)  
- [Getting Started](#getting-started)  
- [Environment Variables](#environment-variables)  
- [Running the Application](#running-the-application)  
- [API Documentation](#api-documentation)  
- [Folder Structure](#folder-structure)  
- [Contributing](#contributing)  
- [License](#license)  
- [Contact](#contact)  

---

## Project Overview

This project implements a modern e-commerce platform that enables users to browse products, filter by categories and attributes, manage a shopping cart, and securely complete purchases. It features a responsive front-end built with React and a robust RESTful API backend using Node.js and Express. The platform supports user authentication via JWT tokens and role-based access control to differentiate between customers and administrators.

The admin panel allows authorized users to manage product catalogs, categories, orders, and user data efficiently. The architecture prioritizes modularity, maintainability, and performance, following industry best practices.

---

## Key Features

- **User Registration & Authentication**: Secure sign-up and login with JWT authentication.
- **Product Browsing & Search**: Browse products with pagination, keyword search, and multi-criteria filtering (category, price range, brand).
- **Product Detail View**: Detailed product information with images, specifications, and customer reviews.
- **Shopping Cart & Checkout**: Add/remove items from cart, update quantities, and proceed through a multi-step checkout process.
- **Order Management**: Users can view their order history and track order status in real-time.
- **Admin Panel**: Manage products, categories, orders, and user roles through a dedicated dashboard.
- **Responsive Design**: Fully optimized UI for desktop, tablet, and mobile devices.
- **Security**: Password hashing with bcrypt, JWT-based authentication, and role-based authorization.
- **RESTful API**: Well-structured API endpoints with input validation and error handling.
- **State Management**: Frontend state managed via Redux for scalable and predictable state updates.
- **Error Handling & Notifications**: User-friendly error messages and confirmation alerts throughout the app.

---

## Technology Stack

| Layer        | Technology / Library          |
|--------------|------------------------------|
| Frontend     | React.js, Redux, React Router, Tailwind CSS, Axios |
| Backend      | Node.js, Express.js           |
| Database     | MongoDB, Mongoose             |
| Authentication | JSON Web Token (JWT), bcrypt |
| Testing      | Jest, Supertest (for API)     |
| Development  | ESLint, Prettier, Nodemon     |
| Version Control | Git & GitHub               |

---

## Getting Started

### Prerequisites

- Node.js (v16 or higher recommended)  
- npm or yarn package manager  
- MongoDB database instance (local or cloud e.g., MongoDB Atlas)  

### Installation

1. **Clone the repository**

```bash
git clone https://github.com/burakcanheyal/e-commerce.git
cd e-commerce
```

2. **Install backend dependencies**

```bash
cd backend
npm install
```

3. **Install frontend dependencies**

```bash
cd ../frontend
npm install
```

---

## Environment Variables

Create a `.env` file in the `backend` directory with the following variables:

```env
PORT=5000
MONGO_URI=your_mongodb_connection_string
JWT_SECRET=your_secret_key
NODE_ENV=development
```

- `PORT`: Port on which the backend server will run.
- `MONGO_URI`: Connection string for MongoDB.
- `JWT_SECRET`: Secret key used for signing JWT tokens.
- `NODE_ENV`: Environment mode (`development` or `production`).

---

## Running the Application

### Backend

Navigate to the backend directory and start the server:

```bash
cd backend
npm run dev
```

The backend API will be available at: `http://localhost:5000`

### Frontend

Navigate to the frontend directory and start the React app:

```bash
cd ../frontend
npm start
```

The frontend will be accessible at: `http://localhost:3000`

---

## API Documentation

Below is an overview of the primary REST API endpoints. Detailed API documentation can be found in the `docs` folder or by using API tools like Postman.

| HTTP Method | Endpoint                 | Description                                | Access         |
|-------------|--------------------------|--------------------------------------------|----------------|
| GET         | `/api/products`           | Retrieve list of all products with filters and pagination | Public         |
| GET         | `/api/products/:id`       | Retrieve detailed information about a specific product | Public         |
| POST        | `/api/users/register`     | Register a new user                        | Public         |
| POST        | `/api/users/login`        | Authenticate user and return JWT token    | Public         |
| GET         | `/api/users/profile`      | Retrieve logged-in user profile            | Private (User) |
| PUT         | `/api/users/profile`      | Update logged-in user profile              | Private (User) |
| POST        | `/api/orders`             | Create a new order                         | Private (User) |
| GET         | `/api/orders/:id`         | Get details of a specific order            | Private (User/Admin) |
| GET         | `/api/orders`             | List all orders (Admin only)               | Private (Admin)|
| POST        | `/api/products`           | Add new product (Admin only)               | Private (Admin)|
| PUT         | `/api/products/:id`       | Update existing product (Admin only)       | Private (Admin)|
| DELETE      | `/api/products/:id`       | Delete a product (Admin only)               | Private (Admin)|

---

## Folder Structure

```
e-commerce/
│
├── backend/                 # Backend source code (Node.js + Express)
│   ├── controllers/         # Route controllers
│   ├── models/              # Mongoose models
│   ├── routes/              # Express route definitions
│   ├── middleware/          # Authentication & error handling middleware
│   ├── config/              # Configuration files (db, environment)
│   ├── utils/               # Utility functions
│   ├── tests/               # Backend test cases
│   ├── server.js            # Backend entry point
│   └── package.json
│
├── frontend/                # Frontend source code (React)
│   ├── public/              # Public assets
│   ├── src/
│   │   ├── components/      # Reusable React components
│   │   ├── pages/           # Page components
│   │   ├── redux/           # Redux slices and store
│   │   ├── services/        # API service calls (Axios)
│   │   ├── styles/          # Tailwind CSS and custom styles
│   │   ├── utils/           # Helper functions
│   │   ├── App.js           # Main React component
│   │   └── index.js         # React entry point
│   └── package.json
│
├── docs/                    # Documentation and API specs
├── README.md                # This file
└── .gitignore
```

---

## Contributing

Contributions are welcome! To contribute:

1. Fork the repository  
2. Create a new feature branch (`git checkout -b feature/your-feature`)  
3. Commit your changes (`git commit -m "Add some feature"`)  
4. Push to the branch (`git push origin feature/your-feature`)  
5. Open a Pull Request describing your changes  

Please ensure your code adheres to the existing style conventions and includes appropriate tests.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Contact

If you have questions, suggestions, or want to collaborate, feel free to reach out:

**Burak Can Heyal**  
Email: burakcanheyal@gmail.com  
GitHub: [https://github.com/burakcanheyal](https://github.com/burakcanheyal)  

---

Thank you for checking out this project! 🚀
