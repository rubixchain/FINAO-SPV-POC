document.addEventListener("DOMContentLoaded", function() {
    const form = document.getElementById("signup-form");

    form.addEventListener("submit", function(event) {
        event.preventDefault();

        const name = form.querySelector('[name="name"]').value;
        const dob = form.querySelector('[name="dob"]').value;
        const email = form.querySelector('[name="email"]').value;
        const mobile = form.querySelector('[name="mobile"]').value; // New field
        const password = form.querySelector('[name="password"]').value;

        // Get selected focus areas and communities
        const publicFocusAreasSelect = form.querySelector('[name="focus_areas_public[]"]');
        const selectedPublicFocusAreas = Array.from(publicFocusAreasSelect.selectedOptions).map(option => option.value);

        const privateFocusAreasSelect = form.querySelector('[name="focus_areas_private[]"]');
        const selectedPrivateFocusAreas = Array.from(privateFocusAreasSelect.selectedOptions).map(option => option.value);

        const publicCommunitiesSelect = form.querySelector('[name="communities_public[]"]');
        const selectedPublicCommunities = Array.from(publicCommunitiesSelect.selectedOptions).map(option => option.value);

        const privateCommunitiesSelect = form.querySelector('[name="communities_private[]"]');
        const selectedPrivateCommunities = Array.from(privateCommunitiesSelect.selectedOptions).map(option => option.value);

        // Create an object to represent the form data
        const formData = {
            name: name,
            dob: dob,
            email: email,
            mobile: mobile, // Added mobile
            password: password,
            public_focus_areas: selectedPublicFocusAreas,
            private_focus_areas: selectedPrivateFocusAreas,
            public_communities: selectedPublicCommunities,
            private_communities: selectedPrivateCommunities
        };

        const xhr = new XMLHttpRequest();
        xhr.open("POST", "submit-signup.php", true);
        xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");

        xhr.onreadystatechange = function() {
            if (xhr.readyState === XMLHttpRequest.DONE) {
                if (xhr.status === 200) {
                    console.log("Server response:", xhr.responseText);

                    // Clear form fields
                    alert("Sign Up successful!");
                    form.reset();
                    window.location.href = "login.html";
                } else {
                    console.error("Error submitting signup data:", xhr.statusText);
                }
            }
        };

        xhr.send(JSON.stringify(formData));
    });
});
