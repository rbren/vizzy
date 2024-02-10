import React from 'react';

type Props = {
  url?: string;
  text?: string;
  onClick?: () => void;
  children?: React.ReactNode;
}

export const Link : React.FC<Props> = ({url='', text='', onClick, children}) => {
  if (url) {
    return (
      <a href={url} target="_blank"
        className="hover:underline cursor-pointer text-secondary">
        {text || children}
      </a>
    );
  }
  return (<a onClick={onClick}
    className="hover:underline cursor-pointer text-secondary">
    {text || children}
  </a>)
}
export default Link;
