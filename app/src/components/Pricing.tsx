import { useState } from 'react';
import { InputWithSubmit } from './ui/InputWithSubmit';
import { Card } from './ui/Card';
import { Alert } from './ui/Alert';

import { request } from '../apiService';

export const Pricing = () => {
  const [loading, setLoading] = useState(false);
  const [email, setEmail] = useState('');
  const [submitted, setSubmitted] = useState(false);
  const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setLoading(true);
    try {
      await request('POST', '/email?email=' + encodeURIComponent(email));
      setSubmitted(true);
    } catch (e) {
      console.error(e);
    }
    setLoading(false);
  }
  return (
    <Card>
    <div className="text-left">
      <h1 className="text-4xl mb-4">Stay in touch!</h1>
      <p className="mb-4">
        Currently Vizzy is free and open source,
        so you'll need to connect your own OpenAI account to use it.
        But if there's sufficient interest, we may consider
        offering Vizzy as a service. Potential features would include:
      </p>
      <ul className="mb-4 list-disc ml-4">
        <li>Private projects</li>
        <li>Build without needing OpenAI account</li>
        <li>User management and access control</li>
      </ul>
      <p className="mb-4">
        If you're interested in these features, or want to stay
        informed as the project develops, send us your email!
      </p>
      <form className="form-control" onSubmit={onSubmit}>
        <InputWithSubmit
          placeholder="Email"
          isLoading={loading}
          value={email}
          onChange={e => setEmail(e.target.value)}
          submitText="Submit"/>
      </form>
    </div>
    { submitted && (
      <Alert level="success">
        Thanks! We'll keep you updated.
      </Alert>
    )}
    </Card>
  );
}
