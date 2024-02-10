import React, { useState, ChangeEvent, FormEvent, useEffect } from 'react';
import { InformationCircleIcon } from '@heroicons/react/24/outline'

import { useModal } from '../context/ModalContext';

import Alert from './ui/Alert';
import Link from './ui/Link';
import Modal from './ui/Modal';
import { getItem, setItem } from '../storage';

export const TOKEN_MODAL_NAME = 'openai-token';

export const SettingsModal: React.FC = () => {
  const { hideModal } = useModal();

  const existingToken = getItem('openai_api_token');
  const [token, setToken] = useState<string>(existingToken || '');
  const [models, setModels] = useState<string[]>([]);
  const [selectedModel, setSelectedModel] = useState<string>('');

  useEffect(() => {
    const savedModel = getItem('openai_model');
    if (savedModel) {
      setSelectedModel(savedModel);
    }
    if (token) {
      fetchModels(token);
    }
  }, []);

  const fetchModels = async (apiToken: string) => {
    const response = await fetch('https://api.openai.com/v1/models', {
      headers: {
        'Authorization': `Bearer ${apiToken}`
      }
    });

    const preferredModels = ['gpt-4-0125-preview', 'gpt-4-1106-preview', 'gpt-4-0613', 'gpt-4', 'gpt-3', 'davinci'];
    if (response.ok) {
      const data = await response.json();
      const ids = data.data.map((model: any) => model.id);
      let sortedModels = ids.sort((id1, id2) => {
        let idx1 = preferredModels.indexOf(id1);
        let idx2 = preferredModels.indexOf(id2);
        if (idx1 !== -1 && idx2 !== -1) {
          return idx1 - idx2;
        } else if (idx1 !== -1) {
          return -1;
        } else if (idx2 !== -1) {
          return 1;
        }
        if (id1.startsWith('gpt-4') && !id2.startsWith('gpt-4')) {
          return -1;
        }
        if (!id1.startsWith('gpt-4') && id2.startsWith('gpt-4')) {
          return 1;
        }
        if (id1.startsWith('gpt-3') && !id2.startsWith('gpt-3')) {
          return -1;
        }
        if (!id1.startsWith('gpt-3') && id2.startsWith('gpt-3')) {
          return 1;
        }
        return id1.length - id2.length;
      });
      setModels(sortedModels);
      setSelectedModel(sortedModels[0]);
    } else {
      alert("Error reaching out to OpenAI. Please check your token");
    }
  };

  const handleSubmit = (e: FormEvent<HTMLFormElement>): void => {
    e.preventDefault();
    setItem('openai_api_token', token);
    if (models.length) {
      setItem('openai_model', selectedModel);
      hideModal();
    } else {
      fetchModels(token);
    }
  };

  const handleModelChange = (e: ChangeEvent<HTMLSelectElement>): void => {
    setSelectedModel(e.target.value);
    setItem('openai_model', e.target.value);
  };

  return (
    <Modal name={TOKEN_MODAL_NAME} title="Connect to OpenAI">
        <p className="mb-2">
          <InformationCircleIcon className="inline-block w-5 h-5 mr-1" />
          You can get an API token at
          <span> </span>
          <Link url="https://platform.openai.com/api-keys" text="platform.openai.com/api-keys">
          </Link>
        </p>
        <p className="text-xs mb-2">
          Your key will only be saved in your browser. It will not be stored on our servers.
        </p>
        <p className="text-xs">
          <Link url="/email">Don't want to use your own key?</Link>
        </p>
        <form onSubmit={handleSubmit}>
          <div className="form-control my-2">
            <label className="label">
              <span className="label-text">API Token</span>
            </label>
            <input
              type="password"
              placeholder="API Token"
              className="input input-bordered"
              value={token}
              onChange={(e: ChangeEvent<HTMLInputElement>) => setToken(e.target.value)}
            />
          </div>
          {models.length > 0 && (
            <div className="form-control mb-2">
              <label className="label">
                <span className="label-text">Select a Model</span>
              </label>
              <select className="select select-bordered" value={selectedModel} onChange={handleModelChange}>
                {models.map(model => (
                  <option key={model} value={model}>{model}</option>
                ))}
              </select>
            </div>
          )}
          <Alert level="warning" robot="money" className="mt-5 mb-4">
            <div className="w-full">
              <p>
                <b>You're responsible for any charges incurred while using this app.</b>{' '}
                Typical usage incurs less than $1 per project, but depends on user behavior and is
                subject to bugs. You're encouraged to{' '}
                <Link url="https://platform.openai.com/usage" text="set spending limits">
                </Link>
                {' '}on your OpenAI account.
              </p>
            </div>
          </Alert>
          <div className="modal-action">
            <button type="submit" className="btn btn-primary">Save</button>
            <button className="btn" onClick={() => hideModal()}>Cancel</button>
          </div>
        </form>
    </Modal>
  );
};

export default SettingsModal;

