import React, { useState } from 'react';
import {
  Box,
  TextField,
  Button,
  FormControlLabel,
  Switch,
  Typography,
  Paper,
} from '@mui/material';
import { ClickHouseConfig } from '../types';

interface Props {
  onConnect: (config: ClickHouseConfig) => Promise<void>;
}

const ClickHouseForm: React.FC<Props> = ({ onConnect }) => {
  const [config, setConfig] = useState<ClickHouseConfig>({
    host: 'localhost',
    port: 9000,
    database: 'default',
    username: 'default',
    password: '',
    jwtToken: '',
    secure: false,
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value, type, checked } = e.target;
    setConfig((prev) => ({
      ...prev,
      [name]: type === 'checkbox' ? checked : type === 'number' ? Number(value) : value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    await onConnect(config);
  };

  return (
    <Paper elevation={3} sx={{ p: 3, maxWidth: 500, mx: 'auto' }}>
      <Typography variant="h6" gutterBottom>
        ClickHouse Connection
      </Typography>
      <Box component="form" onSubmit={handleSubmit} sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
        <TextField
          label="Host"
          name="host"
          value={config.host}
          onChange={handleChange}
          required
        />
        <TextField
          label="Port"
          name="port"
          type="number"
          value={config.port}
          onChange={handleChange}
          required
        />
        <TextField
          label="Database"
          name="database"
          value={config.database}
          onChange={handleChange}
          required
        />
        <TextField
          label="Username"
          name="username"
          value={config.username}
          onChange={handleChange}
          required
        />
        <TextField
          label="Password"
          name="password"
          type="password"
          value={config.password}
          onChange={handleChange}
          required
        />
        <TextField
          label="JWT Token"
          name="jwtToken"
          value={config.jwtToken}
          onChange={handleChange}
        />
        <FormControlLabel
          control={
            <Switch
              name="secure"
              checked={config.secure}
              onChange={handleChange}
            />
          }
          label="Use HTTPS"
        />
        <Button type="submit" variant="contained" color="primary">
          Connect
        </Button>
      </Box>
    </Paper>
  );
};

export default ClickHouseForm; 