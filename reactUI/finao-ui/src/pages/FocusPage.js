import React, { useState } from 'react';
import { Button, TextField, Container, Typography, Select, MenuItem, FormControl, InputLabel, Checkbox, ListItemText, Paper, Grid } from '@mui/material';
import { styled } from '@mui/system';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

const StyledPaper = styled(Paper)(({ theme }) => ({
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    padding: theme.spacing(3),
}));


const FocusPage = () => {
    const [selectedFocusAreas, setSelectedFocusAreas] = useState([]);
    const [selectedCommunities, setSelectedCommunities] = useState([]);
    const [privateData, setPrivateData] = useState({});
    const [userIdData, setUserIdData] = useState({});
    const [whomToShareWith, setWhomToShareWith] = useState('');
    const userId = parseInt(sessionStorage.getItem('UserID'),10);

    console.log(sessionStorage.getItem('UserID'),10)

    const handlePrivateChange = (event, name) => {
        event.stopPropagation();
        setPrivateData(prev => ({ ...prev, [name]: !prev[name] }));
    };

    const handleUserIdChange = (name, value) => {
        setUserIdData(prev => ({ ...prev, [name]: value }));
    };

    const handleCancel = () => {
        window.location.reload();
    };

    const hasPrivateDataSelected = () => {
        return Object.values(privateData).some(value => value === true);
    };

    const generateCurlCommand = (url, method, headers, body) => {
        let curlCmd = `curl -X '${method}' '${url}'`;
    
        for (const header in headers) {
            curlCmd += ` -H '${header}: ${headers[header]}'`;
        }
    
        if (body) {
            curlCmd += ` -d '${JSON.stringify(body)}'`;
        }
    
        return curlCmd;
    };
    

    const addPublicData = async () => {
        const apiUrl = 'http://localhost:8080/addPublicData';
        console.log(userId)
        const apiHeaders = {
            'Content-Type': 'application/json',
        };
        const apiBody = {
            communities: selectedCommunities.join(', '),
            focus_area: selectedFocusAreas.join(', '),
            user_id: userId,
        };
    
        // Print the curl command
        console.log(generateCurlCommand(apiUrl, 'POST', apiHeaders, apiBody));
    
        const response = await fetch(apiUrl, {
            method: 'POST',
            headers: apiHeaders,
            body: JSON.stringify(apiBody),
        });
        
        const rawResponse = await response.text();
        console.log("Raw Response:", rawResponse);
        
        try {
            const data = JSON.parse(rawResponse);
            return data;
        } catch (error) {
            console.error("Error parsing JSON:", error);
        }
    };
    
    const addPrivateData = async () => {
        const apiUrl = 'http://localhost:8080/addPrivateData';
        const apiHeaders = {
            'Content-Type': 'application/json',
            'accept': 'application/json'
        };
        const apiBody = {
            communities: selectedCommunities.join(', '),
            focus_area: selectedFocusAreas.join(', '),
            decrypt_user_id: parseInt(whomToShareWith, 10), 
            user_id: userId
        };
    
        // Print the curl command
        console.log(generateCurlCommand(apiUrl, 'POST', apiHeaders, apiBody));
    
        const response = await fetch(apiUrl, {
            method: 'POST',
            headers: apiHeaders,
            body: JSON.stringify(apiBody),
        });
    
        const data = await response.json();
        return data;
    };
    
    const handleSubmit = async () => {
        try {
            let response;
            if (hasPrivateDataSelected()) {
                response = await addPrivateData();
                if (response && response.status) {
                    toast.success('Private data added successfully');
                } else {
                    toast.error('Error adding private data');
                }
            } else {
                response = await addPublicData();
                if (response && response.status) {
                    toast.success('Public data added successfully');
                } else {
                    toast.error('Error adding public data');
                }
            }
        } catch (error) {
            toast.error('An unexpected error occurred');
        }
    };

      return (
        <Container component="main" maxWidth="sm">
            <StyledPaper elevation={3}>
            <Typography variant="a" align="center" gutterBottom>By default data is public, click on check box to make it private</Typography>
                <Typography variant="h4" align="center" gutterBottom>Select Focus Area</Typography>
                <FormControl fullWidth margin="normal">
                    <InputLabel>Focus Area</InputLabel>
                    <Select
                        multiple
                        value={selectedFocusAreas}
                        onChange={(event) => setSelectedFocusAreas(event.target.value)}
                        MenuProps={{
                            anchorOrigin: {
                                vertical: "bottom",
                                horizontal: "left"
                            },
                            transformOrigin: {
                                vertical: "top",
                                horizontal: "left"
                            },
                            getContentAnchorEl: null
                        }}
                    >
                        {['Aquatics', 'Cooking', 'Sports', 'Mental Health', 'Wealth Management'].map(area => (
                            <MenuItem key={area} value={area}>
                                <Grid container alignItems="center" spacing={2}>
                                    <Grid item>
                                        <Checkbox checked={privateData[area] === true} onClick={(e) => handlePrivateChange(e, area)} />
                                    </Grid>
                                    <Grid item xs>
                                        <ListItemText primary={area} />
                                    </Grid>
                                </Grid>
                            </MenuItem>
                        ))}
                    </Select>
                </FormControl>

                <Typography variant="h4" align="center" style={{ marginTop: '2rem' }} gutterBottom>Select Communities</Typography>
                <FormControl fullWidth margin="normal">
                    <InputLabel>Communities</InputLabel>
                    <Select
                        multiple
                        value={selectedCommunities}
                        onChange={(event) => setSelectedCommunities(event.target.value)}
                        MenuProps={{
                            anchorOrigin: {
                                vertical: "bottom",
                                horizontal: "left"
                            },
                            transformOrigin: {
                                vertical: "top",
                                horizontal: "left"
                            },
                            getContentAnchorEl: null
                        }}
                    >
                        {['Coffee Club', 'Cooking Club', 'Cardio Club', 'Hiking Club', 'New Parenting', 'Finance 101'].map(community => (
                            <MenuItem key={community} value={community}>
                                <Grid container alignItems="center" spacing={2}>
                                    <Grid item>
                                        <Checkbox checked={privateData[community] === true} onClick={(e) => handlePrivateChange(e, community)} />
                                    </Grid>
                                    <Grid item xs>
                                        <ListItemText primary={community} />
                                    </Grid>
                                </Grid>
                            </MenuItem>
                        ))}
                    </Select>
                </FormControl>

                {/* Single "Whom to share" input */}
                {hasPrivateDataSelected() && (
                <TextField
                    fullWidth
                    margin="normal"
                    label="Whom to share data with"
                    value={whomToShareWith}
                    onChange={(e) => setWhomToShareWith(e.target.value)}
                />
                )}

                <Button variant="contained" color="primary" fullWidth style={{ marginTop: '1rem' }} onClick={handleSubmit}>Submit</Button>
                <Button variant="outlined" color="secondary" fullWidth style={{ marginTop: '1rem' }} onClick={handleCancel}>Cancel</Button>
            </StyledPaper>
            <ToastContainer />
        </Container>
    );
}

export default FocusPage;