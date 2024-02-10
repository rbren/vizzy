import { createContext, useContext, useState, useEffect } from 'react';

const ModalContext = createContext<any>({});

export const useModal = () => useContext(ModalContext);

export const showModalGlobal = (name, message='') => {
  window.dispatchEvent(new CustomEvent('showModal', { detail: {name, message} }));
}

export const hideModalGlobal = () => {
  window.dispatchEvent(new CustomEvent('hideModal'));
}

export const ModalProvider = ({ children }) => {
  const [currentModal, setCurrentModal] = useState('');
  const [modalMessage, setModalMessage] = useState('');

  const showModal = (name, message='') => {
    setCurrentModal(name);
    setModalMessage(message);
  }
  const hideModal = () => {
    setCurrentModal('');
  }
  const hideOnEscape = (event) => {
    if (event.key === "Escape") {
      hideModal();
    }
  }

  useEffect(() => {
    const showModalEvent = (e) => showModal(e.detail.name, e.detail.message);
    window.addEventListener('showModal', showModalEvent);
    return () => {
      window.removeEventListener('showModal', showModalEvent);
    };
  }, []);

  useEffect(() => {
    const hideModalEvent = () => hideModal();
    window.addEventListener('hideModal', hideModalEvent);
    return () => {
      window.removeEventListener('hideModal', hideModalEvent);
    };
  }, []);

  useEffect(() => {
    document.addEventListener("keydown", hideOnEscape, false);
    return () => {
      document.removeEventListener("keydown", hideOnEscape, false);
    };
  }, []);

  return (
    <ModalContext.Provider value={{ currentModal, showModal, hideModal, modalMessage }}>
      {children}
    </ModalContext.Provider>
  );
};
