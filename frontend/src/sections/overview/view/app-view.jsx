import { useState, useEffect } from 'react';

import Container from '@mui/material/Container';
import Grid from '@mui/material/Unstable_Grid2';
import Typography from '@mui/material/Typography';

import { accountDefault } from 'src/_mock/account';

import AppTasks from '../app-tasks';
import AppWidgetSummary from '../app-widget-summary';
import AppConversionRates from '../app-conversion-rates';


// ----------------------------------------------------------------------

export default function AppView() {
  const [account, setAccount] = useState(accountDefault);

  useEffect(() => {

    const getAccount = async () => {
      try {
        const response = await window.go.main.App.UserGetAccount();
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

  return (
    <Container maxWidth="xl">
      <Typography variant="h4" sx={{ mb: 5 }}>
        Akao i {account.login} iny ô ? 👋
      </Typography>

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


        <Grid xs={12} md={6} lg={8}>
          <AppConversionRates
            title="Heures consomées"
            subheader="5.2H en moyenne"
            chart={{
              series: [
                { label: 'Lundi', value: 5 },
                { label: 'Mardi', value: 6 },
                { label: 'Mercredi', value: 3 },
                { label: 'Jeudi', value: 1 },
                { label: 'Vendredi', value: 7 },
              ],
            }}
          />
        </Grid>

        <Grid xs={12} md={6} lg={8}>
          <AppTasks
            title="Tasks"
            list={[
              { id: '1', name: 'Creation Logo ABM BLIPP' },
              { id: '2', name: 'Sprint Planning' },
              { id: '3', name: 'Jira Ticket 2' },
              { id: '4', name: 'Configuration Bitbucket' },
              { id: '5', name: 'Daily' },
            ]}
          />
        </Grid>
      </Grid>
    </Container>
  );
}
