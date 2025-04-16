import React, { useState } from 'react';
import {
  Box,
  Button,
  Container,
  Typography,
  Stepper,
  Step,
  StepLabel,
  Alert,
  CircularProgress,
} from '@mui/material';
import ClickHouseForm from '../components/ClickHouseForm';
import FileForm from '../components/FileForm';
import ColumnSelector from '../components/ColumnSelector';
import SourceSelector from '../components/SourceSelector';
import { ClickHouseConfig, FileConfig, Column } from '../types';
import * as api from '../services/api';

const steps = [
  'Select Source',
  'Configure Source',
  'Select Columns',
  'Configure Target',
  'Start Ingestion',
];

const IngestionPage: React.FC = () => {
  const [activeStep, setActiveStep] = useState(0);
  const [source, setSource] = useState<'clickhouse' | 'file'>('clickhouse');
  const [target, setTarget] = useState<'clickhouse' | 'file'>('file');
  const [sourceConfig, setSourceConfig] = useState<ClickHouseConfig | FileConfig | null>(null);
  const [targetConfig, setTargetConfig] = useState<ClickHouseConfig | FileConfig | null>(null);
  const [columns, setColumns] = useState<Column[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [ingestionResult, setIngestionResult] = useState<{ count: number } | null>(null);

  const handleSourceSelect = (selectedSource: 'clickhouse' | 'file') => {
    setSource(selectedSource);
    setTarget(selectedSource === 'clickhouse' ? 'file' : 'clickhouse');
    setActiveStep(1);
  };

  const handleClickHouseConnect = async (config: ClickHouseConfig) => {
    setLoading(true);
    setError(null);
    try {
      await api.connectClickHouse(config);
      const tablesResponse = await api.getTables();
      if (tablesResponse.error) {
        throw new Error(tablesResponse.error);
      }
      // For now, we'll just use the first table
      if (tablesResponse.data && tablesResponse.data.length > 0) {
        const columnsResponse = await api.getColumns(tablesResponse.data[0]);
        if (columnsResponse.error) {
          throw new Error(columnsResponse.error);
        }
        if (columnsResponse.data) {
          setColumns(columnsResponse.data.map(name => ({ name, selected: false })));
        }
      }
      setSourceConfig(config);
      setActiveStep(2);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to connect to ClickHouse');
    } finally {
      setLoading(false);
    }
  };

  const handleFileConfig = async (config: FileConfig) => {
    setLoading(true);
    setError(null);
    try {
      const response = await api.getFileColumns(config);
      if (response.error) {
        throw new Error(response.error);
      }
      if (response.data) {
        setColumns(response.data.map(name => ({ name, selected: false })));
      }
      setSourceConfig(config);
      setActiveStep(2);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to read file');
    } finally {
      setLoading(false);
    }
  };

  const handleColumnToggle = (columnName: string) => {
    setColumns(prev =>
      prev.map(col =>
        col.name === columnName ? { ...col, selected: !col.selected } : col
      )
    );
  };

  const handleSelectAll = () => {
    setColumns(prev => prev.map(col => ({ ...col, selected: true })));
  };

  const handleDeselectAll = () => {
    setColumns(prev => prev.map(col => ({ ...col, selected: false })));
  };

  const handleStartIngestion = async () => {
    if (!sourceConfig || !targetConfig) return;

    setLoading(true);
    setError(null);
    try {
      const selectedColumns = columns.filter(col => col.selected).map(col => col.name);
      const response = await api.ingestData(source, target, {
        sourceConfig,
        targetConfig,
        columns: selectedColumns,
      });

      if (response.error) {
        throw new Error(response.error);
      }

      setIngestionResult(response.data || null);
      setActiveStep(5);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to ingest data');
    } finally {
      setLoading(false);
    }
  };

  const renderStep = () => {
    switch (activeStep) {
      case 0:
        return <SourceSelector onSelect={handleSourceSelect} />;
      case 1:
        return source === 'clickhouse' ? (
          <ClickHouseForm onConnect={handleClickHouseConnect} />
        ) : (
          <FileForm onSubmit={handleFileConfig} />
        );
      case 2:
        return (
          <ColumnSelector
            columns={columns}
            onColumnToggle={handleColumnToggle}
            onSelectAll={handleSelectAll}
            onDeselectAll={handleDeselectAll}
            onSubmit={() => setActiveStep(3)}
          />
        );
      case 3:
        return target === 'clickhouse' ? (
          <ClickHouseForm onConnect={async (config) => {
            setTargetConfig(config);
            setActiveStep(4);
            return Promise.resolve();
          }} />
        ) : (
          <FileForm onSubmit={async (config) => {
            setTargetConfig(config);
            setActiveStep(4);
            return Promise.resolve();
          }} />
        );
      case 4:
        return (
          <Box sx={{ textAlign: 'center' }}>
            <Typography variant="h6" gutterBottom>
              Ready to Start Ingestion
            </Typography>
            <Button
              variant="contained"
              color="primary"
              onClick={handleStartIngestion}
              disabled={loading}
            >
              Start Ingestion
            </Button>
          </Box>
        );
      case 5:
        return (
          <Box sx={{ textAlign: 'center' }}>
            <Typography variant="h6" gutterBottom>
              Ingestion Complete
            </Typography>
            {ingestionResult && (
              <Typography>
                Successfully ingested {ingestionResult.count} records
              </Typography>
            )}
          </Box>
        );
      default:
        return null;
    }
  };

  return (
    <Container maxWidth="md" sx={{ py: 4 }}>
      <Typography variant="h4" gutterBottom align="center">
        Data Ingestion Tool
      </Typography>
      <Stepper activeStep={activeStep} sx={{ mb: 4 }}>
        {steps.map((label) => (
          <Step key={label}>
            <StepLabel>{label}</StepLabel>
          </Step>
        ))}
      </Stepper>
      {error && (
        <Alert severity="error" sx={{ mb: 2 }}>
          {error}
        </Alert>
      )}
      {loading ? (
        <Box sx={{ display: 'flex', justifyContent: 'center', my: 4 }}>
          <CircularProgress />
        </Box>
      ) : (
        renderStep()
      )}
    </Container>
  );
};

export default IngestionPage; 