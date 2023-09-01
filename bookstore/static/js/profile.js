// static/js/profile.js
document.addEventListener("DOMContentLoaded", function () {
    const usernameElement = document.getElementById("username");
    const emailElement = document.getElementById("email");

    // Fetch user information from the backend
    fetch("/api/user/profile", {
        method: "GET",
        headers: {
            "Authorization": "Bearer " + localStorage.getItem("token")
        }
    })
        .then(response => response.json())
        .then(data => {
            usernameElement.textContent = data.username;
            emailElement.textContent = data.email;
            // Update other profile details as needed
        })
        .catch(error => {
            console.error("Error fetching user profile:", error);
        });
});
