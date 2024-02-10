import {
  BrowserRouter as Router,
  Routes,
  Route,
} from "react-router-dom";

import './App.css'

import { Navbar } from './components/Navbar.tsx'
import { Footer } from './components/Footer.tsx'

import { Home } from './components/Home.tsx'
import { Uploader } from './components/Uploader.tsx'
import { Logo } from './components/Logo.tsx'
import { Project } from './components/Project.tsx'
import { Pricing } from './components/Pricing.tsx'
import { EditProject } from './components/EditProject.tsx'

function App() {
  return (
      <Router>
        <Navbar />
        <div id="route">
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/new" element={<Uploader />} />
            <Route path="/logo" element={<Logo />} />
            <Route path="/pricing" element={<Pricing />} />
            <Route path="/projects/:id" element={<Project />} />
            <Route path="/projects/:id/edit" element={<EditProject />} />
          </Routes>
        </div>
        <Footer />
      </Router>
  )
}

export default App
