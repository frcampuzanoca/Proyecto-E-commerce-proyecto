function addToCart(productId) {
    fetch("/add-to-cart", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(productId)
    })
    .then(response => response.json())
    .then(data => {
        if (data.message) {
            alert(data.message);
            location.reload(); // Refresca para actualizar el total del carrito
        } else {
            alert(data.error);
        }
    })
    .catch(error => console.error("Error:", error));
}