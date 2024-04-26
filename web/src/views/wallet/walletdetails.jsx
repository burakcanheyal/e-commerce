import React, { useState } from 'react';
import { Card, CardContent, Typography, Grid } from '@mui/material';

const Wallet = () => {
  const [wallet, setWallet] = useState(null);
  const [operations, setOperations] = useState([]);

  const sampleWallet = {
    balance: 1000,
    status: 1,
    createdAt: '2022-04-10T12:00:00Z',
    updatedAt: '2022-04-11T10:00:00Z',
  };

  const sampleOperations = [
    { id: 1, operationNumber: 'OP001', type: 1, balance: 500, operationDate: '2022-04-10T13:00:00Z' },
    { id: 2, operationNumber: 'OP002', type: 2, balance: 200, operationDate: '2022-04-11T09:00:00Z' },
  ];

  useState(() => {
    setWallet(sampleWallet);
    setOperations(sampleOperations);
  }, []);

  return (
    <Grid container spacing={3}>
      <Grid item xs={12} md={6}>
        <Card>
          <CardContent>
            <Typography variant="h5" gutterBottom>
              Wallet Information
            </Typography>
            {wallet && (
              <div>
                <Typography variant="body1">
                  Balance: {wallet.balance}
                </Typography>
                <Typography variant="body1">
                  Status: {wallet.status === 1 ? 'Active' : 'Inactive'}
                </Typography>
                <Typography variant="body1">
                  Created At: {new Date(wallet.createdAt).toLocaleDateString()}
                </Typography>
                <Typography variant="body1">
                  Updated At: {wallet.updatedAt ? new Date(wallet.updatedAt).toLocaleDateString() : '-'}
                </Typography>
              </div>
            )}
          </CardContent>
        </Card>
      </Grid>
      <Grid item xs={12} md={6}>
        <Card>
          <CardContent>
            <Typography variant="h5" gutterBottom>
              Wallet Operations
            </Typography>
            <table>
              <thead>
              <tr>
                <th>Operation Number</th>
                <th>Type</th>
                <th>Balance</th>
                <th>Date</th>
              </tr>
              </thead>
              <tbody>
              {operations.map(operation => (
                <tr key={operation.id}>
                  <td>{operation.operationNumber}</td>
                  <td>{operation.type === 1 ? 'Deposit' : 'Withdrawal'}</td>
                  <td>{operation.balance}</td>
                  <td>{new Date(operation.operationDate).toLocaleDateString()}</td>
                </tr>
              ))}
              </tbody>
            </table>
          </CardContent>
        </Card>
      </Grid>
    </Grid>
  );
};

export default Wallet;
