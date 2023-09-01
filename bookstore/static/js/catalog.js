// static/js/catalog.js
document.addEventListener("DOMContentLoaded", function () {
    const searchInput = document.getElementById("search-input");
    const filterGenre = document.getElementById("filter-genre");
    const applyFilterButton = document.getElementById("apply-filter");
    const bookList = document.getElementById("book-list");

    applyFilterButton.addEventListener("click", function () {
        const searchTerm = searchInput.value;
        const selectedGenre = filterGenre.value;

        // Fetch book data from the backend based on search and filter criteria
        // Use fetch() API to send GET request to your API endpoint
        // Update the bookList with fetched data
        fetch("/api/books", {
            method: "GET",
            headers: {
                "Authorization": "Bearer " + localStorage.getItem("token")
            }
        });

    });
});
// static/js/catalog.js
document.addEventListener("DOMContentLoaded", function () {
    // Fetch books from the backend
    fetch("/api/books", {
        method: "GET",
        headers: {
            "Authorization": "Bearer " + localStorage.getItem("token")
        }
    })
        .then(response => response.json())
        .then(data => {
            // Process and display fetched books
        })
        .catch(error => {
            console.error("Error fetching books:", error);
        });
});
