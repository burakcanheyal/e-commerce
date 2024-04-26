import React from 'react';
import { Grid, Box } from '@mui/material';

// components
import IntroPage from './components/IntroPage.jsx';
import PageContainer from "../../components/container/PageContainer.jsx";
import SearchMap from './components/addingcities.jsx';


const Dashboard = () => {
    return (
        <PageContainer title="Dashboard" description="this is Dashboard">
            <Box>
                <Grid container spacing={3}>
                    <Grid item xs={12} lg={8}>
                        <IntroPage />
                    </Grid>
                    <Grid item xs={12} lg={8}>
                    <br/><br/><br/><br/><br/>
                    </Grid>
                    <Grid item xs={12}>
                        <SearchMap />
                    </Grid>
                </Grid>
            </Box>
        </PageContainer>
    );
};

export default Dashboard;
