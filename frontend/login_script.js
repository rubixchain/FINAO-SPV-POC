document.addEventListener("DOMContentLoaded", function() {
    const loginForm = document.getElementById("login-form");

    loginForm.addEventListener("submit", function(event) {
        event.preventDefault();

        const email = document.getElementById("login-email").value;
        const password = document.getElementById("login-password").value;

        // Send the login credentials to the server for validation
        validateLogin(email, password);
    });

    function validateLogin(email, password) {
        fetch("validate_login.php", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ email, password })
        })
        .then(response => response.json())
        .then(data => {
            if (data.success) {
                // Redirect to the profile page and pass the email
                window.location.href = "profile.html?email=" + encodeURIComponent(email);
            } else {
                console.log("Login failed. Please check your credentials.");
                alert("Login failed. Please check your credentials");
                loginForm.reset();
            }
        })
        .catch(error => {
            console.error("Error:", error);
        });
    }
});
