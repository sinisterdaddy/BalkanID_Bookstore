// static/js/register.js
document.addEventListener("DOMContentLoaded", function () {
    const registrationForm = document.getElementById("registration-form");
    const registrationStatus = document.getElementById("registration-status");

    registrationForm.addEventListener("submit", function (event) {
        event.preventDefault();

        const username = document.getElementById("username").value;
        const password = document.getElementById("password").value;

        // Perform form validation here

        // Simulate a successful registration for demonstration
        registrationStatus.textContent = "Registration successful!";
    });
});
