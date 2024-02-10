import React, { useState } from 'react';
import { LightBulbIcon } from '@heroicons/react/24/outline'
import { request } from '../apiService';

import { Alert } from './ui/Alert';
import { InputWithSubmit } from './ui/InputWithSubmit';

type DataDescription = {
  type: string;
  title: string;
  description: string;
  data_format: string;
  fields: string[];
  suggested_visualizations: string[];
};

type CreateVisualizationProps = {
  projectId: string;
  projectMetadata: DataDescription;
  onCreate: (viz: any) => void;
};

export const CreateVisualization: React.FC<CreateVisualizationProps> = ({ projectId, projectMetadata, onCreate }) => {
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [visualizationQuery, setVisualizationQuery] = useState<string>('');

  const handleSuggestionClick = (suggestion: string) => {
    setVisualizationQuery(suggestion);
  };

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setIsLoading(true);
    try {
      const response = await request('POST', `/projects/${projectId}/visualizations?prompt=${visualizationQuery}`);
      const viz = await response.json();
      onCreate(viz);
      setVisualizationQuery('');
      setIsLoading(false);
    } catch (error) {
      console.error(error);
      setIsLoading(false);
      return;
    }
  };

  return (
    <div className="w-full">
      <ul className="text-left mb-4">
        {projectMetadata?.suggested_visualizations.map((suggestion, index) => (
          <li key={index} className="font-medium">
            <LightBulbIcon className="inline-block w-4 h-4 mr-1" />
            <a className="text-blue-600 dark:text-blue-500 hover:underline cursor-pointer"
               onClick={() => handleSuggestionClick(suggestion)}>
              {suggestion}
            </a>
          </li>
        ))}
      </ul>

      <form onSubmit={handleSubmit} className="form-control">
        { !isLoading
          ? (<InputWithSubmit
              placeholder="How would you like to visualize the data?"
              submitText="Make it so!"
              value={visualizationQuery}
              onChange={(e) => setVisualizationQuery(e.target.value)} />)
          : (<Alert level="info" robot="glasses" spinner={true} className="mt-4" message="Got it! I'll code this up right away. It'll probably take a minute or two..." />)
        }
      </form>
    </div>
  );
};

export default CreateVisualization;
