'use client';

import React from 'react';
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card';

const Reports: React.FC = () => {
  return (
    <Card className="w-full h-full">
      <CardHeader>
        <CardTitle>Reports</CardTitle>
      </CardHeader>
      <CardContent>
        <div className="space-y-4">
          <div className="p-4 border rounded-lg">
            <h3 className="text-lg font-medium mb-2">Recent Reports</h3>
            {/* Add your reports list or report generator here */}
            <div className="text-sm text-gray-500">
              No reports generated. Click to create a new report.
            </div>
          </div>
          
          <div className="p-4 border rounded-lg">
            <h3 className="text-lg font-medium mb-2">Report Templates</h3>
            {/* Add report templates or types here */}
            <div className="grid grid-cols-2 gap-2">
              <div className="p-2 border rounded">
                <span className="text-sm font-medium">Data Quality</span>
              </div>
              <div className="p-2 border rounded">
                <span className="text-sm font-medium">Rule Execution</span>
              </div>
              <div className="p-2 border rounded">
                <span className="text-sm font-medium">Performance</span>
              </div>
              <div className="p-2 border rounded">
                <span className="text-sm font-medium">Custom Reports</span>
              </div>
            </div>
          </div>

          <div className="p-4 border rounded-lg">
            <h3 className="text-lg font-medium mb-2">Analytics Overview</h3>
            {/* Add analytics or metrics here */}
            <div className="grid grid-cols-3 gap-4">
              <div className="text-center">
                <div className="text-2xl font-bold">0</div>
                <div className="text-sm text-gray-500">Rules Executed</div>
              </div>
              <div className="text-center">
                <div className="text-2xl font-bold">0</div>
                <div className="text-sm text-gray-500">Data Sources</div>
              </div>
              <div className="text-center">
                <div className="text-2xl font-bold">0</div>
                <div className="text-sm text-gray-500">Reports Generated</div>
              </div>
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  );
};

export default Reports; 