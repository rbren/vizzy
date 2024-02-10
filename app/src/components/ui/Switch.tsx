export const Switch = ({checked, onChange, label}) => (
  <label className="label cursor-pointer flex items-center">
    <span className="label-text text-xs mr-1">{label}</span>
    <input
      type="checkbox"
      className="toggle toggle-primary toggle-sm"
      checked={checked}
      onChange={onChange}
    />
  </label>
);


