import React from 'react';
import { useRef, useEffect } from 'react';

type SandboxIframeProps = {
  id: string;
  code: string;
  data?: any;
  width?: string;
  height?: string;
  onMessage?: (data: any) => void;
}

export const SandboxIframe : React.FC<SandboxIframeProps> = ({ code, data, id, width, height, onMessage }) => {
  width = width || '0px';
  height = height || '0px';
  const iframeRef = useRef<any>(null);

  window.addEventListener('message', (event) => {
    if (event.data.id !== id) return;
    if (onMessage) onMessage(event.data);
  });

  useEffect(() => {
    const iframe = iframeRef.current;
    if (!iframe) return;
    iframe.onload = () => {
      iframe.contentWindow.postMessage({ code, data, id }, '*');
    }
  }, [iframeRef.current]);

  useEffect(() => {
    const iframe = iframeRef.current;
    if (!iframe) return;
    iframe.contentWindow.postMessage({ code, data, id }, '*');
  }, [code, data]);

  return (
    <iframe
      ref={iframeRef}
      src="/iframe.html"
      sandbox="allow-scripts"
      scrolling="no"
      style={{ width, height, border: 'none' }}
    />
  );
};

export default SandboxIframe;
