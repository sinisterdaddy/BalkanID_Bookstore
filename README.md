Online Bookstore Application
Welcome to the Online Bookstore Application, a web-based platform for buying and managing books.

Table of Contents
Introduction
Features
Project Setup
Database Configuration
User Authentication and Registration
Data Models
Account Deactivation and Deletion
Book Management
Shopping Cart
Order Management
Book Reviews
Admin Features
API Documentation
Frontend
Styling and Responsiveness
Integration with Backend
Contributing
License
Introduction
The Online Bookstore Application is a full-stack web application built to provide users with an interactive platform for browsing, purchasing, and reviewing books. Admin users have access to additional features for inventory management and user administration.

Features
Secure user registration and authentication
Account deactivation and deletion
Protection against common vulnerabilities
Proper system logging with retention policies
User-friendly book search and filtering
Book downloading and review functionality
Admin capabilities for inventory management
Dockerized application for easy deployment
Unit and workflow testing
Recommendation system for book suggestions (bonus feature)
Project Setup
Initializing the Project
To set up the project, follow these steps:

Clone the repository:

bash
Copy code
git clone https://github.com/your-username/online-bookstore.git
Change into the project directory:

bash
Copy code
cd online-bookstore
Database Configuration
The application uses PostgreSQL as the database. To set up the database:

Create a PostgreSQL database.

Configure the database connection in config/config.yaml.

Run database migrations:

bash
Copy code
go run cmd/migrate/main.go
User Authentication and Registration
User registration and authentication are handled using JSON Web Tokens (JWT). Follow the guidelines in the auth package to configure authentication settings.

Data Models
Data models for users, books, orders, and more are defined using GORM. Refer to the models package for details.

Account Deactivation and Deletion
User accounts can be deactivated or deleted securely. Review the account package for implementation details.

Book Management
CRUD operations for books are implemented. Admins can manage the book inventory through the admin dashboard.

Shopping Cart
Users can add books to their shopping carts, which are stored in the database. The shopping cart is integrated with the checkout process.

Order Management
Users can place orders, view order history, and track the status of their orders.

Book Reviews
Users can leave reviews on books they've purchased. Reviews are stored in the database for reference.

Admin Features
Admins have access to the admin dashboard, where they can manage inventory, user accounts, and perform administrative tasks.

API Documentation
API endpoints are documented using Swagger. To access the documentation, visit /swagger/index.html after running the application.

Frontend
The frontend is implemented using HTML, CSS, and JavaScript. It provides the following pages:

User registration and login forms
User profile page
Book catalog page with search and filtering
Shopping cart page for reviewing and modifying cart items
Checkout page for placing orders
Order history page
Book detail page for viewing book details and leaving reviews
Admin dashboard for managing inventory and users
Styling and Responsiveness
CSS styles are applied to the frontend for a visually appealing design. Media queries ensure responsiveness across different screen sizes.

Integration with Backend
JavaScript is used to make API requests to the backend for data retrieval and manipulation. The fetch() API is utilized to interact with backend endpoints.

Contributing
Contributions to this project are welcome! Please review our contributing guidelines for details on how to get started.
