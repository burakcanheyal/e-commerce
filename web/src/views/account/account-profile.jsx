import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Avatar, Box, Button, Card, CardActions, CardContent, Divider, Typography } from '@mui/material';

const AccountProfile = () => {
  const [userData, setUserData] = useState({
    name: '',
    surname: '',
    country: 'Turkey',
    timezone: 'GMT+3',
    avatar: ''
  });

  useEffect(() => {
    const fetchProfileData = async () => {
      try {
        const accessToken = localStorage.getItem('AccessToken');
        const response = await axios.get('http://localhost:8001/profil/', {
          headers: {
            'Authentication': `${accessToken}`
          }
        });
        const profileData = response.data;
        setUserData({
          name: profileData.name,
          surname: profileData.surname,
          country: 'Turkey',
          timezone: 'GMT+3',
          avatar: profileData.avatar
        });
      } catch (error) {
        console.error('Hesap detayları alınırken hata oluştu:', error);
      }
    };
    fetchProfileData();
  }, []);

  const getInitialLetter = (name) => {
    return name.charAt(0).toUpperCase();
  };

  return (
    <Card>
      <CardContent>
        <Box
          sx={{
            alignItems: 'center',
            display: 'flex',
            flexDirection: 'column',
          }}
        >
          <Avatar
            sx={{
              height: 80,
              width: 80,
              mb: 2,
              color: 'white'
            }}
          >
            {getInitialLetter(userData.name)}
          </Avatar>
          <Typography
            gutterBottom
            variant="h5"
          >
            {userData.name} {userData.surname}
          </Typography>
          <Typography
            color="text.secondary"
            variant="body2"
          >
            {userData.country}
          </Typography>
          <Typography
            color="text.secondary"
            variant="body2"
          >
            {userData.timezone}
          </Typography>
        </Box>
      </CardContent>
      <Divider />
      <CardActions>
        <Button
          fullWidth
          variant="text"
        >
          Upload picture
        </Button>
      </CardActions>
    </Card>
  );
};

export default AccountProfile;
