import * as SecureStore from "expo-secure-store";
import { api } from "./api";
import { PagedResponse } from "@/models/pagedResponse";
import { ItemResponse } from "@/models/item/itemResponse.model";

export const fetchItems = async (page: number, pageSize: number): Promise<PagedResponse<ItemResponse>> => {
  const response = await api.get(`/item`, {
    params: { page, pageSize },
  });

  if (response.data.token) {
    await SecureStore.setItemAsync(process.env.EXPO_PUBLIC_TOKEN_KEY || "token", response.data.token);
  }

  return response.data;
};
