import React from 'react'
import { Spinner } from './Spinner.tsx'

type Props = {
  text?: string
  isLoading?: boolean
  onClick?: any
  className?: string
}


export const SubmitButton : React.FC<Props> = ({ text='', isLoading=false, onClick, className='' }) => (
  <button
      type="submit"
      className={"btn btn-primary " + className}
      disabled={isLoading}
      onClick={onClick} >
    {isLoading ? <Spinner /> : (text || 'Submit')}
  </button>
)

