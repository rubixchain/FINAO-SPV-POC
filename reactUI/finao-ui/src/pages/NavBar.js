import React from 'react';
import { AppBar, Toolbar, Typography, Button, IconButton } from '@mui/material';
import { Home as HomeIcon, ExitToApp as ExitToAppIcon } from '@mui/icons-material';
import { styled } from '@mui/system';
import { useNavigate } from 'react-router-dom';

const StyledAppBar = styled(AppBar)(({ theme }) => ({
    marginBottom: theme.spacing(3),
}));

const Navbar = () => {
    const userId = sessionStorage.getItem('UserID');
    const navigate = useNavigate();

    const handleSignout = () => {
        sessionStorage.removeItem('UserID');
        window.location.reload();
    };

    const handleLogin = () => {
        navigate('/login');
    };

    const navigateHome = () => {
        navigate('/');
    };

    const navigateFocus = () => {
        navigate('/focus');
    };

    return (
        <StyledAppBar position="static">
            <Toolbar>
                <IconButton edge="start" color="inherit" aria-label="home" onClick={navigateHome}>
                    <HomeIcon />
                </IconButton>
                <Typography variant="h6" style={{ flexGrow: 1 }}>
                    Rubix - FINAO PoC
                </Typography>
                <Button color="inherit" onClick={navigateFocus}>Add Data</Button> {/* <-- Add onClick handler */}
                {userId ? (
                    <Button color="inherit" onClick={handleSignout}>
                        <ExitToAppIcon /> Sign Out
                    </Button>
                ) : (
                    <Button color="inherit" onClick={handleLogin}>
                        Login
                    </Button>
                )}
            </Toolbar>
        </StyledAppBar>
    );
};

export default Navbar;
