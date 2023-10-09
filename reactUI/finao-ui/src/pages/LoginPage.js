// LoginPage.js
import React, { useState } from 'react';
import { Container, Typography, TextField, Button, Box } from '@mui/material';

function LoginPage() {
  const [isRegistering, setIsRegistering] = useState(false);

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
            <TextField margin="normal" fullWidth label="Name" variant="outlined" />
            <TextField margin="normal" fullWidth label="Email" variant="outlined" />
            <TextField margin="normal" fullWidth label="Password" type="password" variant="outlined" />
            <TextField margin="normal" fullWidth label="Date of Birth" type="date" InputLabelProps={{ shrink: true }} variant="outlined" />
            <TextField
                label="Mobile No"
                variant="outlined"
                fullWidth
                margin="normal"
              />
            <Button variant="contained" color="primary" fullWidth style={{ marginTop: '1rem' }}>
              Register
            </Button>
            <Button color="secondary" fullWidth style={{ marginTop: '1rem' }} onClick={() => setIsRegistering(false)}>
              Already have an account? Login
            </Button>
          </>
        ) : (
          <>
            <TextField margin="normal" fullWidth label="Email" variant="outlined" />
            <TextField margin="normal" fullWidth label="Password" type="password" variant="outlined" />
            <Button variant="contained" color="primary" fullWidth style={{ marginTop: '1rem' }}>
              Login
            </Button>
            <Button color="secondary" fullWidth style={{ marginTop: '1rem' }} onClick={() => setIsRegistering(true)}>
              Don't have an account? Register
            </Button>
          </>
        )}
      </Box>
    </Container>
  );
}

export default LoginPage;
