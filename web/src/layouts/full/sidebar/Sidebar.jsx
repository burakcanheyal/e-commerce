import React from 'react';
import { useMediaQuery, Box, Drawer } from '@mui/material';
import Logo from '../shared/logo/Logo.jsx';
import SidebarItems from './SidebarItems.jsx';

const Sidebar = (props) => {
    const lgUp = useMediaQuery((theme) => theme.breakpoints.up("lg"));
    const sidebarWidth = '270px';

    if (lgUp) {
        return (
            <Box
                sx={{
                    width: sidebarWidth,
                    flexShrink: 0,
                }}
            >
                <Drawer
                    anchor="left"
                    open={props.isSidebarOpen}
                    variant="permanent"
                    PaperProps={{
                        sx: {
                            width: sidebarWidth,
                            boxSizing: 'border-box',
                        },
                    }}
                >
                    <Box
                        sx={{
                            height: '100%',
                        }}
                    >
                        <Box px={3}>
                            <Logo />
                        </Box>
                        <Box>
                            <SidebarItems />
                        </Box>
                    </Box>
                </Drawer>
            </Box>
        );
    }

    return (
        <Drawer
            anchor="left"
            open={props.isMobileSidebarOpen}
            onClose={props.onSidebarClose}
            variant="temporary"
            PaperProps={{
                sx: {
                    width: sidebarWidth,
                    boxShadow: (theme) => theme.shadows[8],
                },
            }}
        >
            <Box px={2}>
                <Logo />
            </Box>
            <SidebarItems />
        </Drawer>
    );
};

export default Sidebar;
