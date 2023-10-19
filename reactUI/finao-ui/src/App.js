import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import LoginPage from './pages/LoginPage';
import FocusPage from './pages/FocusPage';
import ViewPage from './pages/ViewPage';
import NavBar from './pages/NavBar';
import HomePage from './pages/HomePage';
import DataPage from './pages/DataPage';

function App() {
  return (
    <Router>
      <NavBar />
      <Routes>
        <Route path="/" element={<HomePage/>} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/focus" element={<FocusPage />} />
        <Route path="/view" element={<ViewPage />} />
        <Route path='/data' element={<DataPage />} />
      </Routes>
    </Router>
  );
}

export default App;
