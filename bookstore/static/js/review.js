// static/js/review.js
document.addEventListener("DOMContentLoaded", function () {
    const reviewForm = document.getElementById("review-form");

    reviewForm.addEventListener("submit", function (event) {
        event.preventDefault();

        const rating = document.getElementById("rating").value;
        const comment = document.getElementById("comment").value;

        // Implement review submission logic here
    });
});
