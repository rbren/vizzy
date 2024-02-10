import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

import { useModal } from '../context/ModalContext';
import { getItem, setItem } from '../storage.tsx';
import { request } from '../apiService.tsx';

import { Modal } from './ui/Modal';
import { Alert } from './ui/Alert';
import { SubmitButton } from './ui/SubmitButton';

export const FORK_MODAL_NAME = 'fork';

type ForkModalProps = {
  projectId: string;
}

export const ForkModal: React.FC<ForkModalProps> = ({projectId}) => {
  const { hideModal } = useModal();
  const [ isLoading, setIsLoading ] = useState(false);
  const navigate = useNavigate();

  const forkProject = async () => {
    setIsLoading(true);
    try {
      const resp = await request('POST', `/projects/${projectId}/fork`);
      if (!resp.ok) {
        throw new Error(resp.statusText);
      }
      const { uuid, key } = await resp.json();
      const existingData = JSON.parse(getItem('projects') || '{}');
      existingData[uuid] = key;
      setItem('projects', JSON.stringify(existingData));
      navigate('/projects/' + uuid);
      hideModal();
    } catch (error) {
      // TODO: show error to user
      console.error(error);
    }
    setIsLoading(false);
  }

  return (
    <Modal name={FORK_MODAL_NAME} title="Fork Project">
      <Alert level="info" robot="disguise" message="Seems like this project belongs to someone else. You'll need to fork it before making any changes." />
      <div className="modal-action">
        <SubmitButton isLoading={isLoading} onClick={forkProject} text="Let's do it!" />
        <button className="btn" onClick={() => hideModal()}>Cancel</button>
      </div>
    </Modal>
  );
};

export default ForkModal;

