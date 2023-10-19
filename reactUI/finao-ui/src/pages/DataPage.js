import React, { useState, useEffect } from 'react';
import { Container, Typography, Paper, Box, Divider, Alert } from '@mui/material';
import { styled } from '@mui/system';

const StyledPaper = styled(Paper)(({ theme }) => ({
    marginTop: theme.spacing(4),
    padding: theme.spacing(3),
}));

const DataPage = () => {
    const [privateData, setPrivateData] = useState(null);
    const [publicData, setPublicData] = useState(null);
    const [accessData, setAccessData] = useState([]); // Set initial state to an empty array
    const [error, setError] = useState(null);

    const userId = sessionStorage.getItem('UserID');

    useEffect(() => {
        const fetchData = async (url, setter) => {
            try {
                const response = await fetch(url);
                if (!response.ok) {
                    throw new Error(`Error: ${response.statusText}`);
                }
                const data = await response.json();
                setter(data);
            } catch (err) {
                setError(err.message);
            }
        };

        fetchData(`http://localhost:8080/getAllPrivateDataByID?user_id=${userId}`, setPrivateData);
        fetchData(`http://localhost:8080/getAllPublicDataByID?user_id=${userId}`, setPublicData);
        fetch(`http://localhost:8080/getAllAccessDatabyID?user_id=${userId}`, {
            method: 'GET',
            headers: {
                'accept': 'application/json',
            },
        })
        .then(response => response.json())
        .then(data => {
            if (data && Array.isArray(data)) { // Ensure the data is an array
                setAccessData(data);
            } else {
                setError('Error fetching access data');
            }
        })
        .catch(err => {
            setError('Error fetching access data');
            console.error(err);
        });
    }, [userId]);

    return (
        <Container component="main" maxWidth="md">
            <StyledPaper elevation={3}>
                {error && <Alert severity="error">{error}</Alert>}
                <Typography variant="h5" gutterBottom>
                    Private Data
                </Typography>
                {privateData && (
                    <Box mb={2}>
                        <Typography><strong>Capsule:</strong> {privateData.capsule}</Typography>
                        <Typography><strong>Cipher Text:</strong> {privateData.cipher_text}</Typography>
                    </Box>
                )}
                <Divider />
                <Typography variant="h5" gutterBottom style={{ marginTop: '1rem' }}>
                    Public Data
                </Typography>
                {publicData && ( // Add conditional rendering
                    <Box mb={2}>
                        <Typography><strong>Focus Area:</strong> {publicData.focus_area}</Typography>
                        <Typography><strong>Communities:</strong> {publicData.communities}</Typography>
                    </Box>
                )}
                <Divider />
                <Typography variant="h5" gutterBottom style={{ marginTop: '1rem' }}>
                    Access Data
                </Typography>
                {accessData.length > 0 && accessData.map((data, index) => ( // Check if accessData has items before mapping
                    <Box key={index} mb={2}>
                        <Typography><strong>Capsule:</strong> {data.capsule}</Typography>
                        <Typography><strong>Cipher Text:</strong> {data.cipher_text}</Typography>
                        <Typography><strong>User ID:</strong> {data.user_id}</Typography>
                    </Box>
                ))}
            </StyledPaper>
        </Container>
    );
}

export default DataPage;