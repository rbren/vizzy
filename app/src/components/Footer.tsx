import { Link } from './ui/Link';

export const Footer = () => {

  return (
    <div className="bg-base-100 pt-5 mt-10 border-t-2 border-neutral">
      <span className="flex justify-between md:inline">
        <span className="mr-4">
          <span>Questions? Feedback? </span>
          <Link url="https://github.com/rbren/vizzy">
            Open an issue on GitHub.
          </Link>
        </span>
        <span>
          <span> Have an interesing LLM-based project? </span>
          <Link url="mailto:contact@rbren.io">Get in touch!</Link>
        </span>
      </span>

    </div>
  );
};

export default Footer;
