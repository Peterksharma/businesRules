'use client';

import React from 'react';
import DataSource from './components/DataSource';
import BusinessRules from './components/BusinessRules';
import Reports from './components/Reports';

export default function Home() {
  return (
    <main className="min-h-screen p-8 bg-gray-50">
      <div className="max-w-7xl mx-auto space-y-8">
        {/* Header */}
        <div className="mb-8">
          <h1 className="text-3xl font-bold text-gray-900">Business Rules Engine</h1>
          <p className="text-gray-600 mt-2">
            Manage your data sources, business rules, and generate reports
          </p>
        </div>

        {/* Main Grid Layout */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          {/* Data Source Section */}
          <div className="h-[600px]">
            <DataSource />
          </div>

          {/* Business Rules Section */}
          <div className="h-[600px]">
            <BusinessRules />
          </div>

          {/* Reports Section */}
          <div className="h-[600px]">
            <Reports />
          </div>
        </div>
      </div>
    </main>
  );
}
