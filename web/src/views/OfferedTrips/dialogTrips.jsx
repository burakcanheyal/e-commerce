import React, { useEffect, useState } from 'react';
import { useJsApiLoader, GoogleMap, Marker, InfoWindow, DirectionsRenderer } from '@react-google-maps/api';
import { Typography, TextField, Button, Box, Rating, List, ListItem, ListItemText } from '@mui/material';
import axios from 'axios';

const DialogTrips = ({ places }) => {
  const { isLoaded, loadError } = useJsApiLoader({
    googleMapsApiKey: 'AIzaSyBCOK5PNJk7qVS9ajhD1-0ZmS-hOApa2Vk',
    libraries: ['places'],
  });

  const [map, setMap] = useState(null);
  const [selectedPlace, setSelectedPlace] = useState(null);
  const [directions, setDirections] = useState(null);

  const [feedbacks, setFeedbacks] = useState([
    { user: 'Burak Can', feedback: 'That is a great route!', rating: 5 },
    { user: 'Çağrı', feedback: 'Offered locations are incredible!', rating: 5 }
  ]);

  const [newFeedback, setNewFeedback] = useState('');
  const [newRating, setNewRating] = useState(5);
  const [username, setUsername] = useState('');

  useEffect(() => {
    const fetchProfileDetails = async () => {
      try {
        const accessToken = localStorage.getItem('AccessToken');
        const response = await axios.get('http://localhost:8001/profil/', {
          headers: {
            'Authentication': `${accessToken}`
          }
        });
        setUsername(response.data.username);
      } catch (error) {
        console.error('Error fetching username:', error);
      }
    };
    fetchProfileDetails();
  }, []);

  useEffect(() => {
    if (isLoaded) {
      handleCalculateRoute();
    }
  }, [isLoaded]);

  const renderDirections = () => {
    if (directions) {
      return <DirectionsRenderer directions={directions} />;
    }
    return null;
  };

  const handleCalculateRoute = () => {
    if (places.length < 2) return;

    const waypoints = places.slice(1, places.length).map(place => ({
      location: place.coords,
      stopover: true,
    }));

    const origin = places[0].coords;
    const destination = places[places.length - 1].coords;

    const directionsService = new window.google.maps.DirectionsService();
    directionsService.route(
      {
        origin,
        destination,
        waypoints,
        travelMode: window.google.maps.TravelMode.DRIVING,
      },
      (result, status) => {
        if (status === window.google.maps.DirectionsStatus.OK) {
          setDirections(result);
        } else {
          console.error(`Error fetching directions ${result}`);
        }
      }
    );
  };

  const handleFeedbackSubmit = () => {
    const newFeedbackObj = { user: username, feedback: newFeedback, rating: newRating };
    setFeedbacks([newFeedbackObj, ...feedbacks]);
    setNewFeedback('');
    setNewRating(5);
  };

  if (loadError) return <div>Error: It cannot loaded</div>;
  if (!isLoaded) return <div>Loading...</div>;

  return (
    <div style={{ display: 'flex', flexDirection: 'row', height: '100vh' }}>
      <div style={{ flex: 1, display: 'flex', flexDirection: 'column' }}>
        <div>
          <Typography variant="h4" align="center" gutterBottom>User Feedbacks</Typography>
          <Box sx={{ maxHeight: '50%', overflowY: 'auto', padding: '0 10px' }}>
            {feedbacks.map((feedback, index) => (
              <div key={index} style={{ marginBottom: '10px' }}>
                <Typography variant="body1" gutterBottom>{feedback.user}: {feedback.feedback}</Typography>
                <Rating value={feedback.rating} readOnly />
              </div>
            ))}
          </Box>
          <TextField
            label="Your Feedback"
            variant="outlined"
            fullWidth
            margin="normal"
            value={newFeedback}
            onChange={(e) => setNewFeedback(e.target.value)}
          />
          <Rating
            name="new-feedback-rating"
            value={newRating}
            precision={1}
            onChange={(event, newValue) => setNewRating(newValue)}
          />
          <Button variant="contained" onClick={handleFeedbackSubmit}>Submit Feedback</Button>
        </div>
        <div style={{ marginTop: 'auto' }}>
          <Box sx={{ maxHeight: '50%', overflowY: 'auto' }}>
            <List>
              {places.map((place, index) => (
                <ListItem key={index}>
                  <ListItemText primary={`${index + 1} - ${place.name}`} secondary={place.details} />
                </ListItem>
              ))}
            </List>
          </Box>
        </div>
      </div>
      <div style={{ flex: 2 }}>
        <GoogleMap
          center={places[0].coords}
          zoom={12}
          mapContainerStyle={{ width: '100%', height: '100vh' }}
          options={{
            zoomControl: true,
            mapTypeControl: true,
            streetViewControl: true,
            fullscreenControl: true,
          }}
          onLoad={map => setMap(map)}
        >
          {renderDirections()}

          {selectedPlace && (
            <InfoWindow
              position={selectedPlace.coords}
              onCloseClick={() => setSelectedPlace(null)}
            >
              <div>
                <h2>{selectedPlace.name}</h2>
                <p>{selectedPlace.details}</p>
              </div>
            </InfoWindow>
          )}

        </GoogleMap>
      </div>
    </div>
  );
};

export default DialogTrips;
