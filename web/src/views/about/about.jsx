import React from 'react';
import { Typography, Box } from '@mui/material';
import CTISLogo from '/@/assets/images/logos/30yearsCTISvBlack.png';

const AboutUsPage = () => {
  const handleCTISLink = () => {
    window.open('https://www.ctis.bilkent.edu.tr/ctis_seniorProject.php?id=5010', '_blank');
  };

  return (
    <Box sx={{ textAlign: 'center', paddingTop: 7 }}>
      <Typography variant="body1" sx={{ marginTop: 5 }}>
        Exporia is a revolutionary multi-platform application designed to deliver personalized travel experiences.
        Its core functionality includes allowing users to create customized tour packages based on their unique preferences and personality traits.
        By finding similarities between individuals on different tours, Exporia facilitates the creation of cost-effective, tailor-made itineraries.
      </Typography>
      <Box sx={{ marginTop: 5 }}>
        <img src={CTISLogo} alt="30yearsCTISvBlack Logo" style={{ width: 200, cursor: 'pointer' }} onClick={handleCTISLink} />
      </Box>
    </Box>
  );
};

export default AboutUsPage;
