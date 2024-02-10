import React, { useState, useEffect } from 'react';
import { request, loadFieldsCode } from '../apiService';

import { FieldDetails } from './FieldDetails';
import { SandboxIframe } from './ui/SandboxIframe';
import { Card } from './ui/Card';
import { Alert } from './ui/Alert';
import { Switch } from './ui/Switch';
import { CodeEditor } from './ui/CodeEditor';

type FieldsCodeProps = {
  projectId: string;
  data: string;
};

export const FieldsCode: React.FC<FieldsCodeProps> = ({ projectId, data }) => {
  const [showCode, setShowCode] = useState<boolean>(false);
  const [fieldsCode, setFieldsCode] = useState<any>(null);
  const [fieldsMetadata, setFieldsMetadata] = useState(null);
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(true);

  useEffect(() => {
    const fetchCode = async () => {
      const code = await loadFieldsCode(projectId);
      setFieldsCode(code);
      setIsLoading(false);
    }
    fetchCode();
  }, [projectId]);

  const resetFieldsCode = async () => {
    setError(null);
    setIsLoading(true);
    setFieldsMetadata(null);
    const prevCode = fieldsCode;
    setFieldsCode(null);
    try {
      const code = await loadFieldsCode(projectId, true);
      setFieldsCode(code);
    } catch (e) {
      setFieldsCode(prevCode);
    }
    setIsLoading(false);
  }

  const updateFieldsMetadata = async (metadata) => {
    await request('POST', `/projects/${projectId}/fields-metadata`, metadata, false);
  }

  const onMessage = (msg) => {
    if (msg.error) {
      setError(msg.error);
    } else {
      setFieldsMetadata(msg.metadata);
      updateFieldsMetadata(msg.metadata);
    }
  };

  const formatValue = (value) => {
    if (typeof value === 'number') {
      return parseFloat(value.toFixed(2));
    }
    if (Array.isArray(value)) {
      return value.map(formatValue).join(', ');
    }
    return JSON.stringify(value);
  }

  const learningAlert = (
    <Alert level="info" robot="glasses" spinner={true} message="Hang on while I learn a bit more about the data..." />
  );

  if (isLoading) {
    return learningAlert;
  }

  if (!error && !fieldsMetadata) {
    return (
      <>
        {learningAlert}
        <SandboxIframe id={projectId} code={fieldsCode} data={data} onMessage={onMessage} />
      </>
    );
  }

  return (
    <Card center={false}>
      <div className="flex justify-between items-center w-full">
        <Switch checked={showCode} onChange={() => setShowCode(!showCode)} label="Show code" />
        <button className="btn btn-sm btn-secondary btn-outline w-64 mt-4" onClick={() => resetFieldsCode()}>Something off? Try again</button>
      </div>
      { error && (
        <>
          <Alert level="error" robot="oops">
              <p>
                <span>Oops. Looks like I goofed up. The code threw this error:</span>
                <br />
                <code>{error}</code>
              </p>
          </Alert>
          <button className="btn btn-sm btn-secondary btn-outline my-4" onClick={() => resetFieldsCode()}>Try again</button>
        </>
      )}
      { showCode && (
        <CodeEditor code={fieldsCode} />
      )}
      { !error && !showCode && (
        <>
          <FieldDetails fieldsMetadata={fieldsMetadata} projectId={projectId} />
        </>
      )}
    </Card>
  );
};

export default FieldsCode;

