// static/js/book_detail.js
document.addEventListener("DOMContentLoaded", function () {
    const authorElement = document.getElementById("author");
    const genreElement = document.getElementById("genre");
    const descriptionElement = document.getElementById("description");
    const addToCartButton = document.getElementById("add-to-cart-button");
    const leaveReviewButton = document.getElementById("leave-review-button");

    // Fetch book details from the backend
    fetch("/api/books/book_id", {
        method: "GET",
        headers: {
            "Authorization": "Bearer " + localStorage.getItem("token")
        }
    })
        .then(response => response.json())
        .then(data => {
            authorElement.textContent = data.author;
            genreElement.textContent = data.genre;
            descriptionElement.textContent = data.description;
            // Update other book details as needed
        })
        .catch(error => {
            console.error("Error fetching book details:", error);
        });

    addToCartButton.addEventListener("click", function () {
        // Implement "Add to Cart" functionality here
    });

    leaveReviewButton.addEventListener("click", function () {
        // Redirect to the review page or open a modal for leaving a review
    });
});
