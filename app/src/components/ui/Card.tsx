import React from "react";
import { Spinner } from './Spinner.tsx'

type CardProps = {
  children?: React.ReactNode;
  title?: string;
  isLoading?: boolean;
  className?: string;
  center?: boolean;
}

export const Card : React.FC<CardProps> = ({ children, title, isLoading, className, center=true }) => {
  if (className === undefined) {
    className = "";
  }
  if (center) {
    className += " items-center text-center";
  } else {
    className += " items-start text-left";
  }
  return (
    <div className="card card-compact bordered bg-neutral shadow-lg my-6 -mx-8 sm:mx-0">
      <div className={"card-body " + className}>
        { title && <h2 className="card-title">{title}</h2> }
        {isLoading ? <Spinner /> : children}
      </div>
    </div>
  );
};

export default Card;
