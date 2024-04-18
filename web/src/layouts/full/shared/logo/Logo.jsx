import { Link } from 'react-router-dom';
import exporiaLogo from '/@/assets/images/logos/exporia-high-resolution-logo-transparent.png';
import { styled } from '@mui/material';

const LinkStyled = styled(Link)(() => ({
  height: '70px',
  width: '180px',
  overflow: 'hidden',
  display: 'block',
}));

const Logo = () => {
  return (
    <LinkStyled to="/">
      <img src={exporiaLogo} alt="Exporia Logo" height={43} style={{ marginTop: '20px' }} />
    </LinkStyled>
  )
};

export default Logo;
