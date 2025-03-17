'use client';

import React from 'react';
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card';

const DataSource: React.FC = () => {
  return (
    <Card className="w-full h-full">
      <CardHeader>
        <CardTitle>Data Source</CardTitle>
      </CardHeader>
      <CardContent>
        <div className="space-y-4">
          <div className="p-4 border rounded-lg">
            <h3 className="text-lg font-medium mb-2">Connected Sources</h3>
            {/* Add your data source list or connection form here */}
            <div className="text-sm text-gray-500">
              No data sources connected. Click to add a new source.
            </div>
          </div>
          
          <div className="p-4 border rounded-lg">
            <h3 className="text-lg font-medium mb-2">Source Settings</h3>
            {/* Add settings form or configuration options here */}
            <div className="text-sm text-gray-500">
              Configure your data source settings and parameters.
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  );
};

export default DataSource; 