// static/js/checkout.js
document.addEventListener("DOMContentLoaded", function () {
    const cartSummary = document.getElementById("cart-summary");
    const placeOrderButton = document.getElementById("place-order-button");

    // Fetch user's cart items from the backend and calculate total price
    fetch("/api/cart", {
        method: "GET",
        headers: {
            "Authorization": "Bearer " + localStorage.getItem("token")
        }
    })
        .then(response => response.json())
        .then(data => {
            // Loop through cart items and dynamically create HTML elements
            // Display cart items and total price
        })
        .catch(error => {
            console.error("Error fetching cart items:", error);
        });

    placeOrderButton.addEventListener("click", function () {
        // Implement order placement logic here
    });
});
// static/js/checkout.js
document.addEventListener("DOMContentLoaded", function () {
    const placeOrderButton = document.getElementById("place-order-button");

    placeOrderButton.addEventListener("click", function () {
        // Prepare order data
        const orderData = {
            /* Include order details here */
        };

        // Place order by sending data to the backend
        fetch("/api/orders", {
            method: "POST",
            headers: {
                "Authorization": "Bearer " + localStorage.getItem("token"),
                "Content-Type": "application/json"
            },
            body: JSON.stringify(orderData)
        })
            .then(response => response.json())
            .then(data => {
                // Process response or handle success
            })
            .catch(error => {
                console.error("Error placing order:", error);
            });
    });
});
