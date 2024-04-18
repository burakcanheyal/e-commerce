import React from 'react';
import MenuItems from './MenuItems.jsx';
import { useLocation } from 'react-router';
import { Box, List } from '@mui/material';
import NavItem from './NavItem/index.jsx';
import NavGroup from './NavGroup/NavGroup.jsx';

const SidebarItems = () => {
    const { pathname } = useLocation();
    const pathDirect = pathname;

    return (
        <Box sx={{ px: 3 }}>
            <List sx={{ pt: 0 }} className="sidebarNav">
                {MenuItems.map((item) => {
                    if (item.subheader) {
                        return <NavGroup item={item} key={item.subheader} />;
                    } else {
                        return <NavItem item={item} key={item.id} pathDirect={pathDirect} />;
                    }
                })}
            </List>
        </Box>
    );
};

export default SidebarItems;
