import React, { useState } from 'react';
import axios from 'axios';

const RoutesPage = () => {
  const [tokens, setTokens] = useState(null); // State to store tokens
  const [fetching, setFetching] = useState(false); // State to track fetching status
  const [username, setUsername] = useState(''); // State for username input
  const [password, setPassword] = useState(''); // State for password input

  // Function to handle fetch request
  const fetchData = async () => {
    try {
      // Make a POST request to the backend API with username and password
      const response = await axios.post('http://localhost:8001/login', {
        username,
        password
      });

      // Assuming the API response contains AccessToken and RefreshToken as JSON
      const { AccessToken, RefreshToken } = response.data;

      // Store tokens in state
      setTokens({ AccessToken, RefreshToken });
    } catch (error) {
      console.error('Error fetching data:', error);
    } finally {
      // Reset fetching status after request completes
      setFetching(false);
    }
  };

  // Function to handle button click
  const handleButtonClick = () => {
    // Set fetching status to true
    setFetching(true);

    // Fetch data
    fetchData();
  };

  return (
    <div>
      <h1>Routes Page</h1>

      {/* Username input field */}
      <input
        type="text"
        placeholder="Username"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
      />

      {/* Password input field */}
      <input
        type="password"
        placeholder="Password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />

      {/* Button to trigger fetch request */}
      <button onClick={handleButtonClick} disabled={fetching}>
        {fetching ? 'Fetching...' : 'Send Request'}
      </button>

      {/* Display tokens if available */}
      {tokens && (
        <div>
          <p>Access Token: {tokens.AccessToken}</p>
          <p>Refresh Token: {tokens.RefreshToken}</p>
        </div>
      )}
    </div>
  );
};

export default RoutesPage;
