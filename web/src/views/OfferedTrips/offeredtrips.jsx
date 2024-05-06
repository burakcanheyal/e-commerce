import React, { useState } from 'react';
import {
  Grid,
  Typography,
  Button,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
} from '@mui/material';
import DashboardCard from '../../components/shared/DashboardCard';
import izmirImage from '../../assets/images/izmir.jpg';
import antalyaImage from '../../assets/images/antalya.jpg';
import kapadokyaImage from '../../assets/images/kapadokya.jpg';
import istanbulImage from '../../assets/images/istanbul.jpg';
import ankaraImage from '../../assets/images/ankara.jpg';
import canakkaleImage from '../../assets/images/canakkale.jpg';
import balikesirImage from '../../assets/images/balikesir.jpg';
import samsunImage from '../../assets/images/samsun.jpg';
import eskisehirImage from '../../assets/images/eskisehir.jpg';
import DialogTrips from './dialogTrips.jsx';


const OfferedTrips = () => {
  const [trip, setTrip] = useState('1');
  const [openDialog, setOpenDialog] = useState(false);



  const handleDialogOpen = (tripIndex) => {
    setOpenDialog(true);
    setTrip(tripIndex.toString());
  };

  const handleDialogClose = () => {
    setOpenDialog(false);
  };

  const handleAddToCart = async (trip) => {
    const productId = trip.product_id;
    const accessToken = localStorage.getItem('AccessToken');
    const requestData = {
      product_id: productId,
      quantity: 1
    };

    try {
      const requestOptions = {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authentication': accessToken
        },
        body: JSON.stringify(requestData)
      };
      const response = await fetch('http://localhost:8001/order/', requestOptions);
      if (!response.ok) {
        throw new Error('HTTP error ' + response.status);
      }
      alert('Product added to cart successfully!');
    } catch (error) {
      console.error('Error adding product to cart:', error.message);
      alert('An error occurred while adding product to cart. Please try again later.');
    }
  };

  const tripPackages = [
    {
      product_id: 1,
      name: 'İzmir Tour',
      description: 'Explore the historical sites, coastal areas, and culinary stops in Izmir, the pearl of the Aegean.',
      image: izmirImage,
      places: [
        {id:1, name: 'Alsancak', coords: { lat: 38.4325, lng: 27.1503 }, details: 'Alsancak is a vibrant neighborhood in İzmir known for its entertainment and nightlife.' },
        {id:2, name: 'Konak Square', coords: { lat: 38.4192, lng: 27.1287 }, details: 'Konak Square is a popular public square in İzmir.' },
        {id:3, name: 'Kordon', coords: { lat: 38.4161, lng: 27.1287 }, details: 'Kordon is a promenade along the waterfront of İzmir.' },
        {id:4, name: 'Asansör', coords: { lat: 38.4267, lng: 27.1419 }, details: 'Asansör is a historical elevator in İzmir with panoramic views of the city.' },
        {id:5, name: 'Kemeraltı Market', coords: { lat: 38.4213, lng: 27.1281 }, details: 'Kemeraltı Market is a bustling bazaar in İzmir offering a variety of goods and foods.' },
        {id:6, name: 'Ephesus Ancient City', coords: { lat: 37.9499, lng: 27.3697 }, details: 'Ephesus Ancient City is an archaeological site near Selçuk, İzmir Province, Turkey. It is one of the best-preserved ancient cities in the world.' },
      ],
    },
    {
      product_id: 2,
      name: 'Antalya Tour',
      description: 'Enjoy the sun and sea in Antalya, visit historical sites, and dine with magnificent views.',
      image: antalyaImage,
      places: [
        {id:7, name: 'Konyaaltı Beach', coords: { lat: 36.8636, lng: 30.6956 }, details: 'Konyaaltı Beach is one of the most popular beaches in Antalya, known for its crystal-clear waters and stunning views of the Taurus Mountains.' },
        {id:8, name: 'Old Town (Kaleiçi)', coords: { lat: 36.8854, lng: 30.7057 }, details: 'Old Town, also known as Kaleiçi, is the historic heart of Antalya with narrow cobblestone streets, historic Ottoman houses, and charming cafes.' },
        {id:9, name: 'Düden Waterfalls', coords: { lat: 36.8625, lng: 30.7708 }, details: 'Düden Waterfalls are a series of waterfalls located northeast of Antalya, offering breathtaking natural scenery and walking trails.' },
        {id:10, name: 'Antalya Museum', coords: { lat: 36.8841, lng: 30.7057 }, details: 'Antalya Museum is one of Turkey\'s largest museums, showcasing artifacts from the ancient Lycian, Pamphylian, Hellenistic, Roman, and Byzantine periods.' },
        {id:11,  name: 'Termessos Ancient City', coords: { lat: 37.0179, lng: 30.5214 }, details: 'Termessos is an ancient city located in the Taurus Mountains, known for its well-preserved ruins and stunning mountain views.' },
        {id:12, name: 'Perge Ancient City', coords: { lat: 36.9722, lng: 30.9171 }, details: 'Perge is an ancient city located 15 kilometers east of Antalya, known for its well-preserved Roman ruins, including a theater, stadium, and agora.' },
      ],
    },
    {
      product_id: 3,
      name: 'Cappadocia Tour',
      description: 'Take a balloon tour in Cappadocia, famous for its fairy chimneys and unique geography, explore underground cities, and try local delicacies.',
      image: kapadokyaImage,
      places: [
        {id:13, name: 'Göreme Open Air Museum', coords: { lat: 38.6437, lng: 34.8303 }, details: 'Göreme Open Air Museum is a UNESCO World Heritage Site, famous for its rock-cut churches with stunning frescoes dating back to the 10th century.' },
        {id:14, name: 'Ürgüp', coords: { lat: 38.6303, lng: 34.9140 }, details: 'Ürgüp is a town in Cappadocia known for its cave hotels, vineyards, and spectacular panoramic views of the fairy chimneys.' },
        {id:15, name: 'Love Valley', coords: { lat: 38.6443, lng: 34.8469 }, details: 'Love Valley, named for its phallic-shaped rock formations, offers some of the most iconic views of Cappadocia and is a popular spot for hiking and photography.' },
        {id:16, name: 'Derinkuyu Underground City', coords: { lat: 38.3763, lng: 34.8485 }, details: 'Derinkuyu Underground City is one of the largest and deepest underground cities in Cappadocia, featuring a complex network of tunnels, rooms, and passages.' },
        {id:17, name: 'Hot Air Balloon Ride', coords: { lat: 38.6431, lng: 34.8303 }, details: 'A hot air balloon ride over Cappadocia is a must-do experience, offering panoramic views of the surreal landscape dotted with fairy chimneys, rock formations, and ancient settlements.' },
        {id:18, name: 'Zelve Open Air Museum', coords: { lat: 38.6525, lng: 34.8355 }, details: 'Zelve Open Air Museum is an ancient cave settlement in Cappadocia, showcasing rock-cut churches, monasteries, and dwellings carved into the soft volcanic rock.' },
      ],
    },
    {
      product_id: 4,
      name: 'İstanbul Tour',
      description: 'Discover the vibrant culture, rich history, and stunning landmarks of Istanbul, the crossroads of Europe and Asia.',
      image: istanbulImage,
      places: [
        { id: 19, name: 'Hagia Sophia', coords: { lat: 41.0082, lng: 28.9784 }, details: 'Hagia Sophia is a former Greek Orthodox Christian cathedral, later an Ottoman imperial mosque, and now a museum in Istanbul.' },
        { id: 20, name: 'Topkapı Palace', coords: { lat: 41.0115, lng: 28.9833 }, details: 'Topkapı Palace is a large palace in Istanbul, Turkey, that was the primary residence of the Ottoman sultans for approximately 400 years.' },
        { id: 21, name: 'Blue Mosque', coords: { lat: 41.0054, lng: 28.9760 }, details: 'The Blue Mosque, also known as the Sultan Ahmed Mosque, is a historic mosque located in Istanbul, Turkey.' },
        { id: 22, name: 'Grand Bazaar', coords: { lat: 41.0105, lng: 28.9685 }, details: 'The Grand Bazaar in Istanbul is one of the largest and oldest covered markets in the world, with 61 covered streets and over 4,000 shops.' },
        { id: 23, name: 'Bosphorus Cruise', coords: { lat: 41.0390, lng: 29.0251 }, details: 'A Bosphorus cruise offers breathtaking views of Istanbul\'s skyline, historic landmarks, and waterfront mansions along the Bosphorus Strait.' },
        { id: 24, name: 'Galata Tower', coords: { lat: 41.0256, lng: 28.9747 }, details: 'Galata Tower is a medieval stone tower in the Galata/Karaköy quarter of Istanbul, Turkey.' },
      ],
    },
    {
      product_id: 5,
      name: 'Ankara Tour',
      description: 'Experience the political, cultural, and historical heart of Turkey with a tour of Ankara, its capital city.', // Yeni eklendi
      image: ankaraImage,
      places: [
        { id: 25, name: 'Anıtkabir', coords: { lat: 39.9256, lng: 32.8357 }, details: 'Anıtkabir is the mausoleum of Mustafa Kemal Atatürk, the founder of the Republic of Turkey, located in Ankara.' }, // Yeni eklendi
        { id: 26, name: 'Kocatepe Mosque', coords: { lat: 39.9289, lng: 32.8524 }, details: 'Kocatepe Mosque is the largest mosque in Ankara, known for its imposing architecture and grand interior.' }, // Yeni eklendi
        { id: 27, name: 'Ankara Citadel', coords: { lat: 39.9439, lng: 32.8544 }, details: 'Ankara Citadel is a historic citadel located in Ankara, Turkey, dating back to ancient times and offering panoramic views of the city.' }, // Yeni eklendi
        { id: 28, name: 'Museum of Anatolian Civilizations', coords: { lat: 39.9413, lng: 32.8641 }, details: 'The Museum of Anatolian Civilizations is located in Ankara and houses a rich collection of artifacts from the Neolithic Age to the Byzantine period.' }, // Yeni eklendi
        { id: 29, name: 'Atakule Tower', coords: { lat: 39.8973, lng: 32.8639 }, details: 'Atakule Tower is a prominent landmark in Ankara, offering panoramic views of the city from its observation deck and revolving restaurant.' }, // Yeni eklendi
        { id: 30, name: 'Gençlik Parkı', coords: { lat: 39.9175, lng: 32.8517 }, details: 'Gençlik Parkı is a large public park in Ankara, popular for picnics, leisure activities, and recreational facilities.' }, // Yeni eklendi
      ],
    },
    {
      product_id: 6,
      name: 'Çanakkale Tour',
      description: 'Discover the historical sites and natural beauty of Çanakkale, a city rich in culture and significance.', // Yeni eklendi
      image:canakkaleImage,
      places: [
        { id: 31, name: 'Troy Ancient City', coords: { lat: 39.9572, lng: 26.2386 }, details: 'Troy Ancient City is a UNESCO World Heritage Site, known for its mythological significance as the setting of the Trojan War.' }, // Yeni eklendi
        { id: 32, name: 'Gallipoli Peninsula', coords: { lat: 40.2470, lng: 26.3437 }, details: 'Gallipoli Peninsula is a historic site of significant importance, known for the Gallipoli Campaign during World War I.' }, // Yeni eklendi
        { id: 33, name: 'Çanakkale Martyrs\' Memorial', coords: { lat: 40.1464, lng: 26.4086 }, details: 'Çanakkale Martyrs\' Memorial is a commemorative monument dedicated to the Turkish soldiers who participated in the Battle of Gallipoli.' }, // Yeni eklendi
        { id: 34, name: 'Bozcaada', coords: { lat: 39.8222, lng: 26.0444 }, details: 'Bozcaada is a picturesque island known for its charming villages, vineyards, and pristine beaches.' }, // Yeni eklendi
        { id: 35, name: 'Assos', coords: { lat: 39.4894, lng: 26.3344 }, details: 'Assos is a historical town with ancient ruins, including the Temple of Athena, offering stunning views of the Aegean Sea.' }, // Yeni eklendi
        { id: 36, name: 'Kaz Mountains', coords: { lat: 39.6769, lng: 26.8265 }, details: 'Kaz Mountains, also known as Mount Ida, is a mountain range in northwestern Turkey, famous for its natural beauty and diverse flora.' }, // Yeni eklendi
      ],
    },
    {
      product_id: 7,
      name: 'Balıkesir Tour',
      description: 'Explore the stunning coastline, pristine beaches, and historical sites of Balıkesir, a province on the Aegean coast of Turkey.',
      image: balikesirImage,
      places: [
        { id: 37, name: 'Assos Ancient City', coords: { lat: 39.4891, lng: 26.3329 }, details: 'Assos Ancient City is an archaeological site located in the Çanakkale Province of Turkey. It is known for its Temple of Athena, built on a hill overlooking the Aegean Sea.' },
        { id: 38, name: 'Ayvalık', coords: { lat: 39.3173, lng: 26.6954 }, details: 'Ayvalık is a seaside town known for its olive oil production, historic architecture, and charming streets lined with colorful houses.' },
        { id: 39, name: 'Cunda Island', coords: { lat: 39.4142, lng: 26.9713 }, details: 'Cunda Island, also known as Alibey Island, is a small island connected to the mainland by a causeway. It is famous for its narrow streets, traditional Greek houses, and seafood restaurants.' },
        { id: 40, name: 'Troy Ancient City', coords: { lat: 39.9574, lng: 26.2386 }, details: 'Troy Ancient City, located in the Çanakkale Province, is one of the most famous archaeological sites in the world. It is believed to be the setting of the Trojan War described in Homer\'s epic poems.' },
        { id: 41, name: 'Mount Ida (Kaz Dağı)', coords: { lat: 39.7976, lng: 26.7813 }, details: 'Mount Ida, also known as Kaz Dağı in Turkish, is a mountain range in northwestern Turkey. It is famous for its mythological significance as the birthplace of the Greek god Zeus and for its diverse flora and fauna.' },
        { id: 42, name: 'Edremit', coords: { lat: 39.5996, lng: 26.9291 }, details: 'Edremit is a district in the Balıkesir Province known for its thermal springs, olive groves, and scenic coastline.' },
      ],
    },
    {
      product_id: 8,
      name: 'Samsun Tour',
      description: 'Discover the Black Sea coast and rich history of Samsun, a vibrant city in northern Turkey.',
      image: samsunImage,
      places: [
        { id: 43, name: 'Amisos Hill', coords: { lat: 41.3006, lng: 36.3325 }, details: 'Amisos Hill is a historic site in Samsun, featuring ancient tombs, artifacts, and panoramic views of the city and the Black Sea.' },
        { id: 44, name: 'Bandırma Ferry', coords: { lat: 41.2870, lng: 36.3305 }, details: 'The Bandırma Ferry, docked in Samsun, is a replica of the ship that carried Mustafa Kemal Atatürk from Istanbul to Samsun in 1919, marking the start of the Turkish War of Independence.' },
        { id: 45, name: 'Amazon Village', coords: { lat: 41.3689, lng: 36.2998 }, details: 'The Amazon Village is a cultural complex in Samsun, offering insights into the legendary female warriors of antiquity, the Amazons, through exhibitions, performances, and workshops.' },
        { id: 46, name: 'Samsun Archaeology and Ethnography Museum', coords: { lat: 41.2863, lng: 36.3353 }, details: 'Samsun Archaeology and Ethnography Museum showcases artifacts from prehistoric, Hittite, Phrygian, Roman, Byzantine, and Ottoman periods, providing insights into the region\'s rich history and cultural heritage.' },
        { id: 47, name: 'Gazi Museum', coords: { lat: 41.2869, lng: 36.3325 }, details: 'Gazi Museum in Samsun is dedicated to Mustafa Kemal Atatürk, featuring exhibitions on his life, leadership, and contributions to the founding of the Republic of Turkey.' },
        { id: 48, name: 'İncesu Waterfall', coords: { lat: 41.1892, lng: 36.4338 }, details: 'İncesu Waterfall is a natural wonder near Samsun, offering scenic views, hiking trails, and picnic areas amidst lush greenery.' },
      ],
    },
    {
      product_id: 9,
      name: 'Eskişehir Tour',
      description: 'Experience Eskişehir, a city in northwestern Turkey, known for its Ottoman-era architecture, vibrant arts scene, and thermal hot springs.',
      image: eskisehirImage,
      places: [
        {id:49, name: 'Odunpazarı Historic District', coords: { lat: 39.7755, lng: 30.5211 }, details: 'Odunpazarı is a historic district of Eskişehir, Turkey. It is named after the Ottoman term for "wood market".' },
        {id:50, name: 'Porsuk River', coords: { lat: 39.7736, lng: 30.5206 }, details: 'The Porsuk River is a small river in Anatolia in north-central Turkey, flowing through the city of Eskişehir.' },
        {id:51, name: 'Eskişehir Clock Tower', coords: { lat: 39.7737, lng: 30.5252 }, details: 'The Eskişehir Clock Tower is a clock tower in Eskişehir, Turkey. The tower stands in the city center, in the square that was named after it, on the banks of the Porsuk River.' },
        {id:52, name: 'Lületaşı Museum', coords: { lat: 39.7677, lng: 30.5254 }, details: 'The Lületaşı Museum is a museum in Eskişehir, Turkey, dedicated to the local variety of opal known as lületaşı.' },
      ],
    },
  ];


  return (
    <div>

      <DashboardCard>
        <div>
          <Grid container spacing={5}>
            {tripPackages.map((trip, index) => (
              <Grid item xs={12} sm={6} md={4} key={index}>
                <TripPackage trip={trip} onDialogOpen={() => handleDialogOpen(index + 1)} onAddToCart={() => handleAddToCart(trip)} />
              </Grid>
            ))}
          </Grid>
        </div>
        <Dialog open={openDialog} onClose={handleDialogClose} fullWidth maxWidth="lg" >
          <DialogTitle>Trip Details</DialogTitle>
          <DialogContent dividers style={{ maxHeight: '80vh' }}>
            <DialogTrips places={tripPackages[parseInt(trip) - 1].places} product_id={tripPackages[parseInt(trip) - 1].product_id}/>
          </DialogContent>
          <DialogActions>
            <Button onClick={handleDialogClose} color="primary">Close</Button>
          </DialogActions>
        </Dialog>
      </DashboardCard>
    </div>
  );
};

const TripPackage = ({ trip, onDialogOpen, onAddToCart }) => {
  return (
    <div>
      <img src={trip.image} alt={trip.name} style={{ width: '100%', height: '200px', objectFit: 'cover' }} />
      <div style={{ padding: '10px' }}>
        <Typography variant="h6" style={{ marginBottom: '10px' }}>{trip.name}</Typography>
        <Typography variant="body1" style={{ marginBottom: '10px' }}>{trip.description}</Typography>
        <Button variant="contained" color="primary" onClick={onDialogOpen} style={{ marginRight: '10px' }}>Details</Button>
        <Button variant="contained" color="secondary" onClick={onAddToCart} >Add to Cart</Button>
      </div>
    </div>
  );
};

export default OfferedTrips;
