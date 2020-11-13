import {useState} from 'react';

export const useInputChange = <T extends Object>(): [T, (e: React.ChangeEvent<HTMLInputElement>) => void] => {
  const [input, setInput] = useState<T>(Object)

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInput({
      ...input,
      [e.currentTarget.name]: e.currentTarget.value
    });
  };

  return [input, handleInputChange]
}