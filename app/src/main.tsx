import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'

import { ModalProvider } from './context/ModalContext.tsx'
import { SettingsModal } from './components/SettingsModal.tsx'
import { ErrorModal } from './components/ErrorModal.tsx'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <ModalProvider>
    <SettingsModal />
    <ErrorModal />
    <App />
  </ModalProvider>
)
