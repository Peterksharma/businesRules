'use client';

import React, { useState } from 'react';
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card';
import { Checkbox } from '@/components/ui/checkbox';
import { ChevronDown, ChevronRight } from 'lucide-react';

// Define the transformation rules structure
const transformationRules = {
  transformations: {
    label: 'Transformations',
    items: ['aggregation', 'grouping', 'summarize']
  },
  cleaning: {
    label: 'Cleaning',
    items: [
      'dataFormats',
      'dataMasking',
      'deduplication',
      'missingValues',
      'notNeededData',
      'parsing',
      'spelling'
    ]
  },
  enrichment: {
    label: 'Enrichment',
    items: ['merging', 'toolCalls']
  },
  standardization: {
    label: 'Standardization',
    items: ['normalizeText', 'unifyDataFormat', 'unifyUnits']
  }
};

const BusinessRules: React.FC = () => {
  // State for category expansion
  const [expandedCategories, setExpandedCategories] = useState<Record<string, boolean>>({});
  // State for selected rules
  const [selectedRules, setSelectedRules] = useState<Record<string, boolean>>({});

  // Toggle category expansion
  const toggleCategory = (category: string) => {
    setExpandedCategories(prev => ({
      ...prev,
      [category]: !prev[category]
    }));
  };

  // Toggle rule selection
  const toggleRule = (rule: string) => {
    setSelectedRules(prev => ({
      ...prev,
      [rule]: !prev[rule]
    }));
  };

  // Format rule name for display
  const formatRuleName = (name: string) => {
    return name
      .replace(/([A-Z])/g, ' $1') // Add space before capital letters
      .replace(/^./, str => str.toUpperCase()); // Capitalize first letter
  };

  return (
    <Card className="w-full h-full">
      <CardHeader>
        <CardTitle>Business Rules</CardTitle>
      </CardHeader>
      <CardContent>
        <div className="space-y-4">
          <div className="p-4 border rounded-lg">
            <h3 className="text-lg font-medium mb-4">Data Transformations</h3>
            
            {/* Transformation Rules Checklist */}
            <div className="space-y-4">
              {Object.entries(transformationRules).map(([category, { label, items }]) => (
                <div key={category} className="border rounded-lg p-2">
                  {/* Category Header */}
                  <button
                    onClick={() => toggleCategory(category)}
                    className="w-full flex items-center justify-between p-2 hover:bg-gray-50 rounded"
                  >
                    <span className="font-medium">{label}</span>
                    {expandedCategories[category] ? (
                      <ChevronDown className="h-5 w-5" />
                    ) : (
                      <ChevronRight className="h-5 w-5" />
                    )}
                  </button>

                  {/* Category Items */}
                  {expandedCategories[category] && (
                    <div className="ml-4 mt-2 space-y-2">
                      {items.map(rule => (
                        <div key={rule} className="flex items-center space-x-2">
                          <Checkbox
                            id={rule}
                            checked={selectedRules[rule] || false}
                            onCheckedChange={() => toggleRule(rule)}
                          />
                          <label
                            htmlFor={rule}
                            className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                          >
                            {formatRuleName(rule)}
                          </label>
                        </div>
                      ))}
                    </div>
                  )}
                </div>
              ))}
            </div>
          </div>

          {/* Selected Rules Summary */}
          <div className="p-4 border rounded-lg">
            <h3 className="text-lg font-medium mb-2">Selected Rules</h3>
            <div className="space-y-2">
              {Object.entries(selectedRules)
                .filter(([_, isSelected]) => isSelected)
                .map(([rule]) => (
                  <div key={rule} className="text-sm text-gray-600 flex items-center space-x-2">
                    <span>•</span>
                    <span>{formatRuleName(rule)}</span>
                  </div>
                ))}
              {Object.values(selectedRules).filter(Boolean).length === 0 && (
                <div className="text-sm text-gray-500">
                  No rules selected. Choose from the categories above.
                </div>
              )}
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  );
};

export default BusinessRules; 