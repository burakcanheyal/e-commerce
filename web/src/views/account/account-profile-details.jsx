import React, { useCallback, useEffect, useState } from 'react';
import {
  Box,
  Button,
  Card,
  CardActions,
  CardContent,
  CardHeader,
  Divider,
  TextField,
  Grid,
  IconButton,
  InputAdornment, Typography,
} from '@mui/material';
import VisibilityOutlinedIcon from '@mui/icons-material/VisibilityOutlined';
import VisibilityOffOutlinedIcon from '@mui/icons-material/VisibilityOffOutlined';
import axios from 'axios';

const AccountProfileDetails = () => {
  const [values, setValues] = useState({
    username: '',
    name: '',
    surname: '',
    birthdate: '',
    phone: '',
    email: '',
    password: ''
  });
  const [showPassword, setShowPassword] = useState(false);
  const [errors, setErrors] = useState({});

  const handleChange = useCallback(
    (event) => {
      setValues((prevState) => ({
        ...prevState,
        [event.target.name]: event.target.value
      }));
    },
    []
  );

  const handleSubmit = useCallback(
    async (event) => {
      event.preventDefault();
      try {
        const accessToken = localStorage.getItem('AccessToken');
        const response = await fetch('http://localhost:8001/profil/', {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
            'Authentication': `${accessToken}`
          },
          body: JSON.stringify({
            userName: values.username,
            name: values.name,
            surname: values.surname,
            birthDate: values.birthdate,
            phone: values.phone,
            email: values.email,
            password: values.password
          })
        });
        if (!response.ok) {
          const errorData = await response.json();
          setErrors(errorData);
        } else {
          const data = await response.json();
          console.log(data);
        }
      } catch (error) {
        console.error('Profil bilgileri güncellenirken hata oluştu:', error);
      }
    },
    [values]
  );

  const togglePasswordVisibility = () => {
    setShowPassword((prevShowPassword) => !prevShowPassword);
  };

  useEffect(() => {
    const fetchProfileDetails = async () => {
      try {
        const accessToken = localStorage.getItem('AccessToken');
        const response = await axios.get('http://localhost:8001/profil/', {
          headers: {
            'Authentication': `${accessToken}`
          }
        });
        const profileData = response.data;
        setValues({
          username: profileData.username,
          name: profileData.name,
          surname: profileData.surname,
          birthdate: profileData.birthdate,
          phone: profileData.phone || '',
          email: profileData.email
        });
        console.log(response.data);
      } catch (error) {
        console.error('Hesap detayları alınırken hata oluştu:', error);
      }
    };

    fetchProfileDetails();
  }, []);

  return (
    <form autoComplete="off" noValidate onSubmit={handleSubmit}>
      <Card>
        <CardHeader title="Profile Details" />
        <CardContent sx={{ pt: 2 }}>
          <Box sx={{ m: 1.0 }}>
            <Grid container spacing={3}>
              <Grid xs={12} md={4}>
                <TextField
                  fullWidth
                  label="Username"
                  name="username"
                  onChange={handleChange}
                  required
                  value={values.username}
                />
              </Grid>
              <Grid xs={12} md={6}>
                <TextField
                  fullWidth
                  label="Name"
                  name="name"
                  onChange={handleChange}
                  required
                  value={values.name}
                />
              </Grid>
              <Grid xs={12} md={6}>
                <TextField
                  fullWidth
                  label="Surname"
                  name="surname"
                  onChange={handleChange}
                  required
                  value={values.surname}
                />
              </Grid>
              <Grid xs={12} md={6}>
                <TextField
                  fullWidth
                  label="Birth Date"
                  name="birthdate"
                  onChange={handleChange}
                  required
                  value={values.birthdate}
                />
              </Grid>
              <Grid xs={12} md={6}>
                <TextField
                  fullWidth
                  label="Phone Number"
                  name="phone"
                  onChange={handleChange}
                  type="string"
                  value={values.phone}
                />
              </Grid>
              <Grid xs={12} md={6}>
                <TextField
                  fullWidth
                  label="E-mail Address"
                  name="email"
                  onChange={handleChange}
                  required
                  value={values.email}
                />
              </Grid>
              <Grid xs={12} md={6}>
                <TextField
                  fullWidth
                  label="Password"
                  name="password"
                  onChange={handleChange}
                  required
                  type={showPassword ? 'text' : 'password'}
                  value={values.password}
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
              </Grid>
            </Grid>
            {Object.keys(errors).length > 0 && (
              <Box mt={2}>
                <Typography color="error">
                  {Object.values(errors).join(', ')}
                </Typography>
              </Box>
            )}
          </Box>
        </CardContent>
        <Divider />
        <CardActions sx={{ justifyContent: 'flex-end' }}>
          <Button type="submit" variant="contained">Update Details</Button>

        </CardActions>
      </Card>
    </form>
  );
};

export default AccountProfileDetails;
