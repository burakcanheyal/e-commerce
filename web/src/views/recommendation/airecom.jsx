import React, { useEffect, useState } from 'react';
import { useJsApiLoader, GoogleMap, Marker, InfoWindow } from '@react-google-maps/api';
import axios from 'axios';

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

        // İlk parametrenin konumunu belirle
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
        console.error('Error fetching recommendation data:', error);
      }
    };

    fetchData();
  }, [map]);

  if (loadError) return <div>Error: It cannot loaded</div>;
  if (!isLoaded) return <div>Loading...</div>;

  return (
    <div className="container">
      <h1>AI Recommendation</h1>
      {recommendationData && (
        <ul>
          {recommendationData.trip.map((place, index) => (
            <li key={index}>
              <h2>{place.name}</h2>
              <p>{place.description}</p>
              <p>Latitude: {place.lat}, Longitude: {place.lng}</p>
            </li>
          ))}
        </ul>
      )}
      <div style={{ position: 'relative', flexDirection: 'column', alignItems: 'center', height: '50vh', width: '100%' }}>
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
  );
};

export default RecommendationPage;
