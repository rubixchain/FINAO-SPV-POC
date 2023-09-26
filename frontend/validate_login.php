<?php
session_start(); // Start the session

// Retrieve the POST data
$data = json_decode(file_get_contents("php://input"), true);

// database connection details
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

$email = mysqli_real_escape_string($conn, $data["email"]);
$password = mysqli_real_escape_string($conn, $data["password"]);

//SQL query to check login credentials
$sql = "SELECT * FROM members WHERE email = ? AND password = ?";
$stmt = $conn->prepare($sql);
$stmt->bind_param("ss", $email, $password);
$stmt->execute();

// Fetch the result
$result = $stmt->get_result();

$response = array();

if ($result->num_rows === 1) {
    // Login successful
    $_SESSION["user_email"] = $email; // Store user email in the session
    $response["success"] = true;
} else {
    // Login failed
    $response["success"] = false;
}

// Close the connection
$stmt->close();
$conn->close();

// Send the JSON response
echo json_encode($response);
?>
