<?php
// Increase maximum execution time to 300 seconds (5 minutes)
set_time_limit(300);

// Retrieve the POST data
$data = json_decode(file_get_contents("php://input"), true);

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

$name = mysqli_real_escape_string($conn, $data["name"]);
$dob = mysqli_real_escape_string($conn, $data["dob"]);
$email = mysqli_real_escape_string($conn, $data["email"]);
$mobile = mysqli_real_escape_string($conn, $data["mobile"]); // Added mobile
$password = mysqli_real_escape_string($conn, $data["password"]);

// Get the lists from the data
$publicFocusAreas = $data["public_focus_areas"];
$privateFocusAreas = $data["private_focus_areas"];
$publicCommunities = $data["public_communities"];
$privateCommunities = $data["private_communities"];

// Convert the lists to JSON format and sanitize
$publicFocusAreasJson = json_encode($publicFocusAreas);
$privateFocusAreasJson = json_encode($privateFocusAreas);
$publicCommunitiesJson = json_encode($publicCommunities);
$privateCommunitiesJson = json_encode($privateCommunities);

// cURL to create DID
$curl = curl_init();
curl_setopt($curl, CURLOPT_URL, 'http://localhost:20000/api/createdid');
curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
curl_setopt($curl, CURLOPT_POST, true);
$did_config = array(
    "type" => 0,
    "dir" => "",
    "config" => "",
    "master_did" => "",
    "secret" => "My DID Secret",
    "priv_pwd" => "mypassword",
    "quorum_pwd" => "mypassword",
    "img_file" => "image.png",
    "did_img_file" => "",
    "pub_img_file" => "",
    "priv_img_file" => "",
    "pub_key_file" => "",
    "priv_key_file" => "",
    "quorum_pub_key_file" => "",
    "quorum_priv_key_file" => ""
);
curl_setopt($curl, CURLOPT_POSTFIELDS, array('did_config' => json_encode($did_config)));
$response = curl_exec($curl);

if ($response === false) {
    echo "cURL Error: " . curl_error($curl) . "<br>";
    $createdDid = null;
} else {
    $responseData = json_decode($response, true);
    if ($responseData && isset($responseData["status"]) && $responseData["status"]) {
        // Successfully created DID
        $createdDid = $responseData["result"]["did"];
        echo "DID created successfully: " . $createdDid . "<br>";
    } else {
        echo "Error creating DID: " . ($responseData ? $responseData["message"] : "Unknown error") . "<br>";
        $createdDid = null;
    }
}

curl_close($curl);

// SQL query to insert the data into the database, including DID
$sql = "INSERT INTO members (name, dob, email, mobile, password, focus_areas_private, focus_areas_public, communities_private, communities_public, did) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)";
$stmt = $conn->prepare($sql);
$stmt->bind_param("ssssssssss", $name, $dob, $email, $mobile, $password, $privateFocusAreasJson, $publicFocusAreasJson, $privateCommunitiesJson, $publicCommunitiesJson, $createdDid);

if ($stmt->execute()) {
    echo "User data inserted successfully<br>";
} else {
    echo "Error inserting user data: " . $stmt->error . "<br>";
}

// Close the connection
$stmt->close();
$conn->close();
?>
