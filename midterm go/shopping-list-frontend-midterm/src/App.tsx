import React, { useState } from 'react';
import ItemList from './components/ItemList';
import ItemForm from './components/ItemForm';
import './App.css';

const App: React.FC = () => {
  const [refresh, setRefresh] = useState(false);

  const handleItemCreated = () => {
    setRefresh(!refresh);
  };

  return (
    <div className="App">
      <ItemForm onItemCreated={handleItemCreated} />
      <ItemList key={refresh ? 1 : 0} />
    </div>
  );
};

export default App;
