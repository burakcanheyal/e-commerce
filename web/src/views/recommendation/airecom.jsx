import React, { useEffect, useState } from 'react';
import { useJsApiLoader, GoogleMap, Marker, InfoWindow } from '@react-google-maps/api';
import axios from 'axios';
import { Box, Typography } from '@mui/material';
import { RichTreeView } from '@mui/x-tree-view/RichTreeView';

const RecommendationPage = () => {
  const { isLoaded, loadError } = useJsApiLoader({
    googleMapsApiKey: 'AIzaSyBCOK5PNJk7qVS9ajhD1-0ZmS-hOApa2Vk', // API anahtarınızı buraya ekleyin
    libraries: ['places'],
  });

  const [map, setMap] = useState(null);
  const [selectedCity, setSelectedCity] = useState(null);
  const [recommendationData, setRecommendationData] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const accessToken = localStorage.getItem('AccessToken');
        const response = await axios.get('http://localhost:8001/trips/recommendation/', {
          headers: {
            'Authentication': accessToken
          }
        });

        const data = response.data;
        setRecommendationData(data);

        if (data && data.trip.length > 0) {
          const firstPlace = data.trip[0];
          const selectedLocation = {
            name: firstPlace.name,
            location: { lat: parseFloat(firstPlace.lat), lng: parseFloat(firstPlace.lng) }
          };
          setSelectedCity(selectedLocation);
          if (map) {
            map.panTo(selectedLocation.location);
          }
        }
      } catch (error) {
        if (error.response && error.response.status === 400) {
          setRecommendationData(null);
        }
        console.error('Error fetching recommendation data:', error);
      }
    };

    fetchData();
  }, [map]);

  if (loadError) return <div>Error: It cannot loaded</div>;
  if (!isLoaded) return <div>Loading...</div>;

  return (
    <div className="container">
      <div style={{ display: 'flex', flexDirection: 'row' }}>
        <div style={{ flex: 1 }}>
          <Box sx={{ height: 220, flexGrow: 1, maxWidth: 400 }}>
            {recommendationData ? (
              <RichTreeView items={recommendationData.trip.map((place, index) => ({
                id: `${place.name}-${index}`,
                label: place.name,
                children: [
                  { id: `${place.name}-${index}-desc`, label: `Description: ${place.description}` },
                  { id: `${place.name}-${index}-coords`, label: `Latitude: ${place.lat}, Longitude: ${place.lng}` }
                ]
              }))} />
            ) : (
              <Typography variant="body1">You should first submit the survey questions to be able to see the AI recommendation routes.</Typography>
            )}
          </Box>
        </div>
        <div style={{ flex: 2 }}>
          <div style={{ position: 'relative', flexDirection: 'column', alignItems: 'center', height: '100vh', width: '100%' }}>
            <GoogleMap
              center={selectedCity ? selectedCity.location : { lat: 41.0082, lng: 28.9784 }}
              zoom={9}
              mapContainerStyle={{ width: '100%', height: '100%' }}
              options={{
                zoomControl: true,
                mapTypeControl: true,
                streetViewControl: true,
                fullscreenControl: true,
              }}
              onLoad={(map) => setMap(map)}
            >
              {recommendationData && recommendationData.trip.map((place, index) => (
                <Marker
                  key={index}
                  position={{ lat: parseFloat(place.lat), lng: parseFloat(place.lng) }}
                  onClick={() => {
                    setSelectedCity({ name: place.name, location: { lat: parseFloat(place.lat), lng: parseFloat(place.lng) }});
                  }}
                />
              ))}
              {selectedCity && (
                <InfoWindow position={selectedCity.location}>
                  <div>{selectedCity.name}</div>
                </InfoWindow>
              )}
            </GoogleMap>

          </div>
        </div>
      </div>
    </div>
  );
};

export default RecommendationPage;
