import React from 'react';
import { CssBaseline, ThemeProvider, createTheme } from '@mui/material';
import IngestionPage from './pages/IngestionPage';

const theme = createTheme({
  palette: {
    mode: 'light',
    primary: {
      main: '#1976d2',
    },
    secondary: {
      main: '#dc004e',
    },
  },
});

function App() {
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <IngestionPage />
    </ThemeProvider>
  );
}

export default App;
