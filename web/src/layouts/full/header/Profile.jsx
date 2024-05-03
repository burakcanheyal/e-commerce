import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import axios from 'axios'; // Axios'u ekledim
import {
  Avatar,
  Box,
  Menu,
  Button,
  IconButton,
  MenuItem,
  ListItemIcon,
  ListItemText
} from '@mui/material';
import AccountCircleOutlinedIcon from '@mui/icons-material/AccountCircleOutlined';
import EnhancedEncryptionOutlinedIcon from '@mui/icons-material/EnhancedEncryptionOutlined';

const Profile = () => {
  const [anchorEl2, setAnchorEl2] = useState(null);
  const [username, setUsername] = useState('');

  useEffect(() => {
    const fetchProfileDetails = async () => {
      try {
        const accessToken = localStorage.getItem('AccessToken');
        const response = await axios.get('http://localhost:8001/profil/', {
          headers: {
            'Authentication': `${accessToken}`
          }
        });
        setUsername(response.data.username);
      } catch (error) {
        console.error('Error fetching username:', error);
      }
    };
    fetchProfileDetails();
  }, []);

  const handleClick2 = (event) => {
    setAnchorEl2(event.currentTarget);
  };

  const handleClose2 = () => {
    setAnchorEl2(null);
  };

  const handleLogout = () => {
    localStorage.clear();
    // Redirect to the login page after logout
    return <Navigate to="/auth/login" />;
  };

  return (
    <Box>
      <IconButton
        size="large"
        aria-label="show 11 new notifications"
        color="inherit"
        aria-controls="msgs-menu"
        aria-haspopup="true"
        sx={{
          ...(typeof anchorEl2 === 'object' && {
            color: 'primary.main',
          }),
        }}
        onClick={handleClick2}
      >
        <Avatar
          sx={{
            width: 35,
            height: 35,
          }}
        />
      </IconButton>

      {/* Message Dropdown */}
      <Menu
        id="msgs-menu"
        anchorEl={anchorEl2}
        keepMounted
        open={Boolean(anchorEl2)}
        onClose={handleClose2}
        anchorOrigin={{ horizontal: 'right', vertical: 'bottom' }}
        transformOrigin={{ horizontal: 'right', vertical: 'top' }}
        sx={{
          '& .MuiMenu-paper': {
            width: '200px',
            borderRadius: '5px',
          },
        }}
      >
        <MenuItem component={Link} to="/sample-page">
          <ListItemIcon>
            <AccountCircleOutlinedIcon />
          </ListItemIcon>
          <ListItemText>My Profile ({username})</ListItemText>
        </MenuItem>
        <MenuItem component={Link} to="/updatepassword">
          <ListItemIcon>
            <EnhancedEncryptionOutlinedIcon />
          </ListItemIcon>
          <ListItemText>Update Password</ListItemText>
        </MenuItem>
        <Box mt={1} py={1} px={2}>
          <Button
            to="/auth/login"
            variant="contained"
            color="primary"
            component={Link}
            fullWidth
            onClick={handleLogout}
          >
            Logout
          </Button>
        </Box>
      </Menu>
    </Box>
  );
};

export default Profile;
