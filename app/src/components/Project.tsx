import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

import { request, loadProjectMetadata } from '../apiService';

import { PencilIcon } from '@heroicons/react/24/outline';
import { TableCellsIcon } from '@heroicons/react/24/outline';
import { Card } from "./ui/Card";
import { Alert } from "./ui/Alert";
import { Robot } from "./ui/Robot";
import { Link } from "./ui/Link";
import { CreateVisualization } from "./CreateVisualization";
import { VisualizationList } from "./VisualizationList";
import { FieldsCode } from "./FieldsCode";
import { ForkModal } from "./ForkModal";

import {
  useParams
} from "react-router-dom";

export const Project = () => {
  let { id } = useParams();
  const navigate = useNavigate();

  const [projectData, setProjectData] = useState<string | null>(null);
  const [projectMetadata, setProjectMetadata] = useState<any>(null);
  const [visualizationIds, setVisualizationIds] = useState<string[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(true);

  const fetchMetadata = async () => {
    const metadata = await loadProjectMetadata(id);
    document.title = 'Vizzy - ' + metadata.title;
    setProjectMetadata(metadata);
  }

  const fetchData = async () => {
    const dataResponse = await request('GET', `/projects/${id}/data`);
    if (dataResponse.ok) {
      const data = await dataResponse.text();
      setProjectData(data);
    } else {
      throw new Error("Error fetching data: " + dataResponse.status + " " + dataResponse.statusText);
    }
  }

  const fetchVisualizations = async () => {
    const response = await request('GET', `/projects/${id}/visualizations`);
    if (response.ok) {
      const resp = await response.json();
      setVisualizationIds(resp.ids || []);
    } else {
      throw new Error("Error fetching visualization IDs");
    }
  }

  const fetchAll = async () => {
    setIsLoading(true);
    await Promise.all([fetchMetadata(), fetchData(), fetchVisualizations()]);
    setIsLoading(false);
  }

  useEffect(() => {
    fetchAll();
  }, []);

  if (isLoading) {
    return <Alert level="info" robot="glasses" spinner={true} message="Let's take a look at the data and see what's there..." />
  }
  return (
  <div>
    <ForkModal projectId={id || ''} />
    <div className="text-left">
      <div className="flex grow justify-between">
        <h1 className="text-4xl mb-6">{projectMetadata?.title}</h1>
        <div className="text-right min-w-32">
          <a className="btn btn-outline btn-info mr-4" href={`/api/projects/${id}/data`} title="View data">
            <TableCellsIcon className="h-4 w-4" />
          </a>
          <button className="btn btn-outline btn-secondary" onClick={() => navigate("./edit")} title="Edit details">
            <PencilIcon className="h-4 w-4" />
          </button>
        </div>
      </div>
      <p className="mb-6 text-xl">{projectMetadata?.description}</p>
      <p className="mb-6 text-m">{projectMetadata?.data_format}</p>
      <p className="mb-6 text-m">
        <Robot emotion="thinking" className="mr-2 -mt-2" size={8} />
        <span className="italic">
          <span>Something off? </span>
          <Link text="Edit this information" onClick={() => navigate("./edit")} />
          <span> to help me generate more accurate visualizations.</span>
        </span>
      </p>
    </div>
    { projectMetadata &&
      <div className="text-center">
        <VisualizationList projectId={id || ''} visualizationIds={visualizationIds} data={projectData} onChange={fetchVisualizations} />
        { projectMetadata && <FieldsCode projectId={id || ''} data={projectData || ''} /> }
        <Card title="Create a new visualization">
          <CreateVisualization projectId={id || ''} projectMetadata={projectMetadata} onCreate={fetchVisualizations} />
        </Card>
      </div>
    }
  </div>);
};


