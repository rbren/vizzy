import { useModal } from '../context/ModalContext';

import { Alert } from './ui/Alert';
import { Modal } from './ui/Modal';

export const ERROR_MODAL_NAME = 'error';

export const ErrorModal = () => {
  const { currentModal, modalMessage, hideModal } = useModal();
  if (currentModal !== ERROR_MODAL_NAME) {
    return <></>;
  }

  return (
    <Modal name={ERROR_MODAL_NAME} title="Oops...">
      <Alert level="error" robot="oops" message={modalMessage} />
      <div className="modal-action">
        <button className="btn" onClick={() => hideModal()}>Got it</button>
      </div>
    </Modal>
  );
};

export default ErrorModal;

