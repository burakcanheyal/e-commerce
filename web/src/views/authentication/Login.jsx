import React, { useEffect, useState } from 'react';
import { Link, Navigate } from 'react-router-dom';
import { Grid, Box, Card, Stack, Typography, Button } from '@mui/material';
import { auth, provider } from './auth/AuthLogin.jsx';
import { signInWithPopup } from 'firebase/auth';
import { updateProfile } from 'firebase/auth'; // Firebase'den profil güncelleme işlevini içe aktarın

import AuthLogin from './auth/AuthLogin.jsx';
import PageContainer from "../../components/container/PageContainer.jsx";
import Logo from "../../layouts/full/shared/logo/Logo.jsx";

const Login2 = () => {
    const [userData, setUserData] = useState(null);

    const handleClick = () => {
        signInWithPopup(auth, provider).then((data) => {
            // Google ile giriş yaptıktan sonra kullanıcı verilerini al
            const user = data.user;
            const email = user.email;

            // Firebase'den gelen kullanıcı verilerini kullanarak profil güncelleme işlemini yap
            updateProfile(auth.currentUser, {
                displayName: user.displayName,
                email: user.email,
                // Diğer gerekli bilgileri buraya ekleyebilirsiniz
            }).then(() => {
                setUserData({
                    email: email,
                    // Diğer gerekli bilgileri buraya ekleyebilirsiniz
                });
                localStorage.setItem('email', email);
            }).catch((error) => {
                // Profil güncelleme işlemi başarısız olduysa hata işleme yapabilirsiniz
                console.error('Error updating profile:', error);
            });
        });
    };

    useEffect(() => {
        // Kullanıcının profil bilgilerini localStorage'den al
        const email = localStorage.getItem('email');
        if (email) {
            setUserData({
                email: email,
                // Diğer gerekli bilgileri buradan alabilirsiniz
            });
        }
    }, []);

    if (userData) {
        // Kullanıcı bilgileri mevcutsa, kullanıcıyı ana sayfaya yönlendir
        return <Navigate to="/dashboard" />;
    }

    return (
      <PageContainer title="Login" description="this is Login page">
          <Box
            sx={{
                position: 'relative',
                '&:before': {
                    content: '""',
                    background: 'radial-gradient(#d2f1df, #d3d7fa, #bad8f4)',
                    backgroundSize: '400% 400%',
                    animation: 'gradient 15s ease infinite',
                    position: 'absolute',
                    height: '100%',
                    width: '100%',
                    opacity: '0.3',
                },
            }}
          >
              <Grid container spacing={0} justifyContent="center" sx={{ height: '100vh' }}>
                  <Grid
                    item
                    xs={12}
                    sm={12}
                    lg={4}
                    xl={3}
                    display="flex"
                    justifyContent="center"
                    alignItems="center"
                  >
                      <Card elevation={9} sx={{ p: 4, zIndex: 1, width: '100%', maxWidth: '500px' }}>
                          <Box display="flex" alignItems="center" justifyContent="center">
                              <Logo />
                          </Box>
                          <AuthLogin
                            subtext={
                                <Typography variant="subtitle1" textAlign="center" color="textSecondary" mb={1}>
                                    Your Travel Guide
                                </Typography>
                            }
                            subtitle={
                                <Stack direction="row" spacing={1} justifyContent="center" mt={3}>
                                    <Typography color="textSecondary" variant="h6" fontWeight="500">
                                        New to Exporia?
                                    </Typography>
                                    <Typography
                                      component={Link}
                                      to="/auth/register"
                                      fontWeight="500"
                                      sx={{
                                          textDecoration: 'none',
                                          color: 'primary.main',
                                      }}
                                    >
                                        Create an account
                                    </Typography>
                                </Stack>
                            }
                          />
                          {userData ? (
                            <Button component={Link} to="/home" variant="contained" color="primary" fullWidth>
                                Go to Home
                            </Button>
                          ) : (
                            <Button onClick={handleClick} variant="contained" color="primary" fullWidth>
                                Sign in With Google
                            </Button>
                          )}
                      </Card>
                  </Grid>
              </Grid>
          </Box>
      </PageContainer>
    );
};

export default Login2;
