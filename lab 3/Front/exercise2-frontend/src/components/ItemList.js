import React, { useEffect, useState } from "react";
import axios from "axios";
import ItemForm from "./ItemForm";

const ItemList = ({ token }) => {
  const [items, setItems] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  const fetchItems = async () => {
    setLoading(true);
    setError(null);
    try {
      const response = await axios.get("http://localhost:8080/items", {
        headers: { Authorization: `Bearer ${token}` },
      });
      setItems(response.data || []);
    } catch (err) {
      setError("Failed to load items.");
    } finally {
      setLoading(false);
    }
  };

  const addItem = async (newItem) => {
    setLoading(true);
    setError(null);
    try {
      const response = await axios.post(
        "http://localhost:8080/items",
        newItem,
        { headers: { Authorization: `Bearer ${token}` } }
      );
      setItems((prevItems) => [...prevItems, response.data]);
    } catch (err) {
      setError("Failed to add item.");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchItems();
  }, [token]);

  if (loading) return <p>Loading...</p>;
  if (error) return <p style={{ color: "red" }}>{error}</p>;

  return (
    <div>
      <h2>Items</h2>
      <ItemForm onAdd={addItem} />
      <ul>
        {items.map((item) => (
          <li key={item.id}>
            {item.name} - Quantity: {item.quantity}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default ItemList;
