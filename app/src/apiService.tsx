const BASE_URL = (import.meta.env.VITE_API_URL || '') + '/api';

import { TOKEN_MODAL_NAME } from './components/SettingsModal';
import { FORK_MODAL_NAME } from './components/ForkModal';
import { ERROR_MODAL_NAME } from './components/ErrorModal';
import { showModalGlobal } from './context/ModalContext';
import { getItem } from './storage';

function extractProjectID(path) {
  const regex = /\/projects\/([^\/]+)/;
  const match = path.match(regex);
  return match ? match[1] : null;
}

const ERROR_MESSAGES = {
  413: "That's a bit too much data for me. I can only handle up to 100MB.",
  404: "Looks like I lost track of that item.",
  401: "Looks like you're not allowed to do that.",
  403: "Looks like you're not allowed to do that.",
  500: "Some wires must have gotten crossed somewhere. We'll look into it.",
}

const request = async (method:string, url:string, body?:any, handleErrors?:boolean) => {
  if (handleErrors === undefined) {
    handleErrors = method.toUpperCase() !== 'GET';
  }
  const needsOpenAIKey = method.toUpperCase() !== 'GET';
  const needsProjectKey = method.toUpperCase() !== 'GET' && !url.endsWith('/fork');

  const opts = {method, body, headers: {}};
  if (body && typeof body !== 'string' && !(body instanceof File)) {
    opts.body = JSON.stringify(body);
    opts.headers['Content-Type'] = 'application/json';
  }
  const openAIKey = getItem('openai_api_token');
  if (!openAIKey && needsOpenAIKey) {
    if (handleErrors) showModalGlobal(TOKEN_MODAL_NAME);
    return Promise.reject(new Error('Need an OpenAI API key'));
  }
  if (openAIKey) {
    opts.headers['X-OPENAI-API-KEY'] = openAIKey;
  }
  const openAIModel = getItem('openai_model');
  if (openAIModel) {
    opts.headers['X-OPENAI-MODEL'] = openAIModel;
  }
  const projectID = extractProjectID(url);
  if (projectID) {
    const projects = JSON.parse(getItem('projects') || '{}');
    if (!projects[projectID] && needsProjectKey) {
      if (handleErrors) showModalGlobal(FORK_MODAL_NAME);
      return Promise.reject(new Error('Need a project key'));
    }
    opts.headers['X-PROJECT-KEY'] = projects[projectID];
  }

  const resp = await fetch(`${BASE_URL}${url}`, opts);
  if (handleErrors && (resp.status >= 400 || !resp.ok)) {
    const message = ERROR_MESSAGES[resp.status] || "Something's gone terribly wrong.";
    showModalGlobal(ERROR_MODAL_NAME, message);
    return Promise.reject(new Error(message));
  }
  return resp;
}

const loadProjectMetadata = async (projectId) => {
  let response = await request('GET', `/projects/${projectId}/metadata`);
  if (response.status === 404) {
    console.log(`metadata for project ${projectId} not found, attempting to analyze`);
    response = await request('POST', `/projects/${projectId}/analyze`);
  }
  const metadata = await response.json();
  return metadata;
};

const loadFieldsCode = async (projectId, force=false) => {
  let response = force ? null : await request('GET', `/projects/${projectId}/fields-code`);
  if (!response || response.status === 404) {
    response = await request('POST', `/projects/${projectId}/fields-code`);
  }
  const result = await response.json();
  return result.code;
}

export { request, loadProjectMetadata, loadFieldsCode };
