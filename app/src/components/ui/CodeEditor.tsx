import { useState } from 'react';
import AceEditor from 'react-ace';

// Import required editor themes and modes
import 'ace-builds/src-noconflict/mode-javascript';
import 'ace-builds/src-noconflict/theme-monokai';

type CodeEditorProps = {
  code?: string;
  onChange?: (newCode: string) => void;
};

export const CodeEditor : React.FC<CodeEditorProps> = ({code, onChange}) => {
  const [editedCode, setEditedCode] = useState(code || '');

  const handleCodeChange = (newCode) => {
    setEditedCode(newCode);
    if (onChange) onChange(newCode);
  };

  return (
    <AceEditor
      mode="javascript"
      theme="monokai"
      name="codeEditor"
      value={editedCode}
      onChange={handleCodeChange}
      fontSize={14}
      showPrintMargin={true}
      showGutter={true}
      highlightActiveLine={true}
      setOptions={{
        enableBasicAutocompletion: true,
        enableLiveAutocompletion: true,
        enableSnippets: true,
        showLineNumbers: true,
        tabSize: 2,
      }}
      style={{ width: '100%', height: '500px' }}
    />
  );
};

export default CodeEditor;
