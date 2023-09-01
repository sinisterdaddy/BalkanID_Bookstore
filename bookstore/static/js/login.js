// static/js/login.js
document.addEventListener("DOMContentLoaded", function () {
    const loginForm = document.getElementById("login-form");
    const loginStatus = document.getElementById("login-status");

    loginForm.addEventListener("submit", function (event) {
        event.preventDefault();

        const username = document.getElementById("username").value;
        const password = document.getElementById("password").value;

        // Perform form validation and login logic here

        // Simulate a successful login for demonstration
        loginStatus.textContent = "Login successful!";
    });
});
