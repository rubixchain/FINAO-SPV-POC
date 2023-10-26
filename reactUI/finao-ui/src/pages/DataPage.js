import React, { useState, useEffect } from 'react';
import { Container, Typography, Paper, Box, Divider, Alert, Button } from '@mui/material';
import { styled } from '@mui/system';

import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

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
    const [decryptedPrivateData, setDecryptedPrivateData] = useState({});
    const [decryptedAccessData, setDecryptedAccessData] = useState({});
    const [error, setError] = useState(null);



    const userId = sessionStorage.getItem('UserID');

    const handleButtonClick = async (data, index, section) => {
        try {
            const response = await fetch('http://localhost:8080/decryptData', {
                method: 'POST',
                headers: {
                    'accept': 'application/json',
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    capsule: data.capsule,
                    ciphertext: data.cipher_text,
                    user_id: parseInt(userId, 10),
                })
            });

            if (response.status === 200) {
                const result = await response.json();
                console.log(result);
                if (section === 'private') {
                    setDecryptedPrivateData(prevState => ({ ...prevState, [index]: result }));
                } else if (section === 'access') {
                    setDecryptedAccessData(prevState => ({ ...prevState, [index]: result }));
                }
            } else {
                toast.error('You don\'t have enough permission to decrypt the data');
                /* setError('You don\'t have enough permission to decrypt the data'); */
            }
        } catch (err) {
            console.error(err);
            toast.error('Error decrypting data');
            /* setError('Error decrypting data');
 */
        }
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
    console.log("access Data:", accessData);
    return (
        <Container component="main" maxWidth="md">
            <StyledPaper elevation={3}>
                {error && <Alert severity="error">{error}</Alert>}

                <Typography variant="h5" gutterBottom>
                    My Private Data
                </Typography>

                {privateData && privateData.map((data, index) => (
                    <Box key={index} mb={2}>
                        {decryptedPrivateData[index] ? (<>
                            <Typography><strong>Focus Area:</strong> {decryptedPrivateData[index].focus_area}</Typography>
                            <Typography><strong>Communities:</strong> {decryptedPrivateData[index].communities}</Typography>
                        </>
                        ) : (
                            <>
                                <Typography><strong>Capsule:</strong> {data.capsule}</Typography>
                                <Typography><strong>Cipher Text:</strong> {data.cipher_text}</Typography>
                                <Button variant="outlined" color="primary" onClick={() => handleButtonClick(data, index, 'private')}>Decrypt</Button>                            </>
                        )}
                    </Box>
                ))}

                <Divider />

                <Typography variant="h5" gutterBottom style={{ marginTop: '1rem' }}>
                    My Public Data
                </Typography>

                {publicData && publicData.length > 0 && publicData.map((data, index) => (
                    <Box key={index} mb={2}>
                        <Typography><strong>Focus Area:</strong> {data.focus_area}</Typography>
                        <Typography><strong>Communities:</strong> {data.communities}</Typography>
                    </Box>
                ))}

                <Divider />

                <Typography variant="h5" gutterBottom style={{ marginTop: '1rem' }}>
                    Data Access Given/Received
                </Typography>

                {accessData.length > 0 && accessData.map((data, index) => (
                    <Box key={index} mb={2}>
                        {decryptedAccessData[index] ? (<>
                            <Typography><strong>Focus Area:</strong> {decryptedAccessData[index].focus_area}</Typography>
                            <Typography><strong>Communities:</strong> {decryptedAccessData[index].communities}</Typography>
                        </>
                        ) : (
                            <>
                                <Typography><strong>Capsule:</strong> {data.capsule}</Typography>
                                <Typography><strong>Cipher Text:</strong> {data.cipher_text}</Typography>
                                {/* <Typography><strong>User ID:</strong> {data.user_id}</Typography> */}
                                {data.access_type === "Access Given" ? (<Typography><strong> Access given to User ID:</strong> {data.decrypt_user_id}</Typography>) : (<Typography><strong>Owner User ID:</strong> {data.OwnerUserID}</Typography>)}
                                <Typography><strong>Access Type:</strong> {data.access_type}</Typography>
                                <Button variant="outlined" color="primary" onClick={() => handleButtonClick(data, index, 'access')}>Decrypt</Button>                            </>
                        )}
                    </Box>
                ))}
            </StyledPaper>
            <ToastContainer />
        </Container>
    );
}

export default DataPage;
