import React, { useEffect, useState } from 'react';
import { useJsApiLoader, GoogleMap, Marker, InfoWindow, DirectionsRenderer } from '@react-google-maps/api';
import { Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, TextField, FormGroup, FormControlLabel, Checkbox, Alert } from '@mui/material'; // Import DialogContentText

const SearchMap = () => {
  const { isLoaded, loadError } = useJsApiLoader({
    googleMapsApiKey: 'AIzaSyBCOK5PNJk7qVS9ajhD1-0ZmS-hOApa2Vk',
    libraries: ['places'],
  });

  const [map, setMap] = useState(null);
  const [selectedCity, setSelectedCity] = useState(null);
  const [directions, setDirections] = useState(null);
  const [searchText, setSearchText] = useState('');
  const [autocompleteService, setAutocompleteService] = useState(null);
  const [autocompleteSessionToken, setAutocompleteSessionToken] = useState(null);
  const [autocompleteResults, setAutocompleteResults] = useState([]);
  const [open, setOpen] = useState(false);
  const [showAlert, setShowAlert] = useState(false); // State for alert

  useEffect(() => {
    if (isLoaded) {
      setAutocompleteService(new window.google.maps.places.AutocompleteService());
      setAutocompleteSessionToken(new window.google.maps.places.AutocompleteSessionToken());
    }
  }, [isLoaded]);

  useEffect(() => {
    if (autocompleteService && searchText !== '') {
      autocompleteService.getPlacePredictions(
        {
          input: searchText,
          sessionToken: autocompleteSessionToken,
          componentRestrictions: { country: 'TR' },
          types: ['(cities)']
        },
        (results, status) => {
          if (status === window.google.maps.places.PlacesServiceStatus.OK && results) {
            setAutocompleteResults(results);
          } else {
            setAutocompleteResults([]);
          }
        }
      );
    }
  }, [autocompleteService, searchText, autocompleteSessionToken]);

  useEffect(() => {
    if (selectedCity) {
      setAutocompleteResults([]); // Menüyü kapat
    }
  }, [selectedCity]);

  const handleSearchChange = (event) => {
    setSearchText(event.target.value);
    if (event.target.value === '') {
      setAutocompleteResults([]);
    }
  };

  const handleAutocompleteSelect = (placeId, description) => {
    const service = new window.google.maps.places.PlacesService(map);
    service.getDetails(
      {
        placeId: placeId,
        sessionToken: autocompleteSessionToken,
      },
      (result, status) => {
        if (status === window.google.maps.places.PlacesServiceStatus.OK) {
          const selectedLocation = {
            name: result.name,
            location: {
              lat: result.geometry.location.lat(),
              lng: result.geometry.location.lng(),
            },
          };
          setSelectedCity(selectedLocation);
          setMap((prevMap) => {
            prevMap.panTo(selectedLocation.location);
            return prevMap;
          });
          setSearchText(description);
        }
      }
    );
  };

  const handleMarkerClick = () => {
    setSelectedCity(null);
  };

  const handleClose = () => {
    setOpen(false);
  };

  const handleNewTripClick = () => {
    if (!selectedCity) {
      // Şehir seçilmediyse alert göster
      setShowAlert(true);
      setTimeout(() => setShowAlert(false), 1000);
    } else {
      setOpen(true);
    }
  };

  if (loadError) return <div>Error: It cannot loaded</div>;
  if (!isLoaded) return <div>Loading...</div>;

  return (
    <div style={{ position: 'relative', flexDirection: 'column', alignItems: 'center', height: '100vh', width: '71vw' }}>
      <div style={{ position: 'relative', left: 0, top: 0, height: '100%', width: '100%' }}>
        <TextField
          label="Where do you want to go?"
          value={searchText}
          onChange={handleSearchChange}
          placeholder="Search a city in Turkey..."
          style={{ marginBottom: '20px', padding: '5px', width: '30%' }}
        />
        <ul style={{
          listStyleType: 'none',
          padding: 0,
          margin: 0,
          maxHeight: '200px',
          overflowY: 'scroll',
          width: '50%',
          display: autocompleteResults.length > 0 ? 'block' : 'none'
        }}>
          {autocompleteResults.map((result) => (
            <li
              key={result.place_id}
              onClick={() => handleAutocompleteSelect(result.place_id, result.description)}
              style={{ padding: '5px', cursor: 'pointer' }}
            >
              {result.description}
            </li>
          ))}
        </ul>
        <Button variant="contained" onClick={handleNewTripClick} style={{ marginLeft: '690px', marginBottom: '20px' }}>New Trip</Button>
        {showAlert && !selectedCity && (
          <Alert style={{ position: 'absolute', top: -5, right: '10%'}} severity="warning">
            You should choose city first.
          </Alert>
        )}
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

          {directions && <DirectionsRenderer directions={directions} />}

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
      <Dialog open={open} onClose={handleClose}>
        <DialogTitle>Plan your next adventure</DialogTitle>
        <DialogContent>
          <DialogContentText >
            Select the kind of activities you want to do in <b>{selectedCity ? selectedCity.name : 'selected city'}</b>.
          </DialogContentText>
          <FormGroup>
          <FormControlLabel
              control={<Checkbox />}
              label="Beach"
            />
            <FormControlLabel
              control={<Checkbox />}
              label="Camping"
            />
            <FormControlLabel
              control={<Checkbox />}
              label="Cultural"
            />
            <FormControlLabel
              control={<Checkbox />}
              label="Adventure"
            />
            <FormControlLabel
              control={<Checkbox />}
              label="Road Trips"
            />
          </FormGroup>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose}>Cancel</Button>
          <Button onClick={handleClose}>Create New Trip</Button>
        </DialogActions>
      </Dialog>

    </div>
  );
};

export default SearchMap;
