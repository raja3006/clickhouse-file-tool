import React from 'react';
import {
  Box,
  Checkbox,
  FormControlLabel,
  Typography,
  Paper,
  Button,
} from '@mui/material';
import { Column } from '../types';

interface Props {
  columns: Column[];
  onColumnToggle: (columnName: string) => void;
  onSelectAll: () => void;
  onDeselectAll: () => void;
  onSubmit: () => void;
}

const ColumnSelector: React.FC<Props> = ({
  columns,
  onColumnToggle,
  onSelectAll,
  onDeselectAll,
  onSubmit,
}) => {
  return (
    <Paper elevation={3} sx={{ p: 3, maxWidth: 500, mx: 'auto' }}>
      <Typography variant="h6" gutterBottom>
        Select Columns
      </Typography>
      <Box sx={{ mb: 2 }}>
        <Button
          variant="outlined"
          size="small"
          onClick={onSelectAll}
          sx={{ mr: 1 }}
        >
          Select All
        </Button>
        <Button
          variant="outlined"
          size="small"
          onClick={onDeselectAll}
        >
          Deselect All
        </Button>
      </Box>
      <Box sx={{ maxHeight: 400, overflowY: 'auto', mb: 2 }}>
        {columns.map((column) => (
          <FormControlLabel
            key={column.name}
            control={
              <Checkbox
                checked={column.selected}
                onChange={() => onColumnToggle(column.name)}
              />
            }
            label={column.name}
          />
        ))}
      </Box>
      <Button
        variant="contained"
        color="primary"
        onClick={onSubmit}
        disabled={!columns.some((col) => col.selected)}
      >
        Continue
      </Button>
    </Paper>
  );
};

export default ColumnSelector; 