import React, { useState } from 'react';
import {
  Card,
  CardContent,
  Typography,
  TextField,
  Button,
  List,
  ListItem,
  ListItemText,
  Grid,
} from '@mui/material';
import CurrencyLiraIcon from '@mui/icons-material/CurrencyLira';
import axios from 'axios';

const Wallet = () => {
  const [updateBalance, setUpdateBalance] = useState('');
  const [cartItems, setCartItems] = useState([]);
  const [purchaseHistory, setPurchaseHistory] = useState([]);

  const handleUpdateBalance = async () => {
    try {
      const accessToken = localStorage.getItem('AccessToken');
      const floatBalance = parseFloat(updateBalance); // Parse the updateBalance to float
      const response = await fetch('http://localhost:8001/wallet/', {
        method: 'PUT',
        headers: {
          'Authentication': `${accessToken}`,
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ balance: floatBalance }) // Send floatBalance instead of updateBalance
      });

      if (!response.ok) {
        throw new Error('Failed to update balance');
      }

      console.log('Balance updated:', floatBalance);
      setUpdateBalance('');
    } catch (error) {
      console.error('Error updating balance:', error);
    }
  };

  const handleAddToCart = (item) => {
    setCartItems([...cartItems, item]);
  };

  const handleCompletePurchase = async () => {
    try {
      const accessToken = localStorage.getItem('AccessToken');
      await axios.get(
        'http://localhost:8001/wallet/complete',
        {
          headers: {
            'Authentication': accessToken
          }
        }
      );
      console.log('Purchase completed:', cartItems);
      setPurchaseHistory([...purchaseHistory, ...cartItems]);
      setCartItems([]);
    } catch (error) {
      console.error('Error completing purchase:', error);
    }
  };
  return (
    <Grid container spacing={3}>
      <Grid item xs={12} md={4}>
        <Card>
          <CardContent>
            <Typography variant="h5" gutterBottom>
              <CurrencyLiraIcon fontSize="large" /> Balance
            </Typography>
            <Typography variant="h4" gutterBottom>
              0.00 TL
            </Typography>
            <TextField
              fullWidth
              type="number"
              label="Enter Balance"
              value={updateBalance}
              onChange={(e) => setUpdateBalance(e.target.value)}
            />
            <Button variant="contained" onClick={handleUpdateBalance}>
              Update Balance
            </Button>
          </CardContent>
        </Card>
        <Card style={{ marginTop: '20px' }}>
          <CardContent>
            <Typography variant="h5" gutterBottom>
              Complete Purchase
            </Typography>
            <List>
              {cartItems.map((item, index) => (
                <ListItem key={index}>
                  <ListItemText primary={item} />
                </ListItem>
              ))}
            </List>
            <Button variant="contained" onClick={handleCompletePurchase}>
              Complete Purchase
            </Button>
          </CardContent>
        </Card>
      </Grid>
      <Grid item xs={12} md={9}>
        <Card style={{ marginTop: '20px' }}>
          <CardContent>
            <Typography variant="h5" gutterBottom>
              Old Purchases
            </Typography>
            <List>
              {purchaseHistory.map((purchase, index) => (
                <ListItem key={index}>
                  <ListItemText primary={purchase} />
                </ListItem>
              ))}
            </List>
          </CardContent>
        </Card>
      </Grid>
    </Grid>
  );
};

export default Wallet;
