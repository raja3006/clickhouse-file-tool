import React, { useState } from 'react';
import {
  Box,
  TextField,
  Button,
  Typography,
  Paper,
} from '@mui/material';
import { FileConfig } from '../types';

interface Props {
  onSubmit: (config: FileConfig) => Promise<void>;
}

const FileForm: React.FC<Props> = ({ onSubmit }) => {
  const [config, setConfig] = useState<FileConfig>({
    filePath: '',
    delimiter: ',',
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setConfig((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    await onSubmit(config);
  };

  return (
    <Paper elevation={3} sx={{ p: 3, maxWidth: 500, mx: 'auto' }}>
      <Typography variant="h6" gutterBottom>
        File Configuration
      </Typography>
      <Box component="form" onSubmit={handleSubmit} sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
        <TextField
          label="File Path"
          name="filePath"
          value={config.filePath}
          onChange={handleChange}
          required
          helperText="Enter the path to your flat file"
        />
        <TextField
          label="Delimiter"
          name="delimiter"
          value={config.delimiter}
          onChange={handleChange}
          required
          helperText="Enter the delimiter used in your file (e.g., ',' for CSV)"
        />
        <Button type="submit" variant="contained" color="primary">
          Load File
        </Button>
      </Box>
    </Paper>
  );
};

export default FileForm; 