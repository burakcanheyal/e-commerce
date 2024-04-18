import React, { useEffect, useState } from 'react';
import { useJsApiLoader, GoogleMap, Marker, InfoWindow, DirectionsRenderer } from '@react-google-maps/api';

const DialogTrips = ({ places }) => {
  const { isLoaded, loadError } = useJsApiLoader({
    googleMapsApiKey: 'AIzaSyBCOK5PNJk7qVS9ajhD1-0ZmS-hOApa2Vk', // API anahtarınızı buraya ekleyin
    libraries: ['places'],
  });

  const [map, setMap] = useState(null);
  const [selectedPlace, setSelectedPlace] = useState(null);
  const [directions, setDirections] = useState(null);

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

  if (loadError) return <div>Error: It cannot loaded</div>;
  if (!isLoaded) return <div>Loading...</div>;

  return (
    <div style={{ position: 'relative', flexDirection: 'column', alignItems: 'center', height: '100vh', width: '100vw' }}>
      <div style={{ position: 'absolute', left: 0, top: 0, height: '100%', width: '100%' }}>
        <GoogleMap
          center={places[0].coords}
          zoom={12}
          mapContainerStyle={{ width: '70%', height: '100%' }}
          options={{
            zoomControl: true,
            mapTypeControl: true,
            streetViewControl: true,
            fullscreenControl: true,
          }}
          onLoad={map => setMap(map)}
        >
          {places.map((place, index) => (
            index === 0 ? (
              <Marker
                key={place.name}
                position={place.coords}
                onClick={() => setSelectedPlace(place)}
                label="A"
              />
            ) : index === places.length - 1 ? (
              <Marker
                key={place.name}
                position={place.coords}
                onClick={() => setSelectedPlace(place)}
                label="B"
              />
            ) : null
          ))}

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
