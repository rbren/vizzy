import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

import { request, loadProjectMetadata } from '../apiService';
import { removeProject } from '../storage';

import { SubmitButton } from './ui/SubmitButton';
import { Alert } from './ui/Alert';

import { ForkModal } from "./ForkModal";

import {
  useParams
} from "react-router-dom";

export const EditProject = () => {
  let { id } = useParams();
  const navigate = useNavigate();

  const [projectMetadata, setProjectMetadata] = useState<any>(null);
  const [isLoading, setIsLoading] = useState<boolean>(true);

  const fetchMetadata = async () => {
    const metadata = await loadProjectMetadata(id);
    setProjectMetadata(metadata);
  }

  const deleteProject = async () => {
    if (window.confirm("Are you sure you want to delete this project?")) {
      const response = await request('DELETE', `/projects/${id}`);
      if (response.ok) {
        removeProject(id || '');
        navigate('/');
      } else {
        throw new Error("Error deleting project");
      }
    }
  }

  const fetchAll = async () => {
    setIsLoading(true);
    await Promise.all([fetchMetadata()]);
    setIsLoading(false);
  }

  useEffect(() => {
    fetchAll();
  }, []);

  const saveMetadata = async (e: any) => {
    e.preventDefault();
    await request("POST", `/projects/${id}/metadata`, projectMetadata)
    navigate(`/projects/${id}`);
  }

  if (isLoading) {
    return <Alert level="info" robot="glasses" spinner={true} message="Fetching your current settings..." />
  }
  return (
  <div>
    <ForkModal projectId={id || ''} />
    <form onSubmit={saveMetadata} className="text-left form-control">
      <label>Title</label>
      <input type="text"
          className="input input-bordered mb-4"
          value={projectMetadata?.title}
          onChange={e => setProjectMetadata({...projectMetadata, title: e.target.value})} />
      <label>Description</label>
      <textarea
          className="textarea textarea-bordered mb-4"
          value={projectMetadata?.description}
          onChange={e => setProjectMetadata({...projectMetadata, description: e.target.value})} />
      <label>Data Format</label>
      <textarea
          className="textarea textarea-bordered mb-4"
          value={projectMetadata?.data_format}
          onChange={e => setProjectMetadata({...projectMetadata, data_format: e.target.value})} />
      <div className="items-left">
        <SubmitButton className="mr-4" text="Save"/>
        <button className="btn btn-error btn-outline" onClick={deleteProject}>Delete Project</button>
      </div>
    </form>
  </div>);
};


