import React, { useEffect, useState } from 'react';
import { makeStyles } from '@mui/styles';
import { Drawer, List, ListItem, ListItemText, Typography, Divider, Box, Paper, TextField, Button, Table, TableContainer, TableHead, TableBody, TableRow, TableCell, Dialog, DialogTitle, DialogContent, DialogActions } from '@mui/material';
import axios from 'axios';

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
  redBackground: {
    backgroundColor: 'red',
  },
}));

const Admin = () => {
  const classes = useStyles();
  const [selectedMenu, setSelectedMenu] = useState(null);
  const [messages, setMessages] = useState([]);
  const [trips, setTrips] = useState([]);
  const [users, setUsers] = useState([]);
  const [feedbacks, setFeedbacks] = useState([]);
  const [openDialog, setOpenDialog] = useState(false);
  const [newTrip, setNewTrip] = useState({
    Id: 0,
    Description: '',
    Name: '',
    Status: 0,
    Lat: 0,
    Lng: 0,
    NaturePoint: 0,
    HistoricalPoint: 0,
    AdventurePoint: 0,
    CulturalPoint: 0,
    RelaxPoint: 0
  });

  useEffect(() => {
    const handleMessage = (newMessage) => {
      setMessages([...messages, newMessage]);
    };
    window.addEventListener('message', handleMessage);
    return () => window.removeEventListener('message', handleMessage);
  }, [messages]);

  const handleMenuClick = (menu) => {
    setSelectedMenu(menu === selectedMenu ? null : menu);
    if (menu === 'Trip Management') {
      fetchTrips();
    } else if (menu === 'User Management') {
      fetchUsers();
    } else if (menu === 'Feedback Management') {
      fetchFeedbacks();
    }
  };


  const fetchTrips = async () => {
    try {
      const accessToken = localStorage.getItem('AccessToken');
      const response = await fetch('http://localhost:8001/panel/trip/', {
        headers: {
          'Authentication': `${accessToken}`
        }
      });
      if (response.ok) {
        const data = await response.json();
        setTrips(data);
      } else {
        console.error('Error fetching trips:', response.statusText);
      }
    } catch (error) {
      console.error('Error fetching trips:', error);
    }
  };

  const fetchUsers = async () => {
    try {
      const accessToken = localStorage.getItem('AccessToken');
      const response = await fetch('http://localhost:8001/panel/user/', {
        headers: {
          'Authentication': `${accessToken}`
        }
      });
      if (response.ok) {
        const data = await response.json();
        setUsers(data);
      } else {
        console.error('Error fetching users:', response.statusText);
      }
    } catch (error) {
      console.error('Error fetching users:', error);
    }
  };

  const handleDeleteTrip = async (id) => {
    try {
      const accessToken = localStorage.getItem('AccessToken');
      const response = await fetch('http://localhost:8001/panel/trip/', {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
          'Authentication': accessToken
        },
        body: JSON.stringify({ id: id })
      });
      if (response.ok) {
        fetchTrips();
      } else {
        console.error('Error deleting trip:', response.statusText);
      }
    } catch (error) {
      console.error('Error deleting trip:', error);
    }
  };


  const handleDeleteUser = async (id) => {
    try {
      const accessToken = localStorage.getItem('AccessToken');
      const response = await fetch('http://localhost:8001/panel/user/', {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
          'Authentication': accessToken
        },
        body: JSON.stringify({ id: id })
      });
      if (response.ok) {
        fetchUsers();
      } else {
        console.error('Error deleting user:', response.statusText);
      }
    } catch (error) {
      console.error('Error deleting user:', error);
    }
  };

  const handleCreateTrip = async () => {
    try {
      const accessToken = localStorage.getItem('AccessToken');
      const response = await fetch('http://localhost:8001/panel/trip/', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authentication': accessToken
        },
        body: JSON.stringify(newTrip)
      });
      if (response.ok) {
        setOpenDialog(false);
        setNewTrip({
          Id: 0,
          Description: '',
          Name: '',
          Status: 0,
          Lat: 0,
          Lng: 0,
          NaturePoint: 0,
          HistoricalPoint: 0,
          AdventurePoint: 0,
          CulturalPoint: 0,
          RelaxPoint: 0
        });
        fetchTrips();
      } else {
        console.error('Error creating trip:', response.statusText);
      }
    } catch (error) {
      console.error('Error creating trip:', error);
    }
  };
  const fetchFeedbacks = async () => {
    try {
      const accessToken = localStorage.getItem('AccessToken');
      const response = await fetch('http://localhost:8001/feedback/all/', {
        headers: {
          'Authentication': accessToken
        }
      });
      if (response.ok) {
        const data = await response.json();
        setFeedbacks(data);
      } else {
        console.error('Error fetching feedbacks:', response.statusText);
      }
    } catch (error) {
      console.error('Error fetching feedbacks:', error);
    }
  };
  const handleDeleteFeedback = async (id) => {
    try {
      const accessToken = localStorage.getItem('AccessToken');
      await axios.delete('http://localhost:8001/feedback/', {
        headers: {
          'Authentication': `${accessToken}`,
        },
        data: {
          id: id
        }
      });
    } catch (error) {
      console.error('Error deleting feedback:', error);
    }
    try {
      const accessToken = localStorage.getItem('AccessToken');
      const response = await fetch('http://localhost:8001/feedback/all/', {
        headers: {
          'Authentication': accessToken
        }
      });
      if (response.ok) {
        const data = await response.json();
        setFeedbacks(data);
      } else {
        console.error('Error fetching feedbacks:', response.statusText);
      }
    } catch (error) {
      console.error('Error fetching feedbacks:', error);
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
          {['User Management', 'Trip Management', 'Feedback Management'].map((text) => (
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
                    <TableCell>Name</TableCell>
                    <TableCell>Surname</TableCell>
                    <TableCell>Phone</TableCell>
                    <TableCell className={classes.redBackground}>Status</TableCell> {/* Status sütunu için kırmızı arka plan */}
                    <TableCell>Birth Date</TableCell>
                    <TableCell>Action</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {users.map((user) => (
                    <TableRow key={user.Id} className={user.Status === 9 ? classes.redBackground : null}> {/* Arka planı kırmızı yap */}
                      <TableCell>{user.Id}</TableCell>
                      <TableCell>{user.username}</TableCell>
                      <TableCell>{user.email}</TableCell>
                      <TableCell>{user.name}</TableCell>
                      <TableCell>{user.surname}</TableCell>
                      <TableCell>{user.phone}</TableCell>
                      <TableCell>{user.Status}</TableCell>
                      <TableCell>{user.birth_date}</TableCell>
                      <TableCell>
                        <Button variant="contained" color="error" onClick={() => handleDeleteUser(user.Id)}>
                          Delete
                        </Button>
                      </TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </TableContainer>
          </Paper>
        )}
        {selectedMenu && selectedMenu === 'Trip Management' && (
          <Paper elevation={3} style={{ marginTop: '20px', padding: '20px' }}>
            <Typography variant="h5">Trip Management</Typography>
            <Button variant="contained" color="primary" onClick={() => setOpenDialog(true)}>
              Create Trip
            </Button>
            <Dialog open={openDialog} onClose={() => setOpenDialog(false)} fullWidth maxWidth="lg">
              <DialogTitle>Create New Trip</DialogTitle>
              <DialogContent>
                <TextField
                  label="ID"
                  type="number"
                  value={newTrip.Id}
                  onChange={(e) => setNewTrip({ ...newTrip, Id: parseInt(e.target.value) })}
                  style={{ marginTop: '10px' }}
                  fullWidth
                />
                <TextField
                  label="Name"
                  value={newTrip.Name}
                  onChange={(e) => setNewTrip({ ...newTrip, Name: e.target.value })}
                  style={{ marginTop: '10px' }}
                  fullWidth
                />
                <TextField
                  label="Description"
                  value={newTrip.Description}
                  onChange={(e) => setNewTrip({ ...newTrip, Description: e.target.value })}
                  style={{ marginTop: '10px' }}
                  fullWidth
                />
                <TextField
                  label="Latitude"
                  type="number"
                  value={newTrip.Lat}
                  onChange={(e) => setNewTrip({ ...newTrip, Lat: parseFloat(e.target.value) })}
                  style={{ marginTop: '10px' }}
                  fullWidth
                />
                <TextField
                  label="Longitude"
                  type="number"
                  value={newTrip.Lng}
                  onChange={(e) => setNewTrip({ ...newTrip, Lng: parseFloat(e.target.value) })}
                  style={{ marginTop: '10px' }}
                  fullWidth
                />
              </DialogContent>
              <DialogActions>
                <Button onClick={() => setOpenDialog(false)}>Cancel</Button>
                <Button onClick={handleCreateTrip}>Create</Button>
              </DialogActions>
            </Dialog>
            <TableContainer component={Paper}>
              <Table>
                <TableHead>
                  <TableRow>
                    <TableCell>ID</TableCell>
                    <TableCell>Name</TableCell>
                    <TableCell>Description</TableCell>
                    <TableCell>Latitude</TableCell>
                    <TableCell>Longitude</TableCell>
                    <TableCell className={classes.redBackground}>Status</TableCell> {/* Status sütunu için kırmızı arka plan */}
                    <TableCell>Action</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {trips.map((trip) => (
                    <TableRow key={trip.id} className={trip.status === 9 ? classes.redBackground : null}> {/* Arka planı kırmızı yap */}
                      <TableCell>{trip.id}</TableCell>
                      <TableCell>{trip.name}</TableCell>
                      <TableCell>{trip.description}</TableCell>
                      <TableCell>{trip.lat}</TableCell>
                      <TableCell>{trip.lng}</TableCell>
                      <TableCell>{trip.status}</TableCell>
                      <TableCell>
                        <Button variant="contained" color="error" onClick={() => handleDeleteTrip(trip.id)}>
                          Delete
                        </Button>
                      </TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </TableContainer>
          </Paper>
        )}
        {selectedMenu && selectedMenu === 'Feedback Management' && (
          <Paper elevation={3} style={{ marginTop: '20px', padding: '20px' }}>
            <Typography variant="h5">Feedback Management</Typography>
            <TableContainer component={Paper}>
              <Table>
                <TableHead>
                  <TableRow>
                    <TableCell>ID</TableCell>
                    <TableCell>Description</TableCell>
                    <TableCell>Star</TableCell>
                    <TableCell>Product ID</TableCell>
                    <TableCell>User ID</TableCell>
                    <TableCell>Status</TableCell>
                    <TableCell>User Name</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {feedbacks.map((feedback) => (
                    <TableRow key={feedback.id}>
                      <TableCell>{feedback.id}</TableCell>
                      <TableCell>{feedback.description}</TableCell>
                      <TableCell>{feedback.star}</TableCell>
                      <TableCell>{feedback.product_id}</TableCell>
                      <TableCell>{feedback.user_id}</TableCell>
                      <TableCell>{feedback.status}</TableCell>
                      <TableCell>{feedback.user_name}</TableCell>
                      <TableCell>
                        <Button variant="contained" color="error" onClick={() => handleDeleteFeedback(feedback.id)}>
                          Delete
                        </Button>
                      </TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </TableContainer>
          </Paper>
        )}
      </Box>
    </Box>
  );
};

export default Admin;
