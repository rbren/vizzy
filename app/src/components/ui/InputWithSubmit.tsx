import { SubmitButton } from './SubmitButton'

type InputWithSubmitProps = {
  value?: string,
  onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void,
  placeholder?: string,
  submitText?: string,
  isLoading?: boolean,
}

export const InputWithSubmit: React.FC<InputWithSubmitProps>  = ({ value='', placeholder='', submitText='Submit', isLoading=false, onChange }) => {
  return (
  <>
    <div className="flex grow">
      <input
        type="text"
        placeholder={placeholder}
        className="input rounded-r-none hover:outline-none focus:outline-none focus-within:outline-none grow"
        value={value}
        onChange={onChange} />
      <SubmitButton className="rounded-l-none grow-0" text={submitText} isLoading={isLoading} />
    </div>
  </>
  )
}

