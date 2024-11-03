import React, { useState, useEffect } from "react";
import "./App.css";
import ItemList from "./components/ItemList";
import LoginForm from "./components/LoginForm";

function App() {
  const [token, setToken] = useState(localStorage.getItem("jwtToken") || "");

  const handleLogin = (jwtToken) => {
    setToken(jwtToken);
  };

  const handleLogout = () => {
    localStorage.removeItem("jwtToken");
    setToken("");
  };

  useEffect(() => {
    if (token) {
      localStorage.setItem("jwtToken", token);
    }
  }, [token]);

  return (
    <div className="App">
      <h1>Items App</h1>
      {token ? (
        <>
          <button onClick={handleLogout}>Logout</button>
          <ItemList token={token} />
        </>
      ) : (
        <LoginForm onLogin={handleLogin} />
      )}
    </div>
  );
}

export default App;
