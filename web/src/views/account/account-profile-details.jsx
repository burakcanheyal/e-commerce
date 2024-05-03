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
  Grid
} from '@mui/material';
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
        const data = await response.json();
        console.log(data);
      } catch (error) {
        console.error('Profil bilgileri güncellenirken hata oluştu:', error);
      }
    },
    [values]
  );
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
        <CardHeader title="Detaylar" />
        <CardContent sx={{ pt: 0 }}>
          <Box sx={{ m: -0.5 }}>
            <Grid container spacing={2}>
              <Grid xs={12} md={6}>
                <TextField
                  fullWidth
                  helperText="Kullanıcı adınızı belirtin"
                  label="Kullanıcı Adı"
                  name="username"
                  onChange={handleChange}
                  required
                  value={values.username}
                />
              </Grid>
              <Grid xs={12} md={6}>
                <TextField
                  fullWidth
                  label="Ad"
                  name="name"
                  onChange={handleChange}
                  required
                  value={values.name}
                />
              </Grid>
              <Grid xs={12} md={6}>
                <TextField
                  fullWidth
                  label="Soyad"
                  name="surname"
                  onChange={handleChange}
                  required
                  value={values.surname}
                />
              </Grid>
              <Grid xs={12} md={6}>
                <TextField
                  fullWidth
                  label="Doğum Tarihi"
                  name="birthdate"
                  onChange={handleChange}
                  required
                  value={values.birthdate}
                />
              </Grid>
              <Grid xs={12} md={6}>
                <TextField
                  fullWidth
                  label="Telefon Numarası"
                  name="phone"
                  onChange={handleChange}
                  type="string"
                  value={values.phone}
                />
              </Grid>
              <Grid xs={12} md={6}>
                <TextField
                  fullWidth
                  label="Email Adresi"
                  name="email"
                  onChange={handleChange}
                  required
                  value={values.email}
                />
              </Grid>
              <Grid xs={12} md={6}>
                <TextField
                  fullWidth
                  label="Şifre"
                  name="password"
                  onChange={handleChange}
                  required
                  value={values.password}
                />
              </Grid>
            </Grid>
          </Box>
        </CardContent>
        <Divider />
        <CardActions sx={{ justifyContent: 'flex-end' }}>
          <Button type="submit" variant="contained">Detayları Kaydet</Button>
        </CardActions>
      </Card>
    </form>
  );
};

export default AccountProfileDetails;
