import React, { useState, useEffect } from 'react';
import { Container, Typography, Paper, Box, Divider, Alert, Button } from '@mui/material';
import { styled } from '@mui/system';

const StyledPaper = styled(Paper)(({ theme }) => ({
    marginTop: theme.spacing(4),
    padding: theme.spacing(3),
    overflow: 'auto',
    wordWrap: 'break-word',
}));

const DataPage = () => {
    const [privateData, setPrivateData] = useState([]);
    const [publicData, setPublicData] = useState([]);
    const [accessData, setAccessData] = useState([]);
    const [error, setError] = useState(null);


    const userId = sessionStorage.getItem('UserID');

    const handleButtonClick = (data) => {
        // TODO: Add the logic for the API call here
        console.log(data);
    };

    useEffect(() => {
        const fetchData = async (url, setter) => {
            try {
                const response = await fetch(url);
                if (!response.ok) {
                    throw new Error(`Error: ${response.statusText}`);
                }
                const data = await response.json();
                if (Array.isArray(data)) {
                    setter(data);
                } else {
                    throw new Error('Fetched data is not an array');
                }
            } catch (err) {
                setError(err.message);
            }
        };

        fetchData(`http://localhost:8080/getAllPrivateDataByID?user_id=${userId}`, setPrivateData);
        fetchData(`http://localhost:8080/getAllPublicDataByID?user_id=${userId}`, setPublicData);
        fetchData(`http://localhost:8080/getAllAccessDatabyID?user_id=${userId}`, setAccessData);
    }, [userId]);

    console.log("Public Data:", publicData);
    return (
        <Container component="main" maxWidth="md">
            <StyledPaper elevation={3}>
                {error && <Alert severity="error">{error}</Alert>}
                <Typography variant="h5" gutterBottom>
                    Private Data
                </Typography>
                {privateData && privateData.map((data, index) => (
                    <Box key={index} mb={2}>
                        <Typography><strong>Capsule:</strong> {data.capsule}</Typography>
                        <Typography><strong>Cipher Text:</strong> {data.cipher_text}</Typography>
                        <Button variant="outlined" color="primary" onClick={() => handleButtonClick(data)}>Decrypt</Button>
                    </Box>
                ))}
                <Divider />
                <Typography variant="h5" gutterBottom style={{ marginTop: '1rem' }}>
                    Public Data
                </Typography>
                {publicData && publicData.length > 0 && publicData.map((data, index) => (
                    <Box key={index} mb={2}>
                        <Typography><strong>Focus Area:</strong> {data.focus_area}</Typography>
                        <Typography><strong>Communities:</strong> {data.communities}</Typography>
                    </Box>
                ))}
                <Divider />
                <Typography variant="h5" gutterBottom style={{ marginTop: '1rem' }}>
                    Access Data
                </Typography>
                {accessData.length > 0 && accessData.map((data, index) => (
                    <Box key={index} mb={2}>
                        <Typography><strong>Capsule:</strong> {data.capsule}</Typography>
                        <Typography><strong>Cipher Text:</strong> {data.cipher_text}</Typography>
                        <Typography><strong>User ID:</strong> {data.user_id}</Typography>
                        <Button variant="outlined" color="primary" onClick={() => handleButtonClick(data)}>Decrypt</Button>
                    </Box>
                ))}
            </StyledPaper>
        </Container>
    );
}

export default DataPage;
