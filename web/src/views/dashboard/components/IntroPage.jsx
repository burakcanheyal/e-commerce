import React from 'react';
import { Typography, Button } from '@mui/material';
import { Carousel } from 'react-responsive-carousel';
import 'react-responsive-carousel/lib/styles/carousel.min.css';
import travelimg from '/@/assets/images/travel-photos-6.jpg';
import travelimg2 from '/@/assets/images/travel-photos-7.jpg';
import travelimg3 from '/@/assets/images/travel-photos-8.jpg';
import travelimg4 from '/@/assets/images/travel-photos-14.jpg';
import AddIcon from '@mui/icons-material/Add';
import RemoveRedEyeIcon from '@mui/icons-material/RemoveRedEye';
import { NavLink } from 'react-router-dom';

const IntroPage = () => {
  return (
    <div style={{ display: 'flex', alignItems: 'center' }}>
      <div style={{ flex: 1 }}>
        <Typography variant="h1" style={{ marginLeft:'-10px', marginTop: '-10px' }}>Discover effortless travel planning with us
        </Typography>
        <Typography variant="h3" style={{ marginLeft:'-10px', marginTop: '30px', textAlign: 'left' }}>Build,
          organize, and map your custom itineraries for road trips, powered by our trip planner AI. You can create a trip or see the offered trips based on your survey.
        </Typography>
        <div style={{ marginLeft: '55px', marginTop: '30px', display: 'flex', gap: '50px' }}>
          <Button variant="contained" startIcon={<AddIcon/>} >
            Add Trip
          </Button>
          <Button variant="contained" startIcon={<RemoveRedEyeIcon/>}
                  component={NavLink}
                  to="/offeredtrips"
          >
            View Offered Trips
          </Button>
        </div>
      </div>
      <div style={{ marginRight: '-370px', marginTop: '30px', display: 'flex', gap: '50px' }}>
        <Carousel width={600} showThumbs={false} showStatus={false} infiniteLoop={true} autoPlay={true} interval={5000}
                  transitionTime={100}>
          <div>
            <img src={travelimg} alt="Travel Photo 1" width={500} height={500} />
          </div>
          <div>
            <img src={travelimg2} alt="Travel Photo 2"  width={500} height={500} />
          </div>
          <div>
            <img src={travelimg3} alt="Travel Photo 3"  width={500} height={500} />
          </div>
          <div>
            <img src={travelimg4} alt="Travel Photo 4"  width={500} height={500} />
          </div>
        </Carousel>
      </div>
    </div>
  );
};

export default IntroPage;
