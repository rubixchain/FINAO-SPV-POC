// ViewPage.js
import React from 'react';
import { Container, Typography, Box } from '@mui/material';

function ViewPage() {
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
          View Page Content
        </Typography>
        {/* Add any other components or content you want for the ViewPage here */}
      </Box>
    </Container>
  );
}

export default ViewPage;
