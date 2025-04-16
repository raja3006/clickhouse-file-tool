export interface ClickHouseConfig {
  host: string;
  port: number;
  database: string;
  username: string;
  password: string;
  jwtToken: string;
  secure: boolean;
}

export interface FileConfig {
  filePath: string;
  delimiter: string;
}

export interface Column {
  name: string;
  selected: boolean;
}

export interface Table {
  name: string;
  columns: Column[];
}

export interface IngestionConfig {
  source: 'clickhouse' | 'file';
  target: 'clickhouse' | 'file';
  sourceConfig: ClickHouseConfig | FileConfig;
  targetConfig: ClickHouseConfig | FileConfig;
  selectedColumns: string[];
  tableName?: string;
}

export interface ApiResponse<T> {
  data?: T;
  error?: string;
  message?: string;
} 