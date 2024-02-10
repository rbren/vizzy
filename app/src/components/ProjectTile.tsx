import React, { useState, useEffect } from 'react';
import { Link } from "react-router-dom";

import { request } from '../apiService';

import { Card } from './ui/Card.tsx'

type DataDescription = {
  type: string;
  title: string;
  description: string;
  dataFormat: string;
  fields: string[];
  suggestedVisualizations: string[];
};

type ProjectTileProps = {
  projectId: string;
};

export const ProjectTile: React.FC<ProjectTileProps> = ({ projectId }) => {
  const [projectData, setProjectData] = useState<DataDescription | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(true);

  useEffect(() => {
    const fetchData = async () => {
      const response = await request('GET', `/projects/${projectId}/metadata`);
      const metadata = await response.json();
      setProjectData(metadata);
      setIsLoading(false);
    }
    fetchData();
  }, [projectId]);

  return (
    <Card title={projectData?.title || 'Unknown Project'} isLoading={isLoading}>
      <p>{projectData?.description}</p>
      <div className="card-actions justify-center">
        <Link to={`/projects/${projectId}`} className="btn btn-primary">View Project</Link>
      </div>
    </Card>
  );
};

export default ProjectTile;

