import React, { useEffect, useState } from 'react';
import { Typography, TextField, Button, IconButton, InputAdornment, Box } from '@mui/material';
import VisibilityOutlinedIcon from '@mui/icons-material/VisibilityOutlined';
import VisibilityOffOutlinedIcon from '@mui/icons-material/VisibilityOffOutlined';
import axios from 'axios';

const ChangePasswordPage = () => {
  const [formData, setFormData] = useState({
    username: '',
    password: '',
    new_password: ''
  });
  const [responseMessage, setResponseMessage] = useState('');
  const [accessToken, setAccessToken] = useState('');
  const [showPassword, setShowPassword] = useState(false);
  const [showNewPassword, setShowNewPassword] = useState(false);

  useEffect(() => {
    const fetchProfileDetails = async () => {
      try {
        const accessToken = localStorage.getItem('AccessToken');
        setAccessToken(accessToken);
        const response = await axios.get('http://localhost:8001/profil/', {
          headers: {
            'Authentication': `${accessToken}`
          }
        });
        console.log(response.data);
      } catch (error) {
        console.error('Error:', error);
      }
    };

    fetchProfileDetails();
  }, []);

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch('http://localhost:8001/profil/pass/', {
        method: 'PUT',
        headers: {
          'Authentication': `${accessToken}`,
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(formData)
      });
      if (response.ok) {
        await response.json();
        setResponseMessage('Password changed successfully!');
      } else {
        const errorMessage = await response.json();
        setResponseMessage(errorMessage.message);
      }
    } catch (error) {
      console.error('Error:', error);
    }
  };

  const handleChange = (e) => {
    const { id, value } = e.target;
    setFormData({ ...formData, [id]: value });
  };

  const togglePasswordVisibility = () => {
    setShowPassword((prevShowPassword) => !prevShowPassword);
  };

  const toggleNewPasswordVisibility = () => {
    setShowNewPassword((prevShowNewPassword) => !prevShowNewPassword);
  };

  return (
    <div>
      <Typography variant="h3">Change Password</Typography>
      <form onSubmit={handleSubmit}>
        <Box display="flex" flexDirection="column" alignItems="left">
          <TextField
            id="username"
            label="Username"
            variant="outlined"
            value={formData.username}
            onChange={handleChange}
            style={{ margin: '20px 0', width:'30%' }}
          />
          <TextField
            id="password"
            label="Current Password"
            type={showPassword ? 'text' : 'password'}
            variant="outlined"
            value={formData.password}
            onChange={handleChange}
            style={{ margin: '10px 0', width:'30%' }}
            InputProps={{
              endAdornment: (
                <InputAdornment position="end">
                  <IconButton onClick={togglePasswordVisibility} edge="end">
                    {showPassword ? <VisibilityOutlinedIcon /> : <VisibilityOffOutlinedIcon />}
                  </IconButton>
                </InputAdornment>
              )
            }}
          />
          <TextField
            id="new_password"
            label="New Password"
            type={showNewPassword ? 'text' : 'password'}
            variant="outlined"
            value={formData.new_password}
            onChange={handleChange}
            style={{ margin: '10px 0', width:'30%' }}
            InputProps={{
              endAdornment: (
                <InputAdornment position="end">
                  <IconButton onClick={toggleNewPasswordVisibility} edge="end">
                    {showNewPassword ? <VisibilityOutlinedIcon /> : <VisibilityOffOutlinedIcon />}
                  </IconButton>
                </InputAdornment>
              )
            }}
          />
          <Button
            type="submit"
            variant="contained"
            color="primary"
            style={{ marginTop: '10px',width:'30%' }}
          >
            Change Password
          </Button>
        </Box>
      </form>
      {responseMessage && <Typography style={{margin:'10px 0'}} color="error">{responseMessage}</Typography>}
    </div>
  );
};

export default ChangePasswordPage;
