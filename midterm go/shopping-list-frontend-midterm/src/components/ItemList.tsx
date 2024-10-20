import React, { useEffect, useState } from 'react';
import { Item, getItems, deleteItem, updateItem } from '../services/itemService';

const ItemList: React.FC = () => {
  const [items, setItems] = useState<Item[]>([]);

  useEffect(() => {
    fetchItems();
  }, []);

  const fetchItems = async () => {
    const itemsData = await getItems();
    setItems(itemsData);
  };

  const handleDelete = async (id: number) => {
    await deleteItem(id);
    fetchItems();
  };

  const handleToggleBought = async (id: number, bought: boolean) => {
    await updateItem(id, { bought: !bought });
    fetchItems();
  };

  return (
    <div>
      <h2>Shopping List</h2>
      <ul>
        {items.map((item) => (
          <li key={item.ID}>
            <span style={{ textDecoration: item.bought ? 'line-through' : 'none' }}>
              {item.name}
            </span>
            <button onClick={() => handleToggleBought(item.ID, item.bought)}>
              {item.bought ? 'Unmark' : 'Mark as Bought'}
            </button>
            <button onClick={() => handleDelete(item.ID)}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default ItemList;
