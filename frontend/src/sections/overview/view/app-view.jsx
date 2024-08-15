import { useState, useEffect } from 'react';

import { Button } from '@mui/material';
import Container from '@mui/material/Container';
import Grid from '@mui/material/Unstable_Grid2';
import Typography from '@mui/material/Typography';

import { SessionStatus } from "src/utils/session-status";

import { accountDefault } from 'src/_mock/account';

import AppTasks from '../app-tasks';
import AppWidgetSummary from '../app-widget-summary';
// import AppConversionRates from '../app-conversion-rates';



// ----------------------------------------------------------------------

export default function AppView() {
  const [account, setAccount] = useState(accountDefault);
  const [sessionStatus, setSessionStatus] = useState("");
  // const [updateStatus, setUpdateStatus] = useState(0);


  // account name 
  useEffect(() => {
    const getAccount = async () => {
      try {
        const response = await window.go.app.App.UserGetAccount();
        if (response == null) {
          return;
        }
        setAccount(response);
      } catch (error) {
        console.error('Error getting account:', error);
        setAccount(accountDefault);
      }
    };

    getAccount();
  }, []);


  // status of the session
  useEffect(() => {
    const getSessionStatus = async () => {
      try {
        const response = await window.go.app.App.UserGetStatus();
        if (response == null) {
          return;
        }
        setSessionStatus(response);
      } catch (error) {
        console.error('Error getting session status:', error);
        setSessionStatus(SessionStatus.OFFLINE);
      }
    };

    getSessionStatus();
  }, []);

  return (
    <Container maxWidth="xl">

      <Grid container spacing={0} sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h6" >
          Akao i {account.login.charAt(0).toUpperCase() + account.login.slice(1)} iny ô ? 👋
        </Typography>

        <Typography variant="caption" >
          {sessionStatus} 🚀
        </Typography>

        <Button onClick={() => setSessionStatus("working")} variant="outlined"  >Demarrer la session</Button>
      </Grid>

      <Grid container spacing={3}>

        <Grid xs={12} sm={4} md={4}>
          <AppWidgetSummary
            title="Session en cours"
            total={3.5}
            color="info"
            icon={<img alt="icon" src="/assets/icons/glass/ic_glass_users.png" />}
          />
        </Grid>

        <Grid xs={12} sm={4} md={4}>
          <AppWidgetSummary
            title="Pause du jour"
            total={0.5}
            color="error"
            icon={<img alt="icon" src="/assets/icons/glass/ic_glass_message.png" />}
          />
        </Grid>

        <Grid xs={12} sm={4} md={4}>
          <AppWidgetSummary
            title="Total du mois"
            total={77}
            color="success"
            icon={<img alt="icon" src="/assets/icons/glass/ic_glass_bag.png" />}
          />
        </Grid>

        <Grid xs={12} >
          <AppTasks
            title="Tasks"
            list={[
              { id: '1', name: 'Creation Logo ABM BLIPP' },
            ]}
          />
        </Grid>
      </Grid>
    </Container>
  );
}
