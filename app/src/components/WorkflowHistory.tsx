import { useEffect, useState } from 'react';
import { request } from '../apiService';
import { Link } from './ui/Link';

export const WorkflowHistory = ({onSelect, projectId, visualizationId}) => {
  const [history, setHistory] = useState<any>([]);

  const fetchHistory = async () => {
    const response = await request('GET', `/projects/${projectId}/visualizations/${visualizationId}/versions`);
    if (response.ok) {
      setHistory(await response.json());
    }
  }

  useEffect(() => {
    fetchHistory();
  }, []);

  return (
  <div className="text-left">
  <ul>
    {history.map((item) => (
      <li key={item.version}>
        <span className="min-w-6 inline-block mr-1">{item.version}. </span>
        <Link onClick={() => onSelect(item)}>{item.prompt}</Link>
      </li>
    ))}
  </ul>
  </div>
  )
}
