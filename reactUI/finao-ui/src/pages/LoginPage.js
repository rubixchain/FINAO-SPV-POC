// LoginPage.js
import React, { useState } from 'react';
import { Container, Typography, TextField, Button, Box, Snackbar, LinearProgress } from '@mui/material';
import { useNavigate } from 'react-router-dom';

function LoginPage() {
  const [isRegistering, setIsRegistering] = useState(false);
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [name, setName] = useState('');
  const [dob, setDob] = useState('');
  const [phone, setPhone] = useState('');
  const [openSnackbar, setOpenSnackbar] = useState(false);
  const [snackbarMessage, setSnackbarMessage] = useState('');
  const navigate = useNavigate();


  const handleApiCall = async (url, body) => {
    try {
      const response = await fetch(url, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(body),
      });

      if (!response.ok) {
        throw new Error('Network response was not ok');
      }

      const data = await response.json();
      return data;
    } catch (error) {
      console.error('There was a problem with the fetch operation:', error.message);
      setSnackbarMessage('An error occurred. Please try again.');
      setOpenSnackbar(true);
    }
  };

  const handleSignup = async () => {
    const data = await handleApiCall('http://localhost:8080/signup', {
      date_of_birth: dob,
      email,
      name,
      password,
      phone_number: phone,
    });

    if (data && data.UserID !== 0) {
      sessionStorage.setItem('UserID', data.UserID);
      setSnackbarMessage('User registered successfully');
      setOpenSnackbar(true);
      setTimeout(() => {
        navigate('/focus');
      }, 2000);
    } else if (data) {
      setSnackbarMessage(data.message);
      setOpenSnackbar(true);
    }
  };

  const handleLogin = async () => {
    const data = await handleApiCall('http://localhost:8080/login', {
      email,
      password,
    });

    if (data && data.UserID !== 0) {
      sessionStorage.setItem('UserID', data.UserID);
      setSnackbarMessage('User authenticated successfully');
      setOpenSnackbar(true);
      setTimeout(() => {
        navigate('/focus');
      }, 2000);
    } else if (data) {
      setSnackbarMessage(data.message);
      setOpenSnackbar(true);
    }
  };

  

  return (
    <Container component="main" maxWidth="xs">
      <Box
        display="flex"
        flexDirection="column"
        justifyContent="center"
        alignItems="center"
        minHeight="100vh"
      >
        <Typography variant="h5" align="center">
          {isRegistering ? 'Register' : 'Login'}
        </Typography>
        {isRegistering ? (
          
        <>
        <TextField margin="normal" fullWidth label="Name" variant="outlined" value={name} onChange={(e) => setName(e.target.value)} />
        <TextField margin="normal" fullWidth label="Email" variant="outlined" value={email} onChange={(e) => setEmail(e.target.value)} />
        <TextField margin="normal" fullWidth label="Password" type="password" variant="outlined" value={password} onChange={(e) => setPassword(e.target.value)} />
        <TextField margin="normal" fullWidth label="Date of Birth" type="date" InputLabelProps={{ shrink: true }} variant="outlined" value={dob} onChange={(e) => setDob(e.target.value)} />
        <TextField label="Mobile No" variant="outlined" fullWidth margin="normal" value={phone} onChange={(e) => setPhone(e.target.value)} />
        <Button variant="contained" color="primary" fullWidth style={{ marginTop: '1rem' }} onClick={handleSignup}>
          Register
        </Button>
            <Button color="secondary" fullWidth style={{ marginTop: '1rem' }} onClick={() => setIsRegistering(false)}>
              Already have an account? Login
            </Button>
          </>
        ) : (
          <>
            <TextField margin="normal" fullWidth label="Email" variant="outlined" value={email} onChange={(e) => setEmail(e.target.value)} />
          <TextField margin="normal" fullWidth label="Password" type="password" variant="outlined" value={password} onChange={(e) => setPassword(e.target.value)} />
          <Button variant="contained" color="primary" fullWidth style={{ marginTop: '1rem' }} onClick={handleLogin}>
              Login
            </Button>
            <Button color="secondary" fullWidth style={{ marginTop: '1rem' }} onClick={() => setIsRegistering(true)}>
              Don't have an account? Register
            </Button>
          </>
        )}
      </Box>
      <Snackbar
        open={openSnackbar}
        autoHideDuration={2000}
        onClose={() => setOpenSnackbar(false)}
        message={snackbarMessage}
      />
    </Container>
  );
}

export default LoginPage;
