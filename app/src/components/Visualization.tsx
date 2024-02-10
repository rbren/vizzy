import React, { useState, useEffect } from 'react';
import { TrashIcon } from '@heroicons/react/24/outline';

import { request } from '../apiService';
import { Link } from './ui/Link';
import { Card } from './ui/Card';
import { CodeEditor } from './ui/CodeEditor';
import { SubmitButton } from './ui/SubmitButton';
import { Alert } from './ui/Alert';
import { Switch } from './ui/Switch';
import { Collapse } from './ui/Collapse';
import { InputWithSubmit } from './ui/InputWithSubmit';

import { VisualizationRender } from './VisualizationRender';
import { WorkflowHistory } from './WorkflowHistory';

type VisualizationProps = {
  projectId: string;
  visualizationId: string;
  data: string;
  onChange: () => void;
};

type VisualizationDetails = {
  title: string;
  prompt: string;
  visualization: string;
  version: number;
};

export const Visualization: React.FC<VisualizationProps> = ({ projectId, visualizationId, data, onChange }) => {
  const [visualizationDetails, setVisualizationDetails] = useState<VisualizationDetails | null>(null);
  const [showCode, setShowCode] = useState(false);
  const [visualizationQuery, setVisualizationQuery] = useState<string>('');
  const [code, setCode] = useState<string>('');
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [isModifying, setIsModifying] = useState<boolean>(false);

  const handleDelete = async () => {
    if (window.confirm("Are you sure you want to delete this visualization?")) {
      try {
        await request("DELETE", `/projects/${projectId}/visualizations/${visualizationId}`);
      } catch (error) {
        console.error('Error deleting visualization:', error);
        return;
      }
      onChange();
    }
  };

  const handleRender = (result) => {
    if (!result.success) {
      setVisualizationQuery("Please fix this error: `" + result.error + "`");
    }
  }

  const handleNewVersionSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    return generateNewVersion(visualizationDetails?.visualization !== code);
  }

  const generateNewVersionFromEdits = async () => {
    return generateNewVersion(true);
  }

  const revertCodeChanges = () => {
    setCode(visualizationDetails?.visualization || '');
    setShowCode(false);
  }

  const generateNewVersion = async (useEdits) => {
    setIsModifying(true);
    try {
      let url = `/projects/${projectId}/visualizations/${visualizationId}/versions/${visualizationDetails?.version}`;
      let body:any = null;
      if (useEdits) {
        body = {code};
      } else {
        url += '?prompt=' + visualizationQuery;
      }

      const response = await request('PATCH', url, body);
      if (response.ok) {
        setIsModifying(false);
        setVisualizationQuery('');
        await fetchVisualizationDetails();
      } else {
        console.error('Error generating new version:' + response.status);
        throw new Error('Error generating new version');
      }
    } catch (error) {
      setIsModifying(false);
      console.error('Error fetching visualization data:', error);
    }
  }

  const fetchVisualizationDetails = async (version:any=null) => {
    setIsLoading(true);
    let url = `/projects/${projectId}/visualizations/${visualizationId}`;
    if (version) {
      url += `/versions/${version}`;
    } else {
      url += '/latest';
    }
    try {
      const response = await request('GET', url);
      if (response.ok) {
        const metadata = await response.json();
        setVisualizationDetails(metadata);
        setCode(metadata.visualization);
      }
    } catch (error) {
      console.error('Error fetching visualization data:', error);
    } finally {
      setIsLoading(false);
    }
  };

  const undo = async () => {
    if (visualizationDetails?.prompt !== '[manual edit]') {
      setVisualizationQuery(visualizationDetails?.prompt || '');
    }
    await request('DELETE', `/projects/${projectId}/visualizations/${visualizationId}/versions/${visualizationDetails?.version}`);
    return fetchVisualizationDetails();
  }

  const setHistoricalVersion = async (version) => {
    console.log('set hist', version);
    setCode(version.visualization);
  }

  useEffect(() => {
    fetchVisualizationDetails();
  }, [projectId, visualizationId]);

  return (
    <Card
      title={visualizationDetails?.title || visualizationDetails?.prompt}
      className="min-w-200"
      isLoading={isLoading}>
        <div className="flex justify-between items-center w-full">
          <div className="flex items-center gap-2">
            <Switch label="Edit code" checked={showCode} onChange={() => setShowCode(!showCode)} />
          </div>
          <div className="flex items-center gap-2">
            <button
              onClick={handleDelete}
              className="btn btn-square btn-sm btn-outline btn-error"
            >
              <TrashIcon className="h-4 w-4" />
            </button>

          </div>
        </div>

        <Collapse label='Workflow History'>
          <WorkflowHistory onSelect={setHistoricalVersion} projectId={projectId} visualizationId={visualizationId} />
        </Collapse>
        {!showCode ? (
          <>
            <VisualizationRender data={data} code={code} id={visualizationId} onRender={handleRender}/>
            <div className="w-full mt-4">
              <form onSubmit={handleNewVersionSubmit} className="form-control">
                { !isModifying
                  ? (
                    <div className="input-group w-full mt-4 flex items-center justify-center">
                      { code === visualizationDetails?.visualization ? (
                        <InputWithSubmit
                          placeholder="How would you like to modify this visualization?"
                          submitText="Make it so!"
                          value={visualizationQuery}
                          onChange={(e) => setVisualizationQuery(e.target.value)} />
                      ) : (
                        <div className="mt-4">
                          <SubmitButton className="mr-4" text="Save code" />
                          <button className="btn btn-outline" onClick={revertCodeChanges}>Revert changes</button>
                        </div>
                      )}
                    </div>
                  ) : (
                    <Alert level="info" robot="glasses" spinner={true} className="mt-4" message="On it! I'm making some changes to the code. This will probably take a minute..." />
                  )
                }
              </form>
              <p className="help-text mt-4">
                Totally broken?
                <span> </span>
                { (visualizationDetails?.version || 0) > 1 &&
                  <>
                    <Link onClick={undo} >Undo your last change</Link>
                    <span> or </span>
                  </>}
                <Link onClick={handleDelete}>
                  { (visualizationDetails?.version || 0) > 1 ? 'd' : 'D'}elete
                </Link>
                <span> and start over. </span>
                <span>And help us improve by </span>
                <Link url="https://github.com/rbren/vizzy/edit/main/FailurePatterns.md">submitting a report</Link>
                <span>!</span>
              </p>
            </div>
          </>
        ) : (
          <>
            <CodeEditor code={code} onChange={setCode} />
            <div className="mt-4">
              <SubmitButton isLoading={isModifying} onClick={generateNewVersionFromEdits} text="Save code" className="mr-4" />
              <button className="btn btn-outline" onClick={revertCodeChanges}>Revert changes</button>
            </div>
          </>
        )}
    </Card>
  );
};

export default Visualization;
