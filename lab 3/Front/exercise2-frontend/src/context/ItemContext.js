import React, { createContext, useContext, useState } from "react";
import axios from "axios";

const ItemContext = createContext();

export const useItems = () => useContext(ItemContext);

export const ItemProvider = ({ children }) => {
  const [items, setItems] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  const fetchItems = async () => {
    setLoading(true);
    setError(null);
    try {
      const response = await axios.get("http://localhost:8080/items");
      setItems(response.data || []);
    } catch (err) {
      setError("Failed to load items. Please try again later.");
    } finally {
      setLoading(false);
    }
  };

  const addItem = async (name, quantity) => {
    setLoading(true);
    setError(null);
    try {
      const response = await axios.post("http://localhost:8080/items", {
        name,
        quantity: parseInt(quantity),
      });
      setItems((prevItems) => [...prevItems, response.data]);
    } catch (err) {
      setError("Failed to add item. Please try again.");
    } finally {
      setLoading(false);
    }
  };

  const deleteItem = async (id) => {
    setLoading(true);
    setError(null);
    try {
      await axios.delete(`http://localhost:8080/items/${id}`);
      setItems((prevItems) => prevItems.filter((item) => item.id !== id));
    } catch (err) {
      setError("Failed to delete item. Please try again.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <ItemContext.Provider
      value={{ items, loading, error, fetchItems, addItem, deleteItem }}
    >
      {children}
    </ItemContext.Provider>
  );
};
