// static/js/order_history.js
document.addEventListener("DOMContentLoaded", function () {
    const orderList = document.getElementById("order-list");

    // Fetch user's order history from the backend
    fetch("/api/orders", {
        method: "GET",
        headers: {
            "Authorization": "Bearer " + localStorage.getItem("token")
        }
    })
        .then(response => response.json())
        .then(data => {
            // Loop through order history items and dynamically create HTML elements
            // Display order details
        })
        .catch(error => {
            console.error("Error fetching order history:", error);
        });
});
