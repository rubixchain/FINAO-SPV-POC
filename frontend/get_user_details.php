<?php
session_start(); // Start the session

// If the user email is not set in the session, return an error
if (!isset($_SESSION["user_email"])) {
    echo json_encode(array("success" => false, "message" => "User not authenticated"));
    exit;
}

// Database connection details
$servername = "localhost";
$username = "root";
$password = "";
$dbname = "registration";

// Create a new connection
$conn = new mysqli($servername, $username, $password, $dbname);

// Check the connection
if ($conn->connect_error) {
    die("Connection failed: " . $conn->connect_error);
}

// Fetch user details using the email stored in the session
$email = $_SESSION["user_email"];

$sql = "SELECT name, email, mobile, focus_areas_private, focus_areas_public, communities_private, communities_public FROM members WHERE email = ?";
$stmt = $conn->prepare($sql);
$stmt->bind_param("s", $email);
$stmt->execute();
$result = $stmt->get_result();

if ($result->num_rows === 1) {
    $row = $result->fetch_assoc();
    $userDetails = array(
        "name" => $row["name"],
        "email" => $row["email"],
        "mobile" => $row["mobile"],
        "private_focus_areas" => json_decode($row["focus_areas_private"]),
        "public_focus_areas" => json_decode($row["focus_areas_public"]),
        "private_communities" => json_decode($row["communities_private"]),
        "public_communities" => json_decode($row["communities_public"])
    );

    // Return user details as JSON response
    echo json_encode(array("success" => true, "user" => $userDetails));
} else {
    // User not found
    echo json_encode(array("success" => false, "message" => "User not found"));
}

// Close the connection
$stmt->close();
$conn->close();
?>
