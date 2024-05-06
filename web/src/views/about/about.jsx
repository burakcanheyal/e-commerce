import React, { useState } from 'react';
import { Typography, Box, TextField, Button } from '@mui/material';
import CTISLogo from '/@/assets/images/logos/30yearsCTISvBlack.png';

const AboutUsPage = () => {
  const [formData, setFormData] = useState({
    name: '',
    email: '',
    message: ''
  });

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log(formData);
    sendEmail();
  };

  const sendEmail = () => {

    const emailAddress = 'burakcanheyal@gmail.com';
    const subject = 'Contact Us Form Submission';
    const body = `Name: ${formData.name}\nEmail: ${formData.email}\nMessage: ${formData.message}`;

    const mailToLink = `mailto:${emailAddress}?subject=${encodeURIComponent(subject)}&body=${encodeURIComponent(body)}`;
    window.location.href = mailToLink;
  };

  const handleCTISLink = () => {
    window.open('https://www.ctis.bilkent.edu.tr/ctis_seniorProject.php?id=5010', '_blank');
  };

  return (
    <Box sx={{ textAlign: 'left', paddingTop: 7, overflowY:'hidden' }}>
      <Box sx={{ marginTop: -7 }}>
        <Typography variant="h4" gutterBottom>
          Contact Us
        </Typography>
        <form onSubmit={handleSubmit}>
          <TextField
            name="name"
            label="Name"
            variant="outlined"
            margin="normal"
            style={{ width:'30%' }}
            value={formData.name}
            onChange={handleInputChange}
          />
          <TextField
            name="email"
            label="Email"
            variant="outlined"
            style={{ display:'flex',width:'50%' }}
            value={formData.email}
            onChange={handleInputChange}
          />
          <TextField
            name="message"
            label="Message"
            variant="outlined"
            style={{ display:'flex',width:'60%' }}
            margin="normal"
            multiline
            rows={4}
            value={formData.message}
            onChange={handleInputChange}
          />
          <Button type="submit" variant="contained" color="primary">
            Send
          </Button>
        </form>
      </Box>
      <Box sx={{ marginTop: 5}} style={{textAlign:'center'}}>
        <img src={CTISLogo} alt="30yearsCTISvBlack Logo" style={{ width: 230, cursor: 'pointer' }} onClick={handleCTISLink} />
      </Box>
      <Box sx={{ marginTop:3, textAlign: 'center', paddingBottom: 3}}>
        <Typography variant="body1" color="textSecondary">
          © 2024 Exporia. All rights reserved.
        </Typography>
      </Box>
    </Box>
  );
};

export default AboutUsPage;
