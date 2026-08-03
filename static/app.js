async function calculatePrice() {

    // Get selected food item
    const itemDropdown = document.getElementById("item");
    const selectedOption = itemDropdown.options[itemDropdown.selectedIndex];

    const itemName = selectedOption.value;
    const price = parseFloat(selectedOption.dataset.price);

    // Get selected discount
    const discount = parseFloat(document.getElementById("discount").value);

    // Prepare request object
    const requestData = {
        itemName: itemName,
        price: price,
        discount: discount
    };

    try {

        const response = await fetch("/calculate", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(requestData)
        });

        const result = await response.json();

        document.getElementById("result").innerHTML = `
            <div class="result-card">

                <h4 class="mb-3">Pricing Details</h4>

                <p><strong>Food Item:</strong> ${result.itemName}</p>

                <p><strong>Original Price:</strong> ₹${result.originalPrice.toFixed(2)}</p>

                <p><strong>Discount:</strong> ${result.discountPercent}%</p>

                <p><strong>Discount Amount:</strong> ₹${result.discountAmount.toFixed(2)}</p>

                <hr>

                <h3 class="text-success">
                    Final Price: ₹${result.finalPrice.toFixed(2)}
                </h3>

            </div>
        `;

    } catch (error) {

        alert("Unable to connect to the Go server.");

        console.error(error);

    }

}