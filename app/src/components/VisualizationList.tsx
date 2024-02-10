import { Visualization } from './Visualization';

export const VisualizationList = ({ projectId, visualizationIds, data, onChange }) => {
  return (
    <div>
      {visualizationIds.map(visualizationId => (
        <Visualization
          key={visualizationId}
          projectId={projectId}
          visualizationId={visualizationId}
          data={data}
          onChange={onChange} />
      ))}
    </div>
  );
};

export default VisualizationList;
