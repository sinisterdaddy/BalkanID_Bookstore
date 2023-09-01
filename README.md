<!-- Introduction Section -->
<section id="introduction">
    <h2>Introduction</h2>
    <p>
        The Online Bookstore Application is a full-stack web application built to provide users with an interactive platform for browsing, purchasing, and reviewing books. Admin users have access to additional features for inventory management and user administration.
    </p>
</section>

<!-- Features Section -->
<section id="features">
    <h2>Features</h2>
    <ul>
        <li>Secure user registration and authentication</li>
        <li>Account deactivation and deletion</li>
        <li>Protection against common vulnerabilities</li>
        <li>Proper system logging with retention policies</li>
        <li>User-friendly book search and filtering</li>
        <li>Book downloading and review functionality</li>
        <li>Admin capabilities for inventory management</li>
        <li>Dockerized application for easy deployment</li>
        <li>Unit and workflow testing</li>
        <li>Recommendation system for book suggestions (bonus feature)</li>
    </ul>
</section>

<!-- Project Setup Section -->
<section id="project-setup">
    <h2>Project Setup</h2>
    <p>Initializing the Project</p>
    <pre>
        <code>
            git clone https://github.com/your-username/online-bookstore.git
            cd online-bookstore
        </code>
    </pre>
</section>

<!-- Database Configuration Section -->
<section id="database-configuration">
    <h2>Database Configuration</h2>
    <p>Setting up the database:</p>
    <ol>
        <li>Create a PostgreSQL database.</li>
        <li>Configure the database connection in <code>config/config.yaml</code>.</li>
        <li>Run database migrations:</li>
    </ol>
    <pre>
        <code>
            go run cmd/migrate/main.go
        </code>
    </pre>
</section>

<!-- User Authentication Section -->
<section id="user-authentication">
    <h2>User Authentication and Registration</h2>
    <p>User registration and authentication are handled using JSON Web Tokens (JWT). Follow the guidelines in the <code>auth</code> package to configure authentication settings.</p>
</section>

<!-- Data Models Section -->
<section id="data-models">
    <h2>Data Models</h2>
    <p>Data models for users, books, orders, and more are defined using GORM. Refer to the <code>models</code> package for details.</p>
</section>

<!-- Account Deactivation Section -->
<section id="account-deactivation">
    <h2>Account Deactivation and Deletion</h2>
    <p>User accounts can be deactivated or deleted securely. Review the <code>account</code> package for implementation details.</p>
</section>

<!-- Book Management Section -->
<section id="book-management">
    <h2>Book Management</h2>
    <p>CRUD operations for books are implemented. Admins can manage the book inventory through the admin dashboard.</p>
</section>

<!-- Shopping Cart Section -->
<section id="shopping-cart">
    <h2>Shopping Cart</h2>
    <p>Users can add books to their shopping carts, which are stored in the database. The shopping cart is integrated with the checkout process.</p>
</section>

<!-- Order Management Section -->
<section id="order-management">
    <h2>Order Management</h2>
    <p>Users can place orders, view order history, and track the status of their orders.</p>
</section>

<!-- Book Reviews Section -->
<section id="book-reviews">
    <h2>Book Reviews</h2>
    <p>Users can leave reviews on books they've purchased. Reviews are stored in the database for reference.</p>
</section>

<!-- Admin Features Section -->
<section id="admin-features">
    <h2>Admin Features</h2>
    <p>Admins have access to the admin dashboard, where they can manage inventory, user accounts, and perform administrative tasks.</p>
</section>

<!-- API Documentation Section -->
<section id="api-documentation">
    <h2>API Documentation</h2>
    <p>API endpoints are documented using Swagger. To access the documentation, visit <code>/swagger/index.html</code> after running the application.</p>
</section>

<!-- Frontend Section -->
<section id="frontend">
    <h2>Frontend</h2>
    <p>The frontend is implemented using HTML, CSS, and JavaScript. It provides the following pages:</p>
    <ul>
        <li>User registration and login forms</li>
        <li>User profile page</li>
        <li>Book catalog page with search and filtering</li>
        <li>Shopping cart page for reviewing and modifying cart items</li>
        <li>Checkout page for placing orders</li>
        <li>Order history page</li>
        <li>Book detail page for viewing book details and leaving reviews</li>
        <li>Admin dashboard for managing inventory and users</li>
    </ul>
</section>

<!-- Styling and Responsiveness Section -->
<section id="styling-and-responsiveness">
    <h2>Styling and Responsiveness</h2>
    <p>CSS styles are applied to the frontend for a visually appealing design. Media queries ensure responsiveness across different screen sizes.</p>
</section>

<!-- Integration with Backend Section -->
<section id="integration-with-backend">
    <h2>Integration with Backend</h2>
    <p>JavaScript is used to make API requests to the backend for data retrieval and manipulation. The <code>fetch()</code> API is utilized to interact with backend endpoints.</p>
</section>

<!-- Contributing Section -->
<section id="contributing">
    <h2>Contributing</h2>
    <p>Contributions to this project are welcome! Please review our <a href="CONTRIBUTING.md">contributing guidelines</a> for details on how to get started.</p>
</section>

<!-- License Section -->
<section id="license">
    <h2>License</h2>
    <p>This project is licensed under the MIT License - see the <a href="LICENSE">LICENSE</a> file for details.</p>
</section>
