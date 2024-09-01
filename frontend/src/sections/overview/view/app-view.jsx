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
  const [totalCurrentTime, setTotalCurrentTime] = useState(0.0);


  const updateSessionStatus = async (status) => {
    if (status === "offline") {
      const response = await window.go.app.App.UserStartSession();
      if (response === true) {
        setSessionStatus("working");
      }
    }
    else {
      const response = await window.go.app.App.UserStopSession();
      if (response === true) {
        setSessionStatus("offline");
      }
    }
  }

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


  // total time of the session
  useEffect(() => {
    const getTotalTime = async () => {
      try {
        const response = await window.go.app.App.UserCurrentTotalTime();
        if (response == null) {
          return;
        }
        setTotalCurrentTime(response);
      } catch (error) {
        console.error('Error getting total time:', error);
      }
    };

    getTotalTime();
  }, []);

  return (
    <Container maxWidth="xl">

      <Grid container spacing={0} sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>

        {sessionStatus === "offline" ?
          <Typography variant="h6" >
            Akao i {account.login.charAt(0).toUpperCase() + account.login.slice(1)} iny ô ? 👋
          </Typography>
          :
          <Button variant="contained" color="secondary" >
            Mettre en Pause
          </Button>
        }

        <Typography variant="caption" >
          {sessionStatus} 🚀
        </Typography>

        <Button onClick={() => updateSessionStatus(sessionStatus)} variant="outlined"  >
          {sessionStatus === "offline" ? "Demarrer la session" : "Arrêter la session"}
        </Button>
      </Grid>

      <Grid container spacing={3}>

        <Grid xs={12} sm={4} md={4}>
          <AppWidgetSummary
            title="Session en cours"
            total={totalCurrentTime}
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
