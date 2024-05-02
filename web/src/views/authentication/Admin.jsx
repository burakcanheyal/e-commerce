import React, { useEffect, useState } from 'react';
import { makeStyles } from '@mui/styles';
import { Drawer, List, ListItem, ListItemText, Typography, Divider, Box, Paper, TextField, Button, Table, TableContainer, TableHead, TableBody, TableRow, TableCell } from '@mui/material';

const drawerWidth = 240;

const useStyles = makeStyles((theme) => ({
  drawer: {
    width: drawerWidth,
    flexShrink: 0,
  },
  drawerPaper: {
    width: drawerWidth,
  },
  content: {
    flexGrow: 1,
    padding: theme.spacing(3),
  },
}));

const Admin = () => {
  const classes = useStyles();
  const [selectedMenu, setSelectedMenu] = useState(null);
  const [messages, setMessages] = useState([]);
  const [response, setResponse] = useState('');
  const [users, setUsers] = useState([]);

  useEffect(() => {
    const handleMessage = (newMessage) => {
      setMessages([...messages, newMessage]);
    };
    window.addEventListener('message', handleMessage);
    return () => window.removeEventListener('message', handleMessage);
  }, [messages]);

  const handleMenuClick = (menu) => {
    setSelectedMenu(menu === selectedMenu ? null : menu);
    if (menu === 'User Management') {
      fetchUsers();
    }
  };

  const handleNewResponse = () => {
    if (response.trim() !== '') {
      setMessages([...messages, response]);
      setResponse('');
    }
  };

  const fetchUsers = async () => {
    try {
      const accessToken = localStorage.getItem('AccessToken');
      const response = await fetch('http://localhost:8001/profil/', {
        headers: {
          'Authentication': `${accessToken}`
        }
      });
      if (response.ok) {
        const data = await response.json();
        setUsers(data.users);
      } else {
        console.error('Error fetching users:', response.statusText);
      }
    } catch (error) {
      console.error('Error fetching users:', error);
    }
  };

  return (
    <Box display="flex">
      <Drawer
        className={classes.drawer}
        variant="permanent"
        classes={{
          paper: classes.drawerPaper,
        }}
        anchor="left"
      >
        <List>
          {['User Management', 'Live Response', 'Trip Management'].map((text) => (
            <ListItem button key={text} onClick={() => handleMenuClick(text)}>
              <ListItemText primary={text} />
            </ListItem>
          ))}
        </List>
      </Drawer>
      <Box className={classes.content}>
        <Typography variant="h4">Admin Panel</Typography>
        <Divider />
        {selectedMenu && selectedMenu === 'User Management' && (
          <Paper elevation={3} style={{ marginTop: '20px', padding: '20px' }}>
            <Typography variant="h5">User Management</Typography>
            <TableContainer component={Paper}>
              <Table>
                <TableHead>
                  <TableRow>
                    <TableCell>ID</TableCell>
                    <TableCell>Username</TableCell>
                    <TableCell>Email</TableCell>
                  </TableRow>
                </TableHead>
                {users && users.length > 0 ? (
                  <TableBody>
                    {users.map((user) => (
                      <TableRow key={user.id}>
                        <TableCell>{user.id}</TableCell>
                        <TableCell>{user.username}</TableCell>
                        <TableCell>{user.email}</TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                ) : (
                  <TableBody>
                    <TableRow>
                      <TableCell colSpan={3}>No users found</TableCell>
                    </TableRow>
                  </TableBody>
                )}
              </Table>
            </TableContainer>
          </Paper>
        )}
        {selectedMenu && selectedMenu === 'Live Response' && (
          <Paper elevation={3} style={{ marginTop: '20px', padding: '20px' }}>
            <Typography variant="h5">Live Response</Typography>
            <div>
              {messages.map((message, index) => (
                <div key={index}>
                  <p>{message}</p>
                </div>
              ))}
            </div>
            <TextField
              label="Response"
              multiline
              rows={4}
              value={response}
              onChange={(e) => setResponse(e.target.value)}
            />
            <Button variant="contained" onClick={handleNewResponse}>
              Send Response
            </Button>
          </Paper>
        )}
      </Box>
    </Box>
  );
};

export default Admin;
