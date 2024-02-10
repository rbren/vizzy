export const Collapse = ({label='', children, className=''}) => {
  return (
    <div tabIndex={0} className={`collapse collapse-arrow ${className}`}>
      <input type="checkbox" />
      <div className="collapse-title text-left leading-8 after:left-0">
        {label}
      </div>
      <div className="collapse-content">
        {children}
      </div>
    </div>
  )
}

