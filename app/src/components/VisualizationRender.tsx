import { useState, useEffect } from 'react';

import { Alert } from './ui/Alert';
import { SandboxIframe } from './ui/SandboxIframe';

export const VisualizationRender = ({ code, data, id, onRender }) => {
  onRender = onRender || (() => {});
  const [error, setError] = useState<string>('');

  const onMessage = (data) => {
    if (data.error) setError(data.error);
    onRender(data);
  }

  useEffect(() => {
    setError('');
  }, [code]);

  return (
    <div className="w-full">
      {error ? (
        <Alert level="error" robot="oops">
          <p>
            <span>Oops. Looks like I goofed up. The code threw this error:</span>
            <br />
            <code>{error}</code>
          </p>
        </Alert>
      ) : (
        <SandboxIframe code={code} data={data} id={id} onMessage={onMessage} height="500px" width="100%"/>
      )}
    </div>
  );
};

export default VisualizationRender;
