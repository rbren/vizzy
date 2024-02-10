import React from 'react';

import { Link } from './ui/Link';
import { Robot } from './ui/Robot';

type FieldDetailsProps = {
  projectId: string;
  fieldsMetadata: any;
};

export const FieldDetails: React.FC<FieldDetailsProps> = ({ projectId, fieldsMetadata }) => {
  const formatValue = (value) => {
    if (typeof value === 'number') {
      return parseFloat(value.toFixed(2));
    }
    if (Array.isArray(value)) {
      return value.map(formatValue).join(', ');
    }
    if (typeof value === 'string') {
      if (value.length > 50) {
        return JSON.stringify(value.substring(0, 50) + '...');
      }
    }
    return JSON.stringify(value);
  }

  let allFields = Object.keys(fieldsMetadata);
  let allSubFields = allFields.map(field => Object.keys(fieldsMetadata[field] || {})).flat();
  allSubFields = Array.from(new Set(allSubFields));
  const preferredOrder = ['$dataPoints', 'sampleValue', 'enum', 'average', 'minimum', 'maximum', 'mode', 'mostCommonValue'];
  allSubFields.sort((a, b) => {
    for (let i = 0; i < preferredOrder.length; i++) {
      if (a === preferredOrder[i]) return -1;
      if (b === preferredOrder[i]) return 1;
    }
    return a.localeCompare(b);
  });

  return (
    <div>
      <p className="mb-4 text-lg">
        <Robot emotion="crystal-ball" className="mr-2" />
        <span className="align-middle">
          <Link url={`/api/projects/${projectId}/data`} text="This data set" />
          <span> </span>
          <span>
            has {fieldsMetadata.$dataPoints} entries with the following fields:
          </span>
        </span>
      </p>
      <table>
        <thead>
        <tr>
          <th className="pr-4">Field</th>
          {allSubFields.map((fieldKey) => {
            return (
              <th key={fieldKey + "-key"} className="pr-2">{fieldKey}</th>
            );
          })}
        </tr>
        </thead>
        <tbody>
        {allFields.map(key => {
          if (key === '$dataPoints') return null;
          return (
            <tr key={key}>
              <th className="pr-4">{key}</th>
              {allSubFields.map((fieldKey) => {
                return (
                  <td key={fieldKey + "-key"} className="pr-2">
                    {fieldsMetadata[key][fieldKey] !== undefined && (
                      <span>{formatValue(fieldsMetadata[key][fieldKey])}</span>
                    )}
                  </td>)
              })}
            </tr>
          );
        })}
        </tbody>
      </table>
    </div>
  );
};

export default FieldDetails;

