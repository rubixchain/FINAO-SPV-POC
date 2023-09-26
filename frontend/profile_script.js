document.addEventListener("DOMContentLoaded", function() {
    // Fetch user details from the server using the stored session
    fetch("get_user_details.php")
        .then(response => response.json())
        .then(data => {
            if (data.success) {
                const user = data.user;

                document.getElementById("profile-name").textContent = user.name;
                document.getElementById("profile-email").textContent = user.email;
                document.getElementById("profile-mobile").textContent = user.mobile; // Added mobile

                const publicFocusAreasElement = document.getElementById("profile-public-focus-areas");
                const publicFocusAreasArray = Array.isArray(user.public_focus_areas)
                    ? user.public_focus_areas
                    : JSON.parse(user.public_focus_areas);
                publicFocusAreasElement.textContent = publicFocusAreasArray.join(", ");

                const privateFocusAreasElement = document.getElementById("profile-private-focus-areas");
                const privateFocusAreasArray = Array.isArray(user.private_focus_areas)
                    ? user.private_focus_areas
                    : JSON.parse(user.private_focus_areas);
                privateFocusAreasElement.textContent = privateFocusAreasArray.join(", ");

                const publicCommunitiesElement = document.getElementById("profile-public-communities");
                const publicCommunitiesArray = Array.isArray(user.public_communities)
                    ? user.public_communities
                    : JSON.parse(user.public_communities);
                publicCommunitiesElement.textContent = publicCommunitiesArray.join(", ");

                const privateCommunitiesElement = document.getElementById("profile-private-communities");
                const privateCommunitiesArray = Array.isArray(user.private_communities)
                    ? user.private_communities
                    : JSON.parse(user.private_communities);
                privateCommunitiesElement.textContent = privateCommunitiesArray.join(", ");
            } else {
                console.error("Error fetching user details.");
            }
        })
        .catch(error => {
            console.error("Error:", error);
        });
});
