import React, { useState } from 'react';
import { Box, Typography, Button, Grid, InputAdornment, IconButton } from '@mui/material';
import axios from 'axios';
import CustomTextField from '../../../components/forms/theme-elements/CustomTextField';
import { Stack } from '@mui/system';
import VisibilityOutlinedIcon from '@mui/icons-material/VisibilityOutlined';
import VisibilityOffOutlinedIcon from '@mui/icons-material/VisibilityOffOutlined';

const AuthRegister = ({ title, subtitle, subtext }) => {
    const [formData, setFormData] = useState({
        userName: '',
        password: '',
        email: '',
        name: '',
        surname: '',
        phone: '',
        birthDate: ''
    });
    const [error, setError] = useState('');
    const [showPassword, setShowPassword] = useState(false);
    const togglePasswordVisibility = () => {
        setShowPassword((prevShowPassword) => !prevShowPassword);
    };

    const handleInputChange = (e) => {
        const { id, value } = e.target;
        setFormData({ ...formData, [id]: value });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const response = await axios.post('http://localhost:8001/user/add', formData);
            console.log('Registration successful:', response.data);
            // Add any further actions upon successful registration
            // Yönlendirme işlemi
            window.location.href = "/auth/login";
        } catch (error) {
            console.error('Error during registration:', error);
            if (error.response && error.response.data && error.response.data.message) {
                setError(error.response.data.message);
            } else {
                setError('An error occurred during registration.');
            }
        }
    };

    return (
      <>
          {title ? (
            <Typography fontWeight="700" variant="h2" mb={1}>
                {title}
            </Typography>
          ) : null}

          {subtext}

          <Box>
              <form onSubmit={handleSubmit}>
                  <Grid container spacing={2}>
                      <Grid item xs={12} md={6}>
                          <Stack mb={3}>
                              <Typography variant="subtitle1" fontWeight={600} component="label" htmlFor='userName' mb="5px">Username</Typography>
                              <CustomTextField id="userName" variant="outlined" fullWidth onChange={handleInputChange} value={formData.userName} />

                              <Typography variant="subtitle1" fontWeight={600} component="label" htmlFor='password' mb="5px">Password</Typography>
                              <CustomTextField id="password" variant="outlined" fullWidth type={showPassword ? 'text' : 'password'} InputProps={{
                                  endAdornment: (
                                    <InputAdornment position="end">
                                        <IconButton onClick={togglePasswordVisibility} edge="end">
                                            {showPassword ? <VisibilityOutlinedIcon /> : <VisibilityOffOutlinedIcon />}
                                        </IconButton>
                                    </InputAdornment>
                                  )
                              }} onChange={handleInputChange} value={formData.password} />

                              <Typography variant="subtitle1" fontWeight={600} component="label" htmlFor='email' mb="5px">Email Address</Typography>
                              <CustomTextField id="email" variant="outlined" fullWidth onChange={handleInputChange} value={formData.email} />
                          </Stack>
                      </Grid>
                      <Grid item xs={12} md={6}>
                          <Stack mb={3}>
                              <Typography variant="subtitle1" fontWeight={600} component="label" htmlFor='name' mb="5px">Name</Typography>
                              <CustomTextField id="name" variant="outlined" fullWidth onChange={handleInputChange} value={formData.name} />

                              <Typography variant="subtitle1" fontWeight={600} component="label" htmlFor='surname' mb="5px">Surname</Typography>
                              <CustomTextField id="surname" variant="outlined" fullWidth onChange={handleInputChange} value={formData.surname} />

                              <Typography variant="subtitle1" fontWeight={600} component="label" htmlFor='phone' mb="5px">Phone</Typography>
                              <CustomTextField id="phone" variant="outlined" fullWidth onChange={handleInputChange} value={formData.phone} />

                              <Typography variant="subtitle1" fontWeight={600} component="label" htmlFor='birthDate' mb="5px">Birth Date</Typography>
                              <CustomTextField id="birthDate" variant="outlined" fullWidth onChange={handleInputChange} value={formData.birthDate} />
                          </Stack>
                      </Grid>
                  </Grid>
                  {error && (
                    <Typography variant="body1" color="error" gutterBottom>
                        {error}
                    </Typography>
                  )}
                  <Button color="primary" variant="contained" size="large" fullWidth type="submit">
                      Sign Up
                  </Button>
              </form>
          </Box>
          {subtitle}
      </>
    );
};

export default AuthRegister;
