import React from 'react';

import { useModal } from '../../context/ModalContext';

export const Modal: React.FC<any> = ({name, title, children}) => {
  const { currentModal } = useModal();

  return (
    <div className={`modal ${currentModal === name ? 'modal-open' : ''}`}>
      <div className="modal-box bordered bg-neutral">
        <h3 className="font-bold text-lg">{title}</h3>
        <div className="w-full my-4 text-left">
          {children}
        </div>
      </div>
    </div>
  );
};

export default Modal;

