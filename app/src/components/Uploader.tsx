import React, { useState } from 'react';

import { request } from '../apiService.tsx';

import { SubmitButton } from './ui/SubmitButton.tsx'
import { Alert } from './ui/Alert.tsx'
import { Link } from './ui/Link.tsx'
import { getItem, setItem } from '../storage.tsx';
import { useNavigate } from 'react-router-dom';

function convertToRawGithubLink(githubLink) {
    return githubLink.replace("https://github.com", "https://raw.githubusercontent.com")
                     .replace("/blob", "");
}

export const Uploader: React.FC = () => {
  const navigate = useNavigate();
  const [file, setFile] = useState<File | null>(null);
  const [url, setUrl] = useState<string>('');
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setFile(event.target.files ? event.target.files[0] : null);
  };

  const handleUrlChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    let url = event.target.value;
    if (url.includes("https://github.com") && url.includes("/blob")) {
      url = convertToRawGithubLink(url);
      console.log('set', url);
    }
    setUrl(url);
  };

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setIsLoading(true);

    try {
      let response;
      if (file) {
        response = await request('POST', '/projects', file);
      } else if (url) {
        response = await request('POST', `/projects?url=${encodeURIComponent(url)}`);
      }
      setIsLoading(false);
      if (response && response.ok) {
        const { uuid, key } = await response.json();
        const existingData = JSON.parse(getItem('projects') || '{}');
        existingData[uuid] = key;
        setItem('projects', JSON.stringify(existingData));
        navigate('/projects/' + uuid);
      } else {
        console.error('Error fetching data');
        throw new Error('Error fetching data');
      }
    } catch (e) {
      console.error(e);
      setIsLoading(false);
    }
  };

  return (
    <div className="container mx-auto p-4">
      <form onSubmit={handleSubmit} className="form-control mb-4">
        <div className="input-group">
          <input type="file" onChange={handleFileChange} className="file-input file-input-accent input input-sm input-bordered w-full" />
        </div>
        <div className="divider my-1">OR</div>
        <div className="input-group">
          <input type="text" placeholder="Enter URL" value={url} onChange={handleUrlChange} className="input input-sm input-bordered w-full" />
        </div>
        <div className="mt-4">
          <Alert level="warning" robot="monocle" size="sm">
            <span>
              <span>Watch out: all projects and data are public. </span>
              <Link url="/email"> Learn more about private projects</Link>.
            </span>
          </Alert>
        </div>
        <SubmitButton isLoading={isLoading} className="mt-4 btn-sm" text="Start visualizing" />
      </form>
      <p>
        <span>No data handy? Try </span>
        <Link url="https://github.com/vega/vega-datasets/tree/main/data" text="vega-datasets" />
        <span> or check out the public projects below</span>
      </p>
    </div>
  );
};

export default Uploader;
