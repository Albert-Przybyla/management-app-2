import * as SecureStore from "expo-secure-store";
import { api } from "./api";

export const login = async (email: string, password: string) => {
  const response = await api.post(`/login`, { email, password });
  if (response.data.token) {
    await SecureStore.setItemAsync(process.env.EXPO_PUBLIC_TOKEN_KEY || "token", response.data.token);
  }
  return response.data;
};

export const getProfile = async () => {
  const token = await SecureStore.getItemAsync("token");
  if (token) {
    const response = await api.get(`/me`, {
      headers: { Authorization: `Bearer ${token}` },
    });
    return response.data;
  } else {
    throw new Error("No token found");
  }
};

export const logout = async () => {
  await SecureStore.deleteItemAsync("token");
};
