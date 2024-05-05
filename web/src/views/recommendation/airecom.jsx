import React, { useEffect, useState } from 'react';

const RecommendationPage = () => {
  const [recommendationData, setRecommendationData] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const accessToken = localStorage.getItem('AccessToken');

        const response = await fetch('http://localhost:8001/trips/recommendation/', {
          method: 'GET',
          headers: {
            'Authentication': `${accessToken}`
          }
        });

        if (!response.ok) {
          throw new Error('Network response was not ok');
        }

        const data = await response.json();
        setRecommendationData(data);
      } catch (error) {
        console.error('Error fetching recommendation data:', error);
      }
    };

    fetchData();
  }, []); // Bu etkileşim sadece bir kez çalışacak, bileşen yüklendiğinde

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
    </div>
  );
};

export default RecommendationPage;
