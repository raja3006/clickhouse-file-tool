import axios from 'axios';
import { ClickHouseConfig, FileConfig, ApiResponse } from '../types';

const API_BASE_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080/api';

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

export const connectClickHouse = async (config: ClickHouseConfig): Promise<ApiResponse<void>> => {
  try {
    const response = await api.post('/clickhouse/connect', config);
    return response.data;
  } catch (error) {
    return { error: 'Failed to connect to ClickHouse' };
  }
};

export const getTables = async (): Promise<ApiResponse<string[]>> => {
  try {
    const response = await api.get('/clickhouse/tables');
    return response.data;
  } catch (error) {
    return { error: 'Failed to fetch tables' };
  }
};

export const getColumns = async (table: string): Promise<ApiResponse<string[]>> => {
  try {
    const response = await api.get(`/clickhouse/columns/${table}`);
    return response.data;
  } catch (error) {
    return { error: 'Failed to fetch columns' };
  }
};

export const getFileColumns = async (config: FileConfig): Promise<ApiResponse<string[]>> => {
  try {
    const response = await api.post('/file/columns', config);
    return response.data;
  } catch (error) {
    return { error: 'Failed to fetch file columns' };
  }
};

export const ingestData = async (
  source: 'clickhouse' | 'file',
  target: 'clickhouse' | 'file',
  config: any
): Promise<ApiResponse<{ count: number }>> => {
  try {
    const response = await api.post(`/ingest/${source}/${target}`, config);
    return response.data;
  } catch (error) {
    return { error: 'Failed to ingest data' };
  }
}; 