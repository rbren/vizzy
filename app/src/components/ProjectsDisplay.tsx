import React, { useState, useEffect } from 'react';
import { getItem } from '../storage';

import ProjectTile from './ProjectTile';
import Uploader from './Uploader';
import { Card } from './ui/Card.tsx'

export const ProjectsDisplay: React.FC = () => {
  const [publicProjectIds, setPublicProjectIds] = useState<string[]>([]);
  const [projectIds, setProjectIds] = useState<string[]>([]);

  const fetchPublicProjects = async () => {
    const response = await fetch('/saved-projects.txt');
    if (!response.ok || response.status !== 200) {
      console.error('Failed to fetch public projects');
      return;
    }
    const publicProjects = await response.text();
    setPublicProjectIds(publicProjects.split('\n').filter(p => p));
  }

  useEffect(() => {
    const storedProjects = JSON.parse(getItem('projects') || '{}');
    setProjectIds(Object.keys(storedProjects));
    fetchPublicProjects();
  }, []);

  return (
    <>
      <div className="py-4">
        <Card>
          <Uploader />
        </Card>
      </div>
      { !!projectIds?.length && (
        <div className="mt-12">
          <h2 className="text-2xl font-bold">My Projects</h2>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 py-4">
            {projectIds.map(projectId => (
              <ProjectTile key={projectId} projectId={projectId} />
            ))}
          </div>
        </div>
      )}
      { !!publicProjectIds?.length && (
        <div className="mt-12">
          <h2 className="text-2xl font-bold">Public Project Gallery</h2>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 py-4">
            {publicProjectIds.map(projectId => (
              <ProjectTile key={projectId} projectId={projectId} />
            ))}
          </div>
        </div>
      )}
    </>
  );
};

export default ProjectsDisplay;
