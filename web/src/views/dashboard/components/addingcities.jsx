import React, { useEffect, useState } from 'react';
import { useJsApiLoader, GoogleMap, Marker, InfoWindow, DirectionsRenderer } from '@react-google-maps/api';
import axios from 'axios';

const SearchMap = () => {
  const { isLoaded, loadError } = useJsApiLoader({
    googleMapsApiKey: 'AIzaSyBCOK5PNJk7qVS9ajhD1-0ZmS-hOApa2Vk',
    libraries: ['places'],
  });

  const [map, setMap] = useState(null);
  const [selectedCity, setSelectedCity] = useState(null);

  useEffect(() => {
    const fetchCityRecommendation = async () => {
      try {
        const accessToken = localStorage.getItem('AccessToken');
        const response = await axios.get('http://localhost:8001/trips/recommendation/', {
          headers: {
            'Authentication': accessToken
          }
        });

        const { lat, lng } = response.data;
        const selectedLocation = {
          name: response.data.name,
          location: { lat: parseFloat(lat), lng: parseFloat(lng) }
        };

        setSelectedCity(selectedLocation);
        if (map) {
          map.panTo(selectedLocation.location);
        }
      } catch (error) {
        console.error('Error fetching city recommendation:', error);
      }
    };

    if (isLoaded) {
      fetchCityRecommendation();
    }
  }, [isLoaded, map]);

  const handleMarkerClick = () => {
    setSelectedCity(null);
  };

  if (loadError) return <div>Error: It cannot loaded</div>;
  if (!isLoaded) return <div>Loading...</div>;

  return (
    <div style={{ position: 'relative', flexDirection: 'column', alignItems: 'center', height: '100vh', width: '100%' }}>
      <GoogleMap
        center={selectedCity ? selectedCity.location : { lat: 41.0082, lng: 28.9784 }}
        zoom={10}
        mapContainerStyle={{ width: '100%', height: '100%' }}
        options={{
          zoomControl: true,
          mapTypeControl: true,
          streetViewControl: true,
          fullscreenControl: true,
        }}
        onLoad={(map) => setMap(map)}
      >
        {selectedCity && (
          <Marker
            position={selectedCity.location}
            onClick={handleMarkerClick}
          />
        )}

        {selectedCity && (
          <InfoWindow
            position={selectedCity.location}
            onCloseClick={() => setSelectedCity(null)}
          >
            <div>
              <h2>{selectedCity.name}</h2>
            </div>
          </InfoWindow>
        )}
      </GoogleMap>
    </div>
  );
};

export default SearchMap;
