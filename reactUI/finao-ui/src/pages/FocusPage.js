import React, { useState } from 'react';
import { Button, TextField, Container, Typography, Select, MenuItem, FormControl, InputLabel, Checkbox, ListItemText, Paper, Grid } from '@mui/material';
import { styled } from '@mui/system';

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

    return (
        <Container component="main" maxWidth="sm">
            <StyledPaper elevation={3}>
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
                                    {privateData[area] && 
                                    <Grid item xs={4}>
                                        <TextField
                                            fullWidth
                                            placeholder="Whom to share data"
                                            onChange={(e) => handleUserIdChange(area, e.target.value)}
                                            onClick={(e) => e.stopPropagation()}
                                            onKeyDown={(e) => e.stopPropagation()}
                                        />
                                    </Grid>}
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
                                    {privateData[community] && 
                                    <Grid item xs={4}>
                                        <TextField
                                            fullWidth
                                            placeholder="Whom to share data"
                                            onChange={(e) => handleUserIdChange(community, e.target.value)}
                                            onClick={(e) => e.stopPropagation()}
                                            onKeyDown={(e) => e.stopPropagation()}
                                        />
                                    </Grid>}
                                </Grid>
                            </MenuItem>
                        ))}
                    </Select>
                </FormControl>

                <Button variant="contained" color="primary" fullWidth style={{ marginTop: '1rem' }}>Submit</Button>
                <Button variant="outlined" color="secondary" fullWidth style={{ marginTop: '1rem' }} onClick={handleCancel}>Cancel</Button>
            </StyledPaper>
        </Container>
    );
}

export default FocusPage;
