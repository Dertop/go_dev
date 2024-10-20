import api from './api';

export interface Item {
  ID: number;
  name: string;
  bought: boolean;
}

export const getItems = async () => {
  const response = await api.get<Item[]>('/items');
  return response.data;
};

export const createItem = async (item: Omit<Item, 'ID'>) => {
  const response = await api.post<Item>('/items', item);
  return response.data;
};

export const updateItem = async (id: number, item: Partial<Item>) => {
  const response = await api.put<Item>(`/items/${id}`, item);
  return response.data;
};

export const deleteItem = async (id: number) => {
  await api.delete(`/items/${id}`);
};
