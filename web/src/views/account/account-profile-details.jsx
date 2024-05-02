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
    email: ''
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
    (event) => {
      event.preventDefault();
      // Form gönderimini burada işleyebilirsiniz
    },
    []
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
          birthdate: profileData.birth_date,
          phone: profileData.phone || '',
          email: profileData.email
        });
        console.log(response.data)
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
            </Grid>
          </Box>
        </CardContent>
        <Divider />
        <CardActions sx={{ justifyContent: 'flex-end' }}>
          <Button variant="contained">Detayları Kaydet</Button>
        </CardActions>
      </Card>
    </form>
  );
};

export default AccountProfileDetails;
