import { Spinner } from './Spinner'
import { Robot } from './Robot'

type AlertProps = {
  children?: any,
  level?: string,
  robot?: string,
  size?: string,
  message?: string,
  spinner?: boolean,
  className?: string
}

export const Alert : React.FC<AlertProps>= ({children, level='info', robot='happy2', size='md', message='', spinner=false, className=''}) => {
  className += ` bg-transparent text-inherit border-2 border-${level} rounded-lg alert alert-${level}`
  let robotSize = 12;
  if (size === 'sm') {
    robotSize = 8;
    className += ' p-3';
  }
  return (
  <div role="alert" className={className}>
    <Robot emotion={robot} size={robotSize} />
    { children || (<span>{message}</span>) }
    { spinner && (<Spinner />) }
  </div>
  )
}

export default Alert;
