import { ProjectsDisplay } from './ProjectsDisplay.tsx'
import { Alert } from './ui/Alert.tsx'
import { Link } from './ui/Link.tsx'

export const Home: React.FC = () => {
  return (
    <div>
      <Alert level="info" robot="happy">
      <h1 className="text-left">
      <b className="text-lg">Welcome!</b>
      <p className="mt-2">
        I can help you <b className="text-info">visualize any kind of data</b>.
        Use this form to give me a new file, or check out one of the public projects below.
      </p>
      <p className="mt-2">
        I currently have a <span className="text-success">77.3% success rate</span>.
        Help me improve by <Link url="https://github.com/rbren/vizzy">contributing on GitHub</Link>!
      </p>
      </h1>
      </Alert>
      <ProjectsDisplay />
    </div>
  )
}

