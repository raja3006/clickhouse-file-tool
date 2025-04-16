import React from 'react';
import {
  Box,
  Paper,
  Typography,
  Button,
} from '@mui/material';
import StorageIcon from '@mui/icons-material/Storage';
import InsertDriveFileIcon from '@mui/icons-material/InsertDriveFile';

interface Props {
  onSelect: (source: 'clickhouse' | 'file') => void;
}

const SourceSelector: React.FC<Props> = ({ onSelect }) => {
  return (
    <Paper elevation={3} sx={{ p: 3, maxWidth: 600, mx: 'auto' }}>
      <Typography variant="h6" gutterBottom align="center">
        Select Data Source
      </Typography>
      <Box sx={{ display: 'flex', gap: 3, mt: 2 }}>
        <Box sx={{ flex: 1 }}>
          <Button
            variant="outlined"
            size="large"
            startIcon={<StorageIcon />}
            onClick={() => onSelect('clickhouse')}
            sx={{ width: '100%', height: 100 }}
          >
            <Box sx={{ textAlign: 'center' }}>
              <Typography variant="subtitle1" gutterBottom>
                ClickHouse
              </Typography>
              <Typography variant="body2" color="text.secondary">
                Connect to a ClickHouse database
              </Typography>
            </Box>
          </Button>
        </Box>
        <Box sx={{ flex: 1 }}>
          <Button
            variant="outlined"
            size="large"
            startIcon={<InsertDriveFileIcon />}
            onClick={() => onSelect('file')}
            sx={{ width: '100%', height: 100 }}
          >
            <Box sx={{ textAlign: 'center' }}>
              <Typography variant="subtitle1" gutterBottom>
                Flat File
              </Typography>
              <Typography variant="body2" color="text.secondary">
                Select a local flat file
              </Typography>
            </Box>
          </Button>
        </Box>
      </Box>
    </Paper>
  );
};

export default SourceSelector; 