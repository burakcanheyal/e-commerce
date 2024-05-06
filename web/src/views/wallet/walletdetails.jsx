import React, { useState, useEffect } from 'react';
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
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper
} from '@mui/material';
import CurrencyLiraIcon from '@mui/icons-material/CurrencyLira';
import axios from 'axios';

const Wallet = () => {
  const [updateBalance, setUpdateBalance] = useState('');
  const [cartItems, setCartItems] = useState([]);
  const [purchaseHistory, setPurchaseHistory] = useState([]);
  const [balance, setBalance] = useState(0);
  const [orderData, setOrderData] = useState(null);
  const [completedOrders, setCompletedOrders] = useState([]);

  useEffect(() => {
    const fetchBalance = async () => {
      try {
        const accessToken = localStorage.getItem('AccessToken');
        const response = await axios.get('http://localhost:8001/wallet/', {
          headers: {
            'Authentication': accessToken
          }
        });
        if (response.data && response.data.balance) {
          setBalance(response.data.balance);
        }
      } catch (error) {
        console.error('Error fetching balance:', error);
      }
    };

    fetchBalance();
    handleGetOrders();
    // Bileşen yüklendiğinde otomatik olarak tamamlanmış siparişleri alma
    handleGetCompletedOrders();
  }, []);

  const handleUpdateBalance = async () => {
    try {
      const accessToken = localStorage.getItem('AccessToken');
      const floatBalance = parseFloat(updateBalance);
      const response = await fetch('http://localhost:8001/wallet/', {
        method: 'PUT',
        headers: {
          'Authentication': `${accessToken}`,
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ balance: floatBalance })
      });

      if (!response.ok) {
        throw new Error('Failed to update balance');
      }
      console.log('Balance updated:', floatBalance);
      setUpdateBalance('');
      // Update the balance after successful update
    } catch (error) {
      console.error('Error updating balance:', error);
    }
    try {
      const accessToken = localStorage.getItem('AccessToken');
      const response = await axios.get('http://localhost:8001/wallet/', {
        headers: {
          'Authentication': accessToken
        }
      });
      if (response.data && response.data.balance) {
        setBalance(response.data.balance);
      }
    } catch (error) {
      console.error('Error fetching balance:', error);
    }
  };

  const handleGetOrders = async () => {
    try {
      const accessToken = localStorage.getItem('AccessToken');
      const response = await axios.get('http://localhost:8001/order/', {
        headers: {
          'Authentication': accessToken
        }
      });
      if (response.data) {
        setOrderData(response.data);
      } else {
        setOrderData(null);
        console.log('Shopping cart list is empty');
      }
      console.log('Order data:', response.data);
    } catch (error) {
      console.error('Error completing purchase:', error);
    }
  };

  const handleCompletePurchase = async () => {
    try {
      const accessToken = localStorage.getItem('AccessToken');
      const response = await axios.get('http://localhost:8001/wallet/complete', {
        headers: {
          'Authentication': accessToken
        }
      });
      console.log('Complete purchase response:', response.data);
      if (response.data) {
        setOrderData(response.data);
      }
    } catch (error) {
      console.error('Error completing purchase:', error);
    }

  };

  const handleGetCompletedOrders = async () => {
    try {
      const accessToken = localStorage.getItem('AccessToken');
      const response = await axios.get('http://localhost:8001/wallet/completedOrders', {
        headers: {
          'Authentication': accessToken
        }
      });
      if (response.data) {
        setCompletedOrders(response.data['Siparişler: ']);
      }
      console.log('Completed orders data:', response.data);
    } catch (error) {
      console.error('Error getting completed orders:', error);
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
              {balance.toFixed(2)} TL
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
            {orderData ? (
              <TableContainer component={Paper}>
                <Table>
                  <TableHead>
                    <TableRow>
                      <TableCell>Name</TableCell>
                      <TableCell>Quantity</TableCell>
                      <TableCell>Price</TableCell>
                    </TableRow>
                  </TableHead>
                  <TableBody>
                    {orderData['Siparişler: '].map((order, index) => (
                      <TableRow key={index}>
                        <TableCell>{order.name}</TableCell>
                        <TableCell>{order.quantity}</TableCell>
                        <TableCell>{order.price}</TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </TableContainer>
            ) : (
              <Typography variant="body1" color="textSecondary">
                Shopping cart is empty
              </Typography>
            )}
            <Button variant="contained" onClick={handleCompletePurchase}>
              Complete Purchase
            </Button>
          </CardContent>
        </Card>
      </Grid>
      <Grid item xs={12} md={8}>
        <Card style={{ marginTop: '20px' }}>
          <CardContent>
            <Typography variant="h5" gutterBottom>
              Old Purchases
            </Typography>
            <TableContainer component={Paper}>
              <Table>
                <TableHead>
                  <TableRow>
                    <TableCell>Operation Number</TableCell>
                    <TableCell>Balance</TableCell>
                    <TableCell>Order ID</TableCell>
                    <TableCell>Product Name</TableCell>
                    <TableCell>Order Quantity</TableCell>
                    <TableCell>Seller Name</TableCell>
                    <TableCell>Operation Date</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {completedOrders.map((order, index) => (
                    <TableRow key={index}>
                      <TableCell>{order.OperationNumber}</TableCell>
                      <TableCell>{order.Balance}</TableCell>
                      <TableCell>{order.OrderId}</TableCell>
                      <TableCell>{order.ProductName}</TableCell>
                      <TableCell>{order.OrderQuantity}</TableCell>
                      <TableCell>{order.SellerName}</TableCell>
                      <TableCell>{order.OperationDate}</TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </TableContainer>
          </CardContent>
        </Card>
      </Grid>
    </Grid>
  );
};

export default Wallet;
