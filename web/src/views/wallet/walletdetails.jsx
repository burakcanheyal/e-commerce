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

const Wallet = () => {
  const [updateBalance, setUpdateBalance] = useState('');
  const [cartItems, setCartItems] = useState([]);
  const [purchaseHistory, setPurchaseHistory] = useState([]);

  const handleUpdateBalance = () => {
    // Backend tarafında yapılacak olan balance güncelleme işlemi burada gerçekleştirilecek
    // Örnek olarak: axios.put('http://localhost:8001/updateBalance', { balance: updateBalance });
    console.log('Update Balance:', updateBalance);
    setUpdateBalance('');
  };

  const handleAddToCart = (item) => {
    setCartItems([...cartItems, item]);
  };

  const handleCompletePurchase = () => {
    // Backend tarafında yapılacak olan satın alma işlemi burada gerçekleştirilecek
    // Örnek olarak: axios.post('http://localhost:8001/completePurchase', { cartItems });
    console.log('Complete Purchase:', cartItems);
    setPurchaseHistory([...purchaseHistory, ...cartItems]);
    setCartItems([]);
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
