import { useEffect, useState } from 'react';

const OOPS_ROBOTS = ['cursing', 'dead', 'dizzy', 'explode', 'head-explode', 'melting', 'sick', 'sob', 'wonky'];

const sizes = {
  8: 'w-8 h-8',
  12: 'w-12 h-12',
}
export const Robot = ({emotion, size=12, className=''}) => {
  const [robotImage, setRobotImage] = useState(emotion);
  emotion = emotion || 'normal';
  useEffect(() => {
    if (emotion === 'oops') {
      setRobotImage(OOPS_ROBOTS[Math.floor(Math.random() * OOPS_ROBOTS.length)]);
    } else {
      setRobotImage(emotion);
    }
  }, [emotion]);

  const robotExt = robotImage === 'normal' ? 'svg' : 'png';
  const filename = `/robots/${robotImage}.${robotExt}`;
  className = (className || '') + ' ' + (sizes[size] || sizes[12]) + ' inline-block';
  return <img src={filename} className={className + ' -mt-2 ' + sizes[size] || sizes[12]} />
}

