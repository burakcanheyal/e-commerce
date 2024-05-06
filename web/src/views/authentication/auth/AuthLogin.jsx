import React, { useState } from 'react';
import {
    Box,
    Typography,
    FormGroup,
    FormControlLabel,
    Button,
    Stack,
    Checkbox,
    Dialog,
    DialogTitle,
    DialogContent,
    DialogActions,
    TextField, InputAdornment, IconButton,
} from '@mui/material';
import CustomTextField from '../../../components/forms/theme-elements/CustomTextField';
import VisibilityOutlinedIcon from '@mui/icons-material/VisibilityOutlined';
import VisibilityOffOutlinedIcon from '@mui/icons-material/VisibilityOffOutlined';

const AuthLogin = ({ title, subtitle, subtext, setUserData }) => {
    const [formData, setFormData] = useState({
        username: '',
        password: ''
    });
    const [error, setError] = useState(null);
    const [openDialog, setOpenDialog] = useState(false);
    const [activationData, setActivationData] = useState({
        userName: '',
        code: ''
    });
    const [activationMessage, setActivationMessage] = useState('');
    const [showPassword, setShowPassword] = useState(false);

    const handleInputChange = (e) => {
        const { id, value } = e.target;
        setFormData({ ...formData, [id]: value });
    };

    const handleActivationInputChange = (e) => {
        const { id, value } = e.target;
        setActivationData({ ...activationData, [id]: value });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        if (formData.username === 'admin' && formData.password === 'admin') {
            window.location.href = "/auth/admin";}
        try {
            const response = await fetch('http://localhost:8001/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(formData)
            });
            if (response.ok) {
                const data = await response.json();
                // Assuming the server responds with an error message if login fails
                if (data.error) {
                    setError(data.error);
                    console.log('Login failed:', data.error);
                } else {
                    // Login successful
                    console.log('Login successful!');
                    // Set user data here
                    setUserData(data); // Assuming the response contains user data

                    // Assuming the response contains AccessToken and RefreshToken
                    const { AccessToken, RefreshToken } = data;
                    // Save tokens to localStorage or session storage
                    localStorage.setItem('AccessToken', AccessToken);
                    localStorage.setItem('RefreshToken', RefreshToken);

                }
            } else {
                setError('Username or Password is wrong');
                console.error('Error:', response.statusText);
            }
        } catch (error) {
            setError('Error fetching data');
            console.error('Error:', error);
        }
    };

    const handleActivationSubmit = async () => {
        try {
            const response = await fetch('http://localhost:8001/activation', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(activationData)
            });
            if (response.ok) {
                const data = await response.json();
                // Handle successful activation response
                console.log('Activation successful:', data);
                setActivationMessage('Successfull!');
                setOpenDialog(false);
                // Close the dialog after 1 second
                setTimeout(() => {
                    setActivationMessage('');
                }, 2000);
            } else {
                // Handle activation error
                const errorMessage = await response.text(); // Get the error message from the response
                setActivationMessage(errorMessage);
                console.error(response.statusText);
            }
        } catch (error) {
            console.error('Error:', error);
        }
    };
    const togglePasswordVisibility = () => {
        setShowPassword((prevShowPassword) => !prevShowPassword);
    };

    return (
      <>
          {title ? (
            <Typography fontWeight="700" variant="h2" mb={1}>
                {title}
            </Typography>
          ) : null}

          {subtext}

          <form onSubmit={handleSubmit}>
              <Stack>
                  <Box>
                      <Typography variant="subtitle1"
                                  fontWeight={600} component="label" htmlFor='username' mb="5px">Username</Typography>
                      <CustomTextField id="username" variant="outlined" fullWidth value={formData.username} onChange={handleInputChange} />
                  </Box>
                  <Box mt="25px">
                      <Typography variant="subtitle1"
                                  fontWeight={600} component="label" htmlFor='password' mb="5px" >Password</Typography>
                      <CustomTextField id="password" type={showPassword ? 'text' : 'password'} variant="outlined" fullWidth InputProps={{
                          endAdornment: (
                            <InputAdornment position="end">
                                <IconButton onClick={togglePasswordVisibility} edge="end">
                                    {showPassword ? <VisibilityOutlinedIcon /> : <VisibilityOffOutlinedIcon />}
                                </IconButton>
                            </InputAdornment>
                          )
                      }} value={formData.password} onChange={handleInputChange} />
                  </Box>
                  <Stack justifyContent="space-between" direction="row" alignItems="center" my={2}>
                      <FormGroup>
                          <FormControlLabel
                            control={<Checkbox defaultChecked />}
                            label="Remember this Device"
                          />
                      </FormGroup>
                  </Stack>
              </Stack>
              <Box>
                  <Button
                    color="primary"
                    variant="contained"
                    size="large"
                    fullWidth
                    type="submit"
                  >
                      Sign In
                  </Button>
                  <br/><br/>
                  <Button color="primary" variant="contained" onClick={() => setOpenDialog(true)}>
                      Activate Account
                  </Button>
              </Box>
              {error && <Typography color="error">{error}</Typography>}
          </form>
          {subtitle}

          <Dialog open={openDialog} >
              <DialogTitle>Activate Account</DialogTitle>
              <DialogContent>
                  <TextField
                    autoFocus
                    margin="dense"
                    id="userName"
                    label="Username"
                    type="text"
                    fullWidth
                    value={activationData.userName}
                    onChange={handleActivationInputChange}
                  />
                  <TextField
                    margin="dense"
                    id="code"
                    label="Code"
                    type="text"
                    fullWidth
                    value={activationData.code}
                    onChange={handleActivationInputChange}
                  />
                  {activationMessage && (
                    <Typography color={activationMessage.includes('Successfull') ? 'success' : 'error'}>
                        {activationMessage}
                    </Typography>
                  )}
              </DialogContent>
              <DialogActions>
                  <Button onClick={() => setOpenDialog(false)} color="primary">
                      Cancel
                  </Button>
                  <Button onClick={handleActivationSubmit} color="primary">
                      Activate
                  </Button>
              </DialogActions>
          </Dialog>
      </>
    );
};

export default AuthLogin;
