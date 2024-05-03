import React, { useEffect, useState } from 'react';
import { Typography, TextField, Button } from '@mui/material';
import axios from 'axios';

const ChangePasswordPage = () => {
  const [formData, setFormData] = useState({
    username: '',
    password: '',
    new_password: ''
  });
  const [responseMessage, setResponseMessage] = useState('');
  const [accessToken, setAccessToken] = useState('');

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
        const errorMessage = await response.text();
        setResponseMessage(errorMessage);
      }
    } catch (error) {
      console.error('Error:', error);
    }
  };

  const handleChange = (e) => {
    const { id, value } = e.target;
    setFormData({ ...formData, [id]: value });
  };

  return (
    <div>
      <Typography variant="h3">Change Password</Typography>
      <form onSubmit={handleSubmit}>
        <TextField
          id="username"
          label="Username"
          variant="outlined"
          fullWidth
          value={formData.username}
          onChange={handleChange}
          style={{ margin: '10px 0' }}
        />
        <TextField
          id="password"
          label="Current Password"
          type="password"
          variant="outlined"
          fullWidth
          value={formData.password}
          onChange={handleChange}
          style={{ margin: '10px 0' }}
        />
        <TextField
          id="new_password"
          label="New Password"
          type="password"
          variant="outlined"
          fullWidth
          value={formData.new_password}
          onChange={handleChange}
          style={{ margin: '10px 0' }}
        />
        <Button
          type="submit"
          variant="contained"
          color="primary"
          style={{ marginTop: '10px' }}
        >
          Change Password
        </Button>
      </form>
      {responseMessage && <Typography>{responseMessage}</Typography>}
    </div>
  );
};

export default ChangePasswordPage;
