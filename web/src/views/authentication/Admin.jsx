// Admin.jsx

import React, { useEffect, useState } from 'react';
import { makeStyles } from '@mui/styles';
import { Drawer, List, ListItem, ListItemText, Typography, Divider, Box, Paper, TextField, Button } from '@mui/material';

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

  useEffect(() => {
    const handleMessage = (newMessage) => {
      setMessages([...messages, newMessage]);
    };
    window.addEventListener('message', handleMessage);
    return () => window.removeEventListener('message', handleMessage);
  }, [messages]);

  const handleMenuClick = (menu) => {
    setSelectedMenu(menu === selectedMenu ? null : menu);
  };

  const handleNewResponse = () => {
    if (response.trim() !== '') {
      setMessages([...messages, response]);
      setResponse('');
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
