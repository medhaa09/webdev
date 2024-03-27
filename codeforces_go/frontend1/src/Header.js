import * as React from 'react';
import PropTypes from 'prop-types';
import Toolbar from '@mui/material/Toolbar';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import { useNavigate } from 'react-router-dom';

function Header(props) {
 
  const { sections, title } = props;
  const navigate = useNavigate();
  function logoutHandler(){
    navigate("/user/signup")
    localStorage.removeItem("token")
    localStorage.removeItem("refreshToken")
  }
  function homeHandler(){
    navigate("/")
   
  }

  return (
    <div >
      <Toolbar sx={{ borderBottom: 2, borderColor: 'white' }}>
        <Button size="large" style={{ color:'white'  }}sx={{
        '&:hover': {
            backgroundColor: 'rgba(255, 255, 255, 0.3)' 
        }
    }} onClick={() => homeHandler()}  >Home</Button>
        <Button onClick={() => logoutHandler()}   size="large" style={{color:'white'}} sx={{
        '&:hover': {
            backgroundColor: 'rgba(255, 255, 255, 0.3)' 
        }
    }} >
          Log Out
        </Button>
      </Toolbar>
      <Toolbar
        component="nav"
        variant="dense"
        sx={{ justifyContent: 'space-between', overflowX: 'auto' }}
      >
      </Toolbar>
    </div>
  );
}


export default Header;
