import React, { useState } from 'react';
import { createItem } from '../services/itemService';

const ItemForm: React.FC<{ onItemCreated: () => void }> = ({ onItemCreated }) => {
  const [name, setName] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (name.trim()) {
      await createItem({ name, bought: false });
      setName('');
      onItemCreated();
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="text"
        value={name}
        onChange={(e) => setName(e.target.value)}
        placeholder="Add a new item"
      />
      <button type="submit">Add</button>
    </form>
  );
};

export default ItemForm;
