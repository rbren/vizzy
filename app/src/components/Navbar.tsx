import { Link } from 'react-router-dom';
import { TOKEN_MODAL_NAME } from './SettingsModal.tsx'

import { useModal } from '../context/ModalContext';

export const Navbar = () => {
  const { showModal } = useModal();

  return (
    <div className="navbar bg-base-100 mb-10 border-b-2 border-neutral">
      <div className="navbar-start">
        <Link to="/" className="btn btn-ghost normal-case text-xl -ml-6">
          <img className="max-h-full -mt-2" src="/robots/happy2.png"></img>
          <span className="brand-font hidden sm:inline">Vizzy</span>
        </Link>
      </div>
      <div className="navbar-end">
        <a href="https://github.com/rbren/vizzy" className="btn btn-ghost normal-case text-l px-2 sm:px-4">
          GitHub
        </a>
        <a href="#" onClick={() => showModal(TOKEN_MODAL_NAME)} className="btn btn-ghost normal-case text-l -mr-6 px-2 sm:px-4">
          Settings
        </a>
      </div>
    </div>
  );
};

export default Navbar;
